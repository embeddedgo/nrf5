#!/bin/sh

echo 'For nRF52840 use build-s140-*'
exit 1

GOTARGET=nrf52840
GOTEXT=0x00026000
GOMEM=0x20002000:253952
GOOUT=hex

. $(emgo env GOROOT)/../scripts/build.sh $@

bootversion=1
appversion=1

. $(emgo env GOROOT)/../scripts/nrf5-settings.sh
