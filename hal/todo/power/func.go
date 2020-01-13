package power

import (
	"bits"
	"mmio"

	"nrf5/hal/te"
)

type Task byte

const (
	CONSTLAT Task = 30 // Enable constant latency mode.
	LOWPWR   Task = 31 // Enable low power mode (variable latency).
)

func (t Task) Task() *te.Task { return r().Regs.Task(int(t)) }

type Event byte

const (
	POFWARN    Event = 2 // Power failure warning.
	SLEEPENTER Event = 5 // CPU entered WFI/WFE sleep (nRF52).
	SLEEPEXIT  Event = 6 // CPU exited WFI/WFE sleep (nRF52).
)

func (e Event) Event() *te.Event { return r().Regs.Event(int(e)) }

// ResetReas is a bitfield that describes reset reason.
type ResetReas uint32

const (
	RESETPIN ResetReas = 1 << 0  // Reset from pin-reset.
	DOG      ResetReas = 1 << 1  // Reset from watchdog.
	SREQ     ResetReas = 1 << 2  // Reset from AIRCR.SYSRESETREQ.
	LOCKUP   ResetReas = 1 << 3  // Reset from CPU lock-up.
	OFF      ResetReas = 1 << 16 // Wake up from OFF mode by GPIO DETECT.
	LPCOMP   ResetReas = 1 << 17 // Wake up from OFF mode by LPCOMP ANADETECT.
	DIF      ResetReas = 1 << 18 // Wake up from OFF mode by debug interface.
	NFC      ResetReas = 1 << 19 // Wake up from OFF mode by NFC.
)

// LoadRESETREAS returns reset reason bits.
func LoadRESETREAS() ResetReas {
	return ResetReas(r().resetreas.Load())
}

// ClearRESETREAS clears reset reason bits specified by mask.
func ClearRESETREAS(mask ResetReas) {
	r().resetreas.Store(uint32(mask))
}

// RAMBlocks is a bitfield that describes RAM blocks.
type RAMBlocks byte

const (
	RAMBLOCK0 RAMBlocks = 1 << 0 // RAM block 0.
	RAMBLOCK1 RAMBlocks = 1 << 1 // RAM block 1.
	RAMBLOCK2 RAMBlocks = 1 << 2 // RAM block 2.
	RAMBLOCK3 RAMBlocks = 1 << 3 // RAM block 3.
)

// LoadRAMSTATUS returns bitfield that lists RAM blocks that are powered up.
func LoadRAMSTATUS() RAMBlocks {
	return RAMBlocks(r().ramstatus.Load())
}

// SetSYSTEMOFF sets system into OFF state.
func SetSYSTEMOFF() {
	r().systemoff.Store(1)
}

// POFCon is power failure comparator configuration.
type POFCon byte

const (
	POF       POFCon = 1 << 0  // Set if power failure comparoator is enabled.
	THRESHOLD POFCon = 15 << 1 // Power failure comparator threshold mask.

	// Power failure comparator thresholds.

	V2_1 POFCon = 0 << 1 // Threshold: 2.1 V (nrF51).
	V2_3 POFCon = 1 << 1 // Threshold: 2.3 V (nrF51).
	V2_5 POFCon = 2 << 1 // Threshold: 2.5 V (nrF51).
	V2_7 POFCon = 3 << 1 // Threshold: 2.5 V (nrF51).

	V17 POFCon = 4 << 1  // Threshold: 1.7 V (nRF52).
	V18 POFCon = 5 << 1  // Threshold: 1.8 V (nRF52).
	V19 POFCon = 6 << 1  // Threshold: 1.9 V (nRF52).
	V20 POFCon = 7 << 1  // Threshold: 2.0 V (nRF52).
	V21 POFCon = 8 << 1  // Threshold: 2.1 V (nRF52).
	V22 POFCon = 9 << 1  // Threshold: 2.2 V (nRF52).
	V23 POFCon = 10 << 1 // Threshold: 2.3 V (nRF52).
	V24 POFCon = 11 << 1 // Threshold: 2.4 V (nRF52).
	V25 POFCon = 12 << 1 // Threshold: 2.5 V (nRF52).
	V26 POFCon = 13 << 1 // Threshold: 2.6 V (nRF52).
	V27 POFCon = 14 << 1 // Threshold: 2.7 V (nRF52).
	V28 POFCon = 15 << 1 // Threshold: 2.8 V (nRF52).
)

// LoadPOFCON returns power failure comparator configuration.
func LoadPOFCON() POFCon {
	return POFCon(r().pofcon.Load())
}

// StorePOFCON sets power failure comparator configuration.
func StorePOFCON(pofcon POFCon) {
	r().pofcon.Store(uint32(pofcon))
}

// GPREGRET returns pointer to n-th general purpose retention register. nRF51
// supports one, nRF52 supports two. Only lowest 8 bits can be used.
func GPREGRET(n int) *mmio.U32 {
	return &r().gpregret[n]
}

// LoadRAMON returns configuration of four RAM blocks. On lists RAM blocks
// that are kept on in system ON mode, retain lists RAM blocks that should be
// retained when RAM block is off.
func LoadRAMON() (on, retain RAMBlocks) {
	a := r().ramon.Load()
	b := r().ramonb.Load()
	return RAMBlocks(a | b<<2), RAMBlocks(a>>16 | b>>14)
}

// StoreRAMON sets configuration of four RAM blocks. On lists RAM blocks that
// should be kept on in system ON mode, retain lists RAM blocks that should be
// retained in system off mode.
func StoreRAMON(on, retain RAMBlocks) {
	r().ramon.Store(uint32(on&3) | uint32(retain&3)<<16)
	r().ramonb.Store(uint32(on&12)>>2 | uint32(retain&12)<<14)
}

// LoadRESET reports wheter pin reset is enabled in debug mode (nRF51).
func LoadRESET() bool {
	return r().reset.Load()&1 != 0
}

// StoreRESET enables/disables pin reset in debug mode (nRF51).
func StoreRESET(pinreset bool) {
	r().reset.Store(uint32(bits.One(pinreset)))
}

// LoadDCDCEN reports wheter the DC/DC converter is enabled.
func LoadDCDCEN() bool {
	return r().dcdcen.Load()&1 != 0
}

// StoreDCDCEN enables/disables DC/DC converter.
func StoreDCDCEN(en bool) {
	r().dcdcen.Store(uint32(bits.One(en)))
}

// RAMPower describes power configuration for two sections of RAM block (nRF52).
type RAMPower uint32

const (
	S0POWER     RAMPower = 1 << 0  // Keep RAM section S0 on in system on mode.
	S1POWER     RAMPower = 1 << 1  // Keep RAM section S1 on in system on mode.
	S0RETENTION RAMPower = 1 << 16 // Keep retention of S0 when RAM is off.
	S1RETENTION RAMPower = 1 << 17 // Keep retention of S1 when RAM is off.
)

// LoadRAMPOWER returns power configuration of RAM block n (nRF52).
func LoadRAMPOWER(n int) RAMPower {
	return RAMPower(r().ram[n].power.Load())
}

// LoadRAMPOWER power configuration of RAM block n (nRF52).
func StoreRAMPOWER(n int, val RAMPower) {
	r().ram[n].power.Store(uint32(val))
}

// SetRAMPOWER sets on power configuration of RAM block n according to mask
// (nRF52).
func SetRAMPOWER(n int, mask RAMPower) {
	r().ram[n].powerset.Store(uint32(mask))
}

// ClearRAMPOWER sets off power configuration of RAM block n according to mask
// (nRF52).
func ClearRAMPOWER(n int, mask RAMPower) {
	r().ram[n].powerclr.Store(uint32(mask))
}
