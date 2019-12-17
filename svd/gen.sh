#!/bin/sh

#set -e

cd ../../../embeddedgo/nrf5/p
rm -rf *

svdxgen github.com/embeddedgo/nrf5/p ../svd/*.svd

#for p in gpio; do
#	cd $p
#	xgen *.go
#	cd ..
#done

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

