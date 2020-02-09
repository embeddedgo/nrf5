// Code generated from ../nordic/s140/headers/ble_gattc.h; DO NOT EDIT.

#include "go_asm.h"
#include "textflag.h"

TEXT ·DiscoverGATTCPrimaryServices(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVHU start_handle+2(FP), R1
	MOVW srvc_uuid+4(FP), R2
	SWI $const__GATTC_PRIMARY_SERVICES_DISCOVER
	MOVW R0, ret+8(FP)
	RET

TEXT ·DiscoverGATTCRelationships(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVW handle_range+4(FP), R1
	SWI $const__GATTC_RELATIONSHIPS_DISCOVER
	MOVW R0, ret+8(FP)
	RET

TEXT ·DiscoverGATTCCharacteristics(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVW handle_range+4(FP), R1
	SWI $const__GATTC_CHARACTERISTICS_DISCOVER
	MOVW R0, ret+8(FP)
	RET

TEXT ·DiscoverGATTCDescriptors(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVW handle_range+4(FP), R1
	SWI $const__GATTC_DESCRIPTORS_DISCOVER
	MOVW R0, ret+8(FP)
	RET

TEXT ·ReadGATTCCharValueByUUID(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVW uuid+4(FP), R1
	MOVW handle_range+8(FP), R2
	SWI $const__GATTC_CHAR_VALUE_BY_UUID_READ
	MOVW R0, ret+12(FP)
	RET

TEXT ·ReadGATTC(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVHU handle+2(FP), R1
	MOVHU offset+4(FP), R2
	SWI $const__GATTC_READ
	MOVW R0, ret+8(FP)
	RET

TEXT ·ReadGATTCCharValues(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVW handles+4(FP), R1
	MOVHU handle_count+8(FP), R2
	SWI $const__GATTC_CHAR_VALUES_READ
	MOVW R0, ret+12(FP)
	RET

TEXT ·WriteGATTC(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVW write_params+4(FP), R1
	SWI $const__GATTC_WRITE
	MOVW R0, ret+8(FP)
	RET

TEXT ·ConfirmGATTCHv(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVHU handle+2(FP), R1
	SWI $const__GATTC_HV_CONFIRM
	MOVW R0, ret+4(FP)
	RET

TEXT ·DiscoverGATTCAttrInfo(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVW handle_range+4(FP), R1
	SWI $const__GATTC_ATTR_INFO_DISCOVER
	MOVW R0, ret+8(FP)
	RET

TEXT ·RequestGATTCExchangeMtu(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVHU client_rx_mtu+2(FP), R1
	SWI $const__GATTC_EXCHANGE_MTU_REQUEST
	MOVW R0, ret+4(FP)
	RET
