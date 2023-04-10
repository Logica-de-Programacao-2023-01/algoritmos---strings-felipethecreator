package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "Hello World!"
	var letter string
	var substituition string

	fmt.Println("A frase em questão é ", word)

	fmt.Print("Digite o caractere que deseja substituir na frase: ")
	fmt.Scan(&letter)

	fmt.Print("Agora digite o caracter a ocupar esse lugar: ")
	fmt.Scan(&substituition)

	newWord := strings.ReplaceAll(word, letter, substituition)
	fmt.Print("Sua nova frase é: ", newWord)
}
