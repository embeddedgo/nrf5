// Based on C code by Nordic Semiconductor ASA.
// See LICENSE-NORDIC for original C code license.

// Copyright 2020 Michal Derkacz.

#include "textflag.h"

#define SVC_BASE 0x6C

#define ADDR_SET                  (SVC_BASE + 0)
#define ADDR_GET                  (SVC_BASE + 1)
#define WHITELIST_SET             (SVC_BASE + 2)
#define DEVICE_IDENTITIES_SET     (SVC_BASE + 3)
#define PRIVACY_SET               (SVC_BASE + 4)
#define PRIVACY_GET               (SVC_BASE + 5)
#define ADV_SET_CONFIGURE         (SVC_BASE + 6)
#define ADV_START                 (SVC_BASE + 7)
#define ADV_STOP                  (SVC_BASE + 8)
#define CONN_PARAM_UPDATE         (SVC_BASE + 9)
#define DISCONNECT                (SVC_BASE + 10)
#define TX_POWER_SET              (SVC_BASE + 11)
#define APPEARANCE_SET            (SVC_BASE + 12)
#define APPEARANCE_GET            (SVC_BASE + 13)
#define PPCP_SET                  (SVC_BASE + 14)
#define PPCP_GET                  (SVC_BASE + 15)
#define DEVICE_NAME_SET           (SVC_BASE + 16)
#define DEVICE_NAME_GET           (SVC_BASE + 17)
#define AUTHENTICATE              (SVC_BASE + 18)
#define SEC_PARAMS_REPLY          (SVC_BASE + 19)
#define AUTH_KEY_REPLY            (SVC_BASE + 20)
#define LESC_DHKEY_REPLY          (SVC_BASE + 21)
#define KEYPRESS_NOTIFY           (SVC_BASE + 22)
#define LESC_OOB_DATA_GET         (SVC_BASE + 23)
#define LESC_OOB_DATA_SET         (SVC_BASE + 24)
#define ENCRYPT                   (SVC_BASE + 25)
#define SEC_INFO_REPLY            (SVC_BASE + 26)
#define CONN_SEC_GET              (SVC_BASE + 27)
#define RSSI_START                (SVC_BASE + 28)
#define RSSI_STOP                 (SVC_BASE + 29)
#define SCAN_START                (SVC_BASE + 30)
#define SCAN_STOP                 (SVC_BASE + 31)
#define CONNECT                   (SVC_BASE + 32)
#define CONNECT_CANCEL            (SVC_BASE + 33)
#define RSSI_GET                  (SVC_BASE + 34)
#define PHY_UPDATE                (SVC_BASE + 35)
#define DATA_LENGTH_UPDATE        (SVC_BASE + 36)
#define QOS_CHANNEL_SURVEY_START  (SVC_BASE + 37)
#define QOS_CHANNEL_SURVEY_STOP   (SVC_BASE + 38)
#define ADV_ADDR_GET              (SVC_BASE + 39)
#define NEXT_CONN_EVT_COUNTER_GET (SVC_BASE + 40)
#define CONN_EVT_TRIGGER_START    (SVC_BASE + 41)
#define CONN_EVT_TRIGGER_STOP     (SVC_BASE + 42)

TEXT ·setAddr(SB),NOSPLIT|NOFRAME,$0
	MOVW  addr+0(FP), R0
	SWI   $ADDR_SET
	MOVW  R0, ret+4(FP)
	RET

TEXT ·getAddr(SB),NOSPLIT|NOFRAME,$0
	MOVW  addr+0(FP), R0
	SWI   $ADDR_GET
	MOVW  R0, ret+4(FP)
	RET

TEXT ·getAdvAddr(SB),NOSPLIT|NOFRAME,$0
	MOVBU  adv_handle+0(FP), R0
	MOVW   addr+4(FP), R1
	SWI    $ADV_ADDR_GET
	MOVW   R0, ret+8(FP)
	RET

TEXT ·setWhitelist(SB),NOSPLIT|NOFRAME,$0
	MOVW   pp_wl_addrs+0(FP), R0
	MOVBU  len+4(FP), R1
	SWI    $WHITELIST_SET
	MOVW   R0, ret+8(FP)
	RET

TEXT ·setDeviceIdentities(SB),NOSPLIT|NOFRAME,$0
	MOVW   pp_id_keys+0(FP), R0
	MOVW   pp_local_irks+4(FP), R1
	MOVBU  len+8(FP), R2
	SWI    $DEVICE_IDENTITIES_SET
	MOVW   R0, ret+12(FP)
	RET

TEXT ·setPrivacy(SB),NOSPLIT|NOFRAME,$0
	MOVW  privacy_params+0(FP), R0
	SWI   $PRIVACY_SET
	MOVW  R0, ret+4(FP)
	RET

TEXT ·getPrivacy(SB),NOSPLIT|NOFRAME,$0
	MOVW  privacy_params+0(FP), R0
	SWI   $PRIVACY_GET
	MOVW  R0, ret+4(FP)
	RET

TEXT ·configureAdvSet(SB),NOSPLIT|NOFRAME,$0
	MOVW  adv_handle+0(FP), R0
	MOVW  adv_data+4(FP), R1
	MOVW  adv_params+8(FP), R2
	SWI   $ADV_SET_CONFIGURE
	MOVW  R0, ret+12(FP)
	RET

TEXT ·startAdv(SB),NOSPLIT|NOFRAME,$0
	MOVBU  adv_handle+0(FP), R0
	MOVBU  conn_cfg_tag+1(FP), R1
	SWI    $ADV_START
	MOVW   R0, ret+4(FP)
	RET

TEXT ·stopAdv(SB),NOSPLIT|NOFRAME,$0
	MOVBU  adv_handle+0(FP), R0
	SWI    $ADV_STOP
	MOVW   R0, ret+4(FP)
	RET

TEXT ·updateConnParam(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVW   conn_params+4(FP), R1
	SWI    $CONN_PARAM_UPDATE
	MOVW   R0, ret+8(FP)
	RET

TEXT ·disconnect(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVBU  hci_status_code+2(FP), R1
	SWI    $DISCONNECT
	MOVW   R0, ret+4(FP)
	RET

TEXT ·setTxPower(SB),NOSPLIT|NOFRAME,$0
	MOVBU  role+0(FP), R0
	MOVHU  handle+2(FP), R1
	MOVB   tx_power+4(FP), R2
	SWI    $TX_POWER_SET
	MOVW   R0, ret+8(FP)
	RET

TEXT ·setAppearance(SB),NOSPLIT|NOFRAME,$0
	MOVHU  appearance+0(FP), R0
	SWI    $APPEARANCE_SET
	MOVW   R0, ret+4(FP)
	RET

TEXT ·getAppearance(SB),NOSPLIT|NOFRAME,$0
	MOVW  appearance+0(FP), R0
	SWI   $APPEARANCE_GET
	MOVW  R0, ret+4(FP)
	RET

TEXT ·setPPCP(SB),NOSPLIT|NOFRAME,$0
	MOVW  conn_params+0(FP), R0
	SWI   $PPCP_SET
	MOVW  R0, ret+4(FP)
	RET

TEXT ·getPPCP(SB),NOSPLIT|NOFRAME,$0
	MOVW  connParams+0(FP), R0
	SWI   $PPCP_GET
	MOVW  R0, ret+4(FP)
	RET

TEXT ·setDeviceName(SB),NOSPLIT|NOFRAME,$0
	MOVW   write_perm+0(FP), R0
	MOVW   dev_name+4(FP), R1
	MOVHU  len+8(FP), R2
	SWI    $DEVICE_NAME_SET
	MOVW   R0, ret+12(FP)
	RET

TEXT ·getDeviceName(SB),NOSPLIT|NOFRAME,$0
	MOVW  dev_name+0(FP), R0
	MOVW  len+4(FP), R1
	SWI   $DEVICE_NAME_GET
	MOVW  R0, ret+8(FP)
	RET

TEXT ·authenticate(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVW   sec_params+4(FP), R1
	SWI    $AUTHENTICATE
	MOVW   R0, ret+8(FP)
	RET

TEXT ·replySecParams(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVBU  sec_status+2(FP), R1
	MOVW   sec_params+4(FP), R2
	MOVW   sec_keyset+8(FP), R3
	SWI    $SEC_PARAMS_REPLY
	MOVW   R0, ret+12(FP)
	RET

TEXT ·replyAuthKey(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVBU  key_type+2(FP), R1
	MOVW   key+4(FP), R2
	SWI    $AUTH_KEY_REPLY
	MOVW   R0, ret+8(FP)
	RET

TEXT ·replyLescDhkey(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVW   dhkey+4(FP), R1
	SWI    $LESC_DHKEY_REPLY
	MOVW   R0, ret+8(FP)
	RET

TEXT ·notifyKeypress(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVBU  kp_not+2(FP), R1
	SWI    $KEYPRESS_NOTIFY
	MOVW   R0, ret+4(FP)
	RET

TEXT ·getLescOobData(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVW   pk_own+4(FP), R1
	MOVW   oobd_own+8(FP), R2
	SWI    $LESC_OOB_DATA_GET
	MOVW   R0, ret+12(FP)
	RET

TEXT ·setLescOobData(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVW   oobd_own+4(FP), R1
	MOVW   oobd_peer+8(FP), R2
	SWI    $LESC_OOB_DATA_SET
	MOVW   R0, ret+12(FP)
	RET

TEXT ·encrypt(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVW   master_id+4(FP), R1
	MOVW   enc_info+8(FP), R2
	SWI    $ENCRYPT
	MOVW   R0, ret+12(FP)
	RET

TEXT ·replySecInfo(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVW   enc_info+4(FP), R1
	MOVW   id_info+8(FP), R2
	MOVW   sign_info+12(FP), R3
	SWI    $SEC_INFO_REPLY
	MOVW   R0, ret+16(FP)
	RET

TEXT ·getConnSec(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVW   conn_sec+4(FP), R1
	SWI    $CONN_SEC_GET
	MOVW   R0, ret+8(FP)
	RET

TEXT ·startRssi(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVBU  threshold_dbm+2(FP), R1
	MOVBU  skip_count+3(FP), R2
	SWI    $RSSI_START
	MOVW   R0, ret+4(FP)
	RET

TEXT ·stopRssi(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	SWI    $RSSI_STOP
	MOVW   R0, ret+4(FP)
	RET

TEXT ·getRssi(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVW   rssi+4(FP), R1
	MOVW   ch_index+8(FP), R2
	SWI    $RSSI_GET
	MOVW   R0, ret+12(FP)
	RET

TEXT ·startScan(SB),NOSPLIT|NOFRAME,$0
	MOVW  scan_params+0(FP), R0
	MOVW  adv_report_buffer+4(FP), R1
	SWI   $SCAN_START
	MOVW  R0, ret+8(FP)
	RET

TEXT ·stopScan(SB),NOSPLIT|NOFRAME,$0
	SWI   $SCAN_STOP
	MOVW  R0, ret+0(FP)
	RET

TEXT ·connect(SB),NOSPLIT|NOFRAME,$0
	MOVW   peer_addr+0(FP), R0
	MOVW   scan_params+4(FP), R1
	MOVW   conn_params+8(FP), R2
	MOVBU  conn_cfg_tag+12(FP), R3
	SWI    $CONNECT
	MOVW   R0, ret+16(FP)
	RET

TEXT ·cancelConnect(SB),NOSPLIT|NOFRAME,$0
	SWI   $CONNECT_CANCEL
	MOVW  R0, ret+0(FP)
	RET

TEXT ·updatePhy(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVW   gap_phys+4(FP), R1
	SWI    $PHY_UPDATE
	MOVW   R0, ret+8(FP)
	RET

TEXT ·updateDataLength(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVW   dl_params+4(FP), R1
	MOVW   dl_limitation+8(FP), R2
	SWI    $DATA_LENGTH_UPDATE
	MOVW   R0, ret+12(FP)
	RET

TEXT ·startQosChannelSurvey(SB),NOSPLIT|NOFRAME,$0
	MOVW  interval_us+0(FP), R0
	SWI   $QOS_CHANNEL_SURVEY_START
	MOVW  R0, ret+4(FP)
	RET

TEXT ·stopQosChannelSurvey(SB),NOSPLIT|NOFRAME,$0
	SWI   $QOS_CHANNEL_SURVEY_STOP
	MOVW  R0, ret+0(FP)
	RET

TEXT ·getNextConnEvtCounter(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVW   counter+4(FP), R1
	SWI    $NEXT_CONN_EVT_COUNTER_GET
	MOVW   R0, ret+8(FP)
	RET

TEXT ·startConnEvtTrigger(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	MOVW   params+4(FP), R1
	SWI    $CONN_EVT_TRIGGER_START
	MOVW   R0, ret+8(FP)
	RET

TEXT ·stopConnEvtTrigger(SB),NOSPLIT|NOFRAME,$0
	MOVHU  conn_handle+0(FP), R0
	SWI    $CONN_EVT_TRIGGER_STOP
	MOVW   R0, ret+4(FP)
	RET
