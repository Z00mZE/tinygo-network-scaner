package main

import (
	"time"

	"tinygo.org/x/bluetooth"
)

func main() {
	//println("Initializing Wi-Fi...")
	//if err := espradio.Enable(espradio.Config{}); err != nil {
	//	println("Enable error:", err.Error())
	//	return
	//}

	time.Sleep(2 * time.Second)
	println(`BLE...`)
	adapter := bluetooth.DefaultAdapter
	if err := adapter.Enable(); err != nil {
		println("Start error:", err.Error())
		return
	}
	println("adapter.Enable(): true")

	if scanError := adapter.Scan(bleScanResult); scanError != nil {
		println("Scan error:", scanError.Error())
	}
	//println()
	//println("Scanning Wi-Fi...")
	//
	//for {
	//	aps, err := espradio.Scan()
	//	if err != nil {
	//		println("Scan error:", err.Error())
	//	} else {
	//		println("Found:", len(aps))
	//		for _, ap := range aps {
	//			println("SSID:", ap.SSID, "| RSSI:", ap.RSSI, "dBm")
	//		}
	//	}
	//	break
	//}

}

func bleScanResult(_ *bluetooth.Adapter, result bluetooth.ScanResult) {
	if result.LocalName() != "" {
		println("name:", result.LocalName(), "; RSSI:", result.RSSI, "; Addr: :", result.Address.String())
	}
}
