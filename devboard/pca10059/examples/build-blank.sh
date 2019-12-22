#!/bin/sh

GOTARGET=nrf52840
GOTEXT=0
GOMEM=0x20000000:256K
IRQNAMES=../../../../hal/irq

. ../../../../../scripts/build.sh $@
