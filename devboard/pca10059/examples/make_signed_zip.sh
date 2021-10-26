#!/bin/sh

appversion=1
key=../ble_bootloader.key

name=$(basename $(pwd))

nrfutil pkg generate --hw-version 52 --sd-req 0xCA,0x0100 --key-file $key --application $name.hex --application-version $appversion $name.zip