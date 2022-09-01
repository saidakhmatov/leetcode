package main

import (
	"fmt"
)

func main() {

	key := "the quick brown fox jumps over the lazy dog"
	message := "vkbs bs t suepuv"

	fmt.Println(decodeMessage(key, message))

}

func decodeMessage(key string, message string) string {

	tab := [26]byte{}

	k := byte('a')
	for _, elem := range key {
		if elem != ' ' && tab[elem-'a'] == 0 {
			tab[elem-'a'] = k
			k++
		}
	}

	bs := []byte(message)

	for i, elem := range bs {
		if elem != ' ' {
			bs[i] = tab[elem-'a']
		}
	}
	return string(bs)
}
