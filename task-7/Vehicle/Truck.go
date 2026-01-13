package Vehicle

import "fmt"

type Truck struct {
	vehicle        vehicle
	fuelEfficiency int
	cargoWeight    int
}

func NewTruck(plate string, maxSpeed int, fuelEfficiency int, cargoWeight int) Vehicle {
	truck := &Truck{
		vehicle:        vehicle{licensePlate: plate, maxSpeed: maxSpeed},
		fuelEfficiency: fuelEfficiency,
		cargoWeight:    cargoWeight,
	}
	return truck
}

func (t *Truck) SetLicensePlate(plate string) {
	t.vehicle.licensePlate = plate
}

func (t *Truck) SetMaxSpeed(speed int) {
	t.vehicle.maxSpeed = speed
}
func (t *Truck) GetMaximumSpeed() int {
	return t.vehicle.maxSpeed
}

func (t *Truck) CalculateFuelConsumption(distance int) float64 {
	return float64(distance)/float64(t.fuelEfficiency) + (float64(t.cargoWeight) * 0.05)
}

func (t *Truck) DisplayInfo() {
	fmt.Println("License Plate:", t.vehicle.licensePlate)
	fmt.Println("Maximum Speed:", t.vehicle.maxSpeed)
}

func (t *Truck) GetFuelEfficiency() int {
	return t.fuelEfficiency
}

func (t *Truck) SetFuelEfficiency(efficiency int) {
	t.fuelEfficiency = efficiency
}

func (t *Truck) GetCargoWeight() int {
	return t.cargoWeight
}

func (t *Truck) SetCargoWeight(cargoWeight int) {
	t.cargoWeight = cargoWeight
}
