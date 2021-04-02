package main

import "fmt"

func main() {
	var x int
	var y int
	fmt.Println("Quantos numeros você quer digitar?")
	fmt.Scan(&x)
	var numeros []string
	for i := 0; i < x; i++ {
		fmt.Printf("Digite o numero da posição %d:\t", i+1)
		fmt.Scan(&y)
		if y%3 == 0 {
			numeros = append(numeros, "fizz")

		} else if y%5 == 0 {
			numeros = append(numeros, "buzz")

		} else {
			z := fmt.Sprintf("%v", y)
			numeros = append(numeros, z)
		}
	}
	fmt.Println(numeros)
}
