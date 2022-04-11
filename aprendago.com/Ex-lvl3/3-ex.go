package main

import "fmt"

func main() {
	anoQueNasci := 1999
	anoAtual := 2021
	for anoQueNasci <= anoAtual {
		fmt.Println(anoQueNasci)
		anoQueNasci++
	}
}
