package devices

type Device interface {
	Init()
	Destroy()
}

var Devices = map[string]func(uint8, int) (Device, error) {
	"adxl345": NewAdxl345,
}
