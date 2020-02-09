// Code generated from ../nordic/s140/headers/nrf_soc.h; DO NOT EDIT.

package nrf

import (
	"github.com/embeddedgo/nrf5/p/irq"
	"unsafe"
)

const SOC_SVC_BASE = (0x20)               // Base value for SVCs that are available when the SoftDevice is disabled.
const SOC_SVC_BASE_NOT_AVAILABLE = (0x2C) // Base value for SVCs that are not available when the SoftDevice is disabled.
const RADIO_NOTIFICATION_INACTIVE_GUARANTEED_TIME_US = (62)
const RADIO_MINIMUM_TIMESLOT_LENGTH_EXTENSION_TIME_US = (200)
const RADIO_MAX_EXTENSION_PROCESSING_TIME_US = (20)
const RADIO_MIN_EXTENSION_MARGIN_US = (82)
const SOC_ECB_KEY_LENGTH = (16)                              // ECB key length.
const SOC_ECB_CLEARTEXT_LENGTH = (16)                        // ECB cleartext length.
const SOC_ECB_CIPHERTEXT_LENGTH = (SOC_ECB_CLEARTEXT_LENGTH) // ECB ciphertext length.
const EVT_IRQn = (irq.SWI2_EGU2)                             // SoftDevice Event IRQ number. Used for both protocol events and SoC events.
const RADIO_NOTIFICATION_IRQn = (irq.SWI1_EGU1)              // The radio notification IRQ number.
const RADIO_LENGTH_MIN_US = (100)                            // The shortest allowed radio timeslot, in microseconds.
const RADIO_LENGTH_MAX_US = (100000)                         // The longest allowed radio timeslot, in microseconds.
const RADIO_DISTANCE_MAX_US = (128000000 - 1)                // The longest timeslot distance, in microseconds, allowed for the distance parameter (see @ref nrf_radio_request_normal_t) in the request.
const RADIO_EARLIEST_TIMEOUT_MAX_US = (128000000 - 1)        // The longest timeout, in microseconds, allowed when requesting the earliest possible timeslot.
const RADIO_START_JITTER_US = (2)                            // The maximum jitter in @ref NRF_RADIO_CALLBACK_SIGNAL_TYPE_START relative to the requested start time.
const SOC_SD_PPI_CHANNELS_SD_DISABLED_MSK = ((uint32)(0))
const SOC_SD_PPI_CHANNELS_SD_ENABLED_MSK = ((uint32)((1 << 17) | (1 << 18) | (1 << 19) | (1 << 20) | (1 << 21) | (1 << 22) | (1 << 23) | (1 << 24) | (1 << 25) | (1 << 26) | (1 << 27) | (1 << 28) | (1 << 29) | (1 << 30) | (1 << 31)))
const SOC_SD_PPI_GROUPS_SD_DISABLED_MSK = ((uint32)(0))
const SOC_SD_PPI_GROUPS_SD_ENABLED_MSK = ((uint32)((1 << 4) | (1 << 5)))

const (
	_PPI_CHANNEL_ENABLE_GET               = SOC_SVC_BASE + iota
	_PPI_CHANNEL_ENABLE_SET               = SOC_SVC_BASE + 1
	_PPI_CHANNEL_ENABLE_CLR               = SOC_SVC_BASE + 2
	_PPI_CHANNEL_ASSIGN                   = SOC_SVC_BASE + 3
	_PPI_GROUP_TASK_ENABLE                = SOC_SVC_BASE + 4
	_PPI_GROUP_TASK_DISABLE               = SOC_SVC_BASE + 5
	_PPI_GROUP_ASSIGN                     = SOC_SVC_BASE + 6
	_PPI_GROUP_GET                        = SOC_SVC_BASE + 7
	_FLASH_PAGE_ERASE                     = SOC_SVC_BASE + 8
	_FLASH_WRITE                          = SOC_SVC_BASE + 9
	_PROTECTED_REGISTER_WRITE             = SOC_SVC_BASE + 11
	_MUTEX_NEW                            = SOC_SVC_BASE_NOT_AVAILABLE
	_MUTEX_ACQUIRE                        = SOC_SVC_BASE_NOT_AVAILABLE + 1
	_MUTEX_RELEASE                        = SOC_SVC_BASE_NOT_AVAILABLE + 2
	_RAND_APPLICATION_POOL_CAPACITY_GET   = SOC_SVC_BASE_NOT_AVAILABLE + 3
	_RAND_APPLICATION_BYTES_AVAILABLE_GET = SOC_SVC_BASE_NOT_AVAILABLE + 4
	_RAND_APPLICATION_VECTOR_GET          = SOC_SVC_BASE_NOT_AVAILABLE + 5
	_POWER_MODE_SET                       = SOC_SVC_BASE_NOT_AVAILABLE + 6
	_POWER_SYSTEM_OFF                     = SOC_SVC_BASE_NOT_AVAILABLE + 7
	_POWER_RESET_REASON_GET               = SOC_SVC_BASE_NOT_AVAILABLE + 8
	_POWER_RESET_REASON_CLR               = SOC_SVC_BASE_NOT_AVAILABLE + 9
	_POWER_POF_ENABLE                     = SOC_SVC_BASE_NOT_AVAILABLE + 10
	_POWER_POF_THRESHOLD_SET              = SOC_SVC_BASE_NOT_AVAILABLE + 11
	_POWER_POF_THRESHOLDVDDH_SET          = SOC_SVC_BASE_NOT_AVAILABLE + 12
	_POWER_RAM_POWER_SET                  = SOC_SVC_BASE_NOT_AVAILABLE + 13
	_POWER_RAM_POWER_CLR                  = SOC_SVC_BASE_NOT_AVAILABLE + 14
	_POWER_RAM_POWER_GET                  = SOC_SVC_BASE_NOT_AVAILABLE + 15
	_POWER_GPREGRET_SET                   = SOC_SVC_BASE_NOT_AVAILABLE + 16
	_POWER_GPREGRET_CLR                   = SOC_SVC_BASE_NOT_AVAILABLE + 17
	_POWER_GPREGRET_GET                   = SOC_SVC_BASE_NOT_AVAILABLE + 18
	_POWER_DCDC_MODE_SET                  = SOC_SVC_BASE_NOT_AVAILABLE + 19
	_POWER_DCDC0_MODE_SET                 = SOC_SVC_BASE_NOT_AVAILABLE + 20
	_APP_EVT_WAIT                         = SOC_SVC_BASE_NOT_AVAILABLE + 21
	_CLOCK_HFCLK_REQUEST                  = SOC_SVC_BASE_NOT_AVAILABLE + 22
	_CLOCK_HFCLK_RELEASE                  = SOC_SVC_BASE_NOT_AVAILABLE + 23
	_CLOCK_HFCLK_IS_RUNNING               = SOC_SVC_BASE_NOT_AVAILABLE + 24
	_RADIO_NOTIFICATION_CFG_SET           = SOC_SVC_BASE_NOT_AVAILABLE + 25
	_ECB_BLOCK_ENCRYPT                    = SOC_SVC_BASE_NOT_AVAILABLE + 26
	_ECB_BLOCKS_ENCRYPT                   = SOC_SVC_BASE_NOT_AVAILABLE + 27
	_RADIO_SESSION_OPEN                   = SOC_SVC_BASE_NOT_AVAILABLE + 28
	_RADIO_SESSION_CLOSE                  = SOC_SVC_BASE_NOT_AVAILABLE + 29
	_RADIO_REQUEST                        = SOC_SVC_BASE_NOT_AVAILABLE + 30
	_EVT_GET                              = SOC_SVC_BASE_NOT_AVAILABLE + 31
	_TEMP_GET                             = SOC_SVC_BASE_NOT_AVAILABLE + 32
	_POWER_USBPWRRDY_ENABLE               = SOC_SVC_BASE_NOT_AVAILABLE + 33
	_POWER_USBDETECTED_ENABLE             = SOC_SVC_BASE_NOT_AVAILABLE + 34
	_POWER_USBREMOVED_ENABLE              = SOC_SVC_BASE_NOT_AVAILABLE + 35
	_POWER_USBREGSTATUS_GET               = SOC_SVC_BASE_NOT_AVAILABLE + 36
	_SVC_SOC_LAST                         = SOC_SVC_BASE_NOT_AVAILABLE + 37
)

const (
	MUTEX_FREE = iota
	MUTEX_TAKEN
)

const (
	POWER_MODE_CONSTLAT = iota // Constant latency mode. See power management in the reference manual.
	POWER_MODE_LOWPWR          // Low power mode. See power management in the reference manual.
)

const (
	POWER_THRESHOLD_V17 = 4 + iota // 1.7 Volts power failure threshold.
	POWER_THRESHOLD_V18            // 1.8 Volts power failure threshold.
	POWER_THRESHOLD_V19            // 1.9 Volts power failure threshold.
	POWER_THRESHOLD_V20            // 2.0 Volts power failure threshold.
	POWER_THRESHOLD_V21            // 2.1 Volts power failure threshold.
	POWER_THRESHOLD_V22            // 2.2 Volts power failure threshold.
	POWER_THRESHOLD_V23            // 2.3 Volts power failure threshold.
	POWER_THRESHOLD_V24            // 2.4 Volts power failure threshold.
	POWER_THRESHOLD_V25            // 2.5 Volts power failure threshold.
	POWER_THRESHOLD_V26            // 2.6 Volts power failure threshold.
	POWER_THRESHOLD_V27            // 2.7 Volts power failure threshold.
	POWER_THRESHOLD_V28            // 2.8 Volts power failure threshold.
)

const (
	POWER_THRESHOLDVDDH_V27 = iota // 2.7 Volts power failure threshold.
	POWER_THRESHOLDVDDH_V28        // 2.8 Volts power failure threshold.
	POWER_THRESHOLDVDDH_V29        // 2.9 Volts power failure threshold.
	POWER_THRESHOLDVDDH_V30        // 3.0 Volts power failure threshold.
	POWER_THRESHOLDVDDH_V31        // 3.1 Volts power failure threshold.
	POWER_THRESHOLDVDDH_V32        // 3.2 Volts power failure threshold.
	POWER_THRESHOLDVDDH_V33        // 3.3 Volts power failure threshold.
	POWER_THRESHOLDVDDH_V34        // 3.4 Volts power failure threshold.
	POWER_THRESHOLDVDDH_V35        // 3.5 Volts power failure threshold.
	POWER_THRESHOLDVDDH_V36        // 3.6 Volts power failure threshold.
	POWER_THRESHOLDVDDH_V37        // 3.7 Volts power failure threshold.
	POWER_THRESHOLDVDDH_V38        // 3.8 Volts power failure threshold.
	POWER_THRESHOLDVDDH_V39        // 3.9 Volts power failure threshold.
	POWER_THRESHOLDVDDH_V40        // 4.0 Volts power failure threshold.
	POWER_THRESHOLDVDDH_V41        // 4.1 Volts power failure threshold.
	POWER_THRESHOLDVDDH_V42        // 4.2 Volts power failure threshold.
)

const (
	POWER_DCDC_DISABLE = iota // The DCDC is disabled.
	POWER_DCDC_ENABLE         // The DCDC is enabled.
)

const (
	RADIO_NOTIFICATION_DISTANCE_NONE   = 0 + iota // The event does not have a notification.
	RADIO_NOTIFICATION_DISTANCE_800US             // The distance from the active notification to start of radio activity.
	RADIO_NOTIFICATION_DISTANCE_1740US            // The distance from the active notification to start of radio activity.
	RADIO_NOTIFICATION_DISTANCE_2680US            // The distance from the active notification to start of radio activity.
	RADIO_NOTIFICATION_DISTANCE_3620US            // The distance from the active notification to start of radio activity.
	RADIO_NOTIFICATION_DISTANCE_4560US            // The distance from the active notification to start of radio activity.
	RADIO_NOTIFICATION_DISTANCE_5500US            // The distance from the active notification to start of radio activity.
)

const (
	RADIO_NOTIFICATION_TYPE_NONE            = 0 + iota // The event does not have a radio notification signal.
	RADIO_NOTIFICATION_TYPE_INT_ON_ACTIVE              // Using interrupt for notification when the radio will be enabled.
	RADIO_NOTIFICATION_TYPE_INT_ON_INACTIVE            // Using interrupt for notification when the radio has been disabled.
	RADIO_NOTIFICATION_TYPE_INT_ON_BOTH                // Using interrupt for notification both when the radio will be enabled and disabled.
)

const (
	RADIO_CALLBACK_SIGNAL_TYPE_START            = iota // This signal indicates the start of the radio timeslot.
	RADIO_CALLBACK_SIGNAL_TYPE_TIMER0                  // This signal indicates the NRF_TIMER0 interrupt.
	RADIO_CALLBACK_SIGNAL_TYPE_RADIO                   // This signal indicates the NRF_RADIO interrupt.
	RADIO_CALLBACK_SIGNAL_TYPE_EXTEND_FAILED           // This signal indicates extend action failed.
	RADIO_CALLBACK_SIGNAL_TYPE_EXTEND_SUCCEEDED        // This signal indicates extend action succeeded.
)

const (
	RADIO_SIGNAL_CALLBACK_ACTION_NONE            = iota // Return without action.
	RADIO_SIGNAL_CALLBACK_ACTION_EXTEND                 // Request an extension of the current timeslot. Maximum execution time for this action: @ref NRF_RADIO_MAX_EXTENSION_PROCESSING_TIME_US. This action must be started at least @ref NRF_RADIO_MIN_EXTENSION_MARGIN_US before the end of the timeslot.
	RADIO_SIGNAL_CALLBACK_ACTION_END                    // End the current radio timeslot.
	RADIO_SIGNAL_CALLBACK_ACTION_REQUEST_AND_END        // Request a new radio timeslot and end the current timeslot.
)

const (
	RADIO_HFCLK_CFG_XTAL_GUARANTEED = iota // The SoftDevice will guarantee that the high frequency clock source is the external crystal for the whole duration of the timeslot. This should be the preferred option for events that use the radio or require high timing accuracy. @note The SoftDevice will automatically turn on and off the external crystal, at the beginning and end of the timeslot, respectively. The crystal may also intentionally be left running after the timeslot, in cases where it is needed by the SoftDevice shortly after the end of the timeslot.
	RADIO_HFCLK_CFG_NO_GUARANTEE           // This configuration allows for earlier and tighter scheduling of timeslots. The RC oscillator may be the clock source in part or for the whole duration of the timeslot. The RC oscillator's accuracy must therefore be taken into consideration. @note If the application will use the radio peripheral in timeslots with this configuration, it must make sure that the crystal is running and stable before starting the radio.
)

const (
	RADIO_PRIORITY_HIGH   = iota // High (equal priority as the normal connection priority of the SoftDevice stack(s)).
	RADIO_PRIORITY_NORMAL        // Normal (equal priority as the priority of secondary activities of the SoftDevice stack(s)).
)

const (
	RADIO_REQ_TYPE_EARLIEST = iota // Request radio timeslot as early as possible. This should always be used for the first request in a session.
	RADIO_REQ_TYPE_NORMAL          // Normal radio timeslot request.
)

const (
	EVT_HFCLKSTARTED                         = iota // Event indicating that the HFCLK has started.
	EVT_POWER_FAILURE_WARNING                       // Event indicating that a power failure warning has occurred.
	EVT_FLASH_OPERATION_SUCCESS                     // Event indicating that the ongoing flash operation has completed successfully.
	EVT_FLASH_OPERATION_ERROR                       // Event indicating that the ongoing flash operation has timed out with an error.
	EVT_RADIO_BLOCKED                               // Event indicating that a radio timeslot was blocked.
	EVT_RADIO_CANCELED                              // Event indicating that a radio timeslot was canceled by SoftDevice.
	EVT_RADIO_SIGNAL_CALLBACK_INVALID_RETURN        // Event indicating that a radio timeslot signal callback handler return was invalid.
	EVT_RADIO_SESSION_IDLE                          // Event indicating that a radio timeslot session is idle.
	EVT_RADIO_SESSION_CLOSED                        // Event indicating that a radio timeslot session is closed.
	EVT_POWER_USB_POWER_READY                       // Event indicating that a USB 3.3 V supply is ready.
	EVT_POWER_USB_DETECTED                          // Event indicating that voltage supply is detected on VBUS.
	EVT_POWER_USB_REMOVED                           // Event indicating that voltage supply is removed from VBUS.
	EVT_NUMBER_OF_EVTS
)

type Mutex uint8

type RadioRequestEarliest struct {
	HF        uint8  // High frequency clock source, see @ref NRF_RADIO_HFCLK_CFG.
	Priority  uint8  // The radio timeslot priority, see @ref NRF_RADIO_PRIORITY.
	LengthUs  uint32 // The radio timeslot length (in the range 100 to 100,000] microseconds).
	TimeoutUs uint32 // Longest acceptable delay until the start of the requested timeslot (up to @ref NRF_RADIO_EARLIEST_TIMEOUT_MAX_US microseconds).
}

type RadioRequestNormal struct {
	HF         uint8  // High frequency clock source, see @ref NRF_RADIO_HFCLK_CFG.
	Priority   uint8  // The radio timeslot priority, see @ref NRF_RADIO_PRIORITY.
	DistanceUs uint32 // Distance from the start of the previous radio timeslot (up to @ref NRF_RADIO_DISTANCE_MAX_US microseconds).
	LengthUs   uint32 // The radio timeslot length (in the range [100..100,000] microseconds).
}

//type RadioRequest struct { union }

//type RadioSignalCallbackReturnParam struct { union }

type SocEcbKey [SOC_ECB_KEY_LENGTH]uint8
type SocEcbCleartext [SOC_ECB_CLEARTEXT_LENGTH]uint8
type SocEcbCiphertext [SOC_ECB_CIPHERTEXT_LENGTH]uint8

type EcbHalData struct {
	Key        SocEcbKey        // Encryption key.
	Cleartext  SocEcbCleartext  // Cleartext data.
	Ciphertext SocEcbCiphertext // Ciphertext data.
}

type EcbHalDataBlock struct {
	Key        *SocEcbKey        // Pointer to the Encryption key.
	Cleartext  *SocEcbCleartext  // Pointer to the Cleartext data.
	Ciphertext *SocEcbCiphertext // Pointer to the Ciphertext data.
}

func NewMutex(mutex *Mutex) uint32
func AcquireMutex(mutex *Mutex) uint32
func ReleaseMutex(mutex *Mutex) uint32
func GetRandApplicationPoolCapacity(pool_capacity *uint8) uint32
func GetRandApplicationBytesAvailable(bytes_available *uint8) uint32
func GetRandApplicationVector(buff *uint8, length uint8) uint32
func GetPowerResetReason(reset_reason *uint32) uint32
func ClrPowerResetReason(reset_reason_clr_msk uint32) uint32
func SetPowerMode(power_mode uint8) uint32
func OffPowerSystem() uint32
func EnablePowerPof(pof_enable uint8) uint32
func EnablePowerUsbpwrrdy(usbpwrrdy_enable uint8) uint32
func EnablePowerUsbdetected(usbdetected_enable uint8) uint32
func EnablePowerUsbremoved(usbremoved_enable uint8) uint32
func GetPowerUsbregstatus(usbregstatus *uint32) uint32
func SetPowerPofThreshold(threshold uint8) uint32
func SetPowerPofThresholdvddh(threshold uint8) uint32
func SetPowerRamPower(index uint8, ram_powerset uint32) uint32
func ClrPowerRamPower(index uint8, ram_powerclr uint32) uint32
func GetPowerRamPower(index uint8, ram_power *uint32) uint32
func SetPowerGpregret(gpregret_id uint32, gpregret_msk uint32) uint32
func ClrPowerGpregret(gpregret_id uint32, gpregret_msk uint32) uint32
func GetPowerGpregret(gpregret_id uint32, gpregret *uint32) uint32
func SetPowerDcdcMode(dcdc_mode uint8) uint32
func SetPowerDcdc0Mode(dcdc_mode uint8) uint32
func RequestClockHF() uint32
func ReleaseClockHF() uint32
func IsRunningClockHF(is_running *uint32) uint32
func WaitAppEvt() uint32
func GetPPIChannelEnable(channel_enable *uint32) uint32
func SetPPIChannelEnable(channel_enable_set_msk uint32) uint32
func ClrPPIChannelEnable(channel_enable_clr_msk uint32) uint32
func AssignPPIChannel(channel_num uint8, evt_endpoint unsafe.Pointer, task_endpoint unsafe.Pointer) uint32
func EnablePPIGroupTask(group_num uint8) uint32
func DisablePPIGroupTask(group_num uint8) uint32
func AssignPPIGroup(group_num uint8, channel_msk uint32) uint32
func GetPPIGroup(group_num uint8, channel_msk *uint32) uint32
func SetRadioNotificationCfg(typ uint8, distance uint8) uint32
func EncryptEcbBlock(ecb_data *EcbHalData) uint32
func EncryptEcbBlocks(block_count uint8, data_blocks *EcbHalDataBlock) uint32
func GetEvt(evt_id *uint32) uint32
func GetTemp(temp *int32) uint32
func WriteFlash(dst *uint32, src *uint32, size uint32) uint32
func EraseFlashPage(page_number uint32) uint32
func OpenRadioSession(p_radio_signal_callback unsafe.Pointer) uint32
func CloseRadioSession() uint32
func RequestRadio(request *RadioRequest) uint32
func WriteProtectedRegister(register *uint32, value uint32) uint32
