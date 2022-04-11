package main

import (
	"fmt"
)

func main() {
	anoQueNasci := 1999
	anoAtual := 2021
	for {
		fmt.Println(anoQueNasci)
		anoQueNasci++
		if anoQueNasci > anoAtual {
			break
		}
	}
}
