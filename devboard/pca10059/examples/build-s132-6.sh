#!/bin/sh

GOTARGET=nrf52840
GOTEXT=0x00026000
GOMEM=0x20002000:253952
IRQNAMES=../../../../hal/irq
OUT=hex

. ../../../../../scripts/build.sh $@

bootversion=1
appversion=1

. ../../../../../scripts/nrf5-settings.sh
