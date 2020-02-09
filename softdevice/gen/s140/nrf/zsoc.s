// Code generated from ../nordic/s140/headers/nrf_soc.h; DO NOT EDIT.

#include "go_asm.h"
#include "textflag.h"

TEXT ·NewMutex(SB),NOSPLIT|NOFRAME,$0
	MOVW mutex+0(FP), R0
	SWI $const__MUTEX_NEW
	MOVW R0, ret+4(FP)
	RET

TEXT ·AcquireMutex(SB),NOSPLIT|NOFRAME,$0
	MOVW mutex+0(FP), R0
	SWI $const__MUTEX_ACQUIRE
	MOVW R0, ret+4(FP)
	RET

TEXT ·ReleaseMutex(SB),NOSPLIT|NOFRAME,$0
	MOVW mutex+0(FP), R0
	SWI $const__MUTEX_RELEASE
	MOVW R0, ret+4(FP)
	RET

TEXT ·GetRandApplicationPoolCapacity(SB),NOSPLIT|NOFRAME,$0
	MOVW pool_capacity+0(FP), R0
	SWI $const__RAND_APPLICATION_POOL_CAPACITY_GET
	MOVW R0, ret+4(FP)
	RET

TEXT ·GetRandApplicationBytesAvailable(SB),NOSPLIT|NOFRAME,$0
	MOVW bytes_available+0(FP), R0
	SWI $const__RAND_APPLICATION_BYTES_AVAILABLE_GET
	MOVW R0, ret+4(FP)
	RET

TEXT ·GetRandApplicationVector(SB),NOSPLIT|NOFRAME,$0
	MOVW buff+0(FP), R0
	MOVBU length+4(FP), R1
	SWI $const__RAND_APPLICATION_VECTOR_GET
	MOVW R0, ret+8(FP)
	RET

TEXT ·GetPowerResetReason(SB),NOSPLIT|NOFRAME,$0
	MOVW reset_reason+0(FP), R0
	SWI $const__POWER_RESET_REASON_GET
	MOVW R0, ret+4(FP)
	RET

TEXT ·ClrPowerResetReason(SB),NOSPLIT|NOFRAME,$0
	MOVW reset_reason_clr_msk+0(FP), R0
	SWI $const__POWER_RESET_REASON_CLR
	MOVW R0, ret+4(FP)
	RET

TEXT ·SetPowerMode(SB),NOSPLIT|NOFRAME,$0
	MOVBU power_mode+0(FP), R0
	SWI $const__POWER_MODE_SET
	MOVW R0, ret+4(FP)
	RET

TEXT ·OffPowerSystem(SB),NOSPLIT|NOFRAME,$0
	SWI $const__POWER_SYSTEM_OFF
	MOVW R0, ret+0(FP)
	RET

TEXT ·EnablePowerPof(SB),NOSPLIT|NOFRAME,$0
	MOVBU pof_enable+0(FP), R0
	SWI $const__POWER_POF_ENABLE
	MOVW R0, ret+4(FP)
	RET

TEXT ·EnablePowerUsbpwrrdy(SB),NOSPLIT|NOFRAME,$0
	MOVBU usbpwrrdy_enable+0(FP), R0
	SWI $const__POWER_USBPWRRDY_ENABLE
	MOVW R0, ret+4(FP)
	RET

TEXT ·EnablePowerUsbdetected(SB),NOSPLIT|NOFRAME,$0
	MOVBU usbdetected_enable+0(FP), R0
	SWI $const__POWER_USBDETECTED_ENABLE
	MOVW R0, ret+4(FP)
	RET

TEXT ·EnablePowerUsbremoved(SB),NOSPLIT|NOFRAME,$0
	MOVBU usbremoved_enable+0(FP), R0
	SWI $const__POWER_USBREMOVED_ENABLE
	MOVW R0, ret+4(FP)
	RET

TEXT ·GetPowerUsbregstatus(SB),NOSPLIT|NOFRAME,$0
	MOVW usbregstatus+0(FP), R0
	SWI $const__POWER_USBREGSTATUS_GET
	MOVW R0, ret+4(FP)
	RET

TEXT ·SetPowerPofThreshold(SB),NOSPLIT|NOFRAME,$0
	MOVBU threshold+0(FP), R0
	SWI $const__POWER_POF_THRESHOLD_SET
	MOVW R0, ret+4(FP)
	RET

TEXT ·SetPowerPofThresholdvddh(SB),NOSPLIT|NOFRAME,$0
	MOVBU threshold+0(FP), R0
	SWI $const__POWER_POF_THRESHOLDVDDH_SET
	MOVW R0, ret+4(FP)
	RET

TEXT ·SetPowerRamPower(SB),NOSPLIT|NOFRAME,$0
	MOVBU index+0(FP), R0
	MOVW ram_powerset+4(FP), R1
	SWI $const__POWER_RAM_POWER_SET
	MOVW R0, ret+8(FP)
	RET

TEXT ·ClrPowerRamPower(SB),NOSPLIT|NOFRAME,$0
	MOVBU index+0(FP), R0
	MOVW ram_powerclr+4(FP), R1
	SWI $const__POWER_RAM_POWER_CLR
	MOVW R0, ret+8(FP)
	RET

TEXT ·GetPowerRamPower(SB),NOSPLIT|NOFRAME,$0
	MOVBU index+0(FP), R0
	MOVW ram_power+4(FP), R1
	SWI $const__POWER_RAM_POWER_GET
	MOVW R0, ret+8(FP)
	RET

TEXT ·SetPowerGpregret(SB),NOSPLIT|NOFRAME,$0
	MOVW gpregret_id+0(FP), R0
	MOVW gpregret_msk+4(FP), R1
	SWI $const__POWER_GPREGRET_SET
	MOVW R0, ret+8(FP)
	RET

TEXT ·ClrPowerGpregret(SB),NOSPLIT|NOFRAME,$0
	MOVW gpregret_id+0(FP), R0
	MOVW gpregret_msk+4(FP), R1
	SWI $const__POWER_GPREGRET_CLR
	MOVW R0, ret+8(FP)
	RET

TEXT ·GetPowerGpregret(SB),NOSPLIT|NOFRAME,$0
	MOVW gpregret_id+0(FP), R0
	MOVW gpregret+4(FP), R1
	SWI $const__POWER_GPREGRET_GET
	MOVW R0, ret+8(FP)
	RET

TEXT ·SetPowerDcdcMode(SB),NOSPLIT|NOFRAME,$0
	MOVBU dcdc_mode+0(FP), R0
	SWI $const__POWER_DCDC_MODE_SET
	MOVW R0, ret+4(FP)
	RET

TEXT ·SetPowerDcdc0Mode(SB),NOSPLIT|NOFRAME,$0
	MOVBU dcdc_mode+0(FP), R0
	SWI $const__POWER_DCDC0_MODE_SET
	MOVW R0, ret+4(FP)
	RET

TEXT ·RequestClockHF(SB),NOSPLIT|NOFRAME,$0
	SWI $const__CLOCK_HFCLK_REQUEST
	MOVW R0, ret+0(FP)
	RET

TEXT ·ReleaseClockHF(SB),NOSPLIT|NOFRAME,$0
	SWI $const__CLOCK_HFCLK_RELEASE
	MOVW R0, ret+0(FP)
	RET

TEXT ·IsRunningClockHF(SB),NOSPLIT|NOFRAME,$0
	MOVW is_running+0(FP), R0
	SWI $const__CLOCK_HFCLK_IS_RUNNING
	MOVW R0, ret+4(FP)
	RET

TEXT ·WaitAppEvt(SB),NOSPLIT|NOFRAME,$0
	SWI $const__APP_EVT_WAIT
	MOVW R0, ret+0(FP)
	RET

TEXT ·GetPPIChannelEnable(SB),NOSPLIT|NOFRAME,$0
	MOVW channel_enable+0(FP), R0
	SWI $const__PPI_CHANNEL_ENABLE_GET
	MOVW R0, ret+4(FP)
	RET

TEXT ·SetPPIChannelEnable(SB),NOSPLIT|NOFRAME,$0
	MOVW channel_enable_set_msk+0(FP), R0
	SWI $const__PPI_CHANNEL_ENABLE_SET
	MOVW R0, ret+4(FP)
	RET

TEXT ·ClrPPIChannelEnable(SB),NOSPLIT|NOFRAME,$0
	MOVW channel_enable_clr_msk+0(FP), R0
	SWI $const__PPI_CHANNEL_ENABLE_CLR
	MOVW R0, ret+4(FP)
	RET

TEXT ·AssignPPIChannel(SB),NOSPLIT|NOFRAME,$0
	MOVBU channel_num+0(FP), R0
	MOVW evt_endpoint+4(FP), R1
	MOVW task_endpoint+8(FP), R2
	SWI $const__PPI_CHANNEL_ASSIGN
	MOVW R0, ret+12(FP)
	RET

TEXT ·EnablePPIGroupTask(SB),NOSPLIT|NOFRAME,$0
	MOVBU group_num+0(FP), R0
	SWI $const__PPI_GROUP_TASK_ENABLE
	MOVW R0, ret+4(FP)
	RET

TEXT ·DisablePPIGroupTask(SB),NOSPLIT|NOFRAME,$0
	MOVBU group_num+0(FP), R0
	SWI $const__PPI_GROUP_TASK_DISABLE
	MOVW R0, ret+4(FP)
	RET

TEXT ·AssignPPIGroup(SB),NOSPLIT|NOFRAME,$0
	MOVBU group_num+0(FP), R0
	MOVW channel_msk+4(FP), R1
	SWI $const__PPI_GROUP_ASSIGN
	MOVW R0, ret+8(FP)
	RET

TEXT ·GetPPIGroup(SB),NOSPLIT|NOFRAME,$0
	MOVBU group_num+0(FP), R0
	MOVW channel_msk+4(FP), R1
	SWI $const__PPI_GROUP_GET
	MOVW R0, ret+8(FP)
	RET

TEXT ·SetRadioNotificationCfg(SB),NOSPLIT|NOFRAME,$0
	MOVBU typ+0(FP), R0
	MOVBU distance+1(FP), R1
	SWI $const__RADIO_NOTIFICATION_CFG_SET
	MOVW R0, ret+4(FP)
	RET

TEXT ·EncryptEcbBlock(SB),NOSPLIT|NOFRAME,$0
	MOVW ecb_data+0(FP), R0
	SWI $const__ECB_BLOCK_ENCRYPT
	MOVW R0, ret+4(FP)
	RET

TEXT ·EncryptEcbBlocks(SB),NOSPLIT|NOFRAME,$0
	MOVBU block_count+0(FP), R0
	MOVW data_blocks+4(FP), R1
	SWI $const__ECB_BLOCKS_ENCRYPT
	MOVW R0, ret+8(FP)
	RET

TEXT ·GetEvt(SB),NOSPLIT|NOFRAME,$0
	MOVW evt_id+0(FP), R0
	SWI $const__EVT_GET
	MOVW R0, ret+4(FP)
	RET

TEXT ·GetTemp(SB),NOSPLIT|NOFRAME,$0
	MOVW temp+0(FP), R0
	SWI $const__TEMP_GET
	MOVW R0, ret+4(FP)
	RET

TEXT ·WriteFlash(SB),NOSPLIT|NOFRAME,$0
	MOVW dst+0(FP), R0
	MOVW src+4(FP), R1
	MOVW size+8(FP), R2
	SWI $const__FLASH_WRITE
	MOVW R0, ret+12(FP)
	RET

TEXT ·EraseFlashPage(SB),NOSPLIT|NOFRAME,$0
	MOVW page_number+0(FP), R0
	SWI $const__FLASH_PAGE_ERASE
	MOVW R0, ret+4(FP)
	RET

TEXT ·OpenRadioSession(SB),NOSPLIT|NOFRAME,$0
	MOVW p_radio_signal_callback+0(FP), R0
	SWI $const__RADIO_SESSION_OPEN
	MOVW R0, ret+4(FP)
	RET

TEXT ·CloseRadioSession(SB),NOSPLIT|NOFRAME,$0
	SWI $const__RADIO_SESSION_CLOSE
	MOVW R0, ret+0(FP)
	RET

TEXT ·RequestRadio(SB),NOSPLIT|NOFRAME,$0
	MOVW request+0(FP), R0
	SWI $const__RADIO_REQUEST
	MOVW R0, ret+4(FP)
	RET

TEXT ·WriteProtectedRegister(SB),NOSPLIT|NOFRAME,$0
	MOVW register+0(FP), R0
	MOVW value+4(FP), R1
	SWI $const__PROTECTED_REGISTER_WRITE
	MOVW R0, ret+8(FP)
	RET
