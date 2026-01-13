package main

import (
	"fmt"
	"math"
)

func main() {
	n := 3                     //number of different items
	m := 4                     //number of item to sell
	quantity := []int{4, 1, 5} //quantity of each item

	result := minimumPenalty(n, m, quantity)
	fmt.Println(result)
}

func minimumPenalty(n, m int, quantity []int) int {
	smallestStock := math.MaxInt32
	penalty := 0
	newQuantity := quantity
	smallestIndex := 0
	for i := 0; i < m; i++ {
		for j := 0; j < len(newQuantity); j++ {
			if newQuantity[j] != 0 && newQuantity[j] < smallestStock {
				smallestStock = newQuantity[j]
				smallestIndex = j
			}
		}
		penalty += smallestStock
		newQuantity[smallestIndex] = newQuantity[smallestIndex] - 1
		smallestStock = math.MaxInt32

	}
	return penalty
}
