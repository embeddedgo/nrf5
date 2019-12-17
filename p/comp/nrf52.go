// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build nrf52

// Package comp provides access to the registers of the COMP peripheral.
//
// Instances:
//  COMP  COMP_BASE  -  COMP_LPCOMP*  Comparator
// Registers:
//  0x000 32  TASKS_START   Start comparator
//  0x004 32  TASKS_STOP    Stop comparator
//  0x008 32  TASKS_SAMPLE  Sample comparator value
//  0x100 32  EVENTS_READY  COMP is ready and output is valid
//  0x104 32  EVENTS_DOWN   Downward crossing
//  0x108 32  EVENTS_UP     Upward crossing
//  0x10C 32  EVENTS_CROSS  Downward or upward crossing
//  0x200 32  SHORTS        Shortcuts between local events and tasks
//  0x300 32  INTEN         Enable or disable interrupt
//  0x304 32  INTENSET      Enable interrupt
//  0x308 32  INTENCLR      Disable interrupt
//  0x400 32  RESULT        Compare result
//  0x500 32  ENABLE        COMP enable
//  0x504 32  PSEL          Pin select
//  0x508 32  REFSEL        Reference source select for single-ended mode
//  0x50C 32  EXTREFSEL     External reference select
//  0x530 32  TH            Threshold configuration for hysteresis unit
//  0x534 32  MODE          Mode configuration
//  0x538 32  HYST          Comparator hysteresis enable
// Import:
//  github.com/embeddedgo/nrf5/p/mmap
package comp

const (
	READY_SAMPLE SHORTS = 0x01 << 0 //+ Shortcut between event READY and task SAMPLE
	Disabled     SHORTS = 0x00 << 0 //  Disable shortcut
	Enabled      SHORTS = 0x01 << 0 //  Enable shortcut
	READY_STOP   SHORTS = 0x01 << 1 //+ Shortcut between event READY and task STOP
	Disabled     SHORTS = 0x00 << 1 //  Disable shortcut
	Enabled      SHORTS = 0x01 << 1 //  Enable shortcut
	DOWN_STOP    SHORTS = 0x01 << 2 //+ Shortcut between event DOWN and task STOP
	Disabled     SHORTS = 0x00 << 2 //  Disable shortcut
	Enabled      SHORTS = 0x01 << 2 //  Enable shortcut
	UP_STOP      SHORTS = 0x01 << 3 //+ Shortcut between event UP and task STOP
	Disabled     SHORTS = 0x00 << 3 //  Disable shortcut
	Enabled      SHORTS = 0x01 << 3 //  Enable shortcut
	CROSS_STOP   SHORTS = 0x01 << 4 //+ Shortcut between event CROSS and task STOP
	Disabled     SHORTS = 0x00 << 4 //  Disable shortcut
	Enabled      SHORTS = 0x01 << 4 //  Enable shortcut
)

const (
	READY_SAMPLEn = 0
	READY_STOPn   = 1
	DOWN_STOPn    = 2
	UP_STOPn      = 3
	CROSS_STOPn   = 4
)

const (
	READY    INTEN = 0x01 << 0 //+ Enable or disable interrupt for event READY
	Disabled INTEN = 0x00 << 0 //  Disable
	Enabled  INTEN = 0x01 << 0 //  Enable
	DOWN     INTEN = 0x01 << 1 //+ Enable or disable interrupt for event DOWN
	Disabled INTEN = 0x00 << 1 //  Disable
	Enabled  INTEN = 0x01 << 1 //  Enable
	UP       INTEN = 0x01 << 2 //+ Enable or disable interrupt for event UP
	Disabled INTEN = 0x00 << 2 //  Disable
	Enabled  INTEN = 0x01 << 2 //  Enable
	CROSS    INTEN = 0x01 << 3 //+ Enable or disable interrupt for event CROSS
	Disabled INTEN = 0x00 << 3 //  Disable
	Enabled  INTEN = 0x01 << 3 //  Enable
)

const (
	READYn = 0
	DOWNn  = 1
	UPn    = 2
	CROSSn = 3
)

const (
	READY    INTENSET = 0x01 << 0 //+ Write '1' to enable interrupt for event READY
	Disabled INTENSET = 0x00 << 0 //  Read: Disabled
	Enabled  INTENSET = 0x01 << 0 //  Read: Enabled
	Set      INTENSET = 0x01 << 0 //  Enable
	DOWN     INTENSET = 0x01 << 1 //+ Write '1' to enable interrupt for event DOWN
	Disabled INTENSET = 0x00 << 1 //  Read: Disabled
	Enabled  INTENSET = 0x01 << 1 //  Read: Enabled
	Set      INTENSET = 0x01 << 1 //  Enable
	UP       INTENSET = 0x01 << 2 //+ Write '1' to enable interrupt for event UP
	Disabled INTENSET = 0x00 << 2 //  Read: Disabled
	Enabled  INTENSET = 0x01 << 2 //  Read: Enabled
	Set      INTENSET = 0x01 << 2 //  Enable
	CROSS    INTENSET = 0x01 << 3 //+ Write '1' to enable interrupt for event CROSS
	Disabled INTENSET = 0x00 << 3 //  Read: Disabled
	Enabled  INTENSET = 0x01 << 3 //  Read: Enabled
	Set      INTENSET = 0x01 << 3 //  Enable
)

const (
	READYn = 0
	DOWNn  = 1
	UPn    = 2
	CROSSn = 3
)

const (
	READY    INTENCLR = 0x01 << 0 //+ Write '1' to disable interrupt for event READY
	Disabled INTENCLR = 0x00 << 0 //  Read: Disabled
	Enabled  INTENCLR = 0x01 << 0 //  Read: Enabled
	Clear    INTENCLR = 0x01 << 0 //  Disable
	DOWN     INTENCLR = 0x01 << 1 //+ Write '1' to disable interrupt for event DOWN
	Disabled INTENCLR = 0x00 << 1 //  Read: Disabled
	Enabled  INTENCLR = 0x01 << 1 //  Read: Enabled
	Clear    INTENCLR = 0x01 << 1 //  Disable
	UP       INTENCLR = 0x01 << 2 //+ Write '1' to disable interrupt for event UP
	Disabled INTENCLR = 0x00 << 2 //  Read: Disabled
	Enabled  INTENCLR = 0x01 << 2 //  Read: Enabled
	Clear    INTENCLR = 0x01 << 2 //  Disable
	CROSS    INTENCLR = 0x01 << 3 //+ Write '1' to disable interrupt for event CROSS
	Disabled INTENCLR = 0x00 << 3 //  Read: Disabled
	Enabled  INTENCLR = 0x01 << 3 //  Read: Enabled
	Clear    INTENCLR = 0x01 << 3 //  Disable
)

const (
	READYn = 0
	DOWNn  = 1
	UPn    = 2
	CROSSn = 3
)

const (
	RESULT RESULT = 0x01 << 0 //+ Result of last compare. Decision point SAMPLE task.
	Below  RESULT = 0x00 << 0 //  Input voltage is below the threshold (VIN+ &lt; VIN-)
	Above  RESULT = 0x01 << 0 //  Input voltage is above the threshold (VIN+ &gt; VIN-)
)

const (
	RESULTn = 0
)

const (
	ENABLE   ENABLE = 0x03 << 0 //+ Enable or disable COMP
	Disabled ENABLE = 0x00 << 0 //  Disable
	Enabled  ENABLE = 0x02 << 0 //  Enable
)

const (
	ENABLEn = 0
)

const (
	PSEL         PSEL = 0x07 << 0 //+ Analog pin select
	AnalogInput0 PSEL = 0x00 << 0 //  AIN0 selected as analog input
	AnalogInput1 PSEL = 0x01 << 0 //  AIN1 selected as analog input
	AnalogInput2 PSEL = 0x02 << 0 //  AIN2 selected as analog input
	AnalogInput3 PSEL = 0x03 << 0 //  AIN3 selected as analog input
	AnalogInput4 PSEL = 0x04 << 0 //  AIN4 selected as analog input
	AnalogInput5 PSEL = 0x05 << 0 //  AIN5 selected as analog input
	AnalogInput6 PSEL = 0x06 << 0 //  AIN6 selected as analog input
	AnalogInput7 PSEL = 0x07 << 0 //  AIN7 selected as analog input
)

const (
	PSELn = 0
)

const (
	REFSEL REFSEL = 0x07 << 0 //+ Reference select
	Int1V2 REFSEL = 0x00 << 0 //  VREF = internal 1.2 V reference (VDD &gt;= 1.7 V)
	Int1V8 REFSEL = 0x01 << 0 //  VREF = internal 1.8 V reference (VDD &gt;= VREF + 0.2 V)
	Int2V4 REFSEL = 0x02 << 0 //  VREF = internal 2.4 V reference (VDD &gt;= VREF + 0.2 V)
	VDD    REFSEL = 0x04 << 0 //  VREF = VDD
	ARef   REFSEL = 0x05 << 0 //  VREF = AREF (VDD &gt;= VREF &gt;= AREFMIN)
)

const (
	REFSELn = 0
)

const (
	EXTREFSEL        EXTREFSEL = 0x07 << 0 //+ External analog reference select
	AnalogReference0 EXTREFSEL = 0x00 << 0 //  Use AIN0 as external analog reference
	AnalogReference1 EXTREFSEL = 0x01 << 0 //  Use AIN1 as external analog reference
	AnalogReference2 EXTREFSEL = 0x02 << 0 //  Use AIN2 as external analog reference
	AnalogReference3 EXTREFSEL = 0x03 << 0 //  Use AIN3 as external analog reference
	AnalogReference4 EXTREFSEL = 0x04 << 0 //  Use AIN4 as external analog reference
	AnalogReference5 EXTREFSEL = 0x05 << 0 //  Use AIN5 as external analog reference
	AnalogReference6 EXTREFSEL = 0x06 << 0 //  Use AIN6 as external analog reference
	AnalogReference7 EXTREFSEL = 0x07 << 0 //  Use AIN7 as external analog reference
)

const (
	EXTREFSELn = 0
)

const (
	THDOWN TH = 0x3F << 0 //+ VDOWN = (THDOWN+1)/64*VREF
	THUP   TH = 0x3F << 8 //+ VUP = (THUP+1)/64*VREF
)

const (
	THDOWNn = 0
	THUPn   = 8
)

const (
	SP     MODE = 0x03 << 0 //+ Speed and power modes
	Low    MODE = 0x00 << 0 //  Low-power mode
	Normal MODE = 0x01 << 0 //  Normal mode
	High   MODE = 0x02 << 0 //  High-speed mode
	MAIN   MODE = 0x01 << 8 //+ Main operation modes
	SE     MODE = 0x00 << 8 //  Single-ended mode
	Diff   MODE = 0x01 << 8 //  Differential mode
)

const (
	SPn   = 0
	MAINn = 8
)

const (
	HYST     HYST = 0x01 << 0 //+ Comparator hysteresis
	NoHyst   HYST = 0x00 << 0 //  Comparator hysteresis disabled
	Hyst50mV HYST = 0x01 << 0 //  Comparator hysteresis enabled
)

const (
	HYSTn = 0
)
