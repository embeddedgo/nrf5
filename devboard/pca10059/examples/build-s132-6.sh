#!/bin/sh

echo 'For nRF52840 use build-s140-7.sh'
exit 1

export GOTARGET=nrf52840
export GOTEXT=0x00026000
export GOMEM=0x20002000:253952
export GOSTRIPFN=1
export GOOUT=hex

emgo build $@

bootversion=1
appversion=1

. $(emgo env GOROOT)/../scripts/nrf5-settings.sh
