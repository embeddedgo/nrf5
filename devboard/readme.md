## Support for nRF5 development boards

### Directory structure

Every board directory contains a set of packages (in *board* subdirectory) that provides the interface to the peripherals available on the board (for now the support is modest: only LEDs and buttons).

The board/init package, when imported, configures the whole system for typical usage. If you use any other package from *board* directory the board/init package is imported implicitly to ensure the board is properly configured.

The *examples* subdirectory as the name suggests includes sample code, but also scripts and configuration that help to build, load and debug.

There is also *doc* subdirectory that contain useful information and other resources about this development board.

### Supported boards

[pca10059](pca10059): Nordic nRF52840-Dongle (PCA10059), [website](https://www.nordicsemi.com/Software-and-tools/Development-Kits/nRF52840-Dongle)

![Nordic nRF52840-Dongle](pca10059/doc/board.jpg)

