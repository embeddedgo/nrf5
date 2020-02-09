// Copyright 2020 Michal Derkacz. All rights reserved.

#include "textflag.h"

TEXT ·AppRAMBase(SB),NOSPLIT|NOFRAME,$0
	MOVW  $runtime·ramstart(SB), R0
	MOVW  R0, ret+0(FP)
	RET
