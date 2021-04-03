package main

import (
	"fmt"
	"strconv"
	"strings"
)

func transformToInt(cpf string) string {
	cpf = strings.Replace(cpf, ".", "", -1)
	cpf = strings.Replace(cpf, "-", "", -1)
	cpf = strings.Replace(cpf, ",", "", -1)
	cpf = strings.Replace(cpf, " ", "", -1)
	return cpf
}

func validDigit(partialResult int) int {
	return (partialResult * 10) % 11
}

func validateCpf(cpf string) bool {
	cpf = transformToInt(cpf)
	var cpfInt []int
	for i := 0; i < 11; i++ {
		thisPart := cpf[i : i+1]
		thisPartInt, _ := strconv.Atoi(thisPart)
		cpfInt = append(cpfInt, thisPartInt)
	}
	firstPartPartialResult := cpfInt[0]*10 +
		cpfInt[1]*9 +
		cpfInt[2]*8 +
		cpfInt[3]*7 +
		cpfInt[4]*6 +
		cpfInt[5]*5 +
		cpfInt[6]*4 +
		cpfInt[7]*3 +
		cpfInt[8]*2
	firstPartResult := validDigit(firstPartPartialResult)
	secondPartPartialResult := cpfInt[0]*11 +
		cpfInt[1]*10 +
		cpfInt[2]*9 +
		cpfInt[3]*8 +
		cpfInt[4]*7 +
		cpfInt[5]*6 +
		cpfInt[6]*5 +
		cpfInt[7]*4 +
		cpfInt[8]*3 +
		firstPartResult*2
	secondPartResult := validDigit(secondPartPartialResult)
	if firstPartResult == cpfInt[9] && secondPartResult == cpfInt[10] {
		return true
	}
	return false
}

func main() {
	var cpf string
	fmt.Printf("Digite o cpf para ser validado:  ")
	fmt.Scan(&cpf)
	fmt.Println(validateCpf(cpf))
}
