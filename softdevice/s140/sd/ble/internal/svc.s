// Copyright 2020 Michal Derkacz.

#include "textflag.h"

#define SVC_BASE 0x60

#define CFG_SET (SVC_BASE + 9)

TEXT ·SetCfg(SB),NOSPLIT|NOFRAME,$0
	MOVW  cfg_id+0(FP), R0
	MOVW  cfg+4(FP), R1
	MOVW  app_ram_base+8(FP), R2
	SWI   $CFG_SET
	MOVW  R0, ret+12(FP)
	RET
