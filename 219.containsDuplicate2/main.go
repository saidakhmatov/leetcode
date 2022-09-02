package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 1, 2, 3}
	fmt.Println(containsNearbyDuplicate(nums, 2))
}

func containsNearbyDuplicate(nums []int, k int) bool {

	numsMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if numsMap[int(nums[i])] != 0 && i+1-numsMap[int(nums[i])] <= k {
			return true
		}
		numsMap[int(nums[i])] = i + 1
	}

	return false

}
