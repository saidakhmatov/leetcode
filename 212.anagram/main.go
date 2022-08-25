package main

import (
	"fmt"
)

func main() {

	fmt.Print(isAnagram("anagram", "nagaram"))
}

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	anagramMap := make(map[string]int)

	for i := 0; i < len(s); i++ {
		anagramMap[string(s[i])]++
	}

	for i := 0; i < len(t); i++ {
		anagramMap[string(t[i])]--
	}

	for i := 0; i < len(s); i++ {
		if anagramMap[string(s[i])] != 0 {
			return false
		}
	}

	return true
}
