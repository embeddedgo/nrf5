// Code generated from ../nordic/s140/headers/ble_gap.h; DO NOT EDIT.

package ble

const (
	_GAP_ADDR_SET                  = GAP_SVC_BASE + iota
	_GAP_ADDR_GET                  = GAP_SVC_BASE + 1
	_GAP_WHITELIST_SET             = GAP_SVC_BASE + 2
	_GAP_DEVICE_IDENTITIES_SET     = GAP_SVC_BASE + 3
	_GAP_PRIVACY_SET               = GAP_SVC_BASE + 4
	_GAP_PRIVACY_GET               = GAP_SVC_BASE + 5
	_GAP_ADV_SET_CONFIGURE         = GAP_SVC_BASE + 6
	_GAP_ADV_START                 = GAP_SVC_BASE + 7
	_GAP_ADV_STOP                  = GAP_SVC_BASE + 8
	_GAP_CONN_PARAM_UPDATE         = GAP_SVC_BASE + 9
	_GAP_DISCONNECT                = GAP_SVC_BASE + 10
	_GAP_TX_POWER_SET              = GAP_SVC_BASE + 11
	_GAP_APPEARANCE_SET            = GAP_SVC_BASE + 12
	_GAP_APPEARANCE_GET            = GAP_SVC_BASE + 13
	_GAP_PPCP_SET                  = GAP_SVC_BASE + 14
	_GAP_PPCP_GET                  = GAP_SVC_BASE + 15
	_GAP_DEVICE_NAME_SET           = GAP_SVC_BASE + 16
	_GAP_DEVICE_NAME_GET           = GAP_SVC_BASE + 17
	_GAP_AUTHENTICATE              = GAP_SVC_BASE + 18
	_GAP_SEC_PARAMS_REPLY          = GAP_SVC_BASE + 19
	_GAP_AUTH_KEY_REPLY            = GAP_SVC_BASE + 20
	_GAP_LESC_DHKEY_REPLY          = GAP_SVC_BASE + 21
	_GAP_KEYPRESS_NOTIFY           = GAP_SVC_BASE + 22
	_GAP_LESC_OOB_DATA_GET         = GAP_SVC_BASE + 23
	_GAP_LESC_OOB_DATA_SET         = GAP_SVC_BASE + 24
	_GAP_ENCRYPT                   = GAP_SVC_BASE + 25
	_GAP_SEC_INFO_REPLY            = GAP_SVC_BASE + 26
	_GAP_CONN_SEC_GET              = GAP_SVC_BASE + 27
	_GAP_RSSI_START                = GAP_SVC_BASE + 28
	_GAP_RSSI_STOP                 = GAP_SVC_BASE + 29
	_GAP_SCAN_START                = GAP_SVC_BASE + 30
	_GAP_SCAN_STOP                 = GAP_SVC_BASE + 31
	_GAP_CONNECT                   = GAP_SVC_BASE + 32
	_GAP_CONNECT_CANCEL            = GAP_SVC_BASE + 33
	_GAP_RSSI_GET                  = GAP_SVC_BASE + 34
	_GAP_PHY_UPDATE                = GAP_SVC_BASE + 35
	_GAP_DATA_LENGTH_UPDATE        = GAP_SVC_BASE + 36
	_GAP_QOS_CHANNEL_SURVEY_START  = GAP_SVC_BASE + 37
	_GAP_QOS_CHANNEL_SURVEY_STOP   = GAP_SVC_BASE + 38
	_GAP_ADV_ADDR_GET              = GAP_SVC_BASE + 39
	_GAP_NEXT_CONN_EVT_COUNTER_GET = GAP_SVC_BASE + 40
	_GAP_CONN_EVT_TRIGGER_START    = GAP_SVC_BASE + 41
	_GAP_CONN_EVT_TRIGGER_STOP     = GAP_SVC_BASE + 42
)

const (
	GAP_EVT_CONNECTED                  = GAP_EVT_BASE + iota // Connected to peer. \n See @ref ble_gap_evt_connected_t
	GAP_EVT_DISCONNECTED               = GAP_EVT_BASE + 1    // Disconnected from peer. \n See @ref ble_gap_evt_disconnected_t.
	GAP_EVT_CONN_PARAM_UPDATE          = GAP_EVT_BASE + 2    // Connection Parameters updated. \n See @ref ble_gap_evt_conn_param_update_t.
	GAP_EVT_SEC_PARAMS_REQUEST         = GAP_EVT_BASE + 3    // Request to provide security parameters. \n Reply with @ref sd_ble_gap_sec_params_reply. \n See @ref ble_gap_evt_sec_params_request_t.
	GAP_EVT_SEC_INFO_REQUEST           = GAP_EVT_BASE + 4    // Request to provide security information. \n Reply with @ref sd_ble_gap_sec_info_reply. \n See @ref ble_gap_evt_sec_info_request_t.
	GAP_EVT_PASSKEY_DISPLAY            = GAP_EVT_BASE + 5    // Request to display a passkey to the user. \n In LESC Numeric Comparison, reply with @ref sd_ble_gap_auth_key_reply. \n See @ref ble_gap_evt_passkey_display_t.
	GAP_EVT_KEY_PRESSED                = GAP_EVT_BASE + 6    // Notification of a keypress on the remote device.\n See @ref ble_gap_evt_key_pressed_t
	GAP_EVT_AUTH_KEY_REQUEST           = GAP_EVT_BASE + 7    // Request to provide an authentication key. \n Reply with @ref sd_ble_gap_auth_key_reply. \n See @ref ble_gap_evt_auth_key_request_t.
	GAP_EVT_LESC_DHKEY_REQUEST         = GAP_EVT_BASE + 8    // Request to calculate an LE Secure Connections DHKey. \n Reply with @ref sd_ble_gap_lesc_dhkey_reply. \n See @ref ble_gap_evt_lesc_dhkey_request_t
	GAP_EVT_AUTH_STATUS                = GAP_EVT_BASE + 9    // Authentication procedure completed with status. \n See @ref ble_gap_evt_auth_status_t.
	GAP_EVT_CONN_SEC_UPDATE            = GAP_EVT_BASE + 10   // Connection security updated. \n See @ref ble_gap_evt_conn_sec_update_t.
	GAP_EVT_TIMEOUT                    = GAP_EVT_BASE + 11   // Timeout expired. \n See @ref ble_gap_evt_timeout_t.
	GAP_EVT_RSSI_CHANGED               = GAP_EVT_BASE + 12   // RSSI report. \n See @ref ble_gap_evt_rssi_changed_t.
	GAP_EVT_ADV_REPORT                 = GAP_EVT_BASE + 13   // Advertising report. \n See @ref ble_gap_evt_adv_report_t.
	GAP_EVT_SEC_REQUEST                = GAP_EVT_BASE + 14   // Security Request. \n See @ref ble_gap_evt_sec_request_t.
	GAP_EVT_CONN_PARAM_UPDATE_REQUEST  = GAP_EVT_BASE + 15   // Connection Parameter Update Request. \n Reply with @ref sd_ble_gap_conn_param_update. \n See @ref ble_gap_evt_conn_param_update_request_t.
	GAP_EVT_SCAN_REQ_REPORT            = GAP_EVT_BASE + 16   // Scan request report. \n See @ref ble_gap_evt_scan_req_report_t.
	GAP_EVT_PHY_UPDATE_REQUEST         = GAP_EVT_BASE + 17   // PHY Update Request. \n Reply with @ref sd_ble_gap_phy_update. \n See @ref ble_gap_evt_phy_update_request_t.
	GAP_EVT_PHY_UPDATE                 = GAP_EVT_BASE + 18   // PHY Update Procedure is complete. \n See @ref ble_gap_evt_phy_update_t.
	GAP_EVT_DATA_LENGTH_UPDATE_REQUEST = GAP_EVT_BASE + 19   // Data Length Update Request. \n Reply with @ref sd_ble_gap_data_length_update. \n See @ref ble_gap_evt_data_length_update_request_t.
	GAP_EVT_DATA_LENGTH_UPDATE         = GAP_EVT_BASE + 20   // LL Data Channel PDU payload length updated. \n See @ref ble_gap_evt_data_length_update_t.
	GAP_EVT_QOS_CHANNEL_SURVEY_REPORT  = GAP_EVT_BASE + 21   // Channel survey report. \n See @ref ble_gap_evt_qos_channel_survey_report_t.
	GAP_EVT_ADV_SET_TERMINATED         = GAP_EVT_BASE + 22   // Advertising set terminated. \n See @ref ble_gap_evt_adv_set_terminated_t.
)

const (
	GAP_OPT_CH_MAP                = GAP_OPT_BASE + iota // Channel Map. @ref ble_gap_opt_ch_map_t
	GAP_OPT_LOCAL_CONN_LATENCY    = GAP_OPT_BASE + 1    // Local connection latency. @ref ble_gap_opt_local_conn_latency_t
	GAP_OPT_PASSKEY               = GAP_OPT_BASE + 2    // Set passkey. @ref ble_gap_opt_passkey_t
	GAP_OPT_COMPAT_MODE_1         = GAP_OPT_BASE + 3    // Compatibility mode. @ref ble_gap_opt_compat_mode_1_t
	GAP_OPT_AUTH_PAYLOAD_TIMEOUT  = GAP_OPT_BASE + 4    // Set Authenticated payload timeout. @ref ble_gap_opt_auth_payload_timeout_t
	GAP_OPT_SLAVE_LATENCY_DISABLE = GAP_OPT_BASE + 5    // Disable slave latency. @ref ble_gap_opt_slave_latency_disable_t
)

const (
	GAP_CFG_ROLE_COUNT       = GAP_CFG_BASE + iota // Role count configuration.
	GAP_CFG_DEVICE_NAME      = GAP_CFG_BASE + 1    // Device name configuration.
	GAP_CFG_PPCP_INCL_CONFIG = GAP_CFG_BASE + 2    // Peripheral Preferred Connection Parameters characteristic inclusion configuration.
	GAP_CFG_CAR_INCL_CONFIG  = GAP_CFG_BASE + 3    // Central Address Resolution characteristic inclusion configuration.
)

const (
	GAP_TX_POWER_ROLE_ADV       = 1 + iota // Advertiser role.
	GAP_TX_POWER_ROLE_SCAN_INIT = 2        // Scanner and initiator role.
	GAP_TX_POWER_ROLE_CONN      = 3        // Connection role.
)

const ERROR_GAP_UUID_LIST_MISMATCH = (GAP_ERR_BASE + 0x000)          // UUID list does not contain an integral number of UUIDs.
const ERROR_GAP_DISCOVERABLE_WITH_WHITELIST = (GAP_ERR_BASE + 0x001) // Use of Whitelist not permitted with discoverable advertising.
const ERROR_GAP_INVALID_BLE_ADDR = (GAP_ERR_BASE + 0x002)            // The upper two bits of the address do not correspond to the specified address type.
const ERROR_GAP_WHITELIST_IN_USE = (GAP_ERR_BASE + 0x003)            // Attempt to modify the whitelist while already in use by another operation.
const ERROR_GAP_DEVICE_IDENTITIES_IN_USE = (GAP_ERR_BASE + 0x004)    // Attempt to modify the device identity list while already in use by another operation.
const ERROR_GAP_DEVICE_IDENTITIES_DUPLICATE = (GAP_ERR_BASE + 0x005) // The device identity list contains entries with duplicate identity addresses.
const GAP_ROLE_INVALID = 0x0                                         // Invalid Role.
const GAP_ROLE_PERIPH = 0x1                                          // Peripheral Role.
const GAP_ROLE_CENTRAL = 0x2                                         // Central Role.
const GAP_TIMEOUT_SRC_SCAN = 0x01                                    // Scanning timeout.
const GAP_TIMEOUT_SRC_CONN = 0x02                                    // Connection timeout.
const GAP_TIMEOUT_SRC_AUTH_PAYLOAD = 0x03                            // Authenticated payload timeout.
const GAP_ADDR_TYPE_PUBLIC = 0x00                                    // Public (identity) address.
const GAP_ADDR_TYPE_RANDOM_STATIC = 0x01                             // Random static (identity) address.
const GAP_ADDR_TYPE_RANDOM_PRIVATE_RESOLVABLE = 0x02                 // Random private resolvable address.
const GAP_ADDR_TYPE_RANDOM_PRIVATE_NON_RESOLVABLE = 0x03             // Random private non-resolvable address.
const GAP_ADDR_TYPE_ANONYMOUS = 0x7F                                 // An advertiser may advertise without its address. This type of advertising is called anonymous.
const GAP_DEFAULT_PRIVATE_ADDR_CYCLE_INTERVAL_S = (900)              // 15 minutes.
const GAP_MAX_PRIVATE_ADDR_CYCLE_INTERVAL_S = (41400)                // 11 hours 30 minutes.
const GAP_ADDR_LEN = (6)
const GAP_PRIVACY_MODE_OFF = 0x00             // Device will send and accept its identity address for its own address.
const GAP_PRIVACY_MODE_DEVICE_PRIVACY = 0x01  // Device will send and accept only private addresses for its own address.
const GAP_PRIVACY_MODE_NETWORK_PRIVACY = 0x02 // Device will send and accept only private addresses for its own address, and will not accept a peer using identity address as sender address when the peer IRK is exchanged, non-zero and added to the identity list.
const GAP_POWER_LEVEL_INVALID = 127
const GAP_ADV_SET_HANDLE_NOT_SET = (0xFF)
const GAP_ADV_SET_COUNT_DEFAULT = (1)
const GAP_ADV_SET_COUNT_MAX = (1)
const GAP_ADV_SET_DATA_SIZE_MAX = (31)                                 // Maximum data length for an advertising set. If more advertising data is required, use extended advertising instead.
const GAP_ADV_SET_DATA_SIZE_EXTENDED_MAX_SUPPORTED = (255)             // Maximum supported data length for an extended advertising set.
const GAP_ADV_SET_DATA_SIZE_EXTENDED_CONNECTABLE_MAX_SUPPORTED = (238) // Maximum supported data length for an extended connectable advertising set.
const GAP_ADV_REPORT_SET_ID_NOT_AVAILABLE = 0xFF
const GAP_EVT_ADV_SET_TERMINATED_REASON_TIMEOUT = 0x01                                                                  // Timeout value reached.
const GAP_EVT_ADV_SET_TERMINATED_REASON_LIMIT_REACHED = 0x02                                                            // @ref ble_gap_adv_params_t::max_adv_evts was reached.
const GAP_AD_TYPE_FLAGS = 0x01                                                                                          // Flags for discoverability.
const GAP_AD_TYPE_16BIT_SERVICE_UUID_MORE_AVAILABLE = 0x02                                                              // Partial list of 16 bit service UUIDs.
const GAP_AD_TYPE_16BIT_SERVICE_UUID_COMPLETE = 0x03                                                                    // Complete list of 16 bit service UUIDs.
const GAP_AD_TYPE_32BIT_SERVICE_UUID_MORE_AVAILABLE = 0x04                                                              // Partial list of 32 bit service UUIDs.
const GAP_AD_TYPE_32BIT_SERVICE_UUID_COMPLETE = 0x05                                                                    // Complete list of 32 bit service UUIDs.
const GAP_AD_TYPE_128BIT_SERVICE_UUID_MORE_AVAILABLE = 0x06                                                             // Partial list of 128 bit service UUIDs.
const GAP_AD_TYPE_128BIT_SERVICE_UUID_COMPLETE = 0x07                                                                   // Complete list of 128 bit service UUIDs.
const GAP_AD_TYPE_SHORT_LOCAL_NAME = 0x08                                                                               // Short local device name.
const GAP_AD_TYPE_COMPLETE_LOCAL_NAME = 0x09                                                                            // Complete local device name.
const GAP_AD_TYPE_TX_POWER_LEVEL = 0x0A                                                                                 // Transmit power level.
const GAP_AD_TYPE_CLASS_OF_DEVICE = 0x0D                                                                                // Class of device.
const GAP_AD_TYPE_SIMPLE_PAIRING_HASH_C = 0x0E                                                                          // Simple Pairing Hash C.
const GAP_AD_TYPE_SIMPLE_PAIRING_RANDOMIZER_R = 0x0F                                                                    // Simple Pairing Randomizer R.
const GAP_AD_TYPE_SECURITY_MANAGER_TK_VALUE = 0x10                                                                      // Security Manager TK Value.
const GAP_AD_TYPE_SECURITY_MANAGER_OOB_FLAGS = 0x11                                                                     // Security Manager Out Of Band Flags.
const GAP_AD_TYPE_SLAVE_CONNECTION_INTERVAL_RANGE = 0x12                                                                // Slave Connection Interval Range.
const GAP_AD_TYPE_SOLICITED_SERVICE_UUIDS_16BIT = 0x14                                                                  // List of 16-bit Service Solicitation UUIDs.
const GAP_AD_TYPE_SOLICITED_SERVICE_UUIDS_128BIT = 0x15                                                                 // List of 128-bit Service Solicitation UUIDs.
const GAP_AD_TYPE_SERVICE_DATA = 0x16                                                                                   // Service Data - 16-bit UUID.
const GAP_AD_TYPE_PUBLIC_TARGET_ADDRESS = 0x17                                                                          // Public Target Address.
const GAP_AD_TYPE_RANDOM_TARGET_ADDRESS = 0x18                                                                          // Random Target Address.
const GAP_AD_TYPE_APPEARANCE = 0x19                                                                                     // Appearance.
const GAP_AD_TYPE_ADVERTISING_INTERVAL = 0x1A                                                                           // Advertising Interval.
const GAP_AD_TYPE_LE_BLUETOOTH_DEVICE_ADDRESS = 0x1B                                                                    // LE Bluetooth Device Address.
const GAP_AD_TYPE_LE_ROLE = 0x1C                                                                                        // LE Role.
const GAP_AD_TYPE_SIMPLE_PAIRING_HASH_C256 = 0x1D                                                                       // Simple Pairing Hash C-256.
const GAP_AD_TYPE_SIMPLE_PAIRING_RANDOMIZER_R256 = 0x1E                                                                 // Simple Pairing Randomizer R-256.
const GAP_AD_TYPE_SERVICE_DATA_32BIT_UUID = 0x20                                                                        // Service Data - 32-bit UUID.
const GAP_AD_TYPE_SERVICE_DATA_128BIT_UUID = 0x21                                                                       // Service Data - 128-bit UUID.
const GAP_AD_TYPE_LESC_CONFIRMATION_VALUE = 0x22                                                                        // LE Secure Connections Confirmation Value
const GAP_AD_TYPE_LESC_RANDOM_VALUE = 0x23                                                                              // LE Secure Connections Random Value
const GAP_AD_TYPE_URI = 0x24                                                                                            // URI
const GAP_AD_TYPE_3D_INFORMATION_DATA = 0x3D                                                                            // 3D Information Data.
const GAP_AD_TYPE_MANUFACTURER_SPECIFIC_DATA = 0xFF                                                                     // Manufacturer Specific Data.
const GAP_ADV_FLAG_LE_LIMITED_DISC_MODE = (0x01)                                                                        // LE Limited Discoverable Mode.
const GAP_ADV_FLAG_LE_GENERAL_DISC_MODE = (0x02)                                                                        // LE General Discoverable Mode.
const GAP_ADV_FLAG_BR_EDR_NOT_SUPPORTED = (0x04)                                                                        // BR/EDR not supported.
const GAP_ADV_FLAG_LE_BR_EDR_CONTROLLER = (0x08)                                                                        // Simultaneous LE and BR/EDR, Controller.
const GAP_ADV_FLAG_LE_BR_EDR_HOST = (0x10)                                                                              // Simultaneous LE and BR/EDR, Host.
const GAP_ADV_FLAGS_LE_ONLY_LIMITED_DISC_MODE = (GAP_ADV_FLAG_LE_LIMITED_DISC_MODE | GAP_ADV_FLAG_BR_EDR_NOT_SUPPORTED) // LE Limited Discoverable Mode, BR/EDR not supported.
const GAP_ADV_FLAGS_LE_ONLY_GENERAL_DISC_MODE = (GAP_ADV_FLAG_LE_GENERAL_DISC_MODE | GAP_ADV_FLAG_BR_EDR_NOT_SUPPORTED) // LE General Discoverable Mode, BR/EDR not supported.
const GAP_ADV_INTERVAL_MIN = 0x000020                                                                                   // Minimum Advertising interval in 625 us units, i.e. 20 ms.
const GAP_ADV_INTERVAL_MAX = 0x004000                                                                                   // Maximum Advertising interval in 625 us units, i.e. 10.24 s.
const GAP_SCAN_INTERVAL_MIN = 0x0004                                                                                    // Minimum Scan interval in 625 us units, i.e. 2.5 ms.
const GAP_SCAN_INTERVAL_MAX = 0xFFFF                                                                                    // Maximum Scan interval in 625 us units, i.e. 40,959.375 s.
const GAP_SCAN_WINDOW_MIN = 0x0004                                                                                      // Minimum Scan window in 625 us units, i.e. 2.5 ms.
const GAP_SCAN_WINDOW_MAX = 0xFFFF                                                                                      // Maximum Scan window in 625 us units, i.e. 40,959.375 s.
const GAP_SCAN_TIMEOUT_MIN = 0x0001                                                                                     // Minimum Scan timeout in 10 ms units, i.e 10 ms.
const GAP_SCAN_TIMEOUT_UNLIMITED = 0x0000                                                                               // Continue to scan forever.
const GAP_SCAN_BUFFER_MIN = (31)                                                                                        // Minimum data length for an advertising set.
const GAP_SCAN_BUFFER_MAX = (31)                                                                                        // Maximum data length for an advertising set.
const GAP_SCAN_BUFFER_EXTENDED_MIN = (255)                                                                              // Minimum data length for an extended advertising set.
const GAP_SCAN_BUFFER_EXTENDED_MAX = (1650)                                                                             // Maximum data length for an extended advertising set.
const GAP_SCAN_BUFFER_EXTENDED_MAX_SUPPORTED = (255)                                                                    // Maximum supported data length for an extended advertising set.
const GAP_ADV_TYPE_CONNECTABLE_SCANNABLE_UNDIRECTED = 0x01                                                              // Connectable and scannable undirected advertising events.
const GAP_ADV_TYPE_CONNECTABLE_NONSCANNABLE_DIRECTED_HIGH_DUTY_CYCLE = 0x02                                             // Connectable non-scannable directed advertising events. Advertising interval is less that 3.75 ms. Use this type for fast reconnections. @note Advertising data is not supported.
const GAP_ADV_TYPE_CONNECTABLE_NONSCANNABLE_DIRECTED = 0x03                                                             // Connectable non-scannable directed advertising events. @note Advertising data is not supported.
const GAP_ADV_TYPE_NONCONNECTABLE_SCANNABLE_UNDIRECTED = 0x04                                                           // Non-connectable scannable undirected advertising events.
const GAP_ADV_TYPE_NONCONNECTABLE_NONSCANNABLE_UNDIRECTED = 0x05                                                        // Non-connectable non-scannable undirected advertising events.
const GAP_ADV_TYPE_EXTENDED_CONNECTABLE_NONSCANNABLE_UNDIRECTED = 0x06                                                  // Connectable non-scannable undirected advertising events using extended advertising PDUs.
const GAP_ADV_TYPE_EXTENDED_CONNECTABLE_NONSCANNABLE_DIRECTED = 0x07                                                    // Connectable non-scannable directed advertising events using extended advertising PDUs.
const GAP_ADV_TYPE_EXTENDED_NONCONNECTABLE_SCANNABLE_UNDIRECTED = 0x08                                                  // Non-connectable scannable undirected advertising events using extended advertising PDUs. @note Only scan response data is supported.
const GAP_ADV_TYPE_EXTENDED_NONCONNECTABLE_SCANNABLE_DIRECTED = 0x09                                                    // Non-connectable scannable directed advertising events using extended advertising PDUs. @note Only scan response data is supported.
const GAP_ADV_TYPE_EXTENDED_NONCONNECTABLE_NONSCANNABLE_UNDIRECTED = 0x0A                                               // Non-connectable non-scannable undirected advertising events using extended advertising PDUs.
const GAP_ADV_TYPE_EXTENDED_NONCONNECTABLE_NONSCANNABLE_DIRECTED = 0x0B                                                 // Non-connectable non-scannable directed advertising events using extended advertising PDUs.
const GAP_ADV_FP_ANY = 0x00                                                                                             // Allow scan requests and connect requests from any device.
const GAP_ADV_FP_FILTER_SCANREQ = 0x01                                                                                  // Filter scan requests with whitelist.
const GAP_ADV_FP_FILTER_CONNREQ = 0x02                                                                                  // Filter connect requests with whitelist.
const GAP_ADV_FP_FILTER_BOTH = 0x03                                                                                     // Filter both scan and connect requests with whitelist.
const GAP_ADV_DATA_STATUS_COMPLETE = 0x00                                                                               // All data in the advertising event have been received.
const GAP_ADV_DATA_STATUS_INCOMPLETE_MORE_DATA = 0x01                                                                   // More data to be received. @note This value will only be used if @ref ble_gap_scan_params_t::report_incomplete_evts and @ref ble_gap_adv_report_type_t::extended_pdu are set to true.
const GAP_ADV_DATA_STATUS_INCOMPLETE_TRUNCATED = 0x02                                                                   // Incomplete data. Buffer size insufficient to receive more. @note This value will only be used if @ref ble_gap_adv_report_type_t::extended_pdu is set to true.
const GAP_ADV_DATA_STATUS_INCOMPLETE_MISSED = 0x03                                                                      // Failed to receive the remaining data. @note This value will only be used if @ref ble_gap_adv_report_type_t::extended_pdu is set to true.
const GAP_SCAN_FP_ACCEPT_ALL = 0x00                                                                                     // Accept all advertising packets except directed advertising packets not addressed to this device.
const GAP_SCAN_FP_WHITELIST = 0x01                                                                                      // Accept advertising packets from devices in the whitelist except directed packets not addressed to this device.
const GAP_SCAN_FP_ALL_NOT_RESOLVED_DIRECTED = 0x02                                                                      // Accept all advertising packets specified in @ref BLE_GAP_SCAN_FP_ACCEPT_ALL. In addition, accept directed advertising packets, where the advertiser's address is a resolvable private address that cannot be resolved.
const GAP_SCAN_FP_WHITELIST_NOT_RESOLVED_DIRECTED = 0x03                                                                // Accept all advertising packets specified in @ref BLE_GAP_SCAN_FP_WHITELIST. In addition, accept directed advertising packets, where the advertiser's address is a resolvable private address that cannot be resolved.
const GAP_ADV_TIMEOUT_HIGH_DUTY_MAX = (128)                                                                             // Maximum high duty advertising time in 10 ms units. Corresponds to 1.28 s.
const GAP_ADV_TIMEOUT_LIMITED_MAX = (18000)                                                                             // Maximum advertising time in 10 ms units corresponding to TGAP(lim_adv_timeout) = 180 s in limited discoverable mode.
const GAP_ADV_TIMEOUT_GENERAL_UNLIMITED = (0)                                                                           // Unlimited advertising in general discoverable mode. For high duty cycle advertising, this corresponds to @ref BLE_GAP_ADV_TIMEOUT_HIGH_DUTY_MAX.
const GAP_DISC_MODE_NOT_DISCOVERABLE = 0x00                                                                             // Not discoverable discovery Mode.
const GAP_DISC_MODE_LIMITED = 0x01                                                                                      // Limited Discovery Mode.
const GAP_DISC_MODE_GENERAL = 0x02                                                                                      // General Discovery Mode.
const GAP_IO_CAPS_DISPLAY_ONLY = 0x00                                                                                   // Display Only.
const GAP_IO_CAPS_DISPLAY_YESNO = 0x01                                                                                  // Display and Yes/No entry.
const GAP_IO_CAPS_KEYBOARD_ONLY = 0x02                                                                                  // Keyboard Only.
const GAP_IO_CAPS_NONE = 0x03                                                                                           // No I/O capabilities.
const GAP_IO_CAPS_KEYBOARD_DISPLAY = 0x04                                                                               // Keyboard and Display.
const GAP_AUTH_KEY_TYPE_NONE = 0x00                                                                                     // No key (may be used to reject).
const GAP_AUTH_KEY_TYPE_PASSKEY = 0x01                                                                                  // 6-digit Passkey.
const GAP_AUTH_KEY_TYPE_OOB = 0x02                                                                                      // Out Of Band data.
const GAP_KP_NOT_TYPE_PASSKEY_START = 0x00                                                                              // Passkey entry started.
const GAP_KP_NOT_TYPE_PASSKEY_DIGIT_IN = 0x01                                                                           // Passkey digit entered.
const GAP_KP_NOT_TYPE_PASSKEY_DIGIT_OUT = 0x02                                                                          // Passkey digit erased.
const GAP_KP_NOT_TYPE_PASSKEY_CLEAR = 0x03                                                                              // Passkey cleared.
const GAP_KP_NOT_TYPE_PASSKEY_END = 0x04                                                                                // Passkey entry completed.
const GAP_SEC_STATUS_SUCCESS = 0x00                                                                                     // Procedure completed with success.
const GAP_SEC_STATUS_TIMEOUT = 0x01                                                                                     // Procedure timed out.
const GAP_SEC_STATUS_PDU_INVALID = 0x02                                                                                 // Invalid PDU received.
const GAP_SEC_STATUS_RFU_RANGE1_BEGIN = 0x03                                                                            // Reserved for Future Use range #1 begin.
const GAP_SEC_STATUS_RFU_RANGE1_END = 0x80                                                                              // Reserved for Future Use range #1 end.
const GAP_SEC_STATUS_PASSKEY_ENTRY_FAILED = 0x81                                                                        // Passkey entry failed (user canceled or other).
const GAP_SEC_STATUS_OOB_NOT_AVAILABLE = 0x82                                                                           // Out of Band Key not available.
const GAP_SEC_STATUS_AUTH_REQ = 0x83                                                                                    // Authentication requirements not met.
const GAP_SEC_STATUS_CONFIRM_VALUE = 0x84                                                                               // Confirm value failed.
const GAP_SEC_STATUS_PAIRING_NOT_SUPP = 0x85                                                                            // Pairing not supported.
const GAP_SEC_STATUS_ENC_KEY_SIZE = 0x86                                                                                // Encryption key size.
const GAP_SEC_STATUS_SMP_CMD_UNSUPPORTED = 0x87                                                                         // Unsupported SMP command.
const GAP_SEC_STATUS_UNSPECIFIED = 0x88                                                                                 // Unspecified reason.
const GAP_SEC_STATUS_REPEATED_ATTEMPTS = 0x89                                                                           // Too little time elapsed since last attempt.
const GAP_SEC_STATUS_INVALID_PARAMS = 0x8A                                                                              // Invalid parameters.
const GAP_SEC_STATUS_DHKEY_FAILURE = 0x8B                                                                               // DHKey check failure.
const GAP_SEC_STATUS_NUM_COMP_FAILURE = 0x8C                                                                            // Numeric Comparison failure.
const GAP_SEC_STATUS_BR_EDR_IN_PROG = 0x8D                                                                              // BR/EDR pairing in progress.
const GAP_SEC_STATUS_X_TRANS_KEY_DISALLOWED = 0x8E                                                                      // BR/EDR Link Key cannot be used for LE keys.
const GAP_SEC_STATUS_RFU_RANGE2_BEGIN = 0x8F                                                                            // Reserved for Future Use range #2 begin.
const GAP_SEC_STATUS_RFU_RANGE2_END = 0xFF                                                                              // Reserved for Future Use range #2 end.
const GAP_SEC_STATUS_SOURCE_LOCAL = 0x00                                                                                // Local failure.
const GAP_SEC_STATUS_SOURCE_REMOTE = 0x01                                                                               // Remote failure.
const GAP_CP_MIN_CONN_INTVL_NONE = 0xFFFF                                                                               // No new minimum connection interval specified in connect parameters.
const GAP_CP_MIN_CONN_INTVL_MIN = 0x0006                                                                                // Lowest minimum connection interval permitted, in units of 1.25 ms, i.e. 7.5 ms.
const GAP_CP_MIN_CONN_INTVL_MAX = 0x0C80                                                                                // Highest minimum connection interval permitted, in units of 1.25 ms, i.e. 4 s.
const GAP_CP_MAX_CONN_INTVL_NONE = 0xFFFF                                                                               // No new maximum connection interval specified in connect parameters.
const GAP_CP_MAX_CONN_INTVL_MIN = 0x0006                                                                                // Lowest maximum connection interval permitted, in units of 1.25 ms, i.e. 7.5 ms.
const GAP_CP_MAX_CONN_INTVL_MAX = 0x0C80                                                                                // Highest maximum connection interval permitted, in units of 1.25 ms, i.e. 4 s.
const GAP_CP_SLAVE_LATENCY_MAX = 0x01F3                                                                                 // Highest slave latency permitted, in connection events.
const GAP_CP_CONN_SUP_TIMEOUT_NONE = 0xFFFF                                                                             // No new supervision timeout specified in connect parameters.
const GAP_CP_CONN_SUP_TIMEOUT_MIN = 0x000A                                                                              // Lowest supervision timeout permitted, in units of 10 ms, i.e. 100 ms.
const GAP_CP_CONN_SUP_TIMEOUT_MAX = 0x0C80                                                                              // Highest supervision timeout permitted, in units of 10 ms, i.e. 32 s.
const GAP_DEVNAME_DEFAULT = "nRF5x"                                                                                     // Default device name value.
const GAP_DEVNAME_DEFAULT_LEN = 31                                                                                      // Default number of octets in device name.
const GAP_DEVNAME_MAX_LEN = 248                                                                                         // Maximum number of octets in device name.
const GAP_RSSI_THRESHOLD_INVALID = 0xFF
const GAP_PHY_AUTO = 0x00                                                  // Automatic PHY selection. Refer @ref sd_ble_gap_phy_update for more information.
const GAP_PHY_1MBPS = 0x01                                                 // 1 Mbps PHY.
const GAP_PHY_2MBPS = 0x02                                                 // 2 Mbps PHY.
const GAP_PHY_CODED = 0x04                                                 // Coded PHY.
const GAP_PHY_NOT_SET = 0xFF                                               // PHY is not configured.
const GAP_PHYS_SUPPORTED = (GAP_PHY_1MBPS | GAP_PHY_2MBPS | GAP_PHY_CODED) // All PHYs are supported.
const GAP_SEC_RAND_LEN = 8
const GAP_SEC_KEY_LEN = 16
const GAP_LESC_P256_PK_LEN = 64
const GAP_LESC_DHKEY_LEN = 32
const GAP_PASSKEY_LEN = 6
const GAP_WHITELIST_ADDR_MAX_COUNT = (8)
const GAP_DEVICE_IDENTITIES_MAX_COUNT = (8)
const GAP_CONN_COUNT_DEFAULT = (1)
const GAP_EVENT_LENGTH_MIN = (2)               // Minimum event length, in 1.25 ms units.
const GAP_EVENT_LENGTH_CODED_PHY_MIN = (6)     // The shortest event length in 1.25 ms units supporting LE Coded PHY.
const GAP_EVENT_LENGTH_DEFAULT = (3)           // Default event length, in 1.25 ms units.
const GAP_ROLE_COUNT_PERIPH_DEFAULT = (1)      // Default maximum number of connections concurrently acting as peripherals.
const GAP_ROLE_COUNT_CENTRAL_DEFAULT = (3)     // Default maximum number of connections concurrently acting as centrals.
const GAP_ROLE_COUNT_CENTRAL_SEC_DEFAULT = (1) // Default number of SMP instances shared between all connections acting as centrals.
const GAP_ROLE_COUNT_COMBINED_MAX = (20)       // Maximum supported number of concurrent connections in the peripheral and central roles combined.
const GAP_DATA_LENGTH_AUTO = 0
const GAP_AUTH_PAYLOAD_TIMEOUT_MAX = (48000) // Maximum authenticated payload timeout in 10 ms units, i.e. 8 minutes.
const GAP_AUTH_PAYLOAD_TIMEOUT_MIN = (1)     // Minimum authenticated payload timeout in 10 ms units, i.e. 10 ms.
const GAP_SEC_MODE = 0x00                    // No key (may be used to reject).
const GAP_CHANNEL_COUNT = (40)
const GAP_QOS_CHANNEL_SURVEY_INTERVAL_CONTINUOUS = (0)              // Continuous channel survey.
const GAP_QOS_CHANNEL_SURVEY_INTERVAL_MIN_US = (7500)               // Minimum channel survey interval in microseconds (7.5 ms).
const GAP_QOS_CHANNEL_SURVEY_INTERVAL_MAX_US = (4000000)            // Maximum channel survey interval in microseconds (4 s).
const GAP_CHAR_INCL_CONFIG_INCLUDE = (0)                            // Include the characteristic in the Attribute Table
const GAP_CHAR_INCL_CONFIG_EXCLUDE_WITH_SPACE = (1)                 // Do not include the characteristic in the Attribute table. The SoftDevice will reserve the attribute handles which are otherwise used for this characteristic. By reserving the attribute handles it will be possible to upgrade the SoftDevice without changing handle of the Service Changed characteristic.
const GAP_CHAR_INCL_CONFIG_EXCLUDE_WITHOUT_SPACE = (2)              // Do not include the characteristic in the Attribute table. The SoftDevice will not reserve the attribute handles which are otherwise used for this characteristic.
const GAP_PPCP_INCL_CONFIG_DEFAULT = (GAP_CHAR_INCL_CONFIG_INCLUDE) // Included by default.
const GAP_CAR_INCL_CONFIG_DEFAULT = (GAP_CHAR_INCL_CONFIG_INCLUDE)  // Included by default.

//type GAPAdvProperties struct { bitfield }

//type GAPAdvReportType struct { bitfield }

type GAPAuxPointer struct {
	AuxOffset uint16 // Time offset from the beginning of advertising packet to the auxiliary packet in 100 us units.
	AuxPhy    uint8  // Indicates the PHY on which the auxiliary advertising packet is sent. See @ref BLE_GAP_PHYS.
}

//type GAPAddr struct { bitfield }

type GAPConnParams struct {
	MinConnInterval uint16 // Minimum Connection Interval in 1.25 ms units, see @ref BLE_GAP_CP_LIMITS.
	MaxConnInterval uint16 // Maximum Connection Interval in 1.25 ms units, see @ref BLE_GAP_CP_LIMITS.
	SlaveLatency    uint16 // Slave Latency in number of connection events, see @ref BLE_GAP_CP_LIMITS.
	ConnSupTimeout  uint16 // Connection Supervision Timeout in 10 ms units, see @ref BLE_GAP_CP_LIMITS.
}

//type GAPConnSecMode struct { bitfield }

type GAPConnSec struct {
	SecMode     GAPConnSecMode // Currently active security mode for this connection.
	EncrKeySize uint8          // Length of currently active encryption key, 7 to 16 octets (only applicable for bonding procedures).
}

type GAPIrk struct {
	Irk [GAP_SEC_KEY_LEN]uint8 // Array containing IRK.
}

type GAPChMask [5]uint8

//type GAPAdvParams struct { bitfield }

type GAPAdvData struct {
	AdvData     Data // Advertising data. @note Advertising data can only be specified for a @ref ble_gap_adv_properties_t::type that is allowed to contain advertising data.
	ScanRspData Data // Scan response data. @note Scan response data can only be specified for a @ref ble_gap_adv_properties_t::type that is scannable.
}

//type GAPScanParams struct { bitfield }

type GAPPrivacyParams struct {
	PrivacyMode       uint8   // Privacy mode, see @ref BLE_GAP_PRIVACY_MODES. Default is @ref BLE_GAP_PRIVACY_MODE_OFF.
	PrivateAddrType   uint8   // The private address type must be either @ref BLE_GAP_ADDR_TYPE_RANDOM_PRIVATE_RESOLVABLE or @ref BLE_GAP_ADDR_TYPE_RANDOM_PRIVATE_NON_RESOLVABLE.
	PrivateAddrCycleS uint16  // Private address cycle interval in seconds. Providing an address cycle value of 0 will use the default value defined by @ref BLE_GAP_DEFAULT_PRIVATE_ADDR_CYCLE_INTERVAL_S.
	DeviceIrk         *GAPIrk // When used as input, pointer to IRK structure that will be used as the default IRK. If NULL, the device default IRK will be used. When used as output, pointer to IRK structure where the current default IRK will be written to. If NULL, this argument is ignored. By default, the default IRK is used to generate random private resolvable addresses for the local device unless instructed otherwise.
}

type GAPPhys struct {
	TxPhys uint8 // Preferred transmit PHYs, see @ref BLE_GAP_PHYS.
	RxPhys uint8 // Preferred receive PHYs, see @ref BLE_GAP_PHYS.
}

//type GAPSecKdist struct { bitfield }

//type GAPSecParams struct { bitfield }

//type GAPEncInfo struct { bitfield }

type GAPMasterId struct {
	Ediv uint16                  // Encrypted Diversifier.
	Rand [GAP_SEC_RAND_LEN]uint8 // Random Number.
}

type GAPSignInfo struct {
	Csrk [GAP_SEC_KEY_LEN]uint8 // Connection Signature Resolving Key.
}

type GAPLescP256Pk struct {
	Pk [GAP_LESC_P256_PK_LEN]uint8 // LE Secure Connections Elliptic Curve Diffie-Hellman P-256 Public Key. Stored in the standard SMP protocol format: {X,Y} both in little-endian.
}

type GAPLescDhkey struct {
	Key [GAP_LESC_DHKEY_LEN]uint8 // LE Secure Connections Elliptic Curve Diffie-Hellman Key. Stored in little-endian.
}

type GAPLescOobData struct {
	Addr GAPAddr                // Bluetooth address of the device.
	R    [GAP_SEC_KEY_LEN]uint8 // Random Number.
	C    [GAP_SEC_KEY_LEN]uint8 // Confirm Value.
}

type GAPEvtConnected struct {
	PeerAddr   GAPAddr       // Bluetooth address of the peer device. If the peer_addr resolved: @ref ble_gap_addr_t::addr_id_peer is set to 1 and the address is the device's identity address.
	Role       uint8         // BLE role for this connection, see @ref BLE_GAP_ROLES
	ConnParams GAPConnParams // GAP Connection Parameters.
	AdvHandle  uint8         // Advertising handle in which advertising has ended. This variable is only set if role is set to @ref BLE_GAP_ROLE_PERIPH.
	AdvData    GAPAdvData    // Advertising buffers corresponding to the terminated advertising set. The advertising buffers provided in @ref sd_ble_gap_adv_set_configure are now released. This variable is only set if role is set to @ref BLE_GAP_ROLE_PERIPH.
}

type GAPEvtDisconnected struct {
	Reason uint8 // HCI error code, see @ref BLE_HCI_STATUS_CODES.
}

type GAPEvtConnParamUpdate struct {
	ConnParams GAPConnParams // GAP Connection Parameters.
}

type GAPEvtPhyUpdateRequest struct {
	PeerPreferredPhys GAPPhys // The PHYs the peer prefers to use.
}

type GAPEvtPhyUpdate struct {
	Status uint8 // Status of the procedure, see @ref BLE_HCI_STATUS_CODES.
	TxPhy  uint8 // TX PHY for this connection, see @ref BLE_GAP_PHYS.
	RxPhy  uint8 // RX PHY for this connection, see @ref BLE_GAP_PHYS.
}

type GAPEvtSecParamsRequest struct {
	PeerParams GAPSecParams // Initiator Security Parameters.
}

//type GAPEvtSecInfoRequest struct { bitfield }

//type GAPEvtPasskeyDisplay struct { bitfield }

type GAPEvtKeyPressed struct {
	KpNot uint8 // Keypress notification type, see @ref BLE_GAP_KP_NOT_TYPES.
}

type GAPEvtAuthKeyRequest struct {
	KeyType uint8 // See @ref BLE_GAP_AUTH_KEY_TYPES.
}

//type GAPEvtLescDhkeyRequest struct { bitfield }

//type GAPSecLevels struct { bitfield }

type GAPEncKey struct {
	EncInfo  GAPEncInfo  // Encryption Information.
	MasterId GAPMasterId // Master Identification.
}

type GAPIdKey struct {
	IdInfo     GAPIrk  // Identity Resolving Key.
	IdAddrInfo GAPAddr // Identity Address.
}

type GAPSecKeys struct {
	EncKey  *GAPEncKey     // Encryption Key, or NULL.
	IdKey   *GAPIdKey      // Identity Key, or NULL.
	SignKey *GAPSignInfo   // Signing Key, or NULL.
	Pk      *GAPLescP256Pk // LE Secure Connections P-256 Public Key. When in debug mode the application must use the value defined in the Core Bluetooth Specification v4.2 Vol.3, Part H, Section 2.3.5.6.1
}

type GAPSecKeyset struct {
	KeysOwn  GAPSecKeys // Keys distributed by the local device. For LE Secure Connections the encryption key will be generated locally and will always be stored if bonding.
	KeysPeer GAPSecKeys // Keys distributed by the remote device. For LE Secure Connections, p_enc_key must always be NULL.
}

type GAPDataLengthParams struct {
	MaxTxOctets uint16 // Maximum number of payload octets that a Controller supports for transmission of a single Link Layer Data Channel PDU.
	MaxRxOctets uint16 // Maximum number of payload octets that a Controller supports for reception of a single Link Layer Data Channel PDU.
	MaxTxTimeUs uint16 // Maximum time, in microseconds, that a Controller supports for transmission of a single Link Layer Data Channel PDU.
	MaxRxTimeUs uint16 // Maximum time, in microseconds, that a Controller supports for reception of a single Link Layer Data Channel PDU.
}

type GAPDataLengthLimitation struct {
	TxPayloadLimitedOctets uint16 // If > 0, the requested TX packet length is too long by this many octets.
	RxPayloadLimitedOctets uint16 // If > 0, the requested RX packet length is too long by this many octets.
	TxRxTimeLimitedUs      uint16 // If > 0, the requested combination of TX and RX packet lengths is too long by this many microseconds.
}

//type GAPEvtAuthStatus struct { bitfield }

type GAPEvtConnSecUpdate struct {
	ConnSec GAPConnSec // Connection security level.
}

//type GAPEvtTimeout struct { union }

type GAPEvtRssiChanged struct {
	Rssi    int8  // Received Signal Strength Indication in dBm. @note ERRATA-153 requires the rssi sample to be compensated based on a temperature measurement.
	ChIndex uint8 // Data Channel Index on which the Signal Strength is measured (0-36).
}

type GAPEvtAdvSetTerminated struct {
	Reason                uint8      // Reason for why the advertising set terminated. See @ref BLE_GAP_EVT_ADV_SET_TERMINATED_REASON.
	AdvHandle             uint8      // Advertising handle in which advertising has ended.
	NumCompletedAdvEvents uint8      // If @ref ble_gap_adv_params_t::max_adv_evts was not set to 0, this field indicates the number of completed advertising events.
	AdvData               GAPAdvData // Advertising buffers corresponding to the terminated advertising set. The advertising buffers provided in @ref sd_ble_gap_adv_set_configure are now released.
}

//type GAPEvtAdvReport struct { bitfield }

//type GAPEvtSecRequest struct { bitfield }

type GAPEvtConnParamUpdateRequest struct {
	ConnParams GAPConnParams // GAP Connection Parameters.
}

type GAPEvtScanReqReport struct {
	AdvHandle uint8   // Advertising handle for the advertising set which received the Scan Request
	Rssi      int8    // Received Signal Strength Indication in dBm. @note ERRATA-153 requires the rssi sample to be compensated based on a temperature measurement.
	PeerAddr  GAPAddr // Bluetooth address of the peer device. If the peer_addr resolved: @ref ble_gap_addr_t::addr_id_peer is set to 1 and the address is the device's identity address.
}

type GAPEvtDataLengthUpdateRequest struct {
	PeerParams GAPDataLengthParams // Peer data length parameters.
}

type GAPEvtDataLengthUpdate struct {
	EffectiveParams GAPDataLengthParams // The effective data length parameters.
}

type GAPEvtQosChannelSurveyReport struct {
	ChannelEnergy [GAP_CHANNEL_COUNT]int8 // The measured energy on the Bluetooth Low Energy channels, in dBm, indexed by Channel Index. If no measurement is available for the given channel, channel_energy is set to @ref BLE_GAP_POWER_LEVEL_INVALID.
}

//type GAPEvt struct { union }

type GAPConnCfg struct {
	ConnCount   uint8  // The number of concurrent connections the application can create with this configuration. The default and minimum value is @ref BLE_GAP_CONN_COUNT_DEFAULT.
	EventLength uint16 // The time set aside for this connection on every connection interval in 1.25 ms units. The default value is @ref BLE_GAP_EVENT_LENGTH_DEFAULT, the minimum value is @ref BLE_GAP_EVENT_LENGTH_MIN. The event length and the connection interval are the primary parameters for setting the throughput of a connection. See the SoftDevice Specification for details on throughput.
}

//type GAPCfgRoleCount struct { bitfield }

//type GAPCfgDeviceName struct { bitfield }

type GAPCfgPpcpInclCfg struct {
	IncludeCfg uint8 // Inclusion configuration of the Peripheral Preferred Connection Parameters characteristic. See @ref BLE_GAP_CHAR_INCL_CONFIG. Default is @ref BLE_GAP_PPCP_INCL_CONFIG_DEFAULT.
}

type GAPCfgCarInclCfg struct {
	IncludeCfg uint8 // Inclusion configuration of the Central Address Resolution characteristic. See @ref BLE_GAP_CHAR_INCL_CONFIG. Default is @ref BLE_GAP_CAR_INCL_CONFIG_DEFAULT.
}

//type GAPCfg union

type GAPOptChMap struct {
	ConnHandle uint16   // Connection Handle (only applicable for get)
	ChMap      [5]uint8 // Channel Map (37-bit).
}

type GAPOptLocalConnLatency struct {
	ConnHandle       uint16  // Connection Handle
	RequestedLatency uint16  // Requested local connection latency.
	ActualLatency    *uint16 // Pointer to storage for the actual local connection latency (can be set to NULL to skip return value).
}

//type GAPOptSlaveLatencyDisable struct { bitfield }

type GAPOptPasskey struct {
	Passkey *uint8 // Pointer to 6-digit ASCII string (digit 0..9 only, no NULL termination) passkey to be used during pairing. If this is NULL, the SoftDevice will generate a random passkey if required.
}

//type GAPOptCompatMode1 struct { bitfield }

type GAPOptAuthPayloadTimeout struct {
	ConnHandle         uint16 // Connection Handle
	AuthPayloadTimeout uint16 // Requested timeout in 10 ms unit, see @ref BLE_GAP_AUTH_PAYLOAD_TIMEOUT.
}

//type GAPOpt union

type GAPConnEventTrigger struct {
	PPIChId             uint8  // PPI channel to use. This channel should be regarded as reserved until connection event PPI task triggering is stopped. The PPI channel ID can not be one of the PPI channels reserved by the SoftDevice. See @ref NRF_SOC_SD_PPI_CHANNELS_SD_ENABLED_MSK.
	TaskEndpoint        uint32 // Task Endpoint to trigger.
	ConnEvtCounterStart uint16 // The connection event on which the task triggering should start.
	PeriodInEvents      uint16 // Trigger period. Valid range is [1, 32767]. If the device is in slave role and slave latency is enabled, this parameter should be set to a multiple of (slave latency + 1) to ensure low power operation.
}

// Set own Bluetooth Address.
func SetGAPAddr(addr *GAPAddr) uint32

// Get own Bluetooth Address.
func GetGAPAddr(addr *GAPAddr) uint32

// Get the Address used on air while Advertising.
func GetGAPAdvAddr(adv_handle uint8, addr *GAPAddr) uint32

// Set active whitelist.
func SetGAPWhitelist(pp_wl_addrs **GAPAddr, len uint8) uint32

// Set device identity list.
func SetGAPDeviceIdentities(pp_id_keys **GAPIdKey, pp_local_irks **GAPIrk, len uint8) uint32

// Set Privacy settings
func SetGAPPrivacy(privacy_params *GAPPrivacyParams) uint32

// Get Privacy settings
func GetGAPPrivacy(privacy_params *GAPPrivacyParams) uint32

// Configure an advertising set.
func ConfigureGAPAdvSet(adv_handle *uint8, adv_data *GAPAdvData, adv_params *GAPAdvParams) uint32

// Start Advertising.
func StartGAPAdv(adv_handle uint8, conn_cfg_tag uint8) uint32

// Stop Advertising.
func StopGAPAdv(adv_handle uint8) uint32

// Connection Parameter Update.
func UpdateGAPConnParam(conn_handle uint16, conn_params *GAPConnParams) uint32

// Disconnect.
func DisconnectGAP(conn_handle uint16, hci_status_code uint8) uint32

// Set TX Power.
func SetGAPTxPower(role uint8, handle uint16, tx_power int8) uint32

// Set Appearance.
func SetGAPAppearance(appearance uint16) uint32

// Get Appearance.
func GetGAPAppearance(appearance *uint16) uint32

// Set PPCP.
func SetGAPPpcp(conn_params *GAPConnParams) uint32

// Get PPCP.
func GetGAPPpcp(conn_params *GAPConnParams) uint32

// Set Device Name.
func SetGAPDeviceName(write_perm *GAPConnSecMode, dev_name *uint8, len uint16) uint32

// Get Device Name.
func GetGAPDeviceName(dev_name *uint8, len *uint16) uint32

// Initiate Pairing/Bonding.
func AuthenticateGAP(conn_handle uint16, sec_params *GAPSecParams) uint32

// Reply with Security Parameters.
func ReplyGAPSecParams(conn_handle uint16, sec_status uint8, sec_params *GAPSecParams, sec_keyset *GAPSecKeyset) uint32

// Reply with an authentication key.
func ReplyGAPAuthKey(conn_handle uint16, key_type uint8, key *uint8) uint32

// Reply with an LE Secure Connections DHKey.
func ReplyGAPLescDhkey(conn_handle uint16, dhkey *GAPLescDhkey) uint32

// Notify of a keypress during an authentication procedure.
func NotifyGAPKeypress(conn_handle uint16, kp_not uint8) uint32

// Get the local LE Secure Connections OOB data.
func GetGAPLescOobData(conn_handle uint16, pk_own *GAPLescP256Pk, oobd_own *GAPLescOobData) uint32

// Set the remote LE Secure Connections OOB data.
func SetGAPLescOobData(conn_handle uint16, oobd_own *GAPLescOobData, oobd_peer *GAPLescOobData) uint32

// Initiate encryption procedure.
func EncryptGAP(conn_handle uint16, master_id *GAPMasterId, enc_info *GAPEncInfo) uint32

// Reply with Security Information.
func ReplyGAPSecInfo(conn_handle uint16, enc_info *GAPEncInfo, id_info *GAPIrk, sign_info *GAPSignInfo) uint32

// Obtain connection security level.
func GetGAPConnSec(conn_handle uint16, conn_sec *GAPConnSec) uint32

// Start reporting of changes in RSSI.
func StartGAPRssi(conn_handle uint16, threshold_dbm uint8, skip_count uint8) uint32

// Stop reporting of changes in RSSI.
func StopGAPRssi(conn_handle uint16) uint32

// Get the last RSSI sample.
func GetGAPRssi(conn_handle uint16, rssi *int8, ch_index *uint8) uint32

// Start Scanning.
func StartGAPScan(scan_params *GAPScanParams, adv_report_buffer *Data) uint32

// Stop Scanning.
func StopGAPScan() uint32

// Connect.
func ConnectGAP(peer_addr *GAPAddr, scan_params *GAPScanParams, conn_params *GAPConnParams, conn_cfg_tag uint8) uint32

// Cancel ongoing connection procedure.
func CancelGAPConnect() uint32

// Initiate or respond to a PHY Update Procedure.
func UpdateGAPPhy(conn_handle uint16, gap_phys *GAPPhys) uint32

// Initiate or respond to a Data Length Update Procedure.
func UpdateGAPDataLength(conn_handle uint16, dl_params *GAPDataLengthParams, dl_limitation *GAPDataLengthLimitation) uint32

// Start Quality of Service (QoS) channel survey module.
func StartGAPQosChannelSurvey(interval_us uint32) uint32

// Stop Quality of Service (QoS) channel survey module.
func StopGAPQosChannelSurvey() uint32

// Get the next connection event counter.
func GetGAPNextConnEvtCounter(conn_handle uint16, counter *uint16) uint32

// Start triggering a given task on connection event start.
func StartGAPConnEvtTrigger(conn_handle uint16, params *GAPConnEventTrigger) uint32

// Stop triggering the task configured using @ref sd_ble_gap_conn_evt_trigger_start.
func StopGAPConnEvtTrigger(conn_handle uint16) uint32
