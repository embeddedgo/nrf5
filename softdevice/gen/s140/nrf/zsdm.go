// Code generated from ../nordic/s140/headers/nrf_sdm.h; DO NOT EDIT.

package nrf

import (
	"unsafe"
)

const MBR_SIZE = 0
const MAJOR_VERSION = (7)
const MINOR_VERSION = (0)
const BUGFIX_VERSION = (1)
const VARIANT_ID = 140
const VERSION = (MAJOR_VERSION*1000000 + MINOR_VERSION*1000 + BUGFIX_VERSION)
const SDM_SVC_BASE = 0x10
const UNIQUE_STR_SIZE = 20
const SDM_INFO_FIELD_INVALID = (0)
const SOFTDEVICE_INFO_STRUCT_OFFSET = (0x2000)
const SOFTDEVICE_INFO_STRUCT_ADDRESS = (SOFTDEVICE_INFO_STRUCT_OFFSET + MBR_SIZE)
const INFO_STRUCT_SIZE_OFFSET = (SOFTDEVICE_INFO_STRUCT_OFFSET)
const SIZE_OFFSET = (SOFTDEVICE_INFO_STRUCT_OFFSET + 0x08)
const FWID_OFFSET = (SOFTDEVICE_INFO_STRUCT_OFFSET + 0x0C)
const ID_OFFSET = (SOFTDEVICE_INFO_STRUCT_OFFSET + 0x10)
const VERSION_OFFSET = (SOFTDEVICE_INFO_STRUCT_OFFSET + 0x14)
const UNIQUE_STR_OFFSET = (SOFTDEVICE_INFO_STRUCT_OFFSET + 0x18)
const FLASH_SIZE = 0x26000
const FAULT_ID_SD_RANGE_START = 0x00000000                 // SoftDevice ID range start.
const FAULT_ID_APP_RANGE_START = 0x00001000                // Application ID range start.
const FAULT_ID_SD_ASSERT = (FAULT_ID_SD_RANGE_START + 1)   // SoftDevice assertion. The info parameter is reserved for future used.
const FAULT_ID_APP_MEMACC = (FAULT_ID_APP_RANGE_START + 1) // Application invalid memory access. The info parameter will contain 0x00000000, in case of SoftDevice RAM access violation. In case of SoftDevice peripheral register violation the info parameter will contain the sub-region number of PREGION[0], on whose address range the disallowed write access caused the memory access fault.

const (
	_SOFTDEVICE_ENABLE = SDM_SVC_BASE + iota
	_SOFTDEVICE_DISABLE
	_SOFTDEVICE_IS_ENABLED
	_SOFTDEVICE_VECTOR_TABLE_BASE_SET
	_SVC_SDM_LAST
)

const CLOCK_LF_ACCURACY_250_PPM = (0) // Default: 250 ppm
const CLOCK_LF_ACCURACY_500_PPM = (1) // 500 ppm
const CLOCK_LF_ACCURACY_150_PPM = (2) // 150 ppm
const CLOCK_LF_ACCURACY_100_PPM = (3) // 100 ppm
const CLOCK_LF_ACCURACY_75_PPM = (4)  // 75 ppm
const CLOCK_LF_ACCURACY_50_PPM = (5)  // 50 ppm
const CLOCK_LF_ACCURACY_30_PPM = (6)  // 30 ppm
const CLOCK_LF_ACCURACY_20_PPM = (7)  // 20 ppm
const CLOCK_LF_ACCURACY_10_PPM = (8)  // 10 ppm
const CLOCK_LF_ACCURACY_5_PPM = (9)   // 5 ppm
const CLOCK_LF_ACCURACY_2_PPM = (10)  // 2 ppm
const CLOCK_LF_ACCURACY_1_PPM = (11)  // 1 ppm
const CLOCK_LF_SRC_RC = (0)           // LFCLK RC oscillator.
const CLOCK_LF_SRC_XTAL = (1)         // LFCLK crystal oscillator.
const CLOCK_LF_SRC_SYNTH = (2)        // LFCLK Synthesized from HFCLK.

type ClockLFCfg struct {
	Source     uint8 // LF oscillator clock source, see @ref NRF_CLOCK_LF_SRC.
	RCCtiv     uint8 // Only for ::NRF_CLOCK_LF_SRC_RC: Calibration timer interval in 1/4 second units (nRF52: 1-32). @note To avoid excessive clock drift, 0.5 degrees Celsius is the maximum temperature change allowed in one calibration timer interval. The interval should be selected to ensure this. @note Must be 0 if source is not ::NRF_CLOCK_LF_SRC_RC.
	RCTempCtiv uint8 // Only for ::NRF_CLOCK_LF_SRC_RC: How often (in number of calibration intervals) the RC oscillator shall be calibrated if the temperature hasn't changed. 0: Always calibrate even if the temperature hasn't changed. 1: Only calibrate if the temperature has changed (legacy - nRF51 only). 2-33: Check the temperature and only calibrate if it has changed, however calibration will take place every rc_temp_ctiv intervals in any case. @note Must be 0 if source is not ::NRF_CLOCK_LF_SRC_RC. @note For nRF52, the application must ensure calibration at least once every 8 seconds to ensure +/-500 ppm clock stability. The recommended configuration for ::NRF_CLOCK_LF_SRC_RC on nRF52 is rc_ctiv=16 and rc_temp_ctiv=2. This will ensure calibration at least once every 8 seconds and for temperature changes of 0.5 degrees Celsius every 4 seconds. See the Product Specification for the nRF52 device being used for more information.
	Accuracy   uint8 // External clock accuracy used in the LL to compute timing windows, see @ref NRF_CLOCK_LF_ACCURACY.
}

// ::sd_softdevice_enable
func EnableSoftdevice(clock_lf_cfg *ClockLFCfg, fault_handler unsafe.Pointer) uint32

// ::sd_softdevice_disable
func DisableSoftdevice() uint32

// ::sd_softdevice_is_enabled
func IsEnabledSoftdevice(softdevice_enabled *uint8) uint32

// ::sd_softdevice_vector_table_base_set
func SetSoftdeviceVectorTableBase(address uint32) uint32
