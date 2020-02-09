// Code generated from ../nordic/s140/headers/ble_gap.h; DO NOT EDIT.

#include "go_asm.h"
#include "textflag.h"

TEXT ·SetGAPAddr(SB),NOSPLIT|NOFRAME,$0
	MOVW addr+0(FP), R0
	SWI $const__GAP_ADDR_SET
	MOVW R0, ret+4(FP)
	RET

TEXT ·GetGAPAddr(SB),NOSPLIT|NOFRAME,$0
	MOVW addr+0(FP), R0
	SWI $const__GAP_ADDR_GET
	MOVW R0, ret+4(FP)
	RET

TEXT ·GetGAPAdvAddr(SB),NOSPLIT|NOFRAME,$0
	MOVBU adv_handle+0(FP), R0
	MOVW addr+4(FP), R1
	SWI $const__GAP_ADV_ADDR_GET
	MOVW R0, ret+8(FP)
	RET

TEXT ·SetGAPWhitelist(SB),NOSPLIT|NOFRAME,$0
	MOVW pp_wl_addrs+0(FP), R0
	MOVBU len+4(FP), R1
	SWI $const__GAP_WHITELIST_SET
	MOVW R0, ret+8(FP)
	RET

TEXT ·SetGAPDeviceIdentities(SB),NOSPLIT|NOFRAME,$0
	MOVW pp_id_keys+0(FP), R0
	MOVW pp_local_irks+4(FP), R1
	MOVBU len+8(FP), R2
	SWI $const__GAP_DEVICE_IDENTITIES_SET
	MOVW R0, ret+12(FP)
	RET

TEXT ·SetGAPPrivacy(SB),NOSPLIT|NOFRAME,$0
	MOVW privacy_params+0(FP), R0
	SWI $const__GAP_PRIVACY_SET
	MOVW R0, ret+4(FP)
	RET

TEXT ·GetGAPPrivacy(SB),NOSPLIT|NOFRAME,$0
	MOVW privacy_params+0(FP), R0
	SWI $const__GAP_PRIVACY_GET
	MOVW R0, ret+4(FP)
	RET

TEXT ·ConfigureGAPAdvSet(SB),NOSPLIT|NOFRAME,$0
	MOVW adv_handle+0(FP), R0
	MOVW adv_data+4(FP), R1
	MOVW adv_params+8(FP), R2
	SWI $const__GAP_ADV_SET_CONFIGURE
	MOVW R0, ret+12(FP)
	RET

TEXT ·StartGAPAdv(SB),NOSPLIT|NOFRAME,$0
	MOVBU adv_handle+0(FP), R0
	MOVBU conn_cfg_tag+1(FP), R1
	SWI $const__GAP_ADV_START
	MOVW R0, ret+4(FP)
	RET

TEXT ·StopGAPAdv(SB),NOSPLIT|NOFRAME,$0
	MOVBU adv_handle+0(FP), R0
	SWI $const__GAP_ADV_STOP
	MOVW R0, ret+4(FP)
	RET

TEXT ·UpdateGAPConnParam(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVW conn_params+4(FP), R1
	SWI $const__GAP_CONN_PARAM_UPDATE
	MOVW R0, ret+8(FP)
	RET

TEXT ·DisconnectGAP(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVBU hci_status_code+2(FP), R1
	SWI $const__GAP_DISCONNECT
	MOVW R0, ret+4(FP)
	RET

TEXT ·SetGAPTxPower(SB),NOSPLIT|NOFRAME,$0
	MOVBU role+0(FP), R0
	MOVHU handle+2(FP), R1
	MOVB tx_power+4(FP), R2
	SWI $const__GAP_TX_POWER_SET
	MOVW R0, ret+8(FP)
	RET

TEXT ·SetGAPAppearance(SB),NOSPLIT|NOFRAME,$0
	MOVHU appearance+0(FP), R0
	SWI $const__GAP_APPEARANCE_SET
	MOVW R0, ret+4(FP)
	RET

TEXT ·GetGAPAppearance(SB),NOSPLIT|NOFRAME,$0
	MOVW appearance+0(FP), R0
	SWI $const__GAP_APPEARANCE_GET
	MOVW R0, ret+4(FP)
	RET

TEXT ·SetGAPPpcp(SB),NOSPLIT|NOFRAME,$0
	MOVW conn_params+0(FP), R0
	SWI $const__GAP_PPCP_SET
	MOVW R0, ret+4(FP)
	RET

TEXT ·GetGAPPpcp(SB),NOSPLIT|NOFRAME,$0
	MOVW conn_params+0(FP), R0
	SWI $const__GAP_PPCP_GET
	MOVW R0, ret+4(FP)
	RET

TEXT ·SetGAPDeviceName(SB),NOSPLIT|NOFRAME,$0
	MOVW write_perm+0(FP), R0
	MOVW dev_name+4(FP), R1
	MOVHU len+8(FP), R2
	SWI $const__GAP_DEVICE_NAME_SET
	MOVW R0, ret+12(FP)
	RET

TEXT ·GetGAPDeviceName(SB),NOSPLIT|NOFRAME,$0
	MOVW dev_name+0(FP), R0
	MOVW len+4(FP), R1
	SWI $const__GAP_DEVICE_NAME_GET
	MOVW R0, ret+8(FP)
	RET

TEXT ·AuthenticateGAP(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVW sec_params+4(FP), R1
	SWI $const__GAP_AUTHENTICATE
	MOVW R0, ret+8(FP)
	RET

TEXT ·ReplyGAPSecParams(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVBU sec_status+2(FP), R1
	MOVW sec_params+4(FP), R2
	MOVW sec_keyset+8(FP), R3
	SWI $const__GAP_SEC_PARAMS_REPLY
	MOVW R0, ret+12(FP)
	RET

TEXT ·ReplyGAPAuthKey(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVBU key_type+2(FP), R1
	MOVW key+4(FP), R2
	SWI $const__GAP_AUTH_KEY_REPLY
	MOVW R0, ret+8(FP)
	RET

TEXT ·ReplyGAPLescDhkey(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVW dhkey+4(FP), R1
	SWI $const__GAP_LESC_DHKEY_REPLY
	MOVW R0, ret+8(FP)
	RET

TEXT ·NotifyGAPKeypress(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVBU kp_not+2(FP), R1
	SWI $const__GAP_KEYPRESS_NOTIFY
	MOVW R0, ret+4(FP)
	RET

TEXT ·GetGAPLescOobData(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVW pk_own+4(FP), R1
	MOVW oobd_own+8(FP), R2
	SWI $const__GAP_LESC_OOB_DATA_GET
	MOVW R0, ret+12(FP)
	RET

TEXT ·SetGAPLescOobData(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVW oobd_own+4(FP), R1
	MOVW oobd_peer+8(FP), R2
	SWI $const__GAP_LESC_OOB_DATA_SET
	MOVW R0, ret+12(FP)
	RET

TEXT ·EncryptGAP(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVW master_id+4(FP), R1
	MOVW enc_info+8(FP), R2
	SWI $const__GAP_ENCRYPT
	MOVW R0, ret+12(FP)
	RET

TEXT ·ReplyGAPSecInfo(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVW enc_info+4(FP), R1
	MOVW id_info+8(FP), R2
	MOVW sign_info+12(FP), R3
	SWI $const__GAP_SEC_INFO_REPLY
	MOVW R0, ret+16(FP)
	RET

TEXT ·GetGAPConnSec(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVW conn_sec+4(FP), R1
	SWI $const__GAP_CONN_SEC_GET
	MOVW R0, ret+8(FP)
	RET

TEXT ·StartGAPRssi(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVBU threshold_dbm+2(FP), R1
	MOVBU skip_count+3(FP), R2
	SWI $const__GAP_RSSI_START
	MOVW R0, ret+4(FP)
	RET

TEXT ·StopGAPRssi(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	SWI $const__GAP_RSSI_STOP
	MOVW R0, ret+4(FP)
	RET

TEXT ·GetGAPRssi(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVW rssi+4(FP), R1
	MOVW ch_index+8(FP), R2
	SWI $const__GAP_RSSI_GET
	MOVW R0, ret+12(FP)
	RET

TEXT ·StartGAPScan(SB),NOSPLIT|NOFRAME,$0
	MOVW scan_params+0(FP), R0
	MOVW adv_report_buffer+4(FP), R1
	SWI $const__GAP_SCAN_START
	MOVW R0, ret+8(FP)
	RET

TEXT ·StopGAPScan(SB),NOSPLIT|NOFRAME,$0
	SWI $const__GAP_SCAN_STOP
	MOVW R0, ret+0(FP)
	RET

TEXT ·ConnectGAP(SB),NOSPLIT|NOFRAME,$0
	MOVW peer_addr+0(FP), R0
	MOVW scan_params+4(FP), R1
	MOVW conn_params+8(FP), R2
	MOVBU conn_cfg_tag+12(FP), R3
	SWI $const__GAP_CONNECT
	MOVW R0, ret+16(FP)
	RET

TEXT ·CancelGAPConnect(SB),NOSPLIT|NOFRAME,$0
	SWI $const__GAP_CONNECT_CANCEL
	MOVW R0, ret+0(FP)
	RET

TEXT ·UpdateGAPPhy(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVW gap_phys+4(FP), R1
	SWI $const__GAP_PHY_UPDATE
	MOVW R0, ret+8(FP)
	RET

TEXT ·UpdateGAPDataLength(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVW dl_params+4(FP), R1
	MOVW dl_limitation+8(FP), R2
	SWI $const__GAP_DATA_LENGTH_UPDATE
	MOVW R0, ret+12(FP)
	RET

TEXT ·StartGAPQosChannelSurvey(SB),NOSPLIT|NOFRAME,$0
	MOVW interval_us+0(FP), R0
	SWI $const__GAP_QOS_CHANNEL_SURVEY_START
	MOVW R0, ret+4(FP)
	RET

TEXT ·StopGAPQosChannelSurvey(SB),NOSPLIT|NOFRAME,$0
	SWI $const__GAP_QOS_CHANNEL_SURVEY_STOP
	MOVW R0, ret+0(FP)
	RET

TEXT ·GetGAPNextConnEvtCounter(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVW counter+4(FP), R1
	SWI $const__GAP_NEXT_CONN_EVT_COUNTER_GET
	MOVW R0, ret+8(FP)
	RET

TEXT ·StartGAPConnEvtTrigger(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	MOVW params+4(FP), R1
	SWI $const__GAP_CONN_EVT_TRIGGER_START
	MOVW R0, ret+8(FP)
	RET

TEXT ·StopGAPConnEvtTrigger(SB),NOSPLIT|NOFRAME,$0
	MOVHU conn_handle+0(FP), R0
	SWI $const__GAP_CONN_EVT_TRIGGER_STOP
	MOVW R0, ret+4(FP)
	RET
