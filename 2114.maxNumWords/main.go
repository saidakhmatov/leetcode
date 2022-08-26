package main

// Maximum Number of Words Found in Sentences
// Leetcode Problem #2114

import "fmt"

func main() {

	sentences := []string{"please wait", "continue to fight", "continue to win"}

	fmt.Println(mostWordsFound(sentences))

}

func mostWordsFound(sentences []string) int {

	var max int = 0

	for _, elem := range sentences {

		sum := 0

		for _, elem := range elem {

			if elem == 32 {
				sum++
			}
			if sum > max {
				max = sum
			}

		}
	}
	max++

	return max

}
