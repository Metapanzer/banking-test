package Vehicle

import "fmt"

type Vehicle interface {
	SetLicensePlate(plate string)
	SetMaxSpeed(speed int)
	GetMaximumSpeed() int
	CalculateFuelConsumption(distance int) float64
	DisplayInfo()
}
type vehicle struct {
	licensePlate string
	maxSpeed     int
}

func (v *vehicle) SetLicensePlate(plate string) {
	v.licensePlate = plate
}

func (v *vehicle) SetMaxSpeed(speed int) {
	v.maxSpeed = speed
}
func (v *vehicle) GetMaximumSpeed() int {
	return v.maxSpeed
}

func (v *vehicle) displayInfo() {
	fmt.Println("License Plate:", v.licensePlate)
	fmt.Println("Maximum Speed:", v.maxSpeed)
}
