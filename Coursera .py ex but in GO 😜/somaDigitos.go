package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num string
	totalDigitos := 0
	fmt.Println("Digite um numero inteiro")
	fmt.Scan(&num)
	for len(num) > 0 {
		i, _ := strconv.Atoi(num[:1])
		totalDigitos += i
		num = num[1:]
	}
	fmt.Println(totalDigitos)
}
