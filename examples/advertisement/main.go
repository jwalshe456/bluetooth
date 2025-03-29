package main

import (
	"time"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

func main() {
	must("enable BLE stack", adapter.Enable())
	adv := adapter.DefaultAdvertisement()
	serviceUUID := bluetooth.New32BitUUID(1900)
	userUUID := "11111900-0000-1000-8000-00805f9b34fb"
	must("config adv", adv.Configure(bluetooth.AdvertisementOptions{
		LocalName: "Go",
		ServiceUUIDs: []bluetooth.UUID{
			serviceUUID,
		},
		ServiceData: []bluetooth.ServiceDataElement{
			{UUID: serviceUUID, Data: []byte(userUUID)},
		},
	}))
	must("start adv", adv.Start())

	println("advertising...")
	address, _ := adapter.Address()
	for {
		println("Go Bluetooth /", address.MAC.String())
		time.Sleep(time.Second)
	}
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}
