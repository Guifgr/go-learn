package main

import (
	"fmt"
)

func main() {
	persons := [][]string{
		[]string{"Hérica", "Cadoni", "Ver vídeos"},
		[]string{"Guilherme", "Rocha", "Ficar com a Hérica"},
		[]string{"Gopher", "Gomes", "Ranger os dentes"},
	}
	for _, value := range persons {
		fmt.Println(value)
	}

}
