#!/bin/sh

set -e

cd ../../../embeddedgo/nrf5/hal
hal=$(pwd)
cd ../p
rm -rf *

svdxgen github.com/embeddedgo/nrf5/p ../svd/*.svd

perlscript='
s/\)/"github.com\/embeddedgo\/nrf5\/hal\/periph"\n)/;

s/type (TASK_|EVENT_|PSEL_|INTEN|EVTEN)\w* uint32//sg;

s/type RM?(TASK_|EVENT_|PSEL_|INTEN|EVTEN)\w* .*? }//sg;

s/func \(rm? \*?RM?(TASK_|EVENT_|PSEL_|INTEN|EVTEN)\w*\) .*?}//sg;

s/(\s|\])RTASK_\w+/\1periph.Task/g;

s/(\s|\])REVENT_\w+/\1periph.Event/g;

s/(\s|\])RPSEL_\w+/\1periph.$ENV{signal}/g;

s/\sPSEL_/ SIG_/g;

if ( m/INTENSET/ ) {
	s/\(/(\n"embedded\/rtos"/;

	s/\sINTENSET\s+RINTENSET/ INT periph.Events/g;

	s/\sINTENCLR\s+RINTENCLR\n//g;
	
	s/$/\n
	func (p *Periph) IRQ() rtos.IRQ {
		return rtos.IRQ(uintptr(unsafe.Pointer(p))>>12 & 0x1FF)
	}
	/s;
}

if ( m/EVTENSET/ ) {
	s/\sEVTENSET\s+REVTENSET/ EVT periph.Events/g;

	s/\sEVTENCLR\s+REVTENCLR\n//g;

	s/\sEVTEN\s+REVTEN/_ uint32/g;
}

s/func (\w+)0\(\) \*Periph \{ (.*?) }/func \1(n int) *Periph {
switch n {
case 0: \2/sg;
	
s/func \w+(\d)\(\) \*Periph \{ (.*?) }/case \1: \2/sg;

s/(_BASE\)\)\)\n)\n/\1}
return nil
}\n
/sg;
'

for p in ficr gpio nvmc rtc uart uicr; do
	cd $p
	xgen *.go
	GOOS=noos GOARCH=thumb go build -tags nrf52840
	if [ $p != gpio -a $p != ppi ]; then
		mkdir -p $hal/$p
		for f in *.go; do cp -f $f $hal/$p/periph-$f; done
		signal=Digital
		signal=$signal perl -0pi -e "$perlscript" $hal/$p/periph-x*.go
		gofmt -w $hal/$p/periph-x*.go
	fi
	cd ..
done

perlscript='
s/package irq/$&\n\nimport "embedded\/rtos"/;
s/ = \d/ rtos.IRQ$&/g;
'

cd $hal/irq
rm -f *
cp ../../p/irq/* .
perl -pi -e "$perlscript" *.go