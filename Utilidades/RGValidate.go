package main

import (
	"fmt"
	"strconv"
	"strings"
)

func transformToInt(rg string) string {
	rg = strings.Replace(rg, ".", "", -1)
	rg = strings.Replace(rg, "-", "", -1)
	rg = strings.Replace(rg, ",", "", -1)
	rg = strings.Replace(rg, " ", "", -1)
	return rg
}

func validDigit(partialResult int) int {
	return (partialResult * 10) % 11
}

func validaterg(rg string) bool {
	rg = transformToInt(rg)
	var rgInt []int
	for i := 0; i < 9; i++ {
		thisPart := rg[i : i+1]
		thisPartInt, _ := strconv.Atoi(thisPart)
		rgInt = append(rgInt, thisPartInt)
	}
	if rg[8:] == "x" || rg[8:] == "X" {
		rgInt[8] = 10
	}
	firstPartPartialResult := rgInt[0]*10 +
		rgInt[1]*9 +
		rgInt[2]*8 +
		rgInt[3]*7 +
		rgInt[4]*6 +
		rgInt[5]*5 +
		rgInt[6]*4 +
		rgInt[7]*3 +
		rgInt[8]*2
	firstPartResult := validDigit(firstPartPartialResult)
	if firstPartResult == rgInt[8] {
		return true
	}
	return false
}

func main() {
	var rg string
	fmt.Printf("Digite o rg para ser validado:  ")
	fmt.Scan(&rg)
	fmt.Println(validaterg(rg))
}
