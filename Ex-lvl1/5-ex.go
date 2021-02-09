package main

import (
	"fmt"
)

type hotdog int

var x hotdog
var y int

func main() {
	fmt.Print(x)
	fmt.Printf("\n%T\n", x)
	x = 42
	fmt.Print(x, "\n\n")

	y = int(x)

	fmt.Print(y)
	fmt.Printf("\n%T\n", y)
}
