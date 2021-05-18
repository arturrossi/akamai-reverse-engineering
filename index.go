package main

import (
	"fmt"

	"nike.bot/akamai_reverse_eng/class"
)

func main() {
	bmak := class.BmakConstructor()

	bmak.To()

	sensor_data := bmak.GetSensorData()

	fmt.Println("Test! ", sensor_data)
}
