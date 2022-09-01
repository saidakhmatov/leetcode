package main

import (

	"fmt"

)

func main() {

	fmt.Println(balancedStringSplit("RLRRLLRLRL"))
}

func balancedStringSplit(s string) int {
	
	count := 0
	balance := 0
	for i := 0; i < len(s); i++ {
		
		if s[i] == 'L' {
			balance--
		} else {
			balance++
		}
		
		if balance == 0 {
			count++
		}
	}
	return count
	
}
