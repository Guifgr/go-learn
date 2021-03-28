package main

import "fmt"

func main() {
	var nome, dia, mes, valor string
	fmt.Println("Digite o nome do cliente:")
	fmt.Scan(&nome)
	fmt.Println("Digite o dia de vencimento:")
	fmt.Scan(&dia)
	fmt.Println("Digite o mês de vencimento:")
	fmt.Scan(&mes)
	fmt.Println("Digite o valor da fatura:")
	fmt.Scan(&valor)
	fmt.Println("Olá,", nome, "\nA sua fatura com vencimento em", dia, "de", mes, "no valor de R$", valor, "está fechada.")
}
