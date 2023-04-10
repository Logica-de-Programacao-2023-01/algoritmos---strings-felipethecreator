package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("Digite uma frase: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	line := scanner.Text()

	words := strings.Fields(line)

	count := len(words)

	fmt.Printf("Sua frase tem %d palavras.", count)
}
