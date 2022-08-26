package main

import (
	"fmt"
	"math"
)

func main() {

	sentences := "(1+(2*3)+((8)/4))+1"

	fmt.Println(maxDepth(sentences))

}

func maxDepth(s string) int {

	var res float64
	var num float64

	for _, elem := range s {

		if elem == 40 {
			num += 1
		}
		if elem == 41 {
			num--
		}
		res = math.Max(res, num)
	}

	return int(res)

}
