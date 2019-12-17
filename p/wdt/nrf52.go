// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build nrf52

// Package wdt provides access to the registers of the WDT peripheral.
//
// Instances:
//  WDT  WDT_BASE  -  WDT  Watchdog Timer
// Registers:
//  0x000 32  TASKS_START     Start the watchdog
//  0x100 32  EVENTS_TIMEOUT  Watchdog timeout
//  0x304 32  INTENSET        Enable interrupt
//  0x308 32  INTENCLR        Disable interrupt
//  0x400 32  RUNSTATUS       Run status
//  0x404 32  REQSTATUS       Request status
//  0x504 32  CRV             Counter reload value
//  0x508 32  RREN            Enable register for reload request registers
//  0x50C 32  CONFIG          Configuration register
//  0x600 32  RR[8]           Description collection: Reload request n
// Import:
//  github.com/embeddedgo/nrf5/p/mmap
package wdt

const (
	TIMEOUT  INTENSET = 0x01 << 0 //+ Write '1' to enable interrupt for event TIMEOUT
	Disabled INTENSET = 0x00 << 0 //  Read: Disabled
	Enabled  INTENSET = 0x01 << 0 //  Read: Enabled
	Set      INTENSET = 0x01 << 0 //  Enable
)

const (
	TIMEOUTn = 0
)

const (
	TIMEOUT  INTENCLR = 0x01 << 0 //+ Write '1' to disable interrupt for event TIMEOUT
	Disabled INTENCLR = 0x00 << 0 //  Read: Disabled
	Enabled  INTENCLR = 0x01 << 0 //  Read: Enabled
	Clear    INTENCLR = 0x01 << 0 //  Disable
)

const (
	TIMEOUTn = 0
)

const (
	RUNSTATUS  RUNSTATUS = 0x01 << 0 //+ Indicates whether or not the watchdog is running
	NotRunning RUNSTATUS = 0x00 << 0 //  Watchdog not running
	Running    RUNSTATUS = 0x01 << 0 //  Watchdog is running
)

const (
	RUNSTATUSn = 0
)

const (
	RR0                   REQSTATUS = 0x01 << 0 //+ Request status for RR[0] register
	DisabledOrRequested   REQSTATUS = 0x00 << 0 //  RR[0] register is not enabled, or are already requesting reload
	EnabledAndUnrequested REQSTATUS = 0x01 << 0 //  RR[0] register is enabled, and are not yet requesting reload
	RR1                   REQSTATUS = 0x01 << 1 //+ Request status for RR[1] register
	DisabledOrRequested   REQSTATUS = 0x00 << 1 //  RR[1] register is not enabled, or are already requesting reload
	EnabledAndUnrequested REQSTATUS = 0x01 << 1 //  RR[1] register is enabled, and are not yet requesting reload
	RR2                   REQSTATUS = 0x01 << 2 //+ Request status for RR[2] register
	DisabledOrRequested   REQSTATUS = 0x00 << 2 //  RR[2] register is not enabled, or are already requesting reload
	EnabledAndUnrequested REQSTATUS = 0x01 << 2 //  RR[2] register is enabled, and are not yet requesting reload
	RR3                   REQSTATUS = 0x01 << 3 //+ Request status for RR[3] register
	DisabledOrRequested   REQSTATUS = 0x00 << 3 //  RR[3] register is not enabled, or are already requesting reload
	EnabledAndUnrequested REQSTATUS = 0x01 << 3 //  RR[3] register is enabled, and are not yet requesting reload
	RR4                   REQSTATUS = 0x01 << 4 //+ Request status for RR[4] register
	DisabledOrRequested   REQSTATUS = 0x00 << 4 //  RR[4] register is not enabled, or are already requesting reload
	EnabledAndUnrequested REQSTATUS = 0x01 << 4 //  RR[4] register is enabled, and are not yet requesting reload
	RR5                   REQSTATUS = 0x01 << 5 //+ Request status for RR[5] register
	DisabledOrRequested   REQSTATUS = 0x00 << 5 //  RR[5] register is not enabled, or are already requesting reload
	EnabledAndUnrequested REQSTATUS = 0x01 << 5 //  RR[5] register is enabled, and are not yet requesting reload
	RR6                   REQSTATUS = 0x01 << 6 //+ Request status for RR[6] register
	DisabledOrRequested   REQSTATUS = 0x00 << 6 //  RR[6] register is not enabled, or are already requesting reload
	EnabledAndUnrequested REQSTATUS = 0x01 << 6 //  RR[6] register is enabled, and are not yet requesting reload
	RR7                   REQSTATUS = 0x01 << 7 //+ Request status for RR[7] register
	DisabledOrRequested   REQSTATUS = 0x00 << 7 //  RR[7] register is not enabled, or are already requesting reload
	EnabledAndUnrequested REQSTATUS = 0x01 << 7 //  RR[7] register is enabled, and are not yet requesting reload
)

const (
	RR0n = 0
	RR1n = 1
	RR2n = 2
	RR3n = 3
	RR4n = 4
	RR5n = 5
	RR6n = 6
	RR7n = 7
)

const (
	CRV CRV = 0xFFFFFFFF << 0 //+ Counter reload value in number of cycles of the 32.768 kHz clock
)

const (
	CRVn = 0
)

const (
	RR0      RREN = 0x01 << 0 //+ Enable or disable RR[0] register
	Disabled RREN = 0x00 << 0 //  Disable RR[0] register
	Enabled  RREN = 0x01 << 0 //  Enable RR[0] register
	RR1      RREN = 0x01 << 1 //+ Enable or disable RR[1] register
	Disabled RREN = 0x00 << 1 //  Disable RR[1] register
	Enabled  RREN = 0x01 << 1 //  Enable RR[1] register
	RR2      RREN = 0x01 << 2 //+ Enable or disable RR[2] register
	Disabled RREN = 0x00 << 2 //  Disable RR[2] register
	Enabled  RREN = 0x01 << 2 //  Enable RR[2] register
	RR3      RREN = 0x01 << 3 //+ Enable or disable RR[3] register
	Disabled RREN = 0x00 << 3 //  Disable RR[3] register
	Enabled  RREN = 0x01 << 3 //  Enable RR[3] register
	RR4      RREN = 0x01 << 4 //+ Enable or disable RR[4] register
	Disabled RREN = 0x00 << 4 //  Disable RR[4] register
	Enabled  RREN = 0x01 << 4 //  Enable RR[4] register
	RR5      RREN = 0x01 << 5 //+ Enable or disable RR[5] register
	Disabled RREN = 0x00 << 5 //  Disable RR[5] register
	Enabled  RREN = 0x01 << 5 //  Enable RR[5] register
	RR6      RREN = 0x01 << 6 //+ Enable or disable RR[6] register
	Disabled RREN = 0x00 << 6 //  Disable RR[6] register
	Enabled  RREN = 0x01 << 6 //  Enable RR[6] register
	RR7      RREN = 0x01 << 7 //+ Enable or disable RR[7] register
	Disabled RREN = 0x00 << 7 //  Disable RR[7] register
	Enabled  RREN = 0x01 << 7 //  Enable RR[7] register
)

const (
	RR0n = 0
	RR1n = 1
	RR2n = 2
	RR3n = 3
	RR4n = 4
	RR5n = 5
	RR6n = 6
	RR7n = 7
)

const (
	SLEEP CONFIG = 0x01 << 0 //+ Configure the watchdog to either be paused, or kept running, while the CPU is sleeping
	Pause CONFIG = 0x00 << 0 //  Pause watchdog while the CPU is sleeping
	Run   CONFIG = 0x01 << 0 //  Keep the watchdog running while the CPU is sleeping
	HALT  CONFIG = 0x01 << 3 //+ Configure the watchdog to either be paused, or kept running, while the CPU is halted by the debugger
	Pause CONFIG = 0x00 << 3 //  Pause watchdog while the CPU is halted by the debugger
	Run   CONFIG = 0x01 << 3 //  Keep the watchdog running while the CPU is halted by the debugger
)

const (
	SLEEPn = 0
	HALTn  = 3
)

const (
	RR     RR = 0xFFFFFFFF << 0 //+ Reload request register
	Reload RR = 0x6E524635 << 0 //  Value to request a reload of the watchdog timer
)

const (
	RRn = 0
)
