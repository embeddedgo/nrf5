// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build nrf52840

// Package spim provides access to the registers of the SPIM peripheral.
//
// Instances:
//  SPIM0  SPIM0_BASE  -  SPIM0_SPIS0_TWIM0_TWIS0_SPI0_TWI0*  Serial Peripheral Interface Master with EasyDMA 0
//  SPIM1  SPIM1_BASE  -  SPIM1_SPIS1_TWIM1_TWIS1_SPI1_TWI1*  Serial Peripheral Interface Master with EasyDMA 1
//  SPIM2  SPIM2_BASE  -  SPIM2_SPIS2_SPI2*                   Serial Peripheral Interface Master with EasyDMA 2
//  SPIM3  SPIM3_BASE  -  SPIM3                               Serial Peripheral Interface Master with EasyDMA 3
// Registers:
//  0x010 32  TASK_START        Start SPI transaction
//  0x014 32  TASK_STOP         Stop SPI transaction
//  0x01C 32  TASK_SUSPEND      Suspend SPI transaction
//  0x020 32  TASK_RESUME       Resume SPI transaction
//  0x104 32  EVENT_STOPPED     SPI transaction has stopped
//  0x110 32  EVENT_ENDRX       End of RXD buffer reached
//  0x118 32  EVENT_END         End of RXD buffer and TXD buffer reached
//  0x120 32  EVENT_ENDTX       End of TXD buffer reached
//  0x14C 32  EVENT_STARTED     Transaction started
//  0x200 32  SHORTS            Shortcuts between local events and tasks
//  0x304 32  INTENSET          Enable interrupt
//  0x308 32  INTENCLR          Disable interrupt
//  0x400 32  STALLSTAT         Stall status for EasyDMA RAM accesses. The fields in this register is set to STALL by hardware whenever a stall occurres and can be cleared (set to NOSTALL) by the CPU.
//  0x500 32  ENABLE            Enable SPIM
//  0x508 32  PSEL_SCK          Pin select for SCK
//  0x50C 32  PSEL_MOSI         Pin select for MOSI signal
//  0x510 32  PSEL_MISO         Pin select for MISO signal
//  0x514 32  PSEL_CSN          Pin select for CSN
//  0x524 32  FREQUENCY         SPI frequency. Accuracy depends on the HFCLK source selected.
//  0x534 32  RXD_PTR           Data pointer
//  0x538 32  RXD_MAXCNT        Maximum number of bytes in receive buffer
//  0x53C 32  RXD_AMOUNT        Number of bytes transferred in the last transaction
//  0x540 32  RXD_LIST          EasyDMA list type
//  0x544 32  TXD_PTR           Data pointer
//  0x548 32  TXD_MAXCNT        Number of bytes in transmit buffer
//  0x54C 32  TXD_AMOUNT        Number of bytes transferred in the last transaction
//  0x550 32  TXD_LIST          EasyDMA list type
//  0x554 32  CONFIG            Configuration register
//  0x560 32  IFTIMING_RXDELAY  Sample delay for input serial data on MISO
//  0x564 32  IFTIMING_CSNDUR   Minimum duration between edge of CSN and edge of SCK and minimum duration CSN must stay high between transactions
//  0x568 32  CSNPOL            Polarity of CSN output
//  0x56C 32  PSELDCX           Pin select for DCX signal
//  0x570 32  DCXCNT            DCX configuration
//  0x5C0 32  ORC               Byte transmitted after TXD.MAXCNT bytes have been transmitted in the case when RXD.MAXCNT is greater than TXD.MAXCNT
// Import:
//  github.com/embeddedgo/nrf5/p/mmap
package spim

const (
	END_START SHORTS = 0x01 << 17 //+ Shortcut between event END and task START
)

const (
	END_STARTn = 17
)

const (
	TX STALLSTAT = 0x01 << 0 //+ Stall status for EasyDMA RAM reads
	RX STALLSTAT = 0x01 << 1 //+ Stall status for EasyDMA RAM writes
)

const (
	TXn = 0
	RXn = 1
)

const (
	EN       ENABLE = 0x0F << 0 //+ Enable or disable SPIM
	Disabled ENABLE = 0x00 << 0 //  Disable SPIM
	Enabled  ENABLE = 0x07 << 0 //  Enable SPIM
)

const (
	ENn = 0
)

const (
	FREQ FREQUENCY = 0xFFFFFFFF << 0 //+ SPI master data rate
	K125 FREQUENCY = 0x2000000 << 0  //  125 kbps
	K250 FREQUENCY = 0x4000000 << 0  //  250 kbps
	K500 FREQUENCY = 0x8000000 << 0  //  500 kbps
	M16  FREQUENCY = 0xA000000 << 0  //  16 Mbps
	M1   FREQUENCY = 0x10000000 << 0 //  1 Mbps
	M32  FREQUENCY = 0x14000000 << 0 //  32 Mbps
	M2   FREQUENCY = 0x20000000 << 0 //  2 Mbps
	M4   FREQUENCY = 0x40000000 << 0 //  4 Mbps
	M8   FREQUENCY = 0x80000000 << 0 //  8 Mbps
)

const (
	FREQn = 0
)

const (
	MAXCNT RXD_MAXCNT = 0xFFFF << 0 //+ Maximum number of bytes in receive buffer
)

const (
	MAXCNTn = 0
)

const (
	AMOUNT RXD_AMOUNT = 0xFFFF << 0 //+ Number of bytes transferred in the last transaction
)

const (
	AMOUNTn = 0
)

const (
	LIST      RXD_LIST = 0x03 << 0 //+ List type
	Disabled  RXD_LIST = 0x00 << 0 //  Disable EasyDMA list
	ArrayList RXD_LIST = 0x01 << 0 //  Use array list
)

const (
	LISTn = 0
)

const (
	MAXCNT TXD_MAXCNT = 0xFFFF << 0 //+ Maximum number of bytes in transmit buffer
)

const (
	MAXCNTn = 0
)

const (
	AMOUNT TXD_AMOUNT = 0xFFFF << 0 //+ Number of bytes transferred in the last transaction
)

const (
	AMOUNTn = 0
)

const (
	LIST      TXD_LIST = 0x03 << 0 //+ List type
	Disabled  TXD_LIST = 0x00 << 0 //  Disable EasyDMA list
	ArrayList TXD_LIST = 0x01 << 0 //  Use array list
)

const (
	LISTn = 0
)

const (
	ORDER CONFIG = 0x01 << 0 //+ Bit order
	CPHA  CONFIG = 0x01 << 1 //+ Serial clock (SCK) phase
	CPOL  CONFIG = 0x01 << 2 //+ Serial clock (SCK) polarity
)

const (
	ORDERn = 0
	CPHAn  = 1
	CPOLn  = 2
)

const (
	RXDELAY IFTIMING_RXDELAY = 0x07 << 0 //+ Sample delay for input serial data on MISO. The value specifies the number of 64 MHz clock cycles (15.625 ns) delay from the the sampling edge of SCK (leading edge for CONFIG.CPHA = 0, trailing edge for CONFIG.CPHA = 1) until the input serial data is sampled. As en example, if RXDELAY = 0 and CONFIG.CPHA = 0, the input serial data is sampled on the rising edge of SCK.
)

const (
	RXDELAYn = 0
)

const (
	CSNDUR IFTIMING_CSNDUR = 0xFF << 0 //+ Minimum duration between edge of CSN and edge of SCK and minimum duration CSN must stay high between transactions. The value is specified in number of 64 MHz clock cycles (15.625 ns).
)

const (
	CSNDURn = 0
)

const (
	CSNPOL CSNPOL = 0x01 << 0 //+ Polarity of CSN output
)

const (
	CSNPOLn = 0
)

const (
	PIN     PSELDCX = 0x1F << 0  //+ Pin number
	PORT    PSELDCX = 0x01 << 5  //+ Port number
	CONNECT PSELDCX = 0x01 << 31 //+ Connection
)

const (
	PINn     = 0
	PORTn    = 5
	CONNECTn = 31
)

const (
	DCXCNT DCXCNT = 0x0F << 0 //+ This register specifies the number of command bytes preceding the data bytes. The PSEL.DCX line will be low during transmission of command bytes and high during transmission of data bytes. Value 0xF indicates that all bytes are command bytes.
)

const (
	DCXCNTn = 0
)

const (
	ORC ORC = 0xFF << 0 //+ Byte transmitted after TXD.MAXCNT bytes have been transmitted in the case when RXD.MAXCNT is greater than TXD.MAXCNT.
)

const (
	ORCn = 0
)