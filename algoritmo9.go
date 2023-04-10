package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite algo: ")
	scanner.Scan()
	str := scanner.Text()

	var brotos1, brotos2 string
	fmt.Print("Caractere a ser trocado: ")
	fmt.Scan(&brotos1)
	fmt.Print("Trocar pelo caractere: ")
	fmt.Scan(&brotos2)

	if strings.Contains(str, brotos1) {
		fmt.Println(strings.ReplaceAll(str, brotos1, brotos2)
	} else {
		fmt.Printf("A frase não contém o caractere '%s'", brotos1)
	}
}
