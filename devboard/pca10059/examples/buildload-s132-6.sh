#!/bin/sh

GOTARGET=nrf52
GOTEXT=0x26000
GOMEM=0x20001628:256408
IRQNAMES=../../../../hal/irq
OUT=hex

. ../../../../../scripts/build.sh $@

dfu=usb-serial
port=/dev/ttyACM0

hw=52

sdreq=0xA8,0xAF,0xB7

s132_nrf52_6_0_0_softdevice_hex=0xA8
s132_nrf52_6_1_0_softdevice_hex=0xAF
s132_nrf52_6_1_1_softdevice_hex=0xB7

. ../../../../../scripts/load-nrfutil.sh $@
