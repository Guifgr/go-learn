package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	y := x * x
	x = x * 4
	fmt.Println("perímetro:", x, "- área", y)
}
