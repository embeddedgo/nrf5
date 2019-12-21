#!/bin/sh

GOTARGET=nrf52
GOTEXT=0x27000
GOMEM=0x20001678:256392
IRQNAMES=../../../../hal/irq
OUT=hex

. ../../../../../scripts/build.sh $@

family=NRF52840
bootversion=1
appversion=1

. ../../../../../scripts/nrf5-settings.sh