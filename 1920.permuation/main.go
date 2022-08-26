package main

// Build Array from Permutation
// Leetcode Problem #1920

import (
	"fmt"
)

func main() {

	arr := []int{0, 2, 1, 5, 3, 4}

	fmt.Println(buildArray(arr))
}

func buildArray(nums []int) []int {

	var arr2 []int

	for _, elem := range nums {

		arr2 = append(arr2, nums[elem])

	}
	return arr2

}

// ans = [nums[nums[0]], nums[nums[1]], nums[nums[2]], nums[nums[3]], nums[nums[4]], nums[nums[5]]]
//     = [nums[0], nums[2], nums[1], nums[5], nums[3], nums[4]]
//     = [0,1,2,4,5,3]
