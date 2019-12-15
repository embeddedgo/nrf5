#!/bin/sh

GOTARGET=nrf52
GOTEXT=0x26000
GOMEM=0x20001628:256408

. ../../../../../scripts/build.sh $@
