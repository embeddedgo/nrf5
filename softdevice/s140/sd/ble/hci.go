// Based on C code by Nordic Semiconductor ASA.
// See LICENSE-NORDIC for original C code license.

// Copyright 2020 Michal Derkacz.

package ble

type HCIStatus uint8

const (
	HCISuccess                               HCIStatus = 0x00
	HCIUnknownBTLECommand                    HCIStatus = 0x01
	HCIUnknownConnectionIdentifier           HCIStatus = 0x02
	HCIAuthenticationFailure                 HCIStatus = 0x05
	HCIPinOrKeyMissing                       HCIStatus = 0x06
	HCIMemoryCapacityExceeded                HCIStatus = 0x07
	HCIConnectionTimeout                     HCIStatus = 0x08
	HCICommandDisallowed                     HCIStatus = 0x0C
	HCIInvalidBTLECommandParameters          HCIStatus = 0x12
	HCIRemoteUserTerminatedConnection        HCIStatus = 0x13
	HCIRemoteDevTerminationDueToLowResources HCIStatus = 0x14
	HCIRemoteDevTerminationDUETOPOWEROFF     HCIStatus = 0x15
	HCILocalHostTerminatedConnection         HCIStatus = 0x16
	HCIUnsupportedRemoteFeature              HCIStatus = 0x1A
	HCIInvalidLMPParameters                  HCIStatus = 0x1E
	HCIUnspecifiedError                      HCIStatus = 0x1F
	HCILMPResponseTimeout                    HCIStatus = 0x22
	HCILMPErrorTransactionCollision          HCIStatus = 0x23
	HCILMPPDUNotAllowed                      HCIStatus = 0x24
	HCIInstantPassed                         HCIStatus = 0x28
	HCIPairingWithUnitKeyUnsupported         HCIStatus = 0x29
	HCIDifferentTransactionCollision         HCIStatus = 0x2A
	HCIParameterOutOfMandatoryRange          HCIStatus = 0x30
	HCIControllerBusy                        HCIStatus = 0x3A
	HCIConnIntervalUnacceptable              HCIStatus = 0x3B
	HCIDirectedAdvertiserTimeout             HCIStatus = 0x3C
	HCIConnTerminatedDueToMICFailure         HCIStatus = 0x3D
	HCIConnFailedToBeEstablished             HCIStatus = 0x3E
)
