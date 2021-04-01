package main

import "fmt"

func main() {
	var x int
	fmt.Println("numero")
	fmt.Scan(&x)
	paridade := x % 2
	if paridade == 0 {
		print("par")
	} else {
		print("ímpar")
	}
}
