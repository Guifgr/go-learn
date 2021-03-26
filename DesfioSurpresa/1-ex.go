package main

import (
	"fmt"
)

func main() {
	for x := 33; x <= 122; x++ {
		fmt.Printf("%d\t%#x\t%#U\t%v \n", x, x, x, string(x))
	}
}
