package Vehicle

import "fmt"

type Car struct {
	vehicle        vehicle
	fuelEfficiency int
}

func NewCar(plate string, maxSpeed int, fuelEfficiency int) Vehicle {
	car := &Car{
		vehicle:        vehicle{licensePlate: plate, maxSpeed: maxSpeed},
		fuelEfficiency: fuelEfficiency,
	}
	return car
}

func (c *Car) SetLicensePlate(plate string) {
	c.vehicle.licensePlate = plate
}

func (c *Car) SetMaxSpeed(speed int) {
	c.vehicle.maxSpeed = speed
}
func (c *Car) GetMaximumSpeed() int {
	return c.vehicle.maxSpeed
}

func (c *Car) CalculateFuelConsumption(distance int) float64 {
	return float64(distance) / float64(c.fuelEfficiency)
}

func (c *Car) DisplayInfo() {
	fmt.Println("License Plate:", c.vehicle.licensePlate)
	fmt.Println("Maximum Speed:", c.vehicle.maxSpeed)
}

func (c *Car) GetFuelEfficiency() int {
	return c.fuelEfficiency
}

func (c *Car) SetFuelEfficiency(efficiency int) {
	c.fuelEfficiency = efficiency
}
