#!/bin/sh

export GOTARGET=nrf52840
export GOTEXT=0x00000000
export GOMEM=0x20000000:256K
export GOSTRIPFN=1

emgo build $@
