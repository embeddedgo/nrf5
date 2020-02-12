// Copyright 2020 Michal Derkacz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embedded/rtos"

	"github.com/embeddedgo/nrf5/hal/clock"
	"github.com/embeddedgo/nrf5/hal/irq"
	"github.com/embeddedgo/nrf5/softdevice/s140/sd"
	"github.com/embeddedgo/nrf5/softdevice/s140/sd/ble"
	"github.com/embeddedgo/nrf5/softdevice/s140/sd/ble/gap"
	"github.com/embeddedgo/nrf5/softdevice/s140/sd/ble/gatt"
	"github.com/embeddedgo/nrf5/softdevice/s140/sd/ble/gatt/gatts"
	"github.com/embeddedgo/nrf5/softdevice/s140/sd/sdm"
	"github.com/embeddedgo/nrf5/softdevice/sdutil"

	"github.com/embeddedgo/nrf5/devboard/pca10059/board/buttons"
	"github.com/embeddedgo/nrf5/devboard/pca10059/board/leds"
)

var newEvent rtos.Note

func main() {
	clock.StoreTRACECONFIG(clock.T4MHz, clock.Serial) // enable SWO on P1.00

	print("Setup BLE stack:\n")

	print("- enable softdevice: ")

	lfclkc := sdm.LFCLKC{
		Source:   sdm.LFCLKXTAL,
		Accuracy: sdm.LFCLK20ppm,
	}
	checkStatus(sdm.EnableSoftdevice(&lfclkc, sdFaultHandler))

	// enable SoftDevice event interrupt
	irq.SWI2_EGU2.Enable(rtos.IntPrioLowest)

	appBase := sdutil.AppRAMBase()

	print("- configure connection count: ")

	connCfg := gap.ConnCfg{
		Tag:         1,
		ConnCount:   1,
		EventLength: 6, // in 1.25 ms units
	}
	checkStatus(gap.SetConnCfg(&connCfg, appBase))

	print("- configure role count: ")

	roleCntCfg := gap.RoleCountCfg{
		AdvSetCount:     1,
		PeriphRoleCount: 1,
	}
	checkStatus(gap.SetRoleCountCfg(&roleCntCfg, appBase))

	print("- enable BLE stack: ")

	minBase, err := ble.Enable(appBase)
	status(err)
	ramStart := uintptr(0x2000_0000)
	print("- softdevice RAM: required=", minBase-ramStart, ", reserved=",
		appBase-ramStart, "\n")
	if err != nil {
		blink(leds.Red)
	}

	print("- configure device name: ")

	devName := "Go on nRF52"
	checkStatus(gap.SetDeviceName(gap.SecModeOpen, devName))

	print("- configure connection parameters: ")

	connParams := gap.ConnParams{
		MinConnInterval: 80,  // in 1.25 ms units
		MaxConnInterval: 160, // in 1.25 ms units
		ConnSupTimeout:  400, // in 10 ms units
	}
	checkStatus(gap.SetPPCP(&connParams))

	// TODO: set MTU (see ble_app_uart)

	print("- add LED-Button service base UUID: ")

	uuidBase := ble.UUID128{0x23, 0xD1, 0xBC, 0xEA, 0x5F, 0x78, 0x23, 0x15,
		0xDE, 0xEF, 0x12, 0x12, 0x00, 0x00, 0x00, 0x00}
	uuidType, err := ble.AddVSUUID(&uuidBase)
	checkStatus(err)

	print("- add LED-Button service: ")

	serviceUUID := &ble.UUID{0x1523, uuidType}
	serviceHandle, err := gatts.AddService(gatts.PrimaryService, serviceUUID)
	checkStatus(err)

	print("- add Button characteristic: ")

	meta := gatts.CharMeta{
		Props: gatt.PropRead | gatt.PropNotify,
		CCCDMeta: &gatts.AttrMeta{
			ReadPerm:  gap.SecModeOpen,
			WritePerm: gap.SecModeOpen,
			Options:   gatts.ValStack,
		},
	}
	val := gatts.Attr{
		UUID: &ble.UUID{0x1524, uuidType},
		Meta: &gatts.AttrMeta{
			ReadPerm: gap.SecModeOpen,
			Options:  gatts.ValStack,
		},
		MaxLen: 1, // one byte
	}
	btnHandles, err := gatts.AddCharacteristic(serviceHandle, &meta, &val)
	checkStatus(err)

	print("- add LED characteristic: ")

	meta = gatts.CharMeta{Props: gatt.PropRead | gatt.PropWrite}
	val = gatts.Attr{
		UUID: &ble.UUID{0x1525, uuidType},
		Meta: &gatts.AttrMeta{
			ReadPerm:  gap.SecModeOpen,
			WritePerm: gap.SecModeOpen,
			Options:   gatts.ValStack,
		},
		MaxLen: 1, // one byte
	}
	ledHandles, err := gatts.AddCharacteristic(serviceHandle, &meta, &val)
	checkStatus(err)

	print("- create advertising set: ")

	adata := sdutil.NewAdvData(gap.MaxAdvSetDataSize)
	adata.AppendBytes(gap.Flags, byte(gap.LEOnly|gap.GeneralDiscMode))
	adata.AppendString(gap.CompleteLocalName, devName)

	srdata := sdutil.NewAdvData(gap.MaxAdvSetDataSize)
	srdata.AppendUUIDs(gap.CompleteServiceUUID16, 2, serviceUUID)
	srdata.AppendUUIDs(gap.CompleteServiceUUID128, 16, serviceUUID)

	advData := gap.AdvData{Data: adata.Data(), ScanRspData: srdata.Data()}

	advParams := gap.AdvParams{
		PrimaryPHY: gap.PHY1Mbps,
		Props:      gap.AdvProps{Type: gap.ConnectableScannableUndirected},
		Interval:   64, // in 625 us units
	}

	advHandle, err := gap.ConfigureAdvSet(gap.NewAdvSet, &advData, &advParams)
	checkStatus(err)

	print("- start advertising: ")

	checkStatus(gap.StartAdv(advHandle, connCfg.Tag))

	print("Handling BLE stack events...\n")

	statusLED := leds.Green
	btn := buttons.User.Read()
	btnParams := gatts.HVXParams{btnHandles.Value, gatt.Notification, 0, 1}
	connHandle := ble.ConnInvalid

	for {
		if !newEvent.Sleep(250e6) {
			if b := buttons.User.Read(); b != btn {
				btn = b
				if connHandle != ble.ConnInvalid {
					_, err := gatts.HVX(connHandle, &btnParams, byte(b))
					if err == gatts.ErrSysAttrMissing {
						println(err.Error())
					} else {
						dieErr(err)
					}
				}
			}
			statusLED.Set(statusLED.Get() + 1)
			continue
		}
		for {
			newEvent.Clear()
			evt, err := sdutil.NextEvent()
			if err != nil {
				if err != sd.ErrNotFound {
					dieErr(err)
				}
				break
			}

			println("event id:", evt.EventID())

			switch e := evt.(type) {
			case *gap.Connected:
				connHandle = e.ConnHandle
				statusLED.SetOff()
				statusLED = leds.Blue

			case *gap.Disconnected:
				connHandle = ble.ConnInvalid
				statusLED.SetOff()
				dieErr(gap.StartAdv(advHandle, connCfg.Tag))
				statusLED = leds.Green

			case *gatts.Write:
				data := e.Data()
				switch e.Handle {
				case ledHandles.Value:
					if len(data) == 1 {
						leds.User.Set(int(data[0]))
					}
				}
			}
		}
	}
}

//go:interrupthandler
func SWI2_EGU2_Handler() {
	newEvent.Wakeup()
}

func sdFaultHandler(id, pc, info uint32) {
	print("softdevice fault: id=", id, " pc=", pc, " info=", info, "\n")
}

//// utils

func blink(led leds.LED) {
	for {
		led.SetOn()
		rtos.Nanosleep(50e6)
		led.SetOff()
		rtos.Nanosleep(950e6)
	}
}

func status(err error) bool {
	if err == nil {
		print("OK\n")
		return true
	}
	println(err.Error())
	return false
}

func checkStatus(err error) {
	if status(err) {
		return
	}
	blink(leds.Red)
}

func dieErr(err error) {
	if err == nil {
		return
	}
	println("error:", err.Error())
	blink(leds.Red)
}
