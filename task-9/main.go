package main

import (
	"fmt"
	"math"
)

func main() {
	aliceInitial := 1000
	bobBonus := 800
	aliceRate := 0.05
	bobRate := 0.08

	fmt.Println(minYearsToSurpass(aliceInitial, bobBonus, aliceRate, bobRate))

}

func minYearsToSurpass(aliceInitial int, bobBonus int, aliceRate float64, bobRate float64) int {
	aliceInvestment := float64(aliceInitial)
	bobInvestment := float64(bobBonus)
	gap := aliceInvestment - bobInvestment
	newGap := math.MaxFloat32
	year := 0

	for i := 1; aliceInvestment > bobInvestment; i++ {
		aliceInvestment += aliceInvestment * aliceRate
		bobInvestment += bobInvestment * bobRate

		if bobInvestment > aliceInvestment {
			year = i
		}
		newGap = aliceInvestment - bobInvestment

		gap = newGap
		if newGap >= gap {
			year = -1
			break
		}
	}
	return year
}
