// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func die(f string, a ...any) {
	fmt.Fprintf(os.Stderr, f+"\n", a...)
	os.Exit(1)
}

func dieErr(err error) {
	if err != nil {
		die(err.Error())
	}
}

var nordic = filepath.Join("..", "nordic")

func main() {
	dirs, err := ioutil.ReadDir(nordic)
	dieErr(err)
	for _, d := range dirs {
		dname := d.Name()
		if ok, _ := filepath.Match("s[1,3][1,3,4][0-3]", dname); !ok {
			continue
		}
		headers := filepath.Join(nordic, dname, "headers")
		files, err := ioutil.ReadDir(headers)
		dieErr(err)
		for _, f := range files {
			fname := f.Name()
			if ok, _ := filepath.Match("*.h", fname); !ok {
				continue
			}
			switch fname {
			case "nrf_svc.h", "nrf_nvic.h":
				continue
			}
			// TODO: can use WaitGroup to speed up
			g := NewGen(dname, fname)
			g.Run()
		}
	}
}

type Gen struct {
	lex      *Lexer
	sxx      string
	h        string
	pkg      string
	filename string
	gobuf    *bytes.Buffer
	asmbuf   *bytes.Buffer
	comments map[string]string
	imports  map[string]struct{}
}

func NewGen(sxx, h string) *Gen {
	return &Gen{
		sxx:      sxx,
		h:        h,
		comments: make(map[string]string),
		imports:  make(map[string]struct{}),
	}
}

func doNotEdit(w io.Writer, src string) {
	fmt.Fprintf(w, "// Code generated from %s; DO NOT EDIT.\n\n", src)
}

func (g *Gen) printgo(f string, a ...any) {
	if g.gobuf == nil {
		g.gobuf = new(bytes.Buffer)
	}
	fmt.Fprintf(g.gobuf, f, a...)
}

func (g *Gen) printasm(f string, a ...any) {
	if g.asmbuf == nil {
		g.asmbuf = new(bytes.Buffer)
		doNotEdit(g.asmbuf, g.filename)
		fmt.Fprintf(g.asmbuf, "#include \"go_asm.h\"\n")
		fmt.Fprintf(g.asmbuf, "#include \"textflag.h\"\n")
	}
	fmt.Fprintf(g.asmbuf, f, a...)
}

func (g *Gen) Run() {
	var outname string
	switch {
	case g.h == "ble.h":
		g.pkg = "ble"
		outname = "ble"
	case strings.HasPrefix(g.h, "ble_"):
		g.pkg = "ble"
		outname = g.h[4 : len(g.h)-2]
	case strings.HasPrefix(g.h, "nrf_"):
		g.pkg = "nrf"
		outname = g.h[4 : len(g.h)-2]
	default:
		die("unknown file name: %s", g.h)
	}
	outname = "z" + strings.ReplaceAll(outname, "_", "")
	g.filename = filepath.Join(nordic, g.sxx, "headers", g.h)
	buf, err := ioutil.ReadFile(g.filename)
	dieErr(err)
	g.lex = NewLexer(g.filename, string(buf))
	go g.lex.Run()
	for {
		tok := g.lex.Next()
		if tok.Type == Tend {
			break
		}
		if tok.Type != Tident {
			continue
		}
		switch tok.Str {
		case "typedef":
			g.typedef()
		case "enum":
			g.enum()
		case "#define":
			g.define()
		case "SVCALL":
			g.svcall()
		}
	}
	outdir := filepath.Join(g.sxx, g.pkg)
	os.MkdirAll(outdir, 0777)
	if g.gobuf != nil {
		buf := bytes.NewBuffer(make([]byte, 0, g.gobuf.Len()+256))
		doNotEdit(buf, g.filename)
		fmt.Fprintf(buf, "package %s\n\n", g.pkg)
		if len(g.imports) != 0 {
			fmt.Fprintf(buf, "import (\n")
			for imp := range g.imports {
				if strings.HasPrefix(imp, "nrf5/") {
					imp = "github.com/embeddedgo/" + imp
				}
				fmt.Fprintf(buf, "\t\"%s\"\n", imp)
			}
			fmt.Fprintf(buf, ")\n\n")
		}
		buf.Write(g.gobuf.Bytes())
		src, err := format.Source(buf.Bytes())
		//src := buf.Bytes()
		dieErr(err)
		dieErr(ioutil.WriteFile(filepath.Join(outdir, outname+".go"),
			src, 0666))
	}
	if g.asmbuf != nil {
		dieErr(ioutil.WriteFile(filepath.Join(outdir, outname+".s"),
			g.asmbuf.Bytes(), 0666))
	}
}

func (g *Gen) die(i any) {
	switch v := i.(type) {
	case Token:
		v.Die(g.filename)
	default:
		die("%s: %v", g.filename, v)
	}
}

func (g *Gen) expect(tvs ...any) (tok Token) {
	tok = g.lex.Next()
	i := 0
	for i < len(tvs) {
		typ := tvs[i].(TokenType)
		i++
		k := i
		for ; k < len(tvs); k++ {
			if _, ok := tvs[k].(TokenType); ok {
				break
			}
		}
		if typ != tok.Type {
			i = k
			continue
		}
		if k == i {
			return
		}
		for ; i < k; i++ {
			if tvs[i].(string) == tok.Str {
				return
			}
		}
	}
	g.die(tok)
	return
}

func lineComment(s string) string {
	return "// " + strings.Join(strings.Fields(strings.Trim(s, "/**<")), " ")
}

type seItem struct {
	tokens  []Token
	comment string
}

func nextItem(lex *Lexer) (item seItem) {
loop:
	for {
		tok := lex.Next()
		switch tok.Str {
		case ",":
			tok = lex.Next()
			if tok.Type == Tcomment {
				item.comment = tok.Str
			} else {
				lex.Back(tok)
			}
			break loop
		case "}":
			lex.Back(tok)
			if n := len(item.tokens); n > 0 && item.tokens[n-1].Type == Tcomment {
				item.comment = item.tokens[n-1].Str
				item.tokens = item.tokens[:n-1]
			}
			break loop
		}
		item.tokens = append(item.tokens, tok)
	}
	if len(item.comment) != 0 {
		item.comment = lineComment(item.comment)
	}
	return
}

type field struct{ typ, ptr, name, comment string }

func (g Gen) fields(sep, end string) (fields []field) {
	tok := g.expect(Tident, Tsep, end, Tcomment)
	if tok.Type == Tsep {
		return // f()
	}
	if tok.Str == "void" {
		g.expect(Tsep, end)
		return // f(void)
	}
	if tok.Type == Tcomment {
		tok = g.expect(Tident)
	}
	var unsupported string
loop:
	for {
		switch tok.Str {
		case "union":
			unsupported = "union"
			break loop
		case "const":
			g.expect(Tident, "volatile")
			tok = g.expect(Tident)
		case "volatile":
			tok = g.expect(Tident)
		}
		var f field
		f.typ = g.renameType(tok.Str)
		for {
			tok = g.expect(Tident, Tsep, "*")
			if tok.Type == Tident {
				if tok.Str != "const" {
					break
				}
				tok = g.expect(Tsep, "*")
			}
			f.ptr += "*"
		}
		f.name = tok.Str
		if f.ptr != "" {
			f.name = strings.TrimPrefix(f.name, "p_")
		}
		switch f.name {
		case "type":
			f.name = "typ"
		case "chan":
			f.name = "ch"
		}
		if f.typ == "Void" && f.ptr == "*" {
			f.typ = "unsafe.Pointer"
			f.ptr = ""
			g.imports["unsafe"] = struct{}{}
		}
		fields = append(fields, f)
		last := &fields[len(fields)-1]
		tok = g.expect(Tsep, sep, end, ":", "[")
		if tok.Str == end {
			break
		}
		switch tok.Str {
		case "[":
			n := g.expect(Tident).Str
			g.expect(Tsep, "]")
			last.typ = "[" + g.renameConst(n) + "]" + last.typ
			g.expect(Tsep, sep)
		case ":":
			unsupported = "bitfield"
			break loop
		}
		tok = g.expect(Tident, Tcomment, Tsep, end)
		if tok.Type == Tsep {
			break
		}
		if tok.Type == Tcomment {
			fields[len(fields)-1].comment = lineComment(tok.Str)
			tok = g.expect(Tident, Tsep, end)
			if tok.Type == Tsep {
				break
			}
		}
	}
	if unsupported != "" {
		n := 1
		for {
			switch g.lex.Next().Str {
			case "{":
				n++
			case "}":
				n--
			}
			if n == 0 {
				break
			}
		}
		return []field{{name: unsupported}}
	}
	return
}

func (g *Gen) typedef() {
	typ := g.expect(Tident).Str
	if typ == "volatile" {
		typ = g.expect(Tident).Str
	}
	switch typ {
	case "union":
		for {
			tok := g.lex.Next()
			if tok.Type == Tend {
				g.die(tok)
			}
			if tok.Str == "}" {
				g.printgo(
					"\n//type %s union\n\n",
					g.renameType(g.expect(Tident).Str),
				)
				g.expect(Tsep, ";")
				return
			}
		}
	case "struct":
		g.expect(Tsep, "{")
		fields := g.fields(";", "}")
		name := g.renameType(g.expect(Tident).Str)
		g.expect(Tsep, ";")
		if len(fields) == 1 && fields[0].typ == "" {
			g.printgo("\n//type %s struct { %s }\n\n", name, fields[0].name)
			return
		}
		g.printgo("\ntype %s struct {\n", name)
		for _, f := range fields {
			g.printgo("%s %s%s %s\n", g.renameType(f.name), f.ptr, f.typ,
				f.comment)
		}
		g.printgo("}\n\n")
	default:
		tok := g.expect(Tident, Tsep, "(", "*")
		if tok.Type == Tsep {
			// not supported
			for {
				tok = g.lex.Next()
				if tok.Type == Tend {
					g.die(tok)
				}
				if tok.Str == ";" {
					return
				}
			}
		}
		g.printgo("type %s ", g.renameType(tok.Str))
		if g.expect(Tsep, ";", "[").Str == "[" {
			n := g.expect(Tident)
			g.expect(Tsep, "]")
			g.expect(Tsep, ";")
			g.printgo("[%s]", n)
		}
		g.printgo("%s\n", g.renameType(typ))
	}
}

func (g *Gen) enum() {
	tok := g.expect(Tident, Tsep, "{")
	svc := ""
	if tok.Type == Tident {
		if strings.HasSuffix(tok.Str, "_SVCS") {
			svc = "_"
		}
		g.expect(Tsep, "{")
	}
	g.printgo("\nconst (\n")
	for first := true; ; first = false {
		item := nextItem(g.lex)
		if len(item.tokens) == 0 {
			break
		}
		tok := item.tokens[0]
		if tok.Type != Tident {
			g.die(tok)
		}
		name := svc + g.renameConst(tok.Str)
		if item.comment != "" && svc != "" {
			g.comments[name] = item.comment
			item.comment = ""
		}
		if len(item.tokens) == 1 {
			if first {
				g.printgo("\t%s = iota %s\n", name, item.comment)
			} else {
				g.printgo("\t%s %s\n", name, item.comment)
			}
			continue
		}
		tok = item.tokens[1]
		if tok.Str != "=" {
			g.die(tok)
		}
		tok = item.tokens[2]
		if tok.Type != Tident {
			g.die(tok)
		}
		base := g.renameConst(tok.Str)
		switch len(item.tokens) {
		case 3:
			if first {
				g.printgo("\t%s = %s + iota %s\n", name, base, item.comment)
			} else {
				g.printgo("\t%s = %s %s\n", name, base, item.comment)
			}
		case 5:
			tok = item.tokens[3]
			if tok.Str != "+" {
				g.die(tok)
			}
			tok = item.tokens[4]
			if tok.Type != Tident {
				g.die(tok)
			}
			offset := tok.Str
			g.printgo("\t%s = %s + %s %s\n", name, base, offset, item.comment)
		default:
			g.die("unexpected form of enum item")
		}

	}
	g.expect(Tsep, "}")
	g.printgo(")\n\n")
}

func (g *Gen) define() {
	name := g.lex.NextNewLine()
	if name.Type != Tident {
		g.die(name)
	}
	if strings.HasSuffix(name.Str, "_IRQHandler") {
		for {
			if g.lex.NextSpace().Type == Tnewline {
				return
			}
		}
	}
	tok := g.lex.NextSpace()
	if tok.Type != Tspace {
		for g.lex.NextNewLine().Type != Tnewline {
		}
		return // function-like macro
	}
	tok = g.lex.NextNewLine()
	if tok.Type == Tnewline {
		return // define without value
	}
	g.printgo("const %s =", g.renameConst(name.Str))
	for {
		switch tok.Type {
		case Tnewline:
			g.printgo("\n")
			return
		case Tcomment:
			next := g.lex.NextNewLine()
			if next.Type == Tnewline {
				g.printgo(" %s", lineComment(tok.Str))
			} else {
				g.printgo(" %s", tok.Str)
			}
			g.lex.Back(next)
		default:
			var s string
			if strings.HasSuffix(tok.Str, "_t") {
				s = g.renameType(tok.Str)
			} else {
				s = g.renameConst(tok.String())
			}
			g.printgo(" %s", s)
		}
		tok = g.lex.NextNewLine()
	}
}

func align(offset, align uintptr) uintptr {
	align--
	return (offset + align) &^ align
}

func sizeof(typ, ptr string) (size uintptr, mov string) {
	switch {
	case ptr != "" || typ == "uintptr" || typ == "unsafe.Pointer" ||
		strings.HasSuffix(typ, "int32"):
		return 4, "MOVW"
	case typ == "uint16":
		return 2, "MOVHU"
	case typ == "uint8":
		return 1, "MOVBU"
	case typ == "int16":
		return 2, "MOVH"
	case typ == "int8":
		return 1, "MOVB"
	}
	return 0, ""
}

func (g *Gen) svcall() {
	tok := g.lex.NextSpace()
	if tok.Str != "(" {
		g.die(tok)
	}
	svc := "_" + g.renameConst(g.expect(Tident).Str)
	if comment := g.comments[svc]; comment != "" {
		g.printgo("\n" + comment + "\n")
	}
	g.expect(Tsep, ",")
	ret := g.renameType(g.expect(Tident).Str)
	g.expect(Tsep, ",")
	fname := renameFunc(g.expect(Tident).Str)
	g.expect(Tsep, "(")
	args := g.fields(",", ")")
	g.expect(Tsep, ")")
	g.expect(Tsep, ";")
	g.printgo("func %s(", fname)
	for _, a := range args {
		g.printgo("%s %s%s,", a.name, a.ptr, a.typ)
	}
	g.printgo(") %s\n", ret)
	g.printasm("\nTEXT ·%s(SB),NOSPLIT|NOFRAME,$0\n", fname)
	offset := uintptr(0)
	for i, a := range args {
		n, mov := sizeof(a.typ, a.ptr)
		if n == 0 {
			die(a.typ)
		}
		offset = align(offset, n)
		g.printasm("\t%s %s+%d(FP), R%d\n", mov, a.name, offset, i)
		offset += n
	}
	g.printasm("\tSWI $const_%s\n", svc)
	n, mov := sizeof(ret, "")
	if n == 0 {
		die(ret)
	}
	offset = align(offset, n)
	g.printasm("\t%s R0, ret+%d(FP)\n", mov, offset)
	g.printasm("\tRET\n")
}

//// RENAME

var rename = map[string]string{
	"uuid":    "UUID",
	"vs":      "VS",
	"lf":      "LF",
	"hfclk":   "HF",
	"pa":      "PA",
	"lna":     "LNA",
	"uuid128": "UUID128",
	"gap":     "GAP",
	"gatt":    "GATT",
	"gattc":   "GATTC",
	"gatts":   "GATTS",
	"ppi":     "PPI",
	"gpiote":  "GPIOTE",
	"rc":      "RC",
}

func (g *Gen) renameConst(s string) string {
	s = strings.TrimPrefix(s, "SD_")
	if strings.HasSuffix(s, "_IRQn") && !strings.HasPrefix(s, "__") &&
		s != "EVT_IRQn" && s != "RADIO_NOTIFICATION_IRQn" {
		s = s[:len(s)-5]
		if strings.HasPrefix(s, "SWI") {
			s += "_EGU" + s[3:]
		}
		s = "irq." + s
		g.imports["nrf5/p/irq"] = struct{}{}
	}
	if g.pkg == "nrf" {
		s = strings.TrimPrefix(s, "NRF_")
	} else if g.pkg == "ble" {
		if strings.HasPrefix(s, "NRF_") {
			s = s[4:]
			if !strings.HasSuffix(s, "_ERR_BASE") {
				s = "nrf." + s
				g.imports["nrf5/sd/"+g.sxx+"/nrf"] = struct{}{}
			}
		} else {
			s = strings.TrimPrefix(s, "BLE_")
		}
	}
	if len(s) != 0 && unicode.IsDigit(rune(s[0])) {
		s = strings.TrimRight(s, "uUlL")
	}
	return s
}

func renameTypeFunc(s string) []string {
	s = strings.TrimPrefix(s, "ble_")
	s = strings.TrimPrefix(s, "nrf_")
	parts := strings.Split(s, "_")
	for i, part := range parts {
		if newname, ok := rename[part]; ok {
			parts[i] = newname
		} else {
			parts[i] = strings.ToUpper(part[:1]) + part[1:]
		}
	}
	return parts
}

func (g *Gen) renameType(s string) string {
	s = strings.TrimSuffix(s, "_t")
	if strings.HasPrefix(s, "uint") || strings.HasPrefix(s, "int") {
		return s
	}
	parts := renameTypeFunc(s)
	switch parts[len(parts)-1] {
	case "Callback", "Handler":
		g.imports["unsafe"] = struct{}{}
		return "unsafe.Pointer"
	}
	return strings.Join(parts, "")
}

func renameFunc(s string) string {
	s = strings.TrimPrefix(s, "sd_")
	parts := renameTypeFunc(s)
	action := parts[len(parts)-1]
	parts = parts[:len(parts)-1]
	if len(parts) > 0 && parts[len(parts)-1] == "Is" {
		action = "Is" + action
		parts = parts[:len(parts)-1]
	}
	return action + strings.Join(parts, "")
}

//// LEXER

type TokenType int8

const (
	Tnone TokenType = iota
	Tstring
	Tchar
	Tcomment
	Tsep
	Tspace
	Tnewline
	Tident
	Tend
)

type Token struct {
	Type TokenType
	Str  string
	Line int
}

func (tok Token) String() string {
	switch tok.Type {
	case Tcomment, Tstring:
		s := tok.Str
		if len(s) > 16 {
			s = s[:7] + " ... " + s[len(s)-7:]
		}
		s = fmt.Sprintf("%#v", s)
		if tok.Type == Tcomment {
			s = s[1 : len(s)-1]
			if s[1] == '/' {
				s += "\\n"
			}
		}
		return s
	case Tchar:
		return "'" + tok.Str + "'"
	case Tident, Tsep:
		return tok.Str
	case Tnewline:
		return `\n`
	case Tspace:
		return `\s`
	}
	panic("unknown token type")
}

func (tok Token) Die(filename string) {
	if tok.Type == Tend {
		fmt.Fprintf(os.Stderr, "%s: unexpected end of file", filename)
	} else {
		fmt.Fprintf(
			os.Stderr, "%s:%d: unexpected token: %v", filename, tok.Line, tok,
		)
	}
	os.Exit(1)
}

type Lexer struct {
	ch       chan Token
	back     Token
	n        int
	filename string
	text     string
}

func NewLexer(filename, text string) *Lexer {
	return &Lexer{ch: make(chan Token, 3), n: 1, filename: filename, text: text}
}

func (lex *Lexer) Run() {
	lex.stringComment(lex.text)
	close(lex.ch)
	lex.text = ""
}

func (lex *Lexer) Next() Token {
	if lex.back.Type != Tnone {
		tok := lex.back
		lex.back.Type = Tnone
		return tok
	}
	for {
		tok, ok := <-lex.ch
		if !ok {
			return Token{Type: Tend}
		}
		if tok.Type != Tnewline && tok.Type != Tspace {
			return tok
		}
	}
}

func (lex *Lexer) NextNewLine() Token {
	if lex.back.Type != Tnone {
		tok := lex.back
		lex.back.Type = Tnone
		return tok
	}
	for {
		tok, ok := <-lex.ch
		if !ok {
			return Token{Type: Tend}
		}
		if tok.Type != Tspace {
			return tok
		}
	}
}

func (lex *Lexer) NextSpace() Token {
	if lex.back.Type != Tnone {
		tok := lex.back
		lex.back.Type = Tnone
		return tok
	}
	tok, ok := <-lex.ch
	if !ok {
		return Token{Type: Tend}
	}
	return tok
}

func (lex *Lexer) Back(tok Token) {
	if lex.back.Type != Tnone {
		panic("Lexer.Back called twice")
	}
	lex.back = tok
}

func (lex *Lexer) tok(typ TokenType, str string) {
	lex.ch <- Token{Type: typ, Str: str, Line: lex.n}
}

func (lex *Lexer) countLines(text string) {
	for {
		i := strings.IndexByte(text, '\n')
		if i < 0 {
			return
		}
		lex.n++
		text = text[i+1:]
	}
}

func (lex *Lexer) stringComment(text string) {
	for {
		start := strings.IndexAny(text, `/"'`)
		if start < 0 {
			lex.sepIdent(text)
			return
		}
		end := 0
		switch c := text[start]; c {
		case '/':
			if start+1 == len(text) {
				die("%s:%d: unexpecetd '/'", lex.filename, lex.n)
			}
			if text[start+1] == '*' {
				end = strings.Index(text[start+2:], "*/")
				if end < 0 {
					die("%s:%d: unterminated comment", lex.filename, lex.n)
				}
				end += start + 4
			} else if text[start+1] == '/' {
				end = strings.IndexByte(text[start+2:], '\n')
				if end < 0 {
					end = 0
				}
				end += start + 2
			}
			if end != 0 {
				lex.sepIdent(text[:start])
				lex.tok(Tcomment, text[start:end])
				lex.countLines(text[start:end])
			} else {
				end = start + 1
				lex.sepIdent(text[:end])
			}
			text = text[end:]
		default: // " '
			start++
			end = strings.IndexByte(text[start:], c)
			if end < 0 {
				if c == '"' {
					die("%s:%d: unterminated string", lex.filename, lex.n)
				} else {
					die("%s:%d: unterminated char", lex.filename, lex.n)
				}
			}
			end += start
			lex.sepIdent(text[:start-1])
			lex.tok(Tstring, text[start:end])
			lex.countLines(text[start:end])
			text = text[end+1:]
		}
	}
}

func (lex *Lexer) sepIdent(text string) {
	for len(text) > 0 {
		i := strings.IndexAny(text, " \t\n~!%^&*()-+[]{}=|;:<>,.?/")
		if i < 0 {
			lex.tok(Tident, text)
			return
		}
		n := 1
		if text[i] == '\n' {
			if i > 0 && text[i-1] == '\\' {
				i--
				n = 2
			}
		} else if i+1 < len(text) {
			switch text[i+1] {
			case '=':
				switch text[i] {
				case '!', '%', '^', '&', '*', '-', '=', '+', '/', '<', '>', '|':
					n = 2
				}
			case '+', '-':
				if text[i] == text[i+1] {
					n = 2
				}
			case '<', '>', '|', '&':
				if text[i] == text[i+1] {
					n = 2
					if i+2 < len(text) && text[i+2] == '=' {
						n = 3
					}
				}
			}
		}
		if i > 0 {
			lex.tok(Tident, text[:i])
		}
		sep := text[i : i+n]
		text = text[i+n:]
		space := true
		switch sep {
		case "\n":
			lex.tok(Tnewline, sep)
			lex.n++
		case "\\\n":
			lex.tok(Tspace, sep)
			lex.n++
		case " ", "\t":
			lex.tok(Tspace, sep)
		default:
			if sep == "~" {
				sep = "^"
			}
			lex.tok(Tsep, sep)
			space = false
		}
		if space {
			text = strings.TrimLeft(text, " \t")
		}
	}
}
