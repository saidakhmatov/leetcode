package main

import (
	"fmt"
	"math"
)

//main function
func main() {

	height := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(height))
	
	
}

func maxArea(height []int) int {
		
    var maxArea int = 0
    var left int = 0
    var right int = len(height) - 1

    for left < right {

        var area int = int(math.Min(float64(height[left]), float64(height[right]))) * (right - left)
        if area > maxArea {
            maxArea = area
        }

        if height[left] < height[right] {
            left++
        } else {
            right--
        }

    }

    return maxArea
}
