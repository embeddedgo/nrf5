// Based on C code by Nordic Semiconductor ASA.
// See LICENSE-NORDIC for original C code license.

// Copyright 2020 Michal Derkacz.

#include "textflag.h"

#define SVC_BASE       0x60

#define ENABLE         (SVC_BASE + 0)
#define EVT_GET        (SVC_BASE + 1)
#define UUID_VS_ADD    (SVC_BASE + 2)
#define UUID_DECODE    (SVC_BASE + 3)
#define UUID_ENCODE    (SVC_BASE + 4)
#define VERSION_GET    (SVC_BASE + 5)
#define USER_MEM_REPLY (SVC_BASE + 6)
#define OPT_SET        (SVC_BASE + 7)
#define OPT_GET        (SVC_BASE + 8)
#define CFG_SET        (SVC_BASE + 9)
#define UUID_VS_REMOVE (SVC_BASE + 10)

TEXT ·enable(SB),NOSPLIT|NOFRAME,$0
	MOVW  app_ram_base+0(FP), R0
	SWI   $ENABLE
	MOVW  R0, ret+4(FP)
	RET

TEXT ·setCfg(SB),NOSPLIT|NOFRAME,$0
	MOVW  cfg_id+0(FP), R0
	MOVW  cfg+4(FP), R1
	MOVW  app_ram_base+8(FP), R2
	SWI   $CFG_SET
	MOVW  R0, ret+12(FP)
	RET

TEXT ·getEvt(SB),NOSPLIT|NOFRAME,$0
	MOVW  dest+0(FP), R0
	MOVW  len+4(FP), R1
	SWI   $EVT_GET
	MOVW  R0, ret+8(FP)
	RET

TEXT ·addUUIDVS(SB),NOSPLIT|NOFRAME,$0
	MOVW  vs_uuid+0(FP), R0
	MOVW  uuid_type+4(FP), R1
	SWI   $UUID_VS_ADD
	MOVW  R0, ret+8(FP)
	RET

TEXT ·removeUUIDVS(SB),NOSPLIT|NOFRAME,$0
	MOVW  uuid_type+0(FP), R0
	SWI   $UUID_VS_REMOVE
	MOVW  R0, ret+4(FP)
	RET

TEXT ·decodeUUID(SB),NOSPLIT|NOFRAME,$0
	MOVBU  uuid_le_len+0(FP), R0
	MOVW   uuid_le+4(FP), R1
	MOVW   uuid+8(FP), R2
	SWI    $UUID_DECODE
	MOVW   R0, ret+12(FP)
	RET

TEXT ·encodeUUID(SB),NOSPLIT|NOFRAME,$0
	MOVW  uuid+0(FP), R0
	MOVW  uuid_le_len+4(FP), R1
	MOVW  uuid_le+8(FP), R2
	SWI   $UUID_ENCODE
	MOVW  R0, ret+12(FP)
	RET

TEXT ·getVersion(SB),NOSPLIT|NOFRAME,$0
	MOVW  version+0(FP), R0
	SWI   $VERSION_GET
	MOVW  R0, ret+4(FP)
	RET

TEXT ·replyUserMem(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVW   block+4(FP), R1
	SWI    $USER_MEM_REPLY
	MOVW   R0, ret+8(FP)
	RET

TEXT ·setOpt(SB),NOSPLIT|NOFRAME,$0
	MOVW  opt_id+0(FP), R0
	MOVW  opt+4(FP), R1
	SWI   $OPT_SET
	MOVW  R0, ret+8(FP)
	RET

TEXT ·getOpt(SB),NOSPLIT|NOFRAME,$0
	MOVW  opt_id+0(FP), R0
	MOVW  opt+4(FP), R1
	SWI   $OPT_GET
	MOVW  R0, ret+8(FP)
	RET
