package main

import "fmt"

func main() {

	var algo1, algo2 string
	fmt.Print("Digite algo: ")
	fmt.Scan(&algo1)
	fmt.Print("Digite outro algo: ")
	fmt.Scan(&algo2)

	if algo1 == algo2 {
		fmt.Println("Esses algos são iguais.")
	} else {
		fmt.Println("Não são iguais.")
	}
}
