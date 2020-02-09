// Code generated from ../nordic/s140/headers/ble_l2cap.h; DO NOT EDIT.

package ble

const (
	_L2CAP_CH_SETUP        = L2CAP_SVC_BASE + 0
	_L2CAP_CH_RELEASE      = L2CAP_SVC_BASE + 1
	_L2CAP_CH_RX           = L2CAP_SVC_BASE + 2
	_L2CAP_CH_TX           = L2CAP_SVC_BASE + 3
	_L2CAP_CH_FLOW_CONTROL = L2CAP_SVC_BASE + 4
)

const (
	L2CAP_EVT_CH_SETUP_REQUEST    = L2CAP_EVT_BASE + 0 // L2CAP Channel Setup Request event. \n See @ref ble_l2cap_evt_ch_setup_request_t.
	L2CAP_EVT_CH_SETUP_REFUSED    = L2CAP_EVT_BASE + 1 // L2CAP Channel Setup Refused event. \n See @ref ble_l2cap_evt_ch_setup_refused_t.
	L2CAP_EVT_CH_SETUP            = L2CAP_EVT_BASE + 2 // L2CAP Channel Setup Completed event. \n See @ref ble_l2cap_evt_ch_setup_t.
	L2CAP_EVT_CH_RELEASED         = L2CAP_EVT_BASE + 3 // L2CAP Channel Released event. \n No additional event structure applies.
	L2CAP_EVT_CH_SDU_BUF_RELEASED = L2CAP_EVT_BASE + 4 // L2CAP Channel SDU data buffer released event. \n See @ref ble_l2cap_evt_ch_sdu_buf_released_t.
	L2CAP_EVT_CH_CREDIT           = L2CAP_EVT_BASE + 5 // L2CAP Channel Credit received. \n See @ref ble_l2cap_evt_ch_credit_t.
	L2CAP_EVT_CH_RX               = L2CAP_EVT_BASE + 6 // L2CAP Channel SDU received. \n See @ref ble_l2cap_evt_ch_rx_t.
	L2CAP_EVT_CH_TX               = L2CAP_EVT_BASE + 7 // L2CAP Channel SDU transmitted. \n See @ref ble_l2cap_evt_ch_tx_t.
)

const L2CAP_CH_COUNT_MAX = (64)
const L2CAP_MTU_MIN = (23)
const L2CAP_MPS_MIN = (23)
const L2CAP_CID_INVALID = (0x0000)
const L2CAP_CREDITS_DEFAULT = (1)
const L2CAP_CH_SETUP_REFUSED_SRC_LOCAL = (0x01)             // Local.
const L2CAP_CH_SETUP_REFUSED_SRC_REMOTE = (0x02)            // Remote.
const L2CAP_CH_STATUS_CODE_SUCCESS = (0x0000)               // Success.
const L2CAP_CH_STATUS_CODE_LE_PSM_NOT_SUPPORTED = (0x0002)  // LE_PSM not supported.
const L2CAP_CH_STATUS_CODE_NO_RESOURCES = (0x0004)          // No resources available.
const L2CAP_CH_STATUS_CODE_INSUFF_AUTHENTICATION = (0x0005) // Insufficient authentication.
const L2CAP_CH_STATUS_CODE_INSUFF_AUTHORIZATION = (0x0006)  // Insufficient authorization.
const L2CAP_CH_STATUS_CODE_INSUFF_ENC_KEY_SIZE = (0x0007)   // Insufficient encryption key size.
const L2CAP_CH_STATUS_CODE_INSUFF_ENC = (0x0008)            // Insufficient encryption.
const L2CAP_CH_STATUS_CODE_INVALID_SCID = (0x0009)          // Invalid Source CID.
const L2CAP_CH_STATUS_CODE_SCID_ALLOCATED = (0x000A)        // Source CID already allocated.
const L2CAP_CH_STATUS_CODE_UNACCEPTABLE_PARAMS = (0x000B)   // Unacceptable parameters.
const L2CAP_CH_STATUS_CODE_NOT_UNDERSTOOD = (0x8000)        // Command Reject received instead of LE Credit Based Connection Response.
const L2CAP_CH_STATUS_CODE_TIMEOUT = (0xC000)               // Operation timed out.

type L2capConnCfg struct {
	RxMps       uint16 // The maximum L2CAP PDU payload size, in bytes, that L2CAP shall be able to receive on L2CAP channels on connections with this configuration. The minimum value is @ref BLE_L2CAP_MPS_MIN.
	TxMps       uint16 // The maximum L2CAP PDU payload size, in bytes, that L2CAP shall be able to transmit on L2CAP channels on connections with this configuration. The minimum value is @ref BLE_L2CAP_MPS_MIN.
	RxQueueSize uint8  // Number of SDU data buffers that can be queued for reception per L2CAP channel. The minimum value is one.
	TxQueueSize uint8  // Number of SDU data buffers that can be queued for transmission per L2CAP channel. The minimum value is one.
	ChCount     uint8  // Number of L2CAP channels the application can create per connection with this configuration. The default value is zero, the maximum value is @ref BLE_L2CAP_CH_COUNT_MAX. @note if this parameter is set to zero, all other parameters in @ref ble_l2cap_conn_cfg_t are ignored.
}

type L2capChRxParams struct {
	RxMtu  uint16 // The maximum L2CAP SDU size, in bytes, that L2CAP shall be able to receive on this L2CAP channel. - Must be equal to or greater than @ref BLE_L2CAP_MTU_MIN.
	RxMps  uint16 // The maximum L2CAP PDU payload size, in bytes, that L2CAP shall be able to receive on this L2CAP channel. - Must be equal to or greater than @ref BLE_L2CAP_MPS_MIN. - Must be equal to or less than @ref ble_l2cap_conn_cfg_t::rx_mps.
	SduBuf Data   // SDU data buffer for reception. - If @ref ble_data_t::p_data is non-NULL, initial credits are issued to the peer. - If @ref ble_data_t::p_data is NULL, no initial credits are issued to the peer.
}

type L2capChSetupParams struct {
	RxParams L2capChRxParams // L2CAP channel RX parameters.
	LePsm    uint16          // LE Protocol/Service Multiplexer. Used when requesting setup of an L2CAP channel, ignored otherwise.
	Status   uint16          // Status code, see @ref BLE_L2CAP_CH_STATUS_CODES. Used when replying to a setup request of an L2CAP channel, ignored otherwise.
}

type L2capChTxParams struct {
	TxMtu   uint16 // The maximum L2CAP SDU size, in bytes, that L2CAP is able to transmit on this L2CAP channel.
	PeerMps uint16 // The maximum L2CAP PDU payload size, in bytes, that the peer is able to receive on this L2CAP channel.
	TxMps   uint16 // The maximum L2CAP PDU payload size, in bytes, that L2CAP is able to transmit on this L2CAP channel. This is effective tx_mps, selected by the SoftDevice as MIN( @ref ble_l2cap_ch_tx_params_t::peer_mps, @ref ble_l2cap_conn_cfg_t::tx_mps )
	Credits uint16 // Initial credits given by the peer.
}

type L2capEvtChSetupRequest struct {
	TxParams L2capChTxParams // L2CAP channel TX parameters.
	LePsm    uint16          // LE Protocol/Service Multiplexer.
}

type L2capEvtChSetupRefused struct {
	Source uint8  // Source, see @ref BLE_L2CAP_CH_SETUP_REFUSED_SRCS
	Status uint16 // Status code, see @ref BLE_L2CAP_CH_STATUS_CODES
}

type L2capEvtChSetup struct {
	TxParams L2capChTxParams // L2CAP channel TX parameters.
}

type L2capEvtChSduBufReleased struct {
	SduBuf Data // Returned reception or transmission SDU data buffer. The SoftDevice returns SDU data buffers supplied by the application, which have not yet been returned previously via a @ref BLE_L2CAP_EVT_CH_RX or @ref BLE_L2CAP_EVT_CH_TX event.
}

type L2capEvtChCredit struct {
	Credits uint16 // Additional credits given by the peer.
}

type L2capEvtChRx struct {
	SduLen uint16 // Total SDU length, in bytes.
	SduBuf Data   // SDU data buffer. @note If there is not enough space in the buffer (sdu_buf.len < sdu_len) then the rest of the SDU will be silently discarded by the SoftDevice.
}

type L2capEvtChTx struct {
	SduBuf Data // SDU data buffer.
}

//type L2capEvt struct { union }

// Set up an L2CAP channel.
func SetupL2capCh(conn_handle uint16, local_cid *uint16, params *L2capChSetupParams) uint32

// Release an L2CAP channel.
func ReleaseL2capCh(conn_handle uint16, local_cid uint16) uint32

// Receive an SDU on an L2CAP channel.
func RxL2capCh(conn_handle uint16, local_cid uint16, sdu_buf *Data) uint32

// Transmit an SDU on an L2CAP channel.
func TxL2capCh(conn_handle uint16, local_cid uint16, sdu_buf *Data) uint32

// Advanced SDU reception flow control.
func ControlL2capChFlow(conn_handle uint16, local_cid uint16, credits uint16, credits *uint16) uint32
