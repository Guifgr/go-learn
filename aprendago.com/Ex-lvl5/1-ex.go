package main

import (
	"fmt"
)

type pessoa struct {
	nome            string
	sobrenome       string
	favFlavIcecream []string
}

func main() {
	pessoa1 := pessoa{"Antoniete", "Joaquina", []string{"chocolate"}}
	pessoa2 := pessoa{
		nome:            "Antonio",
		sobrenome:       "da Silva",
		favFlavIcecream: []string{"Chocolate", "Flocos", "feijão"},
	}
	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for _, v := range pessoa1.favFlavIcecream {
		fmt.Println(v)
	}
	fmt.Println()
	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for _, v := range pessoa2.favFlavIcecream {
		fmt.Println(v)
	}
}
