package main

import (
	"fmt"
)

func main() {
	esporteFavorito := "basquete"
	switch esporteFavorito {
	case "basquete":
		fmt.Println("De boa na quadra")
	case "futebol":
		fmt.Println("De boa no campo")
	case "surf":
		fmt.Println("De boa nas ondas")
	default:
		fmt.Println("Chato 😜")
	}
}
