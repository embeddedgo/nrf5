#!/bin/sh

GOTARGET=nrf52840
GOTEXT=0x27000
GOMEM=0x20001678:256392
IRQNAMES=../../../../hal/irq
OUT=hex

. ../../../../../scripts/build.sh $@

bootversion=1
appversion=1

. ../../../../../scripts/nrf5-settings.sh