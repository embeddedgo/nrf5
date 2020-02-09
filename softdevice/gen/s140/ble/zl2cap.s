// Code generated from ../nordic/s140/headers/ble_l2cap.h; DO NOT EDIT.

#include "go_asm.h"
#include "textflag.h"

TEXT ·SetupL2capCh(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVW local_cid+4(FP), R1
	MOVW params+8(FP), R2
	SWI $const__L2CAP_CH_SETUP
	MOVW R0, ret+12(FP)
	RET

TEXT ·ReleaseL2capCh(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVHU local_cid+2(FP), R1
	SWI $const__L2CAP_CH_RELEASE
	MOVW R0, ret+4(FP)
	RET

TEXT ·RxL2capCh(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVHU local_cid+2(FP), R1
	MOVW sdu_buf+4(FP), R2
	SWI $const__L2CAP_CH_RX
	MOVW R0, ret+8(FP)
	RET

TEXT ·TxL2capCh(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVHU local_cid+2(FP), R1
	MOVW sdu_buf+4(FP), R2
	SWI $const__L2CAP_CH_TX
	MOVW R0, ret+8(FP)
	RET

TEXT ·ControlL2capChFlow(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVHU local_cid+2(FP), R1
	MOVHU credits+4(FP), R2
	MOVW credits+8(FP), R3
	SWI $const__L2CAP_CH_FLOW_CONTROL
	MOVW R0, ret+12(FP)
	RET
