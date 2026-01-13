package main

import "fmt"

func main() {
	arr := []int{10, 2, 8, 5}
	k := 2

	result := getDemolitionScore(arr, k)
	fmt.Println(result)
}

func getDemolitionScore(arr []int, k int) int {
	if k == 0 || len(arr) == 0 {
		return 0
	}

	bestScore := 0
	for i := 0; i < len(arr); i++ {
		currentScore := arr[i]

		left := arr[:i]
		right := arr[i+1:]

		var nextArr []int
		if len(left) > len(right) {
			nextArr = left
		} else if len(right) > len(left) {
			nextArr = right
		} else {
			nextArr = right
		}

		weakened := make([]int, 0, len(nextArr))
		for _, item := range nextArr {
			weakenedItem := item - currentScore
			if weakenedItem < 0 {
				weakenedItem = 0
			}
			weakened = append(weakened, weakenedItem)
		}
		score := currentScore + getDemolitionScore(weakened, k-1)
		if score > bestScore {
			bestScore = score
		}
	}
	return bestScore
}
