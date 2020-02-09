// Code generated from ../nordic/s140/headers/ble_gatts.h; DO NOT EDIT.

#include "go_asm.h"
#include "textflag.h"

TEXT ·AddGATTSService(SB),NOSPLIT|NOFRAME,$0
	MOVBU typ+0(FP), R0
	MOVW uuid+4(FP), R1
	MOVW handle+8(FP), R2
	SWI $const__GATTS_SERVICE_ADD
	MOVW R0, ret+12(FP)
	RET

TEXT ·AddGATTSInclude(SB),NOSPLIT|NOFRAME,$0
	MOVHU service_handle+0(FP), R0
	MOVHU inc_srvc_handle+2(FP), R1
	MOVW include_handle+4(FP), R2
	SWI $const__GATTS_INCLUDE_ADD
	MOVW R0, ret+8(FP)
	RET

TEXT ·AddGATTSCharacteristic(SB),NOSPLIT|NOFRAME,$0
	MOVHU service_handle+0(FP), R0
	MOVW char_md+4(FP), R1
	MOVW attr_char_value+8(FP), R2
	MOVW handles+12(FP), R3
	SWI $const__GATTS_CHARACTERISTIC_ADD
	MOVW R0, ret+16(FP)
	RET

TEXT ·AddGATTSDescriptor(SB),NOSPLIT|NOFRAME,$0
	MOVHU char_handle+0(FP), R0
	MOVW attr+4(FP), R1
	MOVW handle+8(FP), R2
	SWI $const__GATTS_DESCRIPTOR_ADD
	MOVW R0, ret+12(FP)
	RET

TEXT ·SetGATTSValue(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVHU handle+2(FP), R1
	MOVW value+4(FP), R2
	SWI $const__GATTS_VALUE_SET
	MOVW R0, ret+8(FP)
	RET

TEXT ·GetGATTSValue(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVHU handle+2(FP), R1
	MOVW value+4(FP), R2
	SWI $const__GATTS_VALUE_GET
	MOVW R0, ret+8(FP)
	RET

TEXT ·HvxGATTS(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVW hvx_params+4(FP), R1
	SWI $const__GATTS_HVX
	MOVW R0, ret+8(FP)
	RET

TEXT ·ChangedGATTSService(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVHU start_handle+2(FP), R1
	MOVHU end_handle+4(FP), R2
	SWI $const__GATTS_SERVICE_CHANGED
	MOVW R0, ret+8(FP)
	RET

TEXT ·ReplyGATTSRwAuthorize(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVW rw_authorize_reply_params+4(FP), R1
	SWI $const__GATTS_RW_AUTHORIZE_REPLY
	MOVW R0, ret+8(FP)
	RET

TEXT ·SetGATTSSysAttr(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVW sys_attr_data+4(FP), R1
	MOVHU len+8(FP), R2
	MOVW flags+12(FP), R3
	SWI $const__GATTS_SYS_ATTR_SET
	MOVW R0, ret+16(FP)
	RET

TEXT ·GetGATTSSysAttr(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVW sys_attr_data+4(FP), R1
	MOVW len+8(FP), R2
	MOVW flags+12(FP), R3
	SWI $const__GATTS_SYS_ATTR_GET
	MOVW R0, ret+16(FP)
	RET

TEXT ·GetGATTSInitialUserHandle(SB),NOSPLIT|NOFRAME,$0
	MOVW handle+0(FP), R0
	SWI $const__GATTS_INITIAL_USER_HANDLE_GET
	MOVW R0, ret+4(FP)
	RET

TEXT ·GetGATTSAttr(SB),NOSPLIT|NOFRAME,$0
	MOVHU handle+0(FP), R0
	MOVW uuid+4(FP), R1
	MOVW md+8(FP), R2
	SWI $const__GATTS_ATTR_GET
	MOVW R0, ret+12(FP)
	RET

TEXT ·ReplyGATTSExchangeMtu(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVHU server_rx_mtu+2(FP), R1
	SWI $const__GATTS_EXCHANGE_MTU_REPLY
	MOVW R0, ret+4(FP)
	RET
