#!/bin/sh

#set -e

cd ../../../embeddedgo/nrf5/p
rm -rf *

svdxgen github.com/embeddedgo/nrf5/p ../svd/*.svd

for p in ficr gpio uicr; do
	cd $p
	xgen *.go
	#GOOS=noos GOARCH=thumb go build -tags nrf52840
	cd ..
done

rm -f ../hal/irq/*

awkscript='{
	gsub("package irq", "package irq\n\nimport \"embedded/rtos\"", $0)
	gsub(" = ", " rtos.IRQ = ", $0)
	print
}'
cd irq
for f in *; do
	awk "$awkscript" $f >../../hal/irq/$f
done

