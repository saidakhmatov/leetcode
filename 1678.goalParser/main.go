package main

// Goal Parser Interpretation
// Leetcode Problem #1678

import (
	"fmt"
	"strings"
)

func main() {
	// Input: command = "(al)G(al)()()G"
	// Output: "alGalooG"

	sentences := "(al)G(al)()()G"

	fmt.Print(interpret(sentences))

}

func interpret(command string) string {

	var max []string

	for i := 0; i < len(command); i++ {

		if command[i] == 41 {
			if command[i-1] == 40 {

				max = append(max, string("o"))
			}
			continue

		} else if command[i] != 41&40 {
			r := rune(command[i])
			max = append(max, string(r))
		} else {
			continue

		}
	}

	res := strings.Join(max, "")

	return res

}
