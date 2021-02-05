package sdc30

import (
	"encoding/binary"
	"fmt"
	"machine"
	"math"
	"time"

	"github.com/sigurn/crc8"
	"tinygo.org/x/drivers"
)

// Device wraps an I2C connection to a SDC30 device.
type Device struct {
	bus     drivers.I2C
	Address uint16
	ready   machine.Pin
}

// New creates a new SDC30 device. The I2C bus must already be configured.
func New(bus drivers.I2C, ready machine.Pin) *Device {
	dev := &Device{
		bus:     bus,
		Address: Address,
		ready:   ready,
	}

	// boot time of sdc30
	time.Sleep(3 * time.Second)

	dev.init()

	return dev
}

func (d *Device) init() {

	d.ready.Configure(machine.PinConfig{Mode: machine.PinInput})

	var fwBytes []byte
	var err error
	for fwBytes, err = d.sendCommand(Firmware, nil); err != nil; {
		println(fmt.Sprintf("get firmware version failed! %v", err))
		time.Sleep(500 * time.Millisecond)
	}
	println(fmt.Sprintf("firmware version: %v.%v", int(fwBytes[0]), int(fwBytes[1])))

	for _, err = d.sendCommand(Start, []byte{0x00, 0x00}); err != nil; {
		println(fmt.Sprintf("start measure failed! %v", err))
		time.Sleep(500 * time.Millisecond)
	}

	println(fmt.Sprintf("started measurement"))
}

func (d *Device) GetMeasurement() (temp, hum, co2 float32, err error) {

	i2cReady := false
	for !i2cReady {
		var rec []byte
		rec, err = d.sendCommand(Ready, nil)
		if err != nil {
			return
		}
		i2cReady = binary.BigEndian.Uint16(rec) == 0x0001
		time.Sleep(time.Second)
	}
	println("measurement ready (i2c)!")

	pinReady := d.ready.Get()
	if pinReady {
		println("measurement ready (pin)!")
	}

	measurement, err := d.sendCommand(Measure, nil)
	if err != nil {
		return
	}

	co2Bytes := measurement[:4]
	tempBytes := measurement[4:8]
	humBytes := measurement[8:]

	//executed command, result: 44a5 82 c35a a3 41b2 21 0494 ff 4211 a3 9c90 41
	//co2 44a5c35a
	//temp 41b20494
	//hum 42119c90

	co2 = math.Float32frombits(binary.BigEndian.Uint32(co2Bytes))
	temp = math.Float32frombits(binary.BigEndian.Uint32(tempBytes))
	hum = math.Float32frombits(binary.BigEndian.Uint32(humBytes))

	println(fmt.Sprintf("co2: %v", co2))
	println(fmt.Sprintf("temp: %v", temp))
	println(fmt.Sprintf("hum: %v", hum))

	return
}

func (d *Device) sendCommand(cmd Command, args []byte) ([]byte, error) {

	if len(args) != cmd.argsLen {
		println("incorrect number of bytes for arg!")
		return nil, fmt.Errorf("incorrect number of bytes for arg")
	}

	cmdBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(cmdBytes, uint16(cmd.command))

	send := make([]byte, 2+cmd.argsLen+cmd.argsLen/2)
	//send[0] = 0xC2 // write header
	send[0] = cmdBytes[0]
	send[1] = cmdBytes[1]
	if len(args) == 2 {
		send[2] = args[0]
		send[3] = args[1]
		send[4] = crc([]byte{args[0], args[1]})
	}

	println(fmt.Sprintf("sending %+x", send))

	receive := make([]byte, cmd.resultLen)

	if err := d.bus.Tx(d.Address, send, receive); err != nil {
		println(fmt.Sprintf("error executing command! %v", err))
		return nil, err
	}

	println(fmt.Sprintf("executed command, result: %+x", receive))

	// TODO fix for longer result!
	// strip crc
	stripped := make([]byte, 0)
	for i := 0; i < cmd.resultLen; i += 3 {
		stripped = append(stripped, receive[i], receive[i+1])
	}
	return stripped, nil

}

func crc(in []byte) byte {
	table := crc8.MakeTable(crc8.Params{
		Poly:   0x31,
		Init:   0xFF,
		RefIn:  false,
		RefOut: false,
		XorOut: 0,
		Check:  0,
		Name:   "",
	})
	crc := crc8.Checksum(in, table)
	return byte(crc)
}
