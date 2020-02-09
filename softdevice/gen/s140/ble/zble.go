// Code generated from ../nordic/s140/headers/ble.h; DO NOT EDIT.

package ble

const (
	_ENABLE = SVC_BASE + iota
	_EVT_GET
	_UUID_VS_ADD
	_UUID_DECODE
	_UUID_ENCODE
	_VERSION_GET
	_USER_MEM_REPLY
	_OPT_SET
	_OPT_GET
	_CFG_SET
	_UUID_VS_REMOVE
)

const (
	EVT_USER_MEM_REQUEST = EVT_BASE + 0 // User Memory request. @ref ble_evt_user_mem_request_t
	EVT_USER_MEM_RELEASE = EVT_BASE + 1 // User Memory release. @ref ble_evt_user_mem_release_t
)

const (
	CONN_CFG_GAP   = CONN_CFG_BASE + 0 // BLE GAP specific connection configuration.
	CONN_CFG_GATTC = CONN_CFG_BASE + 1 // BLE GATTC specific connection configuration.
	CONN_CFG_GATTS = CONN_CFG_BASE + 2 // BLE GATTS specific connection configuration.
	CONN_CFG_GATT  = CONN_CFG_BASE + 3 // BLE GATT specific connection configuration.
	CONN_CFG_L2CAP = CONN_CFG_BASE + 4 // BLE L2CAP specific connection configuration.
)

const (
	COMMON_CFG_VS_UUID = CFG_BASE + iota // Vendor specific base UUID configuration
)

const (
	COMMON_OPT_PA_LNA          = OPT_BASE + 0 // PA and LNA options
	COMMON_OPT_CONN_EVT_EXT    = OPT_BASE + 1 // Extended connection events option
	COMMON_OPT_EXTENDED_RC_CAL = OPT_BASE + 2 // Extended RC calibration option
)

const EVT_PTR_ALIGNMENT = 4
const USER_MEM_TYPE_INVALID = 0x00             // Invalid User Memory Types.
const USER_MEM_TYPE_GATTS_QUEUED_WRITES = 0x01 // User Memory for GATTS queued writes.
const UUID_VS_COUNT_DEFAULT = 10               // Default VS UUID count.
const UUID_VS_COUNT_MAX = 254                  // Maximum VS UUID count.
const CONN_CFG_TAG_DEFAULT = 0                 // Default configuration tag, SoftDevice default connection configuration.

type UserMemBlock struct {
	Mem *uint8 // Pointer to the start of the user memory block.
	Len uint16 // Length in bytes of the user memory block.
}

type EvtUserMemRequest struct {
	Typ uint8 // User memory type, see @ref BLE_USER_MEM_TYPES.
}

type EvtUserMemRelease struct {
	Typ      uint8        // User memory type, see @ref BLE_USER_MEM_TYPES.
	MemBlock UserMemBlock // User memory block
}

//type CommonEvt struct { union }

type EvtHdr struct {
	EvtId  uint16 // Value from a BLE_<module>_EVT series.
	EvtLen uint16 // Length in octets including this header.
}

//type Evt struct { union }

type Version struct {
	VersionNumber    uint8  // Link Layer Version number. See https://www.bluetooth.org/en-us/specification/assigned-numbers/link-layer for assigned values.
	CompanyId        uint16 // Company ID, Nordic Semiconductor's company ID is 89 (0x0059) (https://www.bluetooth.org/apps/content/Default.aspx?doc_id=49708).
	SubversionNumber uint16 // Link Layer Sub Version number, corresponds to the SoftDevice Config ID or Firmware ID (FWID).
}

//type PALNACfg struct { bitfield }

type CommonOptPALNA struct {
	PACfg      PALNACfg // Power Amplifier configuration
	LNACfg     PALNACfg // Low Noise Amplifier configuration
	PPIChIdSet uint8    // PPI channel used for radio pin setting
	PPIChIdClr uint8    // PPI channel used for radio pin clearing
	GPIOTEChId uint8    // GPIOTE channel used for radio pin toggling
}

//type CommonOptConnEvtExt struct { bitfield }

//type CommonOptExtendedRCCal struct { bitfield }

//type CommonOpt union

//type Opt union

//type ConnCfg struct { union }

type CommonCfgVSUUID struct {
	VSUUIDCount uint8 // Number of 128-bit Vendor Specific base UUID bases to allocate memory for. Default value is @ref BLE_UUID_VS_COUNT_DEFAULT. Maximum value is @ref BLE_UUID_VS_COUNT_MAX.
}

//type CommonCfg union

//type Cfg union

// Enable and initialize the BLE stack
func Enable(app_ram_base *uint32) uint32

// Add a configuration to the BLE stack.
func SetCfg(cfg_id uint32, cfg *Cfg, app_ram_base uint32) uint32

// Get an event from the pending events queue.
func GetEvt(dest *uint8, len *uint16) uint32

// Add a Vendor Specific base UUID.
func AddUUIDVS(vs_uuid *UUID128, uuid_type *uint8) uint32

// Remove a Vendor Specific base UUID.
func RemoveUUIDVS(uuid_type *uint8) uint32

// Decode UUID bytes.
func DecodeUUID(uuid_le_len uint8, uuid_le *uint8, uuid *UUID) uint32

// Encode UUID bytes.
func EncodeUUID(uuid *UUID, uuid_le_len *uint8, uuid_le *uint8) uint32

// Get the local version information (company ID, Link Layer Version, Link Layer Subversion).
func GetVersion(version *Version) uint32

// User Memory Reply.
func ReplyUserMem(conn_handle uint16, block *UserMemBlock) uint32

// Set a BLE option.
func SetOpt(opt_id uint32, opt *Opt) uint32

// Get a BLE option.
func GetOpt(opt_id uint32, opt *Opt) uint32
