package main

import (
	"fmt"
)

func main() {

	x := struct {
		nome      string
		sabor     string
		ondeTem   []string
		vaiBemCom map[string]string
	}{
		nome:    "Shitake",
		sabor:   "Shoyu",
		ondeTem: []string{"Aqui em casa", "Restaurante Japones"},
		vaiBemCom: map[string]string{
			"Almoço": "Nana do nana do ludovico",
			"Janta":  "Coquinha marota de gelada",
		},
	}
	fmt.Println(x)
}
