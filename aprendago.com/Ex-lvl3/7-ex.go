package main

import (
	"fmt"
)

func main() {
	ondeEstou := "mar"
	switch {
	case ondeEstou == "praia":
		fmt.Println("De boa na praia")
	case ondeEstou == "campo":
		fmt.Println("De boa no campo")
	case ondeEstou == "mar":
		fmt.Println("De boa nas ondas")
	default:
		fmt.Println("Trabalhando 😢")
	}
}
