package main

import "fmt"

func main() {

	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {

	numsMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		numsMap[int(nums[i])]++
	}

	// fmt.Println(numsMap[int(nums[3])])
	// count := 0

	for i := 0; i < len(nums); i++ {
		if numsMap[int(nums[i])] >= 2 {
			return true
		}
	}

	return false

}
