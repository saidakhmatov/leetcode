package main

// Two Sum
// Leetcode Problem #1

import (
	"fmt"
)

func main() {

	arr := []int{2, 5, 5, 11}
	target := 13

	fmt.Println(twoSum(arr, target))
}

func twoSum(nums []int, target int) []int {

	arr := []int{}

	for i := 0; i < len(nums); i++ {

		for k, x := range nums {

			if i == k {
				continue
			} else if target == nums[i]+x {
				arr = append(arr, i, k)
				return arr
			}
			continue
		}

	}

	return arr

}
