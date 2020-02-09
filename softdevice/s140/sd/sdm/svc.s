// Based on C code by Nordic Semiconductor ASA.
// See LICENSE-NORDIC for original C code license.

// Copyright 2020 Michal Derkacz.

#include "textflag.h"

#define SVC_BASE 0x10

#define SOFTDEVICE_ENABLE                (SVC_BASE + 0)
#define SOFTDEVICE_DISABLE               (SVC_BASE + 1)
#define SOFTDEVICE_IS_ENABLED            (SVC_BASE + 2)
#define SOFTDEVICE_VECTOR_TABLE_BASE_SET (SVC_BASE + 3)

TEXT nrfFaultHandler(SB),NOSPLIT|NOFRAME,$0
	MOVM.DB.W  [R4-R11,R14], (R13)
	MOVW       $0, R3
	MOVM.DB.W  [R0-R3], (R13)  // id, pc, info, 0
	BL         runtime·identcurcpu(SB)
	MOVW       R0, g
	MOVW       ·sdfh(SB), R11
	MOVW       (R11), R0
	BL         (R0)
	ADD        $4, R13
	MOVM.IA.W  (R13), [R4-R11,R15]


TEXT ·enableSoftdevice(SB),NOSPLIT|NOFRAME,$0
	MOVW  lfclkc+0(FP), R0
	MOVW  $nrfFaultHandler(SB), R1
	SWI   $SOFTDEVICE_ENABLE
	MOVW  R0, ret+4(FP)
	RET

TEXT ·disableSoftdevice(SB),NOSPLIT|NOFRAME,$0
	SWI   $SOFTDEVICE_DISABLE
	MOVW  R0, ret+0(FP)
	RET

TEXT ·softdeviceEnabled(SB),NOSPLIT|NOFRAME,$0
	MOVW  softdevice_enabled+0(FP), R0
	SWI   $SOFTDEVICE_IS_ENABLED
	MOVW  R0, ret+4(FP)
	RET

TEXT ·setSoftdeviceVectorTableBase(SB),NOSPLIT|NOFRAME,$0
	MOVW  address+0(FP), R0
	SWI   $SOFTDEVICE_VECTOR_TABLE_BASE_SET
	MOVW  R0, ret+4(FP)
	RET
