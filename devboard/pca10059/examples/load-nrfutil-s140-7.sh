#!/bin/sh

dfu=usb-serial
port=/dev/ttyACM0

hw=52

sdreq=0xCA,0x0100

s140_nrf52_7_0_1_softdevice_hex=0xCA
s140_nrf52_7_2_0_softdevice_hex=0x0100

. $(emgo env GOROOT)/../scripts/load-nrfutil.sh $@
