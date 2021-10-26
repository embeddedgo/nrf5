// Based on C code by Nordic Semiconductor ASA.
// See LICENSE-NORDIC for original C code license.

// Copyright 2020 Michal Derkacz.

package gap

import "github.com/embeddedgo/nrf5/softdevice/s140/sd/ble"

const (
	EvtBase = 0x10
	EvtLast = 0x2F
)

const (
	EvtConnected              = EvtBase + 0  // connected to peer
	EvtDisconnected           = EvtBase + 1  // disconnected from peer
	EvtConnParamUpdate        = EvtBase + 2  // connection parameters updated
	EvtSecParamsReq           = EvtBase + 3  // req. to provide security param
	EvtSecInfoReq             = EvtBase + 4  // req. to provide security info
	EvtPasskeyDisplay         = EvtBase + 5  // req. to display a passkey
	EvtKeyPressed             = EvtBase + 6  // keypress on the remote device
	EvtAuthKeyReq             = EvtBase + 7  // req. to provide an auth. key
	EvtLESCDHKeyReq           = EvtBase + 8  // req. to calculate LESConn DHKey
	EvtAuthStatus             = EvtBase + 9  // auth. procedure completed
	EvtConnSecUpdate          = EvtBase + 10 // connection security update
	EvtTimeout                = EvtBase + 11 // timeout expire
	EvtRSSIChanged            = EvtBase + 12 // RSSI repor
	EvtAdvReport              = EvtBase + 13 // advertising repor
	EvtSecReq                 = EvtBase + 14 // security reques
	EvtConnParamUpdateReq     = EvtBase + 15 // connection param. update req.
	EvtScanReqReport          = EvtBase + 16 // scan request report
	EvtPHYUpdateReq           = EvtBase + 17 // PHY update request
	EvtPHYUpdate              = EvtBase + 18 // PHY update procedure completed
	EvtDataLengthUpdateReq    = EvtBase + 19 // data length update request
	EvtDataLengthUpdate       = EvtBase + 20 // LL data ch. PDU pay.len. update
	EvtQOSChannelSurveyReport = EvtBase + 21 // channel survey report
	EvtADVSetTerminated       = EvtBase + 22 // advertising set terminated
)

type Connected struct {
	ble.EvtHdr
	_          uint16
	PeerAddr   Addr       // bluetooth address of the peer device
	Role       Role       // BLE role for this connection
	ConnParams ConnParams // GAP connection parameters
	AdvHandle  uint8      // advertising handle in which advertising has ended
	AdvData    AdvData    // buffers corresponding to the terminated advert. set
}

type Disconnected struct {
	ble.EvtHdr
	_      uint16
	Reason ble.HCIStatus // HCI error code
}

type ConnParamUpdate struct {
	ble.EvtHdr
	_          uint16
	ConnParams ConnParams // GAP connection parameters
}

type DataLengthUpdateReq struct {
	ble.EvtHdr
	_          uint16
	PeerParams DataLengthParams // Peer data length parameters
}
