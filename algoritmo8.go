package main

import (
	"fmt"
)

func reverseString(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	var bola8 string
	fmt.Print("Digite uma palavra: ")
	fmt.Scan(&bola8)

	fmt.Println(bola8)
	fmt.Println(reverseString(bola8))
}
