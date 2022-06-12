package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Digite um texto:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var text string

	if scanner.Err() == nil {
		text = ""
	} else {
		text = scanner.Text()
	}

	WasSaid, err := say(text)
	if err != nil {
		fmt.Println("Erro:", err.Error())
	}

	switch WasSaid {
	case true:
		fmt.Println("Texto foi dito")
	case false:
		fmt.Println("Texto não foi dito")
	}
}

func say(text string) (bool, error) {
	if len(text) == 0 {
		return false, fmt.Errorf("Texto vazio") // retorna falso e erro
	}

	fmt.Println(text)
	return true, nil // retorna verdadeiro e nenhum erro
}
