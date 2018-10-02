package main

import (
	"Arduino/lib/analogSensor"
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor("COM4")
	lm := analogSensor.NewLmTemperatureSensorDriver(firmataAdaptor, "1")

	work := func() {
		gobot.Every(1000*time.Millisecond, func() {
			value, err := lm.Read()
			if err != nil {
				fmt.Println(err)
			}
			//fmt.Print(err)
			fmt.Println(value / 10)
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{lm},
		work,
	)

	robot.Start()
}
