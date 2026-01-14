package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(minYearsToSurpass(1000, 800, 0.05, 0.08))

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
