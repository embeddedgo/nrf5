#!/bin/sh

GOTARGET=nrf52840
GOTEXT=0x00027000
GOMEM=0x20002800:246K
GOOUT=hex

. ../../../../../scripts/build.sh $@

bootversion=1
appversion=1

. ../../../../../scripts/nrf5-settings.sh
