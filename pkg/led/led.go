package led

import (
	"machine"
	"time"
)

const (
	GREEN = iota
	YELLOW
	RED
)

type LED struct {
	green  machine.Pin
	yellow machine.Pin
	red    machine.Pin
}

func NewLED() *LED {
	led := &LED{
		green:  machine.D10,
		yellow: machine.D11,
		red:    machine.D12,
	}
	led.init()
	return led
}

func (l *LED) init() {
	l.green.Configure(machine.PinConfig{Mode: machine.PinOutput})
	l.yellow.Configure(machine.PinConfig{Mode: machine.PinOutput})
	l.red.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

func (l *LED) Blink(color int, d time.Duration) {
	var pin machine.Pin
	switch color {
	case GREEN:
		pin = l.green
	case YELLOW:
		pin = l.yellow
	case RED:
		pin = l.green
	default:
		panic("invalid led pin!")
	}
	blink(pin, d)
}

func blink(pin machine.Pin, d time.Duration) {
	pin.High()
	time.Sleep(d)
	pin.Low()
}
