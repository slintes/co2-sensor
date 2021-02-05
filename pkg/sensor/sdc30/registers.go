package sdc30

// https://www.sensirion.com/fileadmin/user_upload/customers/sensirion/Dokumente/9.5_CO2/Sensirion_CO2_Sensors_SCD30_Interface_Description.pdf

// The I2C address which this device listens to.
const Address = 0x61

type Register uint16

const (
	StartContMeasurement Register = 0x0010 // arg = pressure in mBar or 0
	StopContMeasurement  Register = 0x0104 // no arg
	MeasurementInterval  Register = 0x4600 // with arg = set interval in sec, without = get
	GetReadyStatus       Register = 0x0202 // return 1 if data ready, else 0
	ReadMeasurement      Register = 0x0300 // no arg
	SelfCalibration      Register = 0x5306 // with arg = set on (1) or off (0), without = get
	ForceCalibration     Register = 0x5204 // arg = CO2 in ppm
	TempOffset           Register = 0x5403 // with arg = set offset in K, without = get
	AltitudeCompensation Register = 0x5102 // with arg = set alt in m above sea level, without = get
	ReadFirmwareVersion  Register = 0xD100 // no arg
	SoftReset            Register = 0xD304 // no arg
)

type Command struct {
	command   Register
	argsLen   int
	resultLen int
}

var (
	Start = Command{
		command:   StartContMeasurement,
		argsLen:   2,
		resultLen: 0,
	}
	Ready = Command{
		command:   GetReadyStatus,
		argsLen:   0,
		resultLen: 3,
	}
	Measure = Command{
		command:   ReadMeasurement,
		argsLen:   0,
		resultLen: 18,
	}
	Firmware = Command{
		command:   ReadFirmwareVersion,
		argsLen:   0,
		resultLen: 3,
	}
)
