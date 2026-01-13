package main

import v "task-7/Vehicle"

func main() {

	Ferrari := v.NewCar("B 5678 FR", 320, 8)
	Fuso := v.NewTruck("A 1234 AI", 120, 8, 15)

	Ferrari.DisplayInfo()
	Fuso.DisplayInfo()
}
