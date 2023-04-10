package main

import (
	"fmt"
	"strings"
)

func main() {
	number := "100.00"

	if strings.Contains(number, ".") {
		fmt.Println("Este é um número com ponto flutuante!")
	} else {
		fmt.Println("Este não é um número com ponto flutuante!")
	}
}
