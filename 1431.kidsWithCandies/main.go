package main

import (
	"fmt"
)

func main() {

	candies := []int{4, 2, 1, 1, 2}
	extraCandies := 1
	fmt.Println(kidsWithCandies(candies, extraCandies))

}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var bool []bool
	max := 0
	for i := 0; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= max {
			bool = append(bool, true)
		} else {
			bool = append(bool, false)
		}
	}
	return bool

}
