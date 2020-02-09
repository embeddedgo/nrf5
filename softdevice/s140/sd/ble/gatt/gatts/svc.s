// Based on C code by Nordic Semiconductor ASA.
// See LICENSE-NORDIC for original C code license.

// Copyright 2020 Michal Derkacz.

#include "textflag.h"

#define SVC_BASE 0xA8

#define SERVICE_ADD             (SVC_BASE + 0)
#define INCLUDE_ADD             (SVC_BASE + 1)
#define CHARACTERISTIC_ADD      (SVC_BASE + 2)
#define DESCRIPTOR_ADD          (SVC_BASE + 3)
#define VALUE_SET               (SVC_BASE + 4)
#define VALUE_GET               (SVC_BASE + 5)
#define HVX                     (SVC_BASE + 6)
#define SERVICE_CHANGED         (SVC_BASE + 7)
#define RW_AUTHORIZE_REPLY      (SVC_BASE + 8)
#define SYS_ATTR_SET            (SVC_BASE + 9)
#define SYS_ATTR_GET            (SVC_BASE + 10)
#define INITIAL_USER_HANDLE_GET (SVC_BASE + 11)
#define ATTR_GET                (SVC_BASE + 12)
#define EXCHANGE_MTU_REPLY      (SVC_BASE + 13)

TEXT ·addService(SB),NOSPLIT|NOFRAME,$0
	MOVBU  typ+0(FP), R0
	MOVW   uuid+4(FP), R1
	MOVW   handle+8(FP), R2
	SWI    $SERVICE_ADD
	MOVW   R0, ret+12(FP)
	RET

TEXT ·addInclude(SB),NOSPLIT|NOFRAME,$0
	MOVHU  service_handle+0(FP), R0
	MOVHU  inc_srvc_handle+2(FP), R1
	MOVW   include_handle+4(FP), R2
	SWI    $INCLUDE_ADD
	MOVW   R0, ret+8(FP)
	RET

TEXT ·addCharacteristic(SB),NOSPLIT|NOFRAME,$0
	MOVHU  service_handle+0(FP), R0
	MOVW   char_md+4(FP), R1
	MOVW   attr_char_value+8(FP), R2
	MOVW   handles+12(FP), R3
	SWI    $CHARACTERISTIC_ADD
	MOVW   R0, ret+16(FP)
	RET

TEXT ·addDescriptor(SB),NOSPLIT|NOFRAME,$0
	MOVHU  char_handle+0(FP), R0
	MOVW   attr+4(FP), R1
	MOVW   handle+8(FP), R2
	SWI    $DESCRIPTOR_ADD
	MOVW   R0, ret+12(FP)
	RET

TEXT ·setValue(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVHU  handle+2(FP), R1
	MOVW   value+4(FP), R2
	SWI    $VALUE_SET
	MOVW   R0, ret+8(FP)
	RET

TEXT ·getValue(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVHU  handle+2(FP), R1
	MOVW   value+4(FP), R2
	SWI    $VALUE_GET
	MOVW   R0, ret+8(FP)
	RET

TEXT ·hvx(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVW   hvx_params+4(FP), R1
	SWI    $HVX
	MOVW   R0, ret+8(FP)
	RET

TEXT ·changedService(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVHU  start_handle+2(FP), R1
	MOVHU  end_handle+4(FP), R2
	SWI    $SERVICE_CHANGED
	MOVW   R0, ret+8(FP)
	RET

TEXT ·replyRwAuthorize(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVW   rw_authorize_reply_params+4(FP), R1
	SWI    $RW_AUTHORIZE_REPLY
	MOVW   R0, ret+8(FP)
	RET

TEXT ·setSysAttr(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVW   sys_attr_data+4(FP), R1
	MOVHU  len+8(FP), R2
	MOVW   flags+12(FP), R3
	SWI    $SYS_ATTR_SET
	MOVW   R0, ret+16(FP)
	RET

TEXT ·getSysAttr(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVW   sys_attr_data+4(FP), R1
	MOVW   len+8(FP), R2
	MOVW   flags+12(FP), R3
	SWI    $SYS_ATTR_GET
	MOVW   R0, ret+16(FP)
	RET

TEXT ·getInitialUserHandle(SB),NOSPLIT|NOFRAME,$0
	MOVW  handle+0(FP), R0
	SWI   $INITIAL_USER_HANDLE_GET
	MOVW  R0, ret+4(FP)
	RET

TEXT ·getAttr(SB),NOSPLIT|NOFRAME,$0
	MOVHU  handle+0(FP), R0
	MOVW   uuid+4(FP), R1
	MOVW   md+8(FP), R2
	SWI    $ATTR_GET
	MOVW   R0, ret+12(FP)
	RET

TEXT ·replyExchangeMTU(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVHU  server_rx_mtu+2(FP), R1
	SWI    $EXCHANGE_MTU_REPLY
	MOVW   R0, ret+4(FP)
	RET
