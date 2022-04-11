package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, value := range slice {
		fmt.Println(value)
	}
	fmt.Printf("%T", slice)
}
