#!/bin/sh

export GOTARGET=nrf52840
export GOTEXT=0x00001000
export GOMEM=0x20000008:262136
export GOSTRIPFN=1
export GOOUT=hex

emgo build $@

bootversion=1
appversion=1

. $(emgo env GOROOT)/../scripts/nrf5-settings.sh
