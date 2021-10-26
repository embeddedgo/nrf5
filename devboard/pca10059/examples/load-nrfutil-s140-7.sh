#!/bin/sh

dfu=usb-serial
port=/dev/ttyACM0

hw=52

sdreq=0xCA

s140_nrf52_7_0_1_softdevice_hex=0xCA
s140_nrf52_7_2_0_softdevice_hex=0x100

. ../../../../../scripts/load-nrfutil.sh $@
