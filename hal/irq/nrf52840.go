// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

//go:build nrf52840

// Package irq provides the list of supported external interrupts.
package irq

import "embedded/rtos"

const (
	POWER_CLOCK                       rtos.IRQ = 0
	RADIO                             rtos.IRQ = 1
	UARTE0_UART0                      rtos.IRQ = 2
	SPIM0_SPIS0_TWIM0_TWIS0_SPI0_TWI0 rtos.IRQ = 3
	SPIM1_SPIS1_TWIM1_TWIS1_SPI1_TWI1 rtos.IRQ = 4
	NFCT                              rtos.IRQ = 5
	GPIOTE                            rtos.IRQ = 6
	SAADC                             rtos.IRQ = 7
	TIMER0                            rtos.IRQ = 8
	TIMER1                            rtos.IRQ = 9
	TIMER2                            rtos.IRQ = 10
	RTC0                              rtos.IRQ = 11
	TEMP                              rtos.IRQ = 12
	RNG                               rtos.IRQ = 13
	ECB                               rtos.IRQ = 14
	CCM_AAR                           rtos.IRQ = 15
	WDT                               rtos.IRQ = 16
	RTC1                              rtos.IRQ = 17
	QDEC                              rtos.IRQ = 18
	COMP_LPCOMP                       rtos.IRQ = 19
	SWI0_EGU0                         rtos.IRQ = 20
	SWI1_EGU1                         rtos.IRQ = 21
	SWI2_EGU2                         rtos.IRQ = 22
	SWI3_EGU3                         rtos.IRQ = 23
	SWI4_EGU4                         rtos.IRQ = 24
	SWI5_EGU5                         rtos.IRQ = 25
	TIMER3                            rtos.IRQ = 26
	TIMER4                            rtos.IRQ = 27
	PWM0                              rtos.IRQ = 28
	PDM                               rtos.IRQ = 29
	MWU                               rtos.IRQ = 32
	PWM1                              rtos.IRQ = 33
	PWM2                              rtos.IRQ = 34
	SPIM2_SPIS2_SPI2                  rtos.IRQ = 35
	RTC2                              rtos.IRQ = 36
	I2S                               rtos.IRQ = 37
	FPU                               rtos.IRQ = 38
	USBD                              rtos.IRQ = 39
	UARTE1                            rtos.IRQ = 40
	QSPI                              rtos.IRQ = 41
	CRYPTOCELL                        rtos.IRQ = 42
	PWM3                              rtos.IRQ = 45
	SPIM3                             rtos.IRQ = 47
)
