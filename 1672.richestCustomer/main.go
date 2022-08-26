package main

// Richest Customer Wealth
// Leetcode Problem #1672

import (
	"fmt"
)

func main() {

	arr := [][]int{{1, 2, 3}, {3, 2, 1, 4, 9}, {1, 2, 3, 4, 5}}

	fmt.Println(maximumWealth(arr))
}

func maximumWealth(accounts [][]int) int {

	var max int = 0

	for _, elem := range accounts {
		sum := 0

		for _, i := range elem {

			sum += i
		}

		if sum > max {
			max = sum
		}

	}

	return max

}
