#!/bin/sh

set -e

cd ../../../embeddedgo/nrf5/hal
hal=$(pwd)
cd ../p
rm -rf *

svdxgen github.com/embeddedgo/nrf5/p ../svd/*.svd

for p in ficr gpio nvmc ppi rtc spi uart uarte uicr; do
	cd $p
	xgen *.go
	GOOS=noos GOARCH=thumb go build -tags nrf52840
	cd ..
done

perlscript='
s/package irq/$&\n\nimport "embedded\/rtos"/;
s/ = \d/ rtos.IRQ$&/g;
'

cd $hal/irq
rm -f *
cp ../../p/irq/* .
perl -pi -e "$perlscript" *.go