// Code generated from ../nordic/s140/headers/ble_types.h; DO NOT EDIT.

package ble

const CONN_HANDLE_INVALID = 0xFFFF                          // Invalid Connection Handle.
const CONN_HANDLE_ALL = 0xFFFE                              // Applies to all Connection Handles.
const UUID_UNKNOWN = 0x0000                                 // Reserved UUID.
const UUID_SERVICE_PRIMARY = 0x2800                         // Primary Service.
const UUID_SERVICE_SECONDARY = 0x2801                       // Secondary Service.
const UUID_SERVICE_INCLUDE = 0x2802                         // Include.
const UUID_CHARACTERISTIC = 0x2803                          // Characteristic.
const UUID_DESCRIPTOR_CHAR_EXT_PROP = 0x2900                // Characteristic Extended Properties Descriptor.
const UUID_DESCRIPTOR_CHAR_USER_DESC = 0x2901               // Characteristic User Description Descriptor.
const UUID_DESCRIPTOR_CLIENT_CHAR_CONFIG = 0x2902           // Client Characteristic Configuration Descriptor.
const UUID_DESCRIPTOR_SERVER_CHAR_CONFIG = 0x2903           // Server Characteristic Configuration Descriptor.
const UUID_DESCRIPTOR_CHAR_PRESENTATION_FORMAT = 0x2904     // Characteristic Presentation Format Descriptor.
const UUID_DESCRIPTOR_CHAR_AGGREGATE_FORMAT = 0x2905        // Characteristic Aggregate Format Descriptor.
const UUID_GATT = 0x1801                                    // Generic Attribute Profile.
const UUID_GATT_CHARACTERISTIC_SERVICE_CHANGED = 0x2A05     // Service Changed Characteristic.
const UUID_GAP = 0x1800                                     // Generic Access Profile.
const UUID_GAP_CHARACTERISTIC_DEVICE_NAME = 0x2A00          // Device Name Characteristic.
const UUID_GAP_CHARACTERISTIC_APPEARANCE = 0x2A01           // Appearance Characteristic.
const UUID_GAP_CHARACTERISTIC_RECONN_ADDR = 0x2A03          // Reconnection Address Characteristic.
const UUID_GAP_CHARACTERISTIC_PPCP = 0x2A04                 // Peripheral Preferred Connection Parameters Characteristic.
const UUID_GAP_CHARACTERISTIC_CAR = 0x2AA6                  // Central Address Resolution Characteristic.
const UUID_GAP_CHARACTERISTIC_RPA_ONLY = 0x2AC9             // Resolvable Private Address Only Characteristic.
const UUID_TYPE_UNKNOWN = 0x00                              // Invalid UUID type.
const UUID_TYPE_BLE = 0x01                                  // Bluetooth SIG UUID (16-bit).
const UUID_TYPE_VENDOR_BEGIN = 0x02                         // Vendor UUID types start at this index (128-bit).
const APPEARANCE_UNKNOWN = 0                                // Unknown.
const APPEARANCE_GENERIC_PHONE = 64                         // Generic Phone.
const APPEARANCE_GENERIC_COMPUTER = 128                     // Generic Computer.
const APPEARANCE_GENERIC_WATCH = 192                        // Generic Watch.
const APPEARANCE_WATCH_SPORTS_WATCH = 193                   // Watch: Sports Watch.
const APPEARANCE_GENERIC_CLOCK = 256                        // Generic Clock.
const APPEARANCE_GENERIC_DISPLAY = 320                      // Generic Display.
const APPEARANCE_GENERIC_REMOTE_CONTROL = 384               // Generic Remote Control.
const APPEARANCE_GENERIC_EYE_GLASSES = 448                  // Generic Eye-glasses.
const APPEARANCE_GENERIC_TAG = 512                          // Generic Tag.
const APPEARANCE_GENERIC_KEYRING = 576                      // Generic Keyring.
const APPEARANCE_GENERIC_MEDIA_PLAYER = 640                 // Generic Media Player.
const APPEARANCE_GENERIC_BARCODE_SCANNER = 704              // Generic Barcode Scanner.
const APPEARANCE_GENERIC_THERMOMETER = 768                  // Generic Thermometer.
const APPEARANCE_THERMOMETER_EAR = 769                      // Thermometer: Ear.
const APPEARANCE_GENERIC_HEART_RATE_SENSOR = 832            // Generic Heart rate Sensor.
const APPEARANCE_HEART_RATE_SENSOR_HEART_RATE_BELT = 833    // Heart Rate Sensor: Heart Rate Belt.
const APPEARANCE_GENERIC_BLOOD_PRESSURE = 896               // Generic Blood Pressure.
const APPEARANCE_BLOOD_PRESSURE_ARM = 897                   // Blood Pressure: Arm.
const APPEARANCE_BLOOD_PRESSURE_WRIST = 898                 // Blood Pressure: Wrist.
const APPEARANCE_GENERIC_HID = 960                          // Human Interface Device (HID).
const APPEARANCE_HID_KEYBOARD = 961                         // Keyboard (HID Subtype).
const APPEARANCE_HID_MOUSE = 962                            // Mouse (HID Subtype).
const APPEARANCE_HID_JOYSTICK = 963                         // Joystick (HID Subtype).
const APPEARANCE_HID_GAMEPAD = 964                          // Gamepad (HID Subtype).
const APPEARANCE_HID_DIGITIZERSUBTYPE = 965                 // Digitizer Tablet (HID Subtype).
const APPEARANCE_HID_CARD_READER = 966                      // Card Reader (HID Subtype).
const APPEARANCE_HID_DIGITAL_PEN = 967                      // Digital Pen (HID Subtype).
const APPEARANCE_HID_BARCODE = 968                          // Barcode Scanner (HID Subtype).
const APPEARANCE_GENERIC_GLUCOSE_METER = 1024               // Generic Glucose Meter.
const APPEARANCE_GENERIC_RUNNING_WALKING_SENSOR = 1088      // Generic Running Walking Sensor.
const APPEARANCE_RUNNING_WALKING_SENSOR_IN_SHOE = 1089      // Running Walking Sensor: In-Shoe.
const APPEARANCE_RUNNING_WALKING_SENSOR_ON_SHOE = 1090      // Running Walking Sensor: On-Shoe.
const APPEARANCE_RUNNING_WALKING_SENSOR_ON_HIP = 1091       // Running Walking Sensor: On-Hip.
const APPEARANCE_GENERIC_CYCLING = 1152                     // Generic Cycling.
const APPEARANCE_CYCLING_CYCLING_COMPUTER = 1153            // Cycling: Cycling Computer.
const APPEARANCE_CYCLING_SPEED_SENSOR = 1154                // Cycling: Speed Sensor.
const APPEARANCE_CYCLING_CADENCE_SENSOR = 1155              // Cycling: Cadence Sensor.
const APPEARANCE_CYCLING_POWER_SENSOR = 1156                // Cycling: Power Sensor.
const APPEARANCE_CYCLING_SPEED_CADENCE_SENSOR = 1157        // Cycling: Speed and Cadence Sensor.
const APPEARANCE_GENERIC_PULSE_OXIMETER = 3136              // Generic Pulse Oximeter.
const APPEARANCE_PULSE_OXIMETER_FINGERTIP = 3137            // Fingertip (Pulse Oximeter subtype).
const APPEARANCE_PULSE_OXIMETER_WRIST_WORN = 3138           // Wrist Worn(Pulse Oximeter subtype).
const APPEARANCE_GENERIC_WEIGHT_SCALE = 3200                // Generic Weight Scale.
const APPEARANCE_GENERIC_OUTDOOR_SPORTS_ACT = 5184          // Generic Outdoor Sports Activity.
const APPEARANCE_OUTDOOR_SPORTS_ACT_LOC_DISP = 5185         // Location Display Device (Outdoor Sports Activity subtype).
const APPEARANCE_OUTDOOR_SPORTS_ACT_LOC_AND_NAV_DISP = 5186 // Location and Navigation Display Device (Outdoor Sports Activity subtype).
const APPEARANCE_OUTDOOR_SPORTS_ACT_LOC_POD = 5187          // Location Pod (Outdoor Sports Activity subtype).
const APPEARANCE_OUTDOOR_SPORTS_ACT_LOC_AND_NAV_POD = 5188  // Location and Navigation Pod (Outdoor Sports Activity subtype).

type UUID128 struct {
	UUID128 [16]uint8 // Little-Endian UUID bytes.
}

type UUID struct {
	UUID uint16 // 16-bit UUID value or octets 12-13 of 128-bit UUID.
	Typ  uint8  // UUID type, see @ref BLE_UUID_TYPES. If type is @ref BLE_UUID_TYPE_UNKNOWN, the value of uuid is undefined.
}

type Data struct {
	Data *uint8 // Pointer to the data buffer provided to/from the application.
	Len  uint16 // Length of the data buffer, in bytes.
}
