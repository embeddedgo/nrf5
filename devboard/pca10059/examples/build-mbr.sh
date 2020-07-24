#!/bin/sh

GOTARGET=nrf52840
GOTEXT=0x00001000
GOMEM=0x20000008:262136
GOOUT=hex

. ../../../../../scripts/build.sh $@

bootversion=1
appversion=1

. ../../../../../scripts/nrf5-settings.sh
