package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo         veiculo
	tracaoNasQuatro bool
}
type sedan struct {
	veiculo    veiculo
	modeloLuxo bool
}

func main() {
	l200 := caminhonete{veiculo{4, "amarelo"}, true}
	sedan := sedan{veiculo{4, "prata"}, true}
	fmt.Println(l200)
	fmt.Println("Tem tração nas quatro:", l200.tracaoNasQuatro)

	fmt.Println(sedan)
	fmt.Println("É de luxo:", sedan.modeloLuxo)
}
