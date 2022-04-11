package main

import (
	"fmt"
)

func main() {
	meImArray := [5]int{1, 2, 3, 4, 5}
	for _, value := range meImArray {
		fmt.Println(value)
	}
	fmt.Printf("%T", meImArray)
}
