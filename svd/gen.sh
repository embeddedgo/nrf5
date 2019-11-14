#!/bin/sh

set -e

cd ../../../embeddedgo/nrf5/p
rm -rf *

svdxgen github.com/embeddedgo/nrf5/p ../svd/*.svd
