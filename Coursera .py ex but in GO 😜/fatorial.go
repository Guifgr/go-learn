package main

import "fmt"

func main() {
	var x int
	fmt.Println("Digite o valor de n:")
	fmt.Scan(&x)
	if x == 0 {
		fmt.Println("1")
	} else {
		total := 1
		for i := x; i > 0; i-- {
			total *= i
		}
		fmt.Println(total)
	}
}
