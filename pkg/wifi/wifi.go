package wifi

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/wifinina"
)

const ssid = ""
const pass = ""

type Wifi struct {
	tx      machine.Pin
	rx      machine.Pin
	spi     machine.SPI
	adaptor *wifinina.Device
}

func NewWifi() *Wifi {
	wifi := &Wifi{
		tx:  machine.NINA_TX,
		rx:  machine.NINA_RX,
		spi: machine.NINA_SPI,
		adaptor: &wifinina.Device{
			SPI:   machine.NINA_SPI,
			CS:    machine.NINA_CS,
			ACK:   machine.NINA_ACK,
			GPIO0: machine.NINA_GPIO0,
			RESET: machine.NINA_RESETN,
		},
	}
	wifi.connect()
	return wifi
}

func (w *Wifi) connect() {
	machine.UART2.Configure(machine.UARTConfig{TX: w.tx, RX: w.rx})

	// Configure SPI for 8Mhz, Mode 0, MSB First
	w.spi.Configure(machine.SPIConfig{
		Frequency: 8 * 1e6,
		SDO:       machine.NINA_SDO,
		SDI:       machine.NINA_SDI,
		SCK:       machine.NINA_SCK,
	})

	w.adaptor.Configure()

	time.Sleep(2 * time.Second)

	w.adaptor.SetPassphrase(ssid, pass)
	for st, _ := w.adaptor.GetConnectionStatus(); st != wifinina.StatusConnected; {
		println("Connection status: " + st.String())
		time.Sleep(1 * time.Second)
		st, _ = w.adaptor.GetConnectionStatus()
	}

	println("Connected.")

	time.Sleep(2 * time.Second)

	ip, _, _, err := w.adaptor.GetIP()
	for ; err != nil; ip, _, _, err = w.adaptor.GetIP() {
		println(err.Error())
		time.Sleep(1 * time.Second)
	}

	println("IP: " + ip.String())
}
