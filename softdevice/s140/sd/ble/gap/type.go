// Based on C code by Nordic Semiconductor ASA.
// See LICENSE-NORDIC for original C code license.

// Copyright 2019 Michal Derkacz. All rights reserved.

package gap

import "github.com/embeddedgo/nrf5/softdevice/s140/sd/ble"

type ConnCfg struct {
	Tag         uint8
	_           uint8
	ConnCount   uint8  // max. number of concurrent connections
	EventLength uint16 // time used on every connection interval in 1.25 ms uints
}

type RoleCountCfg struct {
	AdvSetCount                   uint8 // max. number of advertising sets
	PeriphRoleCount               uint8 // max. number of connections acting as a peripheral
	CentralRoleCount              uint8 // max. number of connections acting as a central
	CentralSecCount               uint8 // number of SMP instances shared between connections acting as a central
	QoSChannelSurveyRoleAvailable bool
}

type ConnSecMode uint8

const (
	// No access.
	SecModeNoAccess ConnSecMode = 0x00 // level=0, mode=0

	// No protection required (open link).
	SecModeOpen ConnSecMode = 0x11 // level=1, mode=1

	// Encryption required, but no MITM protection.
	SecModeEncNoMTIM ConnSecMode = 0x21 // level=2, mode=1

	// Encryption rand MITM protection required.
	SecModeEncWithMTIM ConnSecMode = 0x31 // level=3, mode=1

	// LESC encryption and MITM protection required.
	SecModeLESCWithMTIM_MITM ConnSecMode = 0x41 // level=4, mode1

	// Signing or encryption required, no MITM protection.
	SecModeSignedNoMTIM ConnSecMode = 0x12 // level=1, mode=2

	// Signing or encryption with MITM protection required.
	SecModeSignedWithMTIM ConnSecMode = 0x22 // level=2, mode=2
)

type ConnParams struct {
	MinConnInterval uint16 // min. connection interval in 1.25 ms units
	MaxConnInterval uint16 // max. connection interval in 1.25 ms units
	SlaveLatency    uint16 // slave latency in number of connection events
	ConnSupTimeout  uint16 // connection supervision timeout in 10 ms units
}

const (
	MaxAdvSetDataSize         = 31
	MaxExtendedAdvSetDataSize = 255
)

type AdvData struct {
	Data        ble.Data // advertising data
	ScanRspData ble.Data // scan response data
}

type AdvSetHandle uint8

const NewAdvSet AdvSetHandle = 0xFF

type FilterPolicy uint8

const (
	PolicyAny     FilterPolicy = 0 // allow scan/connect requests from any dev.
	PolicyScanReq FilterPolicy = 1 // filter scan requests with whitelist
	PolicyConnReq FilterPolicy = 2 // filter connect requests with whitelist
	PolicyBoth    FilterPolicy = 3 // filter both scan andconnect requests with whitelist
)

type AdvType uint8

const (
	// Connectable and scannable undirected advertising events.
	ConnectableScannableUndirected AdvType = 1

	// Connectable non-scannable directed advertising events. Advertising
	// interval is less that 3.75 ms. Use this type for fast reconnections.
	// Advertising data is not supported.
	ConnectableNonscannableDirectedHighDutyCycle AdvType = 2

	// Connectable non-scannable directed advertising events. Advertising data
	// is not supported.
	ConnectableNonscannableDirected AdvType = 3

	// Non-connectable scannable undirected advertising events.
	NonconnectableScannableUndirected AdvType = 4

	// Non-connectable non-scannable undirected advertising events.
	NonconnectableNonscannableUndirected AdvType = 5

	// Connectable non-scannable undirected advertising events using extended
	// advertising PDUs.
	ExtendedConnectableNonscannableUndirected AdvType = 6

	// Connectable non-scannable directed advertising events using extended
	// advertising PDUs.
	ExtendedConnectableNonscannableDirected AdvType = 7

	// Non-connectable scannable undirected advertising events using extended
	// advertising PDUs. Only scan response data is supported.
	ExtendedNonconnectableScannableUndirected AdvType = 8

	// Non-connectable scannable directed advertising events using extended
	// advertising PDUs. Only scan response data is supported.
	ExtendedNonconnectableScannableDirected AdvType = 9

	// Non-connectable non-scannable undirected advertising events using
	// extended advertising PDUs.
	ExtendedNonconnectableNonscannableUndirected AdvType = 10

	// Non-connectable non-scannable directed advertising events using extended
	// advertising PDUs.
	ExtendedNonconnectableNonscannableDirected AdvType = 11
)

type AdvOptions uint8

const (
	AdvAnonymous      AdvOptions = 1 << 0 // omit advertiser's address
	AdvIncludeTxPower AdvOptions = 1 << 1 // not supported on this SoftDevice
)

type AdvProps struct {
	Type    AdvType
	Options AdvOptions
}

type AddrType uint8

const (
	PeerID                     AddrType = 1 << 0
	AddrKind                   AddrType = 0x7F << 1
	Public                     AddrType = 0x00 << 1
	RandomStatic               AddrType = 0x01 << 1
	RandomPrivateResolvable    AddrType = 0x02 << 1
	RandomPrivateNonresolvable AddrType = 0x03 << 1
	Anonymous                  AddrType = 0x7F << 1
)

type Addr struct {
	Type  AddrType
	Value [6]byte
}

type ChanMask [5]byte

type PHY uint8

const (
	PHYAuto   PHY = 0
	PHY1Mbps  PHY = 1
	PHY2Mbps  PHY = 2
	PHYCoded  PHY = 4
	PHYNotSet PHY = 255
)

type AdvSet uint8

const (
	AdvSetID      AdvSet = 0x0F << 0
	NotifyScanReq AdvSet = 1 << 4
)

type AdvParams struct {
	Props        AdvProps     // properties of the advertising events
	PeerAddr     *Addr        // address of a known peer or special pseudo-address
	Interval     uint32       // advertising interval in 625 us units
	Duration     uint16       // advertising duration in 10 ms units
	MaxEvts      uint8        // maximum advertising events that can be sent
	ChanMask     ChanMask     // channel mask for primary and secondary adv. channels
	Policy       FilterPolicy // filter policy
	PrimaryPHY   PHY          // for primary advertising channel packets
	SecondaryPHY PHY          // for secondary advertising channel packets
	Set          AdvSet
}

type AdvFlags uint8

const (
	LimitedDiscMode    AdvFlags = 1 << 0 // LE limited discoverable mode
	GeneralDiscMode    AdvFlags = 1 << 1 // LE general discoverable mode
	LEOnly             AdvFlags = 1 << 2 // BR/EDR not supported. */
	DualModeController AdvFlags = 1 << 3 // simultaneous LE and BR/EDR, controller
	DualModeHost       AdvFlags = 1 << 4 // simultaneous LE and BR/EDR, host
)

type AdvDataType uint8

const (

	// Flags for discoverability.
	Flags AdvDataType = 0x01

	// Partial list of 16-bit service UUIDs.
	PartialServiceUUID16 AdvDataType = 0x02

	// Complete list of 16-bit service UUIDs.
	CompleteServiceUUID16 AdvDataType = 0x03

	// Partial list of 32-bit service UUIDs.
	PartialServiceUUID32 AdvDataType = 0x04

	// Complete list of 32-bit service UUIDs.
	CompleteServiceUUID32 AdvDataType = 0x05

	// Partial list of 128-bit service UUIDs.
	PartialServiceUUID128 AdvDataType = 0x06

	// Complete list of 128-bit service UUIDs.
	CompleteServiceUUID128 AdvDataType = 0x07

	// Short local device name.
	ShortLocalName AdvDataType = 0x08

	// Complete local device name.
	CompleteLocalName AdvDataType = 0x09

	// Transmit power level.
	TxPowerLevel AdvDataType = 0x0A

	// Class of device.
	ClassOfDevice AdvDataType = 0x0D

	// Simple Pairing Hash C.
	SimplePairingHashC AdvDataType = 0x0E

	// Simple Pairing Randomizer R.
	SimplePairingRandomizerR AdvDataType = 0x0F

	// Security Manager TK Value.
	SecurityManagerTKValue AdvDataType = 0x10

	// Security Manager Out Of Band Flags.
	SecurityManagerOOBFlags AdvDataType = 0x11

	// Slave Connection Interval Range
	SlaveConnectionIntervalRange AdvDataType = 0x12

	// List of 16-bit Service Solicitation UUIDs.
	SolicitedServiceUUID16 AdvDataType = 0x14

	// List of 128-bit Service Solicitation UUIDs.
	SolicitedServiceUUID128 AdvDataType = 0x15

	// Service Data - 16-bit UUID.
	ServiceData AdvDataType = 0x16

	// Public Target Address.
	PublicTargetAddress AdvDataType = 0x17

	// Random Target Address.
	RandomTargetAddress AdvDataType = 0x18

	// Appearance.
	Appearance AdvDataType = 0x19

	// Advertising Interval.
	AdvertisingInterval AdvDataType = 0x1A

	// LE Bluetooth Device Address.
	LEBluetoothDeviceAddress AdvDataType = 0x1B

	// LE Role.
	LERole AdvDataType = 0x1C

	// Simple Pairing Hash C-256.
	SimplePairingHashC256 AdvDataType = 0x1D

	// Simple Pairing Randomizer R-256.
	SimplePairingRandomizerR256 AdvDataType = 0x1E

	// Service Data - 32-bit UUID.
	ServiceDataUUID32 AdvDataType = 0x20

	// Service Data - 128-bit UUID.
	ServiceDataUUID128 AdvDataType = 0x21

	// LE Secure Connections Confirmation Value.
	LESC_ConfirmationValue AdvDataType = 0x22

	// LE Secure Connections Random Value.
	LESCRandomValue AdvDataType = 0x23

	// URI
	URI AdvDataType = 0x24

	// 3D Information Data.
	InformationData3D AdvDataType = 0x3D

	// Manufacturer Specific Data.
	ManufacturerSpecificData AdvDataType = 0xFF
)

type Role uint8

const (
	RoleInvalid Role = 0 // invalid role
	RolePeriph  Role = 1 // peripheral role
	RoleCentral Role = 2 // central
)

// DataLengthParams contains maximum length/time that a Controller supports for
// single Link Layer Data Channel PDU.
type DataLengthParams struct {
	MaxTxOctets uint16
	MaxRxOctets uint16
	MaxTxTimeUs uint16
	MaxRxTimeUs uint16
}

type DataLengthLimitation struct {
	TxPayloadLimitedOctets uint16 // If > 0, the requested TX packet length is too long by this many octets.
	RxPayloadLimitedOctets uint16 // If > 0, the requested RX packet length is too long by this many octets.
	TxRxTimeLimitedUs      uint16 // If > 0, the requested combination of TX and RX packet lengths is too long by this many microseconds.
}
