package sensor

import (
	"machine"
	"time"

	"github.com/slintes/co2-sensor/pkg/sensor/sdc30"
)

type Sensor struct {
	sdc30 *sdc30.Device
	Temp  int
	Hum   int
	CO2   int
}

func New() *Sensor {
	machine.I2C0.Configure(machine.I2CConfig{Frequency: machine.TWI_FREQ_100KHZ})
	sdc30 := sdc30.New(machine.I2C0, machine.D8)
	s := &Sensor{
		sdc30: sdc30,
	}

	for {
		sdc30.GetMeasurement()
		time.Sleep(5 * time.Second)
	}

	return s
}
