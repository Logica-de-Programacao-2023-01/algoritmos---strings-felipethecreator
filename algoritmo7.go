package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var phrase string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Entre com uma string: ")

	scanner.Scan()
	phrase = scanner.Text()

	if strings.ContainsAny(phrase, "0123456789") {
		fmt.Println("Contém ao menos um número!")
	} else {
		fmt.Println("Não contém números!")
	}

}
