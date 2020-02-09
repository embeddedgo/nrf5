// Code generated from ../nordic/s140/headers/ble_gatts.h; DO NOT EDIT.

package ble

const (
	_GATTS_SERVICE_ADD = GATTS_SVC_BASE + iota
	_GATTS_INCLUDE_ADD
	_GATTS_CHARACTERISTIC_ADD
	_GATTS_DESCRIPTOR_ADD
	_GATTS_VALUE_SET
	_GATTS_VALUE_GET
	_GATTS_HVX
	_GATTS_SERVICE_CHANGED
	_GATTS_RW_AUTHORIZE_REPLY
	_GATTS_SYS_ATTR_SET
	_GATTS_SYS_ATTR_GET
	_GATTS_INITIAL_USER_HANDLE_GET
	_GATTS_ATTR_GET
	_GATTS_EXCHANGE_MTU_REPLY
)

const (
	GATTS_EVT_WRITE                = GATTS_EVT_BASE + iota // Write operation performed. \n See @ref ble_gatts_evt_write_t.
	GATTS_EVT_RW_AUTHORIZE_REQUEST                         // Read/Write Authorization request. \n Reply with @ref sd_ble_gatts_rw_authorize_reply. \n See @ref ble_gatts_evt_rw_authorize_request_t.
	GATTS_EVT_SYS_ATTR_MISSING                             // A persistent system attribute access is pending. \n Respond with @ref sd_ble_gatts_sys_attr_set. \n See @ref ble_gatts_evt_sys_attr_missing_t.
	GATTS_EVT_HVC                                          // Handle Value Confirmation. \n See @ref ble_gatts_evt_hvc_t.
	GATTS_EVT_SC_CONFIRM                                   // Service Changed Confirmation. \n No additional event structure applies.
	GATTS_EVT_EXCHANGE_MTU_REQUEST                         // Exchange MTU Request. \n Reply with @ref sd_ble_gatts_exchange_mtu_reply. \n See @ref ble_gatts_evt_exchange_mtu_request_t.
	GATTS_EVT_TIMEOUT                                      // Peer failed to respond to an ATT request in time. \n See @ref ble_gatts_evt_timeout_t.
	GATTS_EVT_HVN_TX_COMPLETE                              // Handle Value Notification transmission complete. \n See @ref ble_gatts_evt_hvn_tx_complete_t.
)

const (
	GATTS_CFG_SERVICE_CHANGED = GATTS_CFG_BASE + iota // Service changed configuration.
	GATTS_CFG_ATTR_TAB_SIZE                           // Attribute table size configuration.
)

const ERROR_GATTS_INVALID_ATTR_TYPE = (GATTS_ERR_BASE + 0x000) // Invalid attribute type.
const ERROR_GATTS_SYS_ATTR_MISSING = (GATTS_ERR_BASE + 0x001)  // System Attributes missing.
const GATTS_FIX_ATTR_LEN_MAX = (510)                           // Maximum length for fixed length Attribute Values.
const GATTS_VAR_ATTR_LEN_MAX = (512)                           // Maximum length for variable length Attribute Values.
const GATTS_SRVC_TYPE_INVALID = 0x00                           // Invalid Service Type.
const GATTS_SRVC_TYPE_PRIMARY = 0x01                           // Primary Service.
const GATTS_SRVC_TYPE_SECONDARY = 0x02                         // Secondary Type.
const GATTS_ATTR_TYPE_INVALID = 0x00                           // Invalid Attribute Type.
const GATTS_ATTR_TYPE_PRIM_SRVC_DECL = 0x01                    // Primary Service Declaration.
const GATTS_ATTR_TYPE_SEC_SRVC_DECL = 0x02                     // Secondary Service Declaration.
const GATTS_ATTR_TYPE_INC_DECL = 0x03                          // Include Declaration.
const GATTS_ATTR_TYPE_CHAR_DECL = 0x04                         // Characteristic Declaration.
const GATTS_ATTR_TYPE_CHAR_VAL = 0x05                          // Characteristic Value.
const GATTS_ATTR_TYPE_DESC = 0x06                              // Descriptor.
const GATTS_ATTR_TYPE_OTHER = 0x07                             // Other, non-GATT specific type.
const GATTS_OP_INVALID = 0x00                                  // Invalid Operation.
const GATTS_OP_WRITE_REQ = 0x01                                // Write Request.
const GATTS_OP_WRITE_CMD = 0x02                                // Write Command.
const GATTS_OP_SIGN_WRITE_CMD = 0x03                           // Signed Write Command.
const GATTS_OP_PREP_WRITE_REQ = 0x04                           // Prepare Write Request.
const GATTS_OP_EXEC_WRITE_REQ_CANCEL = 0x05                    // Execute Write Request: Cancel all prepared writes.
const GATTS_OP_EXEC_WRITE_REQ_NOW = 0x06                       // Execute Write Request: Immediately execute all prepared writes.
const GATTS_VLOC_INVALID = 0x00                                // Invalid Location.
const GATTS_VLOC_STACK = 0x01                                  // Attribute Value is located in stack memory, no user memory is required.
const GATTS_VLOC_USER = 0x02                                   // Attribute Value is located in user memory. This requires the user to maintain a valid buffer through the lifetime of the attribute, since the stack will read and write directly to the memory using the pointer provided in the APIs. There are no alignment requirements for the buffer.
const GATTS_AUTHORIZE_TYPE_INVALID = 0x00                      // Invalid Type.
const GATTS_AUTHORIZE_TYPE_READ = 0x01                         // Authorize a Read Operation.
const GATTS_AUTHORIZE_TYPE_WRITE = 0x02                        // Authorize a Write Request Operation.
const GATTS_SYS_ATTR_FLAG_SYS_SRVCS = (1 << 0)                 // Restrict system attributes to system services only.
const GATTS_SYS_ATTR_FLAG_USR_SRVCS = (1 << 1)                 // Restrict system attributes to user services only.
const GATTS_SERVICE_CHANGED_DEFAULT = (1)                      // Default is to include the Service Changed characteristic in the Attribute Table.
const GATTS_ATTR_TAB_SIZE_MIN = (248)                          // Minimum Attribute Table size
const GATTS_ATTR_TAB_SIZE_DEFAULT = (1408)                     // Default Attribute Table size.
const GATTS_HVN_TX_QUEUE_SIZE_DEFAULT = 1                      // Default number of Handle Value Notifications that can be queued for transmission.

type GATTSConnCfg struct {
	HvnTxQueueSize uint8 // Minimum guaranteed number of Handle Value Notifications that can be queued for transmission. The default value is @ref BLE_GATTS_HVN_TX_QUEUE_SIZE_DEFAULT
}

//type GATTSAttrMd struct { bitfield }

type GATTSAttr struct {
	UUID     *UUID        // Pointer to the attribute UUID.
	AttrMd   *GATTSAttrMd // Pointer to the attribute metadata structure.
	InitLen  uint16       // Initial attribute value length in bytes.
	InitOffs uint16       // Initial attribute value offset in bytes. If different from zero, the first init_offs bytes of the attribute value will be left uninitialized.
	MaxLen   uint16       // Maximum attribute value length in bytes, see @ref BLE_GATTS_ATTR_LENS_MAX for maximum values.
	Value    *uint8       // Pointer to the attribute data. Please note that if the @ref BLE_GATTS_VLOC_USER value location is selected in the attribute metadata, this will have to point to a buffer that remains valid through the lifetime of the attribute. This excludes usage of automatic variables that may go out of scope or any other temporary location. The stack may access that memory directly without the application's knowledge. For writable characteristics, this value must not be a location in flash memory.
}

type GATTSValue struct {
	Len    uint16 // Length in bytes to be written or read. Length in bytes written or read after successful return.
	Offset uint16 // Attribute value offset.
	Value  *uint8 // Pointer to where value is stored or will be stored. If value is stored in user memory, only the attribute length is updated when p_value == NULL. Set to NULL when reading to obtain the complete length of the attribute value
}

type GATTSCharPf struct {
	Format    uint8  // Format of the value, see @ref BLE_GATT_CPF_FORMATS.
	Exponent  int8   // Exponent for integer data types.
	Unit      uint16 // Unit from Bluetooth Assigned Numbers.
	NameSpace uint8  // Namespace from Bluetooth Assigned Numbers, see @ref BLE_GATT_CPF_NAMESPACES.
	Desc      uint16 // Namespace description from Bluetooth Assigned Numbers, see @ref BLE_GATT_CPF_NAMESPACES.
}

type GATTSCharMd struct {
	CharProps           GATTCharProps    // Characteristic Properties.
	CharExtProps        GATTCharExtProps // Characteristic Extended Properties.
	CharUserDesc        *uint8           // Pointer to a UTF-8 encoded string (non-NULL terminated), NULL if the descriptor is not required.
	CharUserDescMaxSize uint16           // The maximum size in bytes of the user description descriptor.
	CharUserDescSize    uint16           // The size of the user description, must be smaller or equal to char_user_desc_max_size.
	CharPf              *GATTSCharPf     // Pointer to a presentation format structure or NULL if the CPF descriptor is not required.
	UserDescMd          *GATTSAttrMd     // Attribute metadata for the User Description descriptor, or NULL for default values.
	CccdMd              *GATTSAttrMd     // Attribute metadata for the Client Characteristic Configuration Descriptor, or NULL for default values.
	SccdMd              *GATTSAttrMd     // Attribute metadata for the Server Characteristic Configuration Descriptor, or NULL for default values.
}

type GATTSCharHandles struct {
	ValueHandle    uint16 // Handle to the characteristic value.
	UserDescHandle uint16 // Handle to the User Description descriptor, or @ref BLE_GATT_HANDLE_INVALID if not present.
	CccdHandle     uint16 // Handle to the Client Characteristic Configuration Descriptor, or @ref BLE_GATT_HANDLE_INVALID if not present.
	SccdHandle     uint16 // Handle to the Server Characteristic Configuration Descriptor, or @ref BLE_GATT_HANDLE_INVALID if not present.
}

type GATTSHvxParams struct {
	Handle uint16  // Characteristic Value Handle.
	Typ    uint8   // Indication or Notification, see @ref BLE_GATT_HVX_TYPES.
	Offset uint16  // Offset within the attribute value.
	Len    *uint16 // Length in bytes to be written, length in bytes written after return.
	Data   *uint8  // Actual data content, use NULL to use the current attribute value.
}

//type GATTSAuthorizeParams struct { bitfield }

//type GATTSRwAuthorizeReplyParams struct { union }

//type GATTSCfgServiceChanged struct { bitfield }

type GATTSCfgAttrTabSize struct {
	AttrTabSize uint32 // Attribute table size. Default is @ref BLE_GATTS_ATTR_TAB_SIZE_DEFAULT, minimum is @ref BLE_GATTS_ATTR_TAB_SIZE_MIN.
}

//type GATTSCfg union

type GATTSEvtWrite struct {
	Handle       uint16   // Attribute Handle.
	UUID         UUID     // Attribute UUID.
	Op           uint8    // Type of write operation, see @ref BLE_GATTS_OPS.
	AuthRequired uint8    // Writing operation deferred due to authorization requirement. Application may use @ref sd_ble_gatts_value_set to finalize the writing operation.
	Offset       uint16   // Offset for the write operation.
	Len          uint16   // Length of the received data.
	Data         [1]uint8 // Received data. @note This is a variable length array. The size of 1 indicated is only a placeholder for compilation. See @ref sd_ble_evt_get for more information on how to use event structures with variable length array members.
}

type GATTSEvtRead struct {
	Handle uint16 // Attribute Handle.
	UUID   UUID   // Attribute UUID.
	Offset uint16 // Offset for the read operation.
}

//type GATTSEvtRwAuthorizeRequest struct { union }

type GATTSEvtSysAttrMissing struct {
	Hint uint8 // Hint (currently unused).
}

type GATTSEvtHvc struct {
	Handle uint16 // Attribute Handle.
}

type GATTSEvtExchangeMtuRequest struct {
	ClientRxMtu uint16 // Client RX MTU size.
}

type GATTSEvtTimeout struct {
	Src uint8 // Timeout source, see @ref BLE_GATT_TIMEOUT_SOURCES.
}

type GATTSEvtHvnTxComplete struct {
	Count uint8 // Number of notification transmissions completed.
}

//type GATTSEvt struct { union }

// Add a service.
func AddGATTSService(typ uint8, uuid *UUID, handle *uint16) uint32

// Add an included service.
func AddGATTSInclude(service_handle uint16, inc_srvc_handle uint16, include_handle *uint16) uint32

// Add a characteristic.
func AddGATTSCharacteristic(service_handle uint16, char_md *GATTSCharMd, attr_char_value *GATTSAttr, handles *GATTSCharHandles) uint32

// Add a generic attribute.
func AddGATTSDescriptor(char_handle uint16, attr *GATTSAttr, handle *uint16) uint32

// Set an attribute value.
func SetGATTSValue(conn_handle uint16, handle uint16, value *GATTSValue) uint32

// Get an attribute value.
func GetGATTSValue(conn_handle uint16, handle uint16, value *GATTSValue) uint32

// Handle Value Notification or Indication.
func HvxGATTS(conn_handle uint16, hvx_params *GATTSHvxParams) uint32

// Perform a Service Changed Indication to one or more peers.
func ChangedGATTSService(conn_handle uint16, start_handle uint16, end_handle uint16) uint32

// Reply to an authorization request for a read or write operation on one or more attributes.
func ReplyGATTSRwAuthorize(conn_handle uint16, rw_authorize_reply_params *GATTSRwAuthorizeReplyParams) uint32

// Set the persistent system attributes for a connection.
func SetGATTSSysAttr(conn_handle uint16, sys_attr_data *uint8, len uint16, flags uint32) uint32

// Retrieve the persistent system attributes.
func GetGATTSSysAttr(conn_handle uint16, sys_attr_data *uint8, len *uint16, flags uint32) uint32

// Retrieve the first valid user handle.
func GetGATTSInitialUserHandle(handle *uint16) uint32

// Retrieve the UUID and/or metadata of an attribute.
func GetGATTSAttr(handle uint16, uuid *UUID, md *GATTSAttrMd) uint32

// Reply to Exchange MTU Request.
func ReplyGATTSExchangeMtu(conn_handle uint16, server_rx_mtu uint16) uint32
