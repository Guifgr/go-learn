package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Println("Primeiro número:")
	fmt.Scan(&x)
	fmt.Println("Segundo número:")
	fmt.Scan(&y)
	fmt.Println("Terceiro número:")
	fmt.Scan(&z)
	if x <= y && y <= z {
		fmt.Println("Crescente")
	} else {
		fmt.Println("Não está em ordem crescente")
	}
}
