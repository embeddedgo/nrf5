#!/bin/sh

GOTARGET=nrf52840
GOTEXT=0x26000
GOMEM=0x20001628:256408
IRQNAMES=../../../../hal/irq
OUT=hex

. ../../../../../scripts/build.sh $@

bootversion=1
appversion=1

. ../../../../../scripts/nrf5-settings.sh