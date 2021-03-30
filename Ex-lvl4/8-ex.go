package main

import (
	"fmt"
)

func main() {
	persons := map[string][]string{
		"Guilherme": []string{"Hérica", "Chocolate", "Ver vídeos"},
		"Hérica":    []string{"Guilherme", "Chocolate", "Ver videos"},
		"Gopher":    []string{"Compilar", "Identar", "Ranger os dentes"},
	}
	for key, value := range persons {
		fmt.Println(key)
		for _, hobby := range value {
			fmt.Println(hobby)
		}
		fmt.Println()
	}
}
