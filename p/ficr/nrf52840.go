// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build nrf52840

// Package ficr provides access to the registers of the FICR peripheral.
//
// Instances:
//
//	FICR  FICR_BASE  -  -  Factory information configuration registers
//
// Registers:
//
//	0x010 32  CODEPAGESIZE      Code memory page size
//	0x014 32  CODESIZE          Code memory size
//	0x060 32  DEVICEID[2]       Device identifier
//	0x080 32  ER[4]             Encryption root, word n
//	0x090 32  IR[4]             Identity Root, word n
//	0x0A0 32  DEVICEADDRTYPE    Device address type
//	0x0A4 32  DEVICEADDR[2]     Device address n
//	0x100 32  INFO_PART         Part code
//	0x104 32  INFO_VARIANT      Build code (hardware version and production configuration)
//	0x108 32  INFO_PACKAGE      Package option
//	0x10C 32  INFO_RAM          RAM variant
//	0x110 32  INFO_FLASH        Flash variant
//	0x350 32  PRODTEST[3]       Production test signature n
//	0x404 32  TEMP_A[6]         Slope definition
//	0x41C 32  TEMP_B[6]         Y-intercept
//	0x434 32  TEMP_T[5]         Segment end
//	0x450 32  NFC_TAGHEADER0    Default header for NFC tag. Software can read these values to populate NFCID1_3RD_LAST, NFCID1_2ND_LAST and NFCID1_LAST.
//	0x454 32  NFC_TAGHEADER1    Default header for NFC tag. Software can read these values to populate NFCID1_3RD_LAST, NFCID1_2ND_LAST and NFCID1_LAST.
//	0x458 32  NFC_TAGHEADER2    Default header for NFC tag. Software can read these values to populate NFCID1_3RD_LAST, NFCID1_2ND_LAST and NFCID1_LAST.
//	0x45C 32  NFC_TAGHEADER3    Default header for NFC tag. Software can read these values to populate NFCID1_3RD_LAST, NFCID1_2ND_LAST and NFCID1_LAST.
//	0xC00 32  TRNG90B_BYTES     Amount of bytes for the required entropy bits
//	0xC04 32  TRNG90B_RCCUTOFF  Repetition counter cutoff
//	0xC08 32  TRNG90B_APCUTOFF  Adaptive proportion cutoff
//	0xC0C 32  TRNG90B_STARTUP   Amount of bytes for the startup tests
//	0xC10 32  TRNG90B_ROSC1     Sample count for ring oscillator 1
//	0xC14 32  TRNG90B_ROSC2     Sample count for ring oscillator 2
//	0xC18 32  TRNG90B_ROSC3     Sample count for ring oscillator 3
//	0xC1C 32  TRNG90B_ROSC4     Sample count for ring oscillator 4
//
// Import:
//
//	github.com/embeddedgo/nrf5/p/mmap
package ficr

const (
	PUBLIC DEVICEADDRTYPE = 0x01 << 0 //+ Device address type
)

const (
	PUBLICn = 0
)

const (
	PART        INFO_PART = 0xFFFFFFFF << 0 //+ Part code
	N52840      INFO_PART = 0x52840 << 0    //  nRF52840
	Unspecified INFO_PART = 0xFFFFFFFF << 0 //  Unspecified
)

const (
	PARTn = 0
)

const (
	VARIANT INFO_VARIANT = 0xFFFFFFFF << 0 //+ Build code (hardware version and production configuration). Encoded as ASCII.
	AAAA    INFO_VARIANT = 0x41414141 << 0 //  AAAA
	AAAB    INFO_VARIANT = 0x41414142 << 0 //  AAAB
	AABA    INFO_VARIANT = 0x41414241 << 0 //  AABA
	AABB    INFO_VARIANT = 0x41414242 << 0 //  AABB
	AACA    INFO_VARIANT = 0x41414341 << 0 //  AACA
	BAAA    INFO_VARIANT = 0x42414141 << 0 //  BAAA
	CAAA    INFO_VARIANT = 0x43414141 << 0 //  CAAA
	Vunspec INFO_VARIANT = 0xFFFFFFFF << 0 //  Unspecified
)

const (
	VARIANTn = 0
)

const (
	PACKAGE INFO_PACKAGE = 0xFFFFFFFF << 0 //+ Package option
	QI      INFO_PACKAGE = 0x2004 << 0     //  QIxx - 73-pin aQFN
	Punspec INFO_PACKAGE = 0xFFFFFFFF << 0 //  Unspecified
)

const (
	PACKAGEn = 0
)

const (
	RAM     INFO_RAM = 0xFFFFFFFF << 0 //+ RAM variant
	R16K    INFO_RAM = 0x10 << 0       //  16 kByte RAM
	R32K    INFO_RAM = 0x20 << 0       //  32 kByte RAM
	R64K    INFO_RAM = 0x40 << 0       //  64 kByte RAM
	R128K   INFO_RAM = 0x80 << 0       //  128 kByte RAM
	R256K   INFO_RAM = 0x100 << 0      //  256 kByte RAM
	Runspec INFO_RAM = 0xFFFFFFFF << 0 //  Unspecified
)

const (
	RAMn = 0
)

const (
	FLASH   INFO_FLASH = 0xFFFFFFFF << 0 //+ Flash variant
	F128K   INFO_FLASH = 0x80 << 0       //  128 kByte FLASH
	F256K   INFO_FLASH = 0x100 << 0      //  256 kByte FLASH
	F512K   INFO_FLASH = 0x200 << 0      //  512 kByte FLASH
	F1024K  INFO_FLASH = 0x400 << 0      //  1 MByte FLASH
	F2048K  INFO_FLASH = 0x800 << 0      //  2 MByte FLASH
	Funspec INFO_FLASH = 0xFFFFFFFF << 0 //  Unspecified
)

const (
	FLASHn = 0
)

const (
	PTEST   PRODTEST = 0xFFFFFFFF << 0 //+ Production test signature n
	Done    PRODTEST = 0xBB42319F << 0 //  Production tests done
	NotDone PRODTEST = 0xFFFFFFFF << 0 //  Production tests not done
)

const (
	PTESTn = 0
)

const (
	A TEMP_A = 0xFFF << 0 //+ A (slope definition) register.
)

const (
	An = 0
)

const (
	B TEMP_B = 0x3FFF << 0 //+ B (y-intercept)
)

const (
	Bn = 0
)

const (
	T TEMP_T = 0xFF << 0 //+ T (segment end) register
)

const (
	Tn = 0
)

const (
	MFGID NFC_TAGHEADER0 = 0xFF << 0  //+ Default Manufacturer ID: Nordic Semiconductor ASA has ICM 0x5F
	UD1   NFC_TAGHEADER0 = 0xFF << 8  //+ Unique identifier byte 1
	UD2   NFC_TAGHEADER0 = 0xFF << 16 //+ Unique identifier byte 2
	UD3   NFC_TAGHEADER0 = 0xFF << 24 //+ Unique identifier byte 3
)

const (
	MFGIDn = 0
	UD1n   = 8
	UD2n   = 16
	UD3n   = 24
)

const (
	UD4 NFC_TAGHEADER1 = 0xFF << 0  //+ Unique identifier byte 4
	UD5 NFC_TAGHEADER1 = 0xFF << 8  //+ Unique identifier byte 5
	UD6 NFC_TAGHEADER1 = 0xFF << 16 //+ Unique identifier byte 6
	UD7 NFC_TAGHEADER1 = 0xFF << 24 //+ Unique identifier byte 7
)

const (
	UD4n = 0
	UD5n = 8
	UD6n = 16
	UD7n = 24
)

const (
	UD8  NFC_TAGHEADER2 = 0xFF << 0  //+ Unique identifier byte 8
	UD9  NFC_TAGHEADER2 = 0xFF << 8  //+ Unique identifier byte 9
	UD10 NFC_TAGHEADER2 = 0xFF << 16 //+ Unique identifier byte 10
	UD11 NFC_TAGHEADER2 = 0xFF << 24 //+ Unique identifier byte 11
)

const (
	UD8n  = 0
	UD9n  = 8
	UD10n = 16
	UD11n = 24
)

const (
	UD12 NFC_TAGHEADER3 = 0xFF << 0  //+ Unique identifier byte 12
	UD13 NFC_TAGHEADER3 = 0xFF << 8  //+ Unique identifier byte 13
	UD14 NFC_TAGHEADER3 = 0xFF << 16 //+ Unique identifier byte 14
	UD15 NFC_TAGHEADER3 = 0xFF << 24 //+ Unique identifier byte 15
)

const (
	UD12n = 0
	UD13n = 8
	UD14n = 16
	UD15n = 24
)
