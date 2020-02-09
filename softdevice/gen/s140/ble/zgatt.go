// Code generated from ../nordic/s140/headers/ble_gatt.h; DO NOT EDIT.

package ble

const GATT_ATT_MTU_DEFAULT = 23
const GATT_HANDLE_INVALID = 0x0000
const GATT_HANDLE_START = 0x0001
const GATT_HANDLE_END = 0xFFFF
const GATT_TIMEOUT_SRC_PROTOCOL = 0x00                   // ATT Protocol timeout.
const GATT_OP_INVALID = 0x00                             // Invalid Operation.
const GATT_OP_WRITE_REQ = 0x01                           // Write Request.
const GATT_OP_WRITE_CMD = 0x02                           // Write Command.
const GATT_OP_SIGN_WRITE_CMD = 0x03                      // Signed Write Command.
const GATT_OP_PREP_WRITE_REQ = 0x04                      // Prepare Write Request.
const GATT_OP_EXEC_WRITE_REQ = 0x05                      // Execute Write Request.
const GATT_EXEC_WRITE_FLAG_PREPARED_CANCEL = 0x00        // Cancel prepared write.
const GATT_EXEC_WRITE_FLAG_PREPARED_WRITE = 0x01         // Execute prepared write.
const GATT_HVX_INVALID = 0x00                            // Invalid Operation.
const GATT_HVX_NOTIFICATION = 0x01                       // Handle Value Notification.
const GATT_HVX_INDICATION = 0x02                         // Handle Value Indication.
const GATT_STATUS_SUCCESS = 0x0000                       // Success.
const GATT_STATUS_UNKNOWN = 0x0001                       // Unknown or not applicable status.
const GATT_STATUS_ATTERR_INVALID = 0x0100                // ATT Error: Invalid Error Code.
const GATT_STATUS_ATTERR_INVALID_HANDLE = 0x0101         // ATT Error: Invalid Attribute Handle.
const GATT_STATUS_ATTERR_READ_NOT_PERMITTED = 0x0102     // ATT Error: Read not permitted.
const GATT_STATUS_ATTERR_WRITE_NOT_PERMITTED = 0x0103    // ATT Error: Write not permitted.
const GATT_STATUS_ATTERR_INVALID_PDU = 0x0104            // ATT Error: Used in ATT as Invalid PDU.
const GATT_STATUS_ATTERR_INSUF_AUTHENTICATION = 0x0105   // ATT Error: Authenticated link required.
const GATT_STATUS_ATTERR_REQUEST_NOT_SUPPORTED = 0x0106  // ATT Error: Used in ATT as Request Not Supported.
const GATT_STATUS_ATTERR_INVALID_OFFSET = 0x0107         // ATT Error: Offset specified was past the end of the attribute.
const GATT_STATUS_ATTERR_INSUF_AUTHORIZATION = 0x0108    // ATT Error: Used in ATT as Insufficient Authorization.
const GATT_STATUS_ATTERR_PREPARE_QUEUE_FULL = 0x0109     // ATT Error: Used in ATT as Prepare Queue Full.
const GATT_STATUS_ATTERR_ATTRIBUTE_NOT_FOUND = 0x010A    // ATT Error: Used in ATT as Attribute not found.
const GATT_STATUS_ATTERR_ATTRIBUTE_NOT_LONG = 0x010B     // ATT Error: Attribute cannot be read or written using read/write blob requests.
const GATT_STATUS_ATTERR_INSUF_ENC_KEY_SIZE = 0x010C     // ATT Error: Encryption key size used is insufficient.
const GATT_STATUS_ATTERR_INVALID_ATT_VAL_LENGTH = 0x010D // ATT Error: Invalid value size.
const GATT_STATUS_ATTERR_UNLIKELY_ERROR = 0x010E         // ATT Error: Very unlikely error.
const GATT_STATUS_ATTERR_INSUF_ENCRYPTION = 0x010F       // ATT Error: Encrypted link required.
const GATT_STATUS_ATTERR_UNSUPPORTED_GROUP_TYPE = 0x0110 // ATT Error: Attribute type is not a supported grouping attribute.
const GATT_STATUS_ATTERR_INSUF_RESOURCES = 0x0111        // ATT Error: Insufficient resources.
const GATT_STATUS_ATTERR_RFU_RANGE1_BEGIN = 0x0112       // ATT Error: Reserved for Future Use range #1 begin.
const GATT_STATUS_ATTERR_RFU_RANGE1_END = 0x017F         // ATT Error: Reserved for Future Use range #1 end.
const GATT_STATUS_ATTERR_APP_BEGIN = 0x0180              // ATT Error: Application range begin.
const GATT_STATUS_ATTERR_APP_END = 0x019F                // ATT Error: Application range end.
const GATT_STATUS_ATTERR_RFU_RANGE2_BEGIN = 0x01A0       // ATT Error: Reserved for Future Use range #2 begin.
const GATT_STATUS_ATTERR_RFU_RANGE2_END = 0x01DF         // ATT Error: Reserved for Future Use range #2 end.
const GATT_STATUS_ATTERR_RFU_RANGE3_BEGIN = 0x01E0       // ATT Error: Reserved for Future Use range #3 begin.
const GATT_STATUS_ATTERR_RFU_RANGE3_END = 0x01FC         // ATT Error: Reserved for Future Use range #3 end.
const GATT_STATUS_ATTERR_CPS_WRITE_REQ_REJECTED = 0x01FC // ATT Common Profile and Service Error: Write request rejected.
const GATT_STATUS_ATTERR_CPS_CCCD_CONFIG_ERROR = 0x01FD  // ATT Common Profile and Service Error: Client Characteristic Configuration Descriptor improperly configured.
const GATT_STATUS_ATTERR_CPS_PROC_ALR_IN_PROG = 0x01FE   // ATT Common Profile and Service Error: Procedure Already in Progress.
const GATT_STATUS_ATTERR_CPS_OUT_OF_RANGE = 0x01FF       // ATT Common Profile and Service Error: Out Of Range.
const GATT_CPF_FORMAT_RFU = 0x00                         // Reserved For Future Use.
const GATT_CPF_FORMAT_BOOLEAN = 0x01                     // Boolean.
const GATT_CPF_FORMAT_2BIT = 0x02                        // Unsigned 2-bit integer.
const GATT_CPF_FORMAT_NIBBLE = 0x03                      // Unsigned 4-bit integer.
const GATT_CPF_FORMAT_UINT8 = 0x04                       // Unsigned 8-bit integer.
const GATT_CPF_FORMAT_UINT12 = 0x05                      // Unsigned 12-bit integer.
const GATT_CPF_FORMAT_UINT16 = 0x06                      // Unsigned 16-bit integer.
const GATT_CPF_FORMAT_UINT24 = 0x07                      // Unsigned 24-bit integer.
const GATT_CPF_FORMAT_UINT32 = 0x08                      // Unsigned 32-bit integer.
const GATT_CPF_FORMAT_UINT48 = 0x09                      // Unsigned 48-bit integer.
const GATT_CPF_FORMAT_UINT64 = 0x0A                      // Unsigned 64-bit integer.
const GATT_CPF_FORMAT_UINT128 = 0x0B                     // Unsigned 128-bit integer.
const GATT_CPF_FORMAT_SINT8 = 0x0C                       // Signed 2-bit integer.
const GATT_CPF_FORMAT_SINT12 = 0x0D                      // Signed 12-bit integer.
const GATT_CPF_FORMAT_SINT16 = 0x0E                      // Signed 16-bit integer.
const GATT_CPF_FORMAT_SINT24 = 0x0F                      // Signed 24-bit integer.
const GATT_CPF_FORMAT_SINT32 = 0x10                      // Signed 32-bit integer.
const GATT_CPF_FORMAT_SINT48 = 0x11                      // Signed 48-bit integer.
const GATT_CPF_FORMAT_SINT64 = 0x12                      // Signed 64-bit integer.
const GATT_CPF_FORMAT_SINT128 = 0x13                     // Signed 128-bit integer.
const GATT_CPF_FORMAT_FLOAT32 = 0x14                     // IEEE-754 32-bit floating point.
const GATT_CPF_FORMAT_FLOAT64 = 0x15                     // IEEE-754 64-bit floating point.
const GATT_CPF_FORMAT_SFLOAT = 0x16                      // IEEE-11073 16-bit SFLOAT.
const GATT_CPF_FORMAT_FLOAT = 0x17                       // IEEE-11073 32-bit FLOAT.
const GATT_CPF_FORMAT_DUINT16 = 0x18                     // IEEE-20601 format.
const GATT_CPF_FORMAT_UTF8S = 0x19                       // UTF-8 string.
const GATT_CPF_FORMAT_UTF16S = 0x1A                      // UTF-16 string.
const GATT_CPF_FORMAT_STRUCT = 0x1B                      // Opaque Structure.
const GATT_CPF_NAMESPACE_BTSIG = 0x01                    // Bluetooth SIG defined Namespace.
const GATT_CPF_NAMESPACE_DESCRIPTION_UNKNOWN = 0x0000    // Namespace Description Unknown.

type GATTConnCfg struct {
	AttMtu uint16 // Maximum size of ATT packet the SoftDevice can send or receive. The default and minimum value is @ref BLE_GATT_ATT_MTU_DEFAULT. @mscs @mmsc{@ref BLE_GATTC_MTU_EXCHANGE} @mmsc{@ref BLE_GATTS_MTU_EXCHANGE} @endmscs
}

//type GATTCharProps struct { bitfield }

//type GATTCharExtProps struct { bitfield }
