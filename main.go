package main // import "github.com/slintes/co2-sensor"

import (
	"machine"
	"time"

	"github.com/slintes/co2-sensor/pkg/led"
	"github.com/slintes/co2-sensor/pkg/wifi"
)

func main() {

	// for USB debugging
	machine.UART1.Configure(machine.UARTConfig{TX: machine.PA22, RX: machine.PA23})

	// flash a LED
	lights := led.NewLED()
	lights.Blink(led.GREEN, 500*time.Millisecond)
	time.Sleep(200 * time.Millisecond)
	lights.Blink(led.GREEN, 500*time.Millisecond)
	time.Sleep(200 * time.Millisecond)
	lights.Blink(led.GREEN, 500*time.Millisecond)

	// some time for connecting to USB for debugging
	time.Sleep(5 * time.Second)

	// connect to WIFI
	_ = wifi.NewWifi()

	// TODO run forever
	time.Sleep(10 * time.Minute)

}
