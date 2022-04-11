package main

import (
	"fmt"
)

func main() {
	persons := map[string][]string{
		"Guilherme": {"Hérica", "Chocolate", "Ver vídeos"},
		"Hérica":    {"Guilherme", "Chocolate", "Ver videos"},
		"Gopher":    {"Compilar", "Identar", "Ranger os dentes"},
	}

	persons["c#"] = []string{"imitar java", "Microsoft", "Ser muito foda"}
	for key, value := range persons {
		fmt.Println(key)
		for _, hobby := range value {
			fmt.Println(hobby)
		}
		fmt.Println()
	}
}
