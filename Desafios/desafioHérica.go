package main

import "fmt"

func recursiveILuvU(contagem int) {
	if contagem < 1000 {
		contagem += 1
		fmt.Println("Eu te amo Hérica", contagem)
		recursiveILuvU(contagem)
	}
}

func main() {
	recursiveILuvU(0)
}
