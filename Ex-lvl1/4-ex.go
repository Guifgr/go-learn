package main

import (
	"fmt"
)

type hotdog int

var x hotdog

func main() {
	fmt.Print(x)
	fmt.Printf("\n%T\n", x)
	x = 42
	fmt.Print(x)
}
