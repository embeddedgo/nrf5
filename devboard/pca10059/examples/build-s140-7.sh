#!/bin/sh

export GOTARGET=nrf52840
export GOTEXT=0x00027000
export GOMEM=0x20002800:246K
export GOSTRIPFN=1
export GOOUT=hex

emgo build $@

bootversion=1
appversion=1

. $(emgo env GOROOT)/../scripts/nrf5-settings.sh
