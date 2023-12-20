#!/bin/sh

set -e

gdb_cmd=gdb
if command -v arm-none-eabi-gdb; then
	gdb_cmd=arm-none-eabi-gdb
elif command -v gdb-multiarch; then
	gdb_cmd=gdb-multiarch
fi

$gdb_cmd \
	-ex 'set pagination off' \
	-ex 'set confirm off' \
	-ex 'target extended-remote /dev/ttyACM0' \
	-ex 'monitor connect_srst enable' \
	-ex 'monitor swdp_scan' \
	-ex 'attach 1' \
	-ex "load $(basename $(pwd))-settings.hex" \
	-ex "load $(basename $(pwd)).hex" \
	-ex 'kill' \
	-ex 'quit'
