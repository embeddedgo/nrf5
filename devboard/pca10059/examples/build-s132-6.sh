#!/bin/sh

GOTARGET=nrf52
GOTEXT=0x26000
GOMEM=0x20001628:256408
IRQNAMES=../../../../hal/irq
OUT=hex

. ../../../../../scripts/build.sh $@
