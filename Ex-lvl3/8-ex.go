package main

import (
	"fmt"
)

func main() {
	ondeEstou := "mar"
	switch ondeEstou {
	case "praia":
		fmt.Println("De boa na praia")
	case "campo":
		fmt.Println("De boa no campo")
	case "mar":
		fmt.Println("De boa nas ondas")
	default:
		fmt.Println("Trabalhando 😢")
	}
}
