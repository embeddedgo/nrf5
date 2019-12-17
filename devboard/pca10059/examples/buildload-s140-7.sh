#!/bin/sh

GOTARGET=nrf52
GOTEXT=0x27000
GOMEM=0x20001678:256392
IRQNAMES=../../../../hal/irq
OUT=hex

. ../../../../../scripts/build.sh $@

dfu=usb-serial
port=/dev/ttyACM0

hw=52

sdreq=0xCA

s140_nrf52_7_0_1_softdevice_hex=0xCA

. ../../../../../scripts/load-nrfutil.sh $@
