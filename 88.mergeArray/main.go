package main

// Merge Sorted Array
// Leetcode Problem #88

import (
	"fmt"
)

func main() {

	a := []int{1, 2, 3, 0, 0, 0}
	b := []int{2, 5, 6}
	m := 3
	n := 3

	merge(a, m, b, n)

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}

	// for index := 0; index < n; index++ {
	// 	nums1 = append(nums1, nums2[index])
	// }
	i := m - 1
	j := n - 1
	k := m + n - 1

	for ; i >= 0 && j >= 0; k-- {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
	for ; j >= 0; k-- {
		nums1[k] = nums2[j]
		j--
	}
	fmt.Println(nums1)
}