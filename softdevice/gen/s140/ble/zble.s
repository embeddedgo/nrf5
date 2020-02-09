// Code generated from ../nordic/s140/headers/ble.h; DO NOT EDIT.

#include "go_asm.h"
#include "textflag.h"

TEXT ·Enable(SB),NOSPLIT|NOFRAME,$0
	MOVW app_ram_base+0(FP), R0
	SWI $const__ENABLE
	MOVW R0, ret+4(FP)
	RET

TEXT ·SetCfg(SB),NOSPLIT|NOFRAME,$0
	MOVW cfg_id+0(FP), R0
	MOVW cfg+4(FP), R1
	MOVW app_ram_base+8(FP), R2
	SWI $const__CFG_SET
	MOVW R0, ret+12(FP)
	RET

TEXT ·GetEvt(SB),NOSPLIT|NOFRAME,$0
	MOVW dest+0(FP), R0
	MOVW len+4(FP), R1
	SWI $const__EVT_GET
	MOVW R0, ret+8(FP)
	RET

TEXT ·AddUUIDVS(SB),NOSPLIT|NOFRAME,$0
	MOVW vs_uuid+0(FP), R0
	MOVW uuid_type+4(FP), R1
	SWI $const__UUID_VS_ADD
	MOVW R0, ret+8(FP)
	RET

TEXT ·RemoveUUIDVS(SB),NOSPLIT|NOFRAME,$0
	MOVW uuid_type+0(FP), R0
	SWI $const__UUID_VS_REMOVE
	MOVW R0, ret+4(FP)
	RET

TEXT ·DecodeUUID(SB),NOSPLIT|NOFRAME,$0
	MOVBU uuid_le_len+0(FP), R0
	MOVW uuid_le+4(FP), R1
	MOVW uuid+8(FP), R2
	SWI $const__UUID_DECODE
	MOVW R0, ret+12(FP)
	RET

TEXT ·EncodeUUID(SB),NOSPLIT|NOFRAME,$0
	MOVW uuid+0(FP), R0
	MOVW uuid_le_len+4(FP), R1
	MOVW uuid_le+8(FP), R2
	SWI $const__UUID_ENCODE
	MOVW R0, ret+12(FP)
	RET

TEXT ·GetVersion(SB),NOSPLIT|NOFRAME,$0
	MOVW version+0(FP), R0
	SWI $const__VERSION_GET
	MOVW R0, ret+4(FP)
	RET

TEXT ·ReplyUserMem(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVW block+4(FP), R1
	SWI $const__USER_MEM_REPLY
	MOVW R0, ret+8(FP)
	RET

TEXT ·SetOpt(SB),NOSPLIT|NOFRAME,$0
	MOVW opt_id+0(FP), R0
	MOVW opt+4(FP), R1
	SWI $const__OPT_SET
	MOVW R0, ret+8(FP)
	RET

TEXT ·GetOpt(SB),NOSPLIT|NOFRAME,$0
	MOVW opt_id+0(FP), R0
	MOVW opt+4(FP), R1
	SWI $const__OPT_GET
	MOVW R0, ret+8(FP)
	RET
