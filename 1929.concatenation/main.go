package main

// Concatenation-of-array
// Leetcode Problem #1929

import "fmt"

func main() {

	array := []int{1, 3, 5}

	fmt.Print(getConcatenation(array))

}

func getConcatenation(nums []int) []int {

	x := len(nums)

	var arr []int

	for i := 0; i < 2; i++ {

		for i := 0; i < x; i++ {
			arr = append(arr, nums[i])
		}

	}
	return arr

}
