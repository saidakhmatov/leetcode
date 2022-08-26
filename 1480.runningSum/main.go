package main

// Running Sum of 1d Array
// Leetcode Problem #1480

import (
	"fmt"
)

func main() {

	arr := []int{1, 2, 3, 4}

	fmt.Println(runningSum(arr))
}

func runningSum(nums []int) []int {

	arr2 := []int{}

	sum := 0

	for _, elem := range nums {

		sum = sum + elem

		arr2 = append(arr2, sum)

	}
	return arr2

}
