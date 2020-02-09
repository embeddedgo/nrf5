// Code generated from ../nordic/s140/headers/nrf_error.h; DO NOT EDIT.

package nrf

const ERROR_BASE_NUM = (0x0)                              // Global error base
const ERROR_SDM_BASE_NUM = (0x1000)                       // SDM error base
const ERROR_SOC_BASE_NUM = (0x2000)                       // SoC error base
const ERROR_STK_BASE_NUM = (0x3000)                       // STK error base
const SUCCESS = (ERROR_BASE_NUM + 0)                      // Successful command
const ERROR_SVC_HANDLER_MISSING = (ERROR_BASE_NUM + 1)    // SVC handler is missing
const ERROR_SOFTDEVICE_NOT_ENABLED = (ERROR_BASE_NUM + 2) // SoftDevice has not been enabled
const ERROR_INTERNAL = (ERROR_BASE_NUM + 3)               // Internal Error
const ERROR_NO_MEM = (ERROR_BASE_NUM + 4)                 // No Memory for operation
const ERROR_NOT_FOUND = (ERROR_BASE_NUM + 5)              // Not found
const ERROR_NOT_SUPPORTED = (ERROR_BASE_NUM + 6)          // Not supported
const ERROR_INVALID_PARAM = (ERROR_BASE_NUM + 7)          // Invalid Parameter
const ERROR_INVALID_STATE = (ERROR_BASE_NUM + 8)          // Invalid state, operation disallowed in this state
const ERROR_INVALID_LENGTH = (ERROR_BASE_NUM + 9)         // Invalid Length
const ERROR_INVALID_FLAGS = (ERROR_BASE_NUM + 10)         // Invalid Flags
const ERROR_INVALID_DATA = (ERROR_BASE_NUM + 11)          // Invalid Data
const ERROR_DATA_SIZE = (ERROR_BASE_NUM + 12)             // Invalid Data size
const ERROR_TIMEOUT = (ERROR_BASE_NUM + 13)               // Operation timed out
const ERROR_NULL = (ERROR_BASE_NUM + 14)                  // Null Pointer
const ERROR_FORBIDDEN = (ERROR_BASE_NUM + 15)             // Forbidden Operation
const ERROR_INVALID_ADDR = (ERROR_BASE_NUM + 16)          // Bad Memory Address
const ERROR_BUSY = (ERROR_BASE_NUM + 17)                  // Busy
const ERROR_CONN_COUNT = (ERROR_BASE_NUM + 18)            // Maximum connection count exceeded.
const ERROR_RESOURCES = (ERROR_BASE_NUM + 19)             // Not enough resources for operation
