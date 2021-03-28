package main

import "fmt"

func main() {
	var x int
	fmt.Println("Digite um numero inteiro")
	fmt.Scan(&x)
	x = x / 10
	fmt.Println(x)
}
