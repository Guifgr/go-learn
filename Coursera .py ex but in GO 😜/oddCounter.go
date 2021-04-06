package main

import "fmt"

func main() {
	var contador int
	fmt.Printf("Digite o valor de n: ")
	fmt.Scan(&contador)
	numeros := 0
	for contador > 0 {
		if (numeros % 2) != 0 {
			fmt.Println(numeros)
			contador -= 1
			numeros += 1
		} else {
			numeros += 1
		}
	}
}
