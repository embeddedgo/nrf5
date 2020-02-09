// Code generated from ../nordic/s140/headers/ble_gattc.h; DO NOT EDIT.

package ble

const (
	_GATTC_PRIMARY_SERVICES_DISCOVER = GATTC_SVC_BASE + iota
	_GATTC_RELATIONSHIPS_DISCOVER
	_GATTC_CHARACTERISTICS_DISCOVER
	_GATTC_DESCRIPTORS_DISCOVER
	_GATTC_ATTR_INFO_DISCOVER
	_GATTC_CHAR_VALUE_BY_UUID_READ
	_GATTC_READ
	_GATTC_CHAR_VALUES_READ
	_GATTC_WRITE
	_GATTC_HV_CONFIRM
	_GATTC_EXCHANGE_MTU_REQUEST
)

const (
	GATTC_EVT_PRIM_SRVC_DISC_RSP        = GATTC_EVT_BASE + iota // Primary Service Discovery Response event. \n See @ref ble_gattc_evt_prim_srvc_disc_rsp_t.
	GATTC_EVT_REL_DISC_RSP                                      // Relationship Discovery Response event. \n See @ref ble_gattc_evt_rel_disc_rsp_t.
	GATTC_EVT_CHAR_DISC_RSP                                     // Characteristic Discovery Response event. \n See @ref ble_gattc_evt_char_disc_rsp_t.
	GATTC_EVT_DESC_DISC_RSP                                     // Descriptor Discovery Response event. \n See @ref ble_gattc_evt_desc_disc_rsp_t.
	GATTC_EVT_ATTR_INFO_DISC_RSP                                // Attribute Information Response event. \n See @ref ble_gattc_evt_attr_info_disc_rsp_t.
	GATTC_EVT_CHAR_VAL_BY_UUID_READ_RSP                         // Read By UUID Response event. \n See @ref ble_gattc_evt_char_val_by_uuid_read_rsp_t.
	GATTC_EVT_READ_RSP                                          // Read Response event. \n See @ref ble_gattc_evt_read_rsp_t.
	GATTC_EVT_CHAR_VALS_READ_RSP                                // Read multiple Response event. \n See @ref ble_gattc_evt_char_vals_read_rsp_t.
	GATTC_EVT_WRITE_RSP                                         // Write Response event. \n See @ref ble_gattc_evt_write_rsp_t.
	GATTC_EVT_HVX                                               // Handle Value Notification or Indication event. \n Confirm indication with @ref sd_ble_gattc_hv_confirm. \n See @ref ble_gattc_evt_hvx_t.
	GATTC_EVT_EXCHANGE_MTU_RSP                                  // Exchange MTU Response event. \n See @ref ble_gattc_evt_exchange_mtu_rsp_t.
	GATTC_EVT_TIMEOUT                                           // Timeout event. \n See @ref ble_gattc_evt_timeout_t.
	GATTC_EVT_WRITE_CMD_TX_COMPLETE                             // Write without Response transmission complete. \n See @ref ble_gattc_evt_write_cmd_tx_complete_t.
)

const ERROR_GATTC_PROC_NOT_PERMITTED = (GATTC_ERR_BASE + 0x000) // Procedure not Permitted.
const GATTC_ATTR_INFO_FORMAT_16BIT = 1                          // 16-bit Attribute Information Format.
const GATTC_ATTR_INFO_FORMAT_128BIT = 2                         // 128-bit Attribute Information Format.
const GATTC_WRITE_CMD_TX_QUEUE_SIZE_DEFAULT = 1                 // Default number of Write without Response that can be queued for transmission.

type GATTCConnCfg struct {
	WriteCmdTxQueueSize uint8 // The guaranteed minimum number of Write without Response that can be queued for transmission. The default value is @ref BLE_GATTC_WRITE_CMD_TX_QUEUE_SIZE_DEFAULT
}

type GATTCHandleRange struct {
	StartHandle uint16 // Start Handle.
	EndHandle   uint16 // End Handle.
}

type GATTCService struct {
	UUID        UUID             // Service UUID.
	HandleRange GATTCHandleRange // Service Handle Range.
}

type GATTCInclude struct {
	Handle       uint16       // Include Handle.
	IncludedSrvc GATTCService // Handle of the included service.
}

//type GATTCChar struct { bitfield }

type GATTCDesc struct {
	Handle uint16 // Descriptor Handle.
	UUID   UUID   // Descriptor UUID.
}

type GATTCWriteParams struct {
	WriteOp uint8  // Write Operation to be performed, see @ref BLE_GATT_WRITE_OPS.
	Flags   uint8  // Flags, see @ref BLE_GATT_EXEC_WRITE_FLAGS.
	Handle  uint16 // Handle to the attribute to be written.
	Offset  uint16 // Offset in bytes. @note For WRITE_CMD and WRITE_REQ, offset must be 0.
	Len     uint16 // Length of data in bytes.
	Value   *uint8 // Pointer to the value data.
}

type GATTCAttrInfo16 struct {
	Handle uint16 // Attribute handle.
	UUID   UUID   // 16-bit Attribute UUID.
}

type GATTCAttrInfo128 struct {
	Handle uint16  // Attribute handle.
	UUID   UUID128 // 128-bit Attribute UUID.
}

type GATTCEvtPrimSrvcDiscRsp struct {
	Count    uint16          // Service count.
	Services [1]GATTCService // Service data. @note This is a variable length array. The size of 1 indicated is only a placeholder for compilation. See @ref sd_ble_evt_get for more information on how to use event structures with variable length array members.
}

type GATTCEvtRelDiscRsp struct {
	Count    uint16          // Include count.
	Includes [1]GATTCInclude // Include data. @note This is a variable length array. The size of 1 indicated is only a placeholder for compilation. See @ref sd_ble_evt_get for more information on how to use event structures with variable length array members.
}

type GATTCEvtCharDiscRsp struct {
	Count uint16       // Characteristic count.
	Chars [1]GATTCChar // Characteristic data. @note This is a variable length array. The size of 1 indicated is only a placeholder for compilation. See @ref sd_ble_evt_get for more information on how to use event structures with variable length array members.
}

type GATTCEvtDescDiscRsp struct {
	Count uint16       // Descriptor count.
	Descs [1]GATTCDesc // Descriptor data. @note This is a variable length array. The size of 1 indicated is only a placeholder for compilation. See @ref sd_ble_evt_get for more information on how to use event structures with variable length array members.
}

//type GATTCEvtAttrInfoDiscRsp struct { union }

type GATTCHandleValue struct {
	Handle uint16 // Attribute Handle.
	Value  *uint8 // Pointer to the Attribute Value, length is available in @ref ble_gattc_evt_char_val_by_uuid_read_rsp_t::value_len.
}

type GATTCEvtCharValByUUIDReadRsp struct {
	Count       uint16   // Handle-Value Pair Count.
	ValueLen    uint16   // Length of the value in Handle-Value(s) list.
	HandleValue [1]uint8 // Handle-Value(s) list. To iterate through the list use @ref sd_ble_gattc_evt_char_val_by_uuid_read_rsp_iter. @note This is a variable length array. The size of 1 indicated is only a placeholder for compilation. See @ref sd_ble_evt_get for more information on how to use event structures with variable length array members.
}

type GATTCEvtReadRsp struct {
	Handle uint16   // Attribute Handle.
	Offset uint16   // Offset of the attribute data.
	Len    uint16   // Attribute data length.
	Data   [1]uint8 // Attribute data. @note This is a variable length array. The size of 1 indicated is only a placeholder for compilation. See @ref sd_ble_evt_get for more information on how to use event structures with variable length array members.
}

type GATTCEvtCharValsReadRsp struct {
	Len    uint16   // Concatenated Attribute values length.
	Values [1]uint8 // Attribute values. @note This is a variable length array. The size of 1 indicated is only a placeholder for compilation. See @ref sd_ble_evt_get for more information on how to use event structures with variable length array members.
}

type GATTCEvtWriteRsp struct {
	Handle  uint16   // Attribute Handle.
	WriteOp uint8    // Type of write operation, see @ref BLE_GATT_WRITE_OPS.
	Offset  uint16   // Data offset.
	Len     uint16   // Data length.
	Data    [1]uint8 // Data. @note This is a variable length array. The size of 1 indicated is only a placeholder for compilation. See @ref sd_ble_evt_get for more information on how to use event structures with variable length array members.
}

type GATTCEvtHvx struct {
	Handle uint16   // Handle to which the HVx operation applies.
	Typ    uint8    // Indication or Notification, see @ref BLE_GATT_HVX_TYPES.
	Len    uint16   // Attribute data length.
	Data   [1]uint8 // Attribute data. @note This is a variable length array. The size of 1 indicated is only a placeholder for compilation. See @ref sd_ble_evt_get for more information on how to use event structures with variable length array members.
}

type GATTCEvtExchangeMtuRsp struct {
	ServerRxMtu uint16 // Server RX MTU size.
}

type GATTCEvtTimeout struct {
	Src uint8 // Timeout source, see @ref BLE_GATT_TIMEOUT_SOURCES.
}

type GATTCEvtWriteCmdTxComplete struct {
	Count uint8 // Number of write without response transmissions completed.
}

//type GATTCEvt struct { union }

// Primary Service Discovery.
func DiscoverGATTCPrimaryServices(conn_handle uint16, start_handle uint16, srvc_uuid *UUID) uint32

// Relationship Discovery.
func DiscoverGATTCRelationships(conn_handle uint16, handle_range *GATTCHandleRange) uint32

// Characteristic Discovery.
func DiscoverGATTCCharacteristics(conn_handle uint16, handle_range *GATTCHandleRange) uint32

// Characteristic Descriptor Discovery.
func DiscoverGATTCDescriptors(conn_handle uint16, handle_range *GATTCHandleRange) uint32

// Read Characteristic Value by UUID.
func ReadGATTCCharValueByUUID(conn_handle uint16, uuid *UUID, handle_range *GATTCHandleRange) uint32

// Generic read.
func ReadGATTC(conn_handle uint16, handle uint16, offset uint16) uint32

// Read multiple Characteristic Values.
func ReadGATTCCharValues(conn_handle uint16, handles *uint16, handle_count uint16) uint32

// Generic write.
func WriteGATTC(conn_handle uint16, write_params *GATTCWriteParams) uint32

// Handle Value Confirmation.
func ConfirmGATTCHv(conn_handle uint16, handle uint16) uint32

// Attribute Information Discovery.
func DiscoverGATTCAttrInfo(conn_handle uint16, handle_range *GATTCHandleRange) uint32

// Exchange MTU Request.
func RequestGATTCExchangeMtu(conn_handle uint16, client_rx_mtu uint16) uint32
