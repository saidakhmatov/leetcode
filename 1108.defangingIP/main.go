package main

// Defanging an IP Address
// Leetcode Problem #1108

import (
	"fmt"
	"strings"
)

func main() {

	ip := "1.1.1.1"

	fmt.Println(defangIPaddr(ip))

}

func defangIPaddr(ip string) string {

	// Input: address = "1.1.1.1"
	// Output: "1[.]1[.]1[.]1"

	var arr []string
	for _, elem := range ip {
		if elem == 46 {
			arr = append(arr, "[.]")
		} else {
			arr = append(arr, string(elem))
		}
	}
	res := strings.Join(arr, "")

	return res

}
