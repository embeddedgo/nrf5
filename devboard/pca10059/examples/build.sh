#!/bin/sh

GOTARGET=nrf52
GOTEXT=0
GOMEM=0x20000000:256K

. ../../../../../scripts/build.sh $@
