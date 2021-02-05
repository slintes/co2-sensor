package main

import (
	"machine"
	"time"
)

func main() {

	green := machine.D10
	yellow := machine.D11
	red := machine.D12

	green.Configure(machine.PinConfig{Mode: machine.PinOutput})
	yellow.Configure(machine.PinConfig{Mode: machine.PinOutput})
	red.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		green.High()
		time.Sleep(time.Millisecond * 500)
		green.Low()
		time.Sleep(time.Millisecond * 500)

		yellow.High()
		time.Sleep(time.Millisecond * 500)
		yellow.Low()
		time.Sleep(time.Millisecond * 500)

		red.High()
		time.Sleep(time.Millisecond * 500)
		red.Low()
		time.Sleep(time.Millisecond * 500)
	}
}
