package analogSensor

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/aio"
)

var _ gobot.Driver = (*LmTemperatureSensorDriver)(nil)

// LmTemperatureSensorDriver represents a Temperature Sensor
// The temperature is reported in degree Celsius
type LmTemperatureSensorDriver struct {
	name        string
	pin         string
	halt        chan bool
	temperature float64
	interval    time.Duration
	connection  aio.AnalogReader
	gobot.Eventer
}

// NewLmTemperatureSensorDriver returns a new LmTemperatureSensorDriver with a polling interval of
// 10 Milliseconds given an AnalogReader and pin.
//
// Optionally accepts:
// 	time.Duration: Interval at which the TemperatureSensor is polled for new information
//
// Adds the following API Commands:
// 	"Read" - See AnalogSensor.Read
func NewLmTemperatureSensorDriver(a aio.AnalogReader, pin string, v ...time.Duration) *LmTemperatureSensorDriver {
	d := &LmTemperatureSensorDriver{
		name:       gobot.DefaultName("LmTemperatureSensor"),
		connection: a,
		pin:        pin,
		Eventer:    gobot.NewEventer(),
		interval:   10 * time.Millisecond,
		halt:       make(chan bool),
	}

	if len(v) > 0 {
		d.interval = v[0]
	}

	// d.AddEvent(Data)
	// d.AddEvent(Error)

	return d
}

// Start starts the LmTemperatureSensorDriver and reads the Sensor at the given interval.
// Emits the Events:
//	Data int - Event is emitted on change and represents the current temperature in celsius from the sensor.
//	Error error - Event is emitted on error reading from the sensor.
func (a *LmTemperatureSensorDriver) Start() (err error) {
	// thermistor := 3975.0
	a.temperature = 0

	go func() {
		for {
			// rawValue, _ := a.Read()

			// resistance := float64(1023.0-rawValue) * 10000 / float64(rawValue)
			// newValue := 1/(math.Log(resistance/10000.0)/thermistor+1/298.15) - 273.15

			// newValue := (float64)(5.0*rawValue*100.0) / 1024.0
			// newValue := (float64(rawValue))
			a.temperature = 1000
			// if err != nil {
			// 	// a.Publish(Error, err)
			// } else if newValue != a.temperature && newValue != -1 {
			// 	a.temperature = newValue
			// 	// a.Publish(Data, a.temperature)
			// }
			select {
			case <-time.After(a.interval):
			case <-a.halt:
				return
			}
		}
	}()
	return
}

// Halt stops polling the analog sensor for new information
func (a *LmTemperatureSensorDriver) Halt() (err error) {
	a.halt <- true
	return
}

// Name returns the LmTemperatureSensorDrivers name
func (a *LmTemperatureSensorDriver) Name() string { return a.name }

// SetName sets the LmTemperatureSensorDrivers name
func (a *LmTemperatureSensorDriver) SetName(n string) { a.name = n }

// Pin returns the LmTemperatureSensorDrivers pin
func (a *LmTemperatureSensorDriver) Pin() string { return a.pin }

// Connection returns the LmTemperatureSensorDrivers Connection
func (a *LmTemperatureSensorDriver) Connection() gobot.Connection {
	return a.connection.(gobot.Connection)
}

// Read returns the current Temperature from the Sensor
func (a *LmTemperatureSensorDriver) Temperature() (val float64) {
	return a.temperature
}

// Read returns the raw reading from the Sensor
func (a *LmTemperatureSensorDriver) Read() (val int, err error) {
	return a.connection.AnalogRead(a.Pin())
}
