//Escreva um programa que receba uma string e remova todos
//os espaços em branco. Informe ao usuário o resultado.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Informe uma palavra ou frase: ")
	scanner.Scan()
	frase := scanner.Text()

	result := strings.ReplaceAll(frase, " ", "")

	fmt.Println("Resultado:", result)
}
