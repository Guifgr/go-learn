package main

import "fmt"

type pessoa struct {
	nome            string
	sobrenome       string
	favFlavIcecream []string
}

func main() {
	newMap := make(map[string]pessoa)

	newMap["Joaquina"] = pessoa{"Antoniete", "Joaquina", []string{"chocolate"}}

	newMap["da Silva"] = pessoa{
		nome:            "Antonio",
		sobrenome:       "da Silva",
		favFlavIcecream: []string{"Chocolate", "Flocos", "feijão"},
	}

	for _, v := range newMap {
		fmt.Println(v.nome, v.sobrenome)
		for _, value := range v.favFlavIcecream {
			fmt.Println(value)
		}
		fmt.Println()
	}
}
