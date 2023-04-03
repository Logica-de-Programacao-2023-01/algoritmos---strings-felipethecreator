//Escreva um programa que receba uma string e converta todas as
//letras minúsculas para maiúsculas. Informe ao usuário o resultado.

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

	AllCaps := strings.ToUpper(frase)
	fmt.Println(AllCaps)
}
