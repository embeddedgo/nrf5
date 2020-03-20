#!/bin/sh

GOTARGET=nrf52840
GOTEXT=0x00000000
GOMEM=0x20000000:256K

. ../../../../../scripts/build.sh $@
