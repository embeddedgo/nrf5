// Code generated from ../nordic/s140/headers/nrf_sdm.h; DO NOT EDIT.

#include "go_asm.h"
#include "textflag.h"

TEXT ·EnableSoftdevice(SB),NOSPLIT|NOFRAME,$0
	MOVW clock_lf_cfg+0(FP), R0
	MOVW fault_handler+4(FP), R1
	SWI $const__SOFTDEVICE_ENABLE
	MOVW R0, ret+8(FP)
	RET

TEXT ·DisableSoftdevice(SB),NOSPLIT|NOFRAME,$0
	SWI $const__SOFTDEVICE_DISABLE
	MOVW R0, ret+0(FP)
	RET

TEXT ·IsEnabledSoftdevice(SB),NOSPLIT|NOFRAME,$0
	MOVW softdevice_enabled+0(FP), R0
	SWI $const__SOFTDEVICE_IS_ENABLED
	MOVW R0, ret+4(FP)
	RET

TEXT ·SetSoftdeviceVectorTableBase(SB),NOSPLIT|NOFRAME,$0
	MOVW address+0(FP), R0
	SWI $const__SOFTDEVICE_VECTOR_TABLE_BASE_SET
	MOVW R0, ret+4(FP)
	RET
