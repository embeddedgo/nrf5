// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build nrf52840

// Package uart provides access to the registers of the UART0 peripheral.
//
// Instances:
//  UART0  UART0_BASE  -  UARTE0_UART0*  Universal Asynchronous Receiver/Transmitter
// Registers:
//  0x000 32  TASK_STARTRX  Start UART receiver
//  0x004 32  TASK_STOPRX   Stop UART receiver
//  0x008 32  TASK_STARTTX  Start UART transmitter
//  0x00C 32  TASK_STOPTX   Stop UART transmitter
//  0x01C 32  TASK_SUSPEND  Suspend UART
//  0x100 32  EVENT_CTS     CTS is activated (set low). Clear To Send.
//  0x104 32  EVENT_NCTS    CTS is deactivated (set high). Not Clear To Send.
//  0x108 32  EVENT_RXDRDY  Data received in RXD
//  0x11C 32  EVENT_TXDRDY  Data sent from TXD
//  0x124 32  EVENT_ERROR   Error detected
//  0x144 32  EVENT_RXTO    Receiver timeout
//  0x200 32  SHORTS        Shortcuts between local events and tasks
//  0x304 32  INTENSET      Enable interrupt
//  0x308 32  INTENCLR      Disable interrupt
//  0x480 32  ERRORSRC      Error source
//  0x500 32  ENABLE        Enable UART
//  0x508 32  PSEL_RTS      Pin select for RTS
//  0x50C 32  PSEL_TXD      Pin select for TXD
//  0x510 32  PSEL_CTS      Pin select for CTS
//  0x514 32  PSEL_RXD      Pin select for RXD
//  0x518 32  RXD           RXD register
//  0x51C 32  TXD           TXD register
//  0x524 32  BAUDRATE      Baud rate. Accuracy depends on the HFCLK source selected.
//  0x56C 32  CONFIG        Configuration of parity and hardware flow control
// Import:
//  github.com/embeddedgo/nrf5/p/mmap
package uart

const (
	CTS_STARTRX SHORTS = 0x01 << 3 //+ Shortcut between event CTS and task STARTRX
	NCTS_STOPRX SHORTS = 0x01 << 4 //+ Shortcut between event NCTS and task STOPRX
)

const (
	CTS_STARTRXn = 3
	NCTS_STOPRXn = 4
)

const (
	EOVERRUN ERRORSRC = 0x01 << 0 //+ Overrun error
	EPARITY  ERRORSRC = 0x01 << 1 //+ Parity error
	EFRAMING ERRORSRC = 0x01 << 2 //+ Framing error occurred
	EBREAK   ERRORSRC = 0x01 << 3 //+ Break condition
)

const (
	EOVERRUNn = 0
	EPARITYn  = 1
	EFRAMINGn = 2
	EBREAKn   = 3
)

const (
	EN       ENABLE = 0x0F << 0 //+ Enable or disable UART
	Disabled ENABLE = 0x00 << 0 //  Disable UART
	Enabled  ENABLE = 0x04 << 0 //  Enable UART
)

const (
	ENn = 0
)

const (
	BR         BAUDRATE = 0xFFFFFFFF << 0 //+ Baud rate
	Baud1200   BAUDRATE = 0x4F000 << 0    //  1200 baud (actual rate: 1205)
	Baud2400   BAUDRATE = 0x9D000 << 0    //  2400 baud (actual rate: 2396)
	Baud4800   BAUDRATE = 0x13B000 << 0   //  4800 baud (actual rate: 4808)
	Baud9600   BAUDRATE = 0x275000 << 0   //  9600 baud (actual rate: 9598)
	Baud14400  BAUDRATE = 0x3B0000 << 0   //  14400 baud (actual rate: 14414)
	Baud19200  BAUDRATE = 0x4EA000 << 0   //  19200 baud (actual rate: 19208)
	Baud28800  BAUDRATE = 0x75F000 << 0   //  28800 baud (actual rate: 28829)
	Baud31250  BAUDRATE = 0x800000 << 0   //  31250 baud
	Baud38400  BAUDRATE = 0x9D5000 << 0   //  38400 baud (actual rate: 38462)
	Baud56000  BAUDRATE = 0xE50000 << 0   //  56000 baud (actual rate: 55944)
	Baud57600  BAUDRATE = 0xEBF000 << 0   //  57600 baud (actual rate: 57762)
	Baud76800  BAUDRATE = 0x13A9000 << 0  //  76800 baud (actual rate: 76923)
	Baud115200 BAUDRATE = 0x1D7E000 << 0  //  115200 baud (actual rate: 115942)
	Baud230400 BAUDRATE = 0x3AFB000 << 0  //  230400 baud (actual rate: 231884)
	Baud250000 BAUDRATE = 0x4000000 << 0  //  250000 baud
	Baud460800 BAUDRATE = 0x75F7000 << 0  //  460800 baud (actual rate: 470588)
	Baud921600 BAUDRATE = 0xEBED000 << 0  //  921600 baud (actual rate: 941176)
	Baud1M     BAUDRATE = 0x10000000 << 0 //  1Mega baud
)

const (
	BRn = 0
)

const (
	HWFC     CONFIG = 0x01 << 0 //+ Hardware flow control
	PARITY   CONFIG = 0x07 << 1 //+ Parity
	Excluded CONFIG = 0x00 << 1 //  Exclude parity bit
	Included CONFIG = 0x07 << 1 //  Include parity bit
	STOP     CONFIG = 0x01 << 4 //+ Stop bits
)

const (
	HWFCn   = 0
	PARITYn = 1
	STOPn   = 4
)
