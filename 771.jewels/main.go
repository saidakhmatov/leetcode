package main

import (
	"fmt"
)

func main() {


	fmt.Println(numJewelsInStones("AAB", "AAABB"))

}

func numJewelsInStones(jewels string, stones string) int {
	jewelsMap := make(map[string]int)
	for i := 0; i < len(jewels); i++ {
		jewelsMap[string(jewels[i])]++
	}

	count := 0

	for i := 0; i < len(stones); i++ {
		if jewelsMap[string(stones[i])] > 0 {
			count++
		}
	}

	return count

}
