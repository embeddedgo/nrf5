// DO NOT EDIT THIS FILE. GENERATED BY svdxgen.

// +build nrf52840

// Package ppi provides access to the registers of the PPI peripheral.
//
// Instances:
//  PPI  PPI_BASE  -  -  Programmable Peripheral Interconnect
// Registers:
//  0x000 32  TASK_CHG{EN,DIS}[6]  Enable channel group n; Disable channel group n
//  0x500 32  CHEN                 Channel enable register
//  0x504 32  CHENSET              Channel enable set register
//  0x508 32  CHENCLR              Channel enable clear register
//  0x510 32  CH{EEP,TEP}[20]      Channel n event end-point; Channel n task end-point
//  0x800 32  CHG[6]               Channel group n
//  0x910 32  FORK_TEP[32]         Channel n task end-point
// Import:
//  github.com/embeddedgo/nrf5/p/mmap
package ppi
