package main

import (
	"time"

	"tinygo.org/x/espradio"
)

func main() {
	time.Sleep(2 * time.Second)
	println("Initializing Wi-Fi...")
	if err := espradio.Enable(espradio.Config{}); err != nil {
		println("Enable error:", err.Error())
		return
	}

	if err := espradio.Start(); err != nil {
		println("Start error:", err.Error())
		return
	}

	for {
		println()
		println("Scanning Wi-Fi...")

		aps, err := espradio.Scan()
		if err != nil {
			println("Scan error:", err.Error())
		} else {
			println("Found:", len(aps))

			for _, ap := range aps {
				println("SSID:", ap.SSID, "| RSSI:", ap.RSSI, "dBm")
			}
		}

		time.Sleep(5 * time.Second)
	}
}
