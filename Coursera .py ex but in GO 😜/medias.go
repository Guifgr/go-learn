package main

import "fmt"

func main() {
	var total, grade int
	for i := 0; i < 4; i++ {
		fmt.Println("Digite a nota №:", i+1)
		fmt.Scan(&grade)
		total += grade
	}
	fmt.Println("Sua média é", total/4)
}
