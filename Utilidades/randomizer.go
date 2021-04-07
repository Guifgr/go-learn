package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randNum() int {
	min := 10
	max := 30
	return rand.Intn(max-min+1) + min
}

func createTheArrayOfRandNum(howMany int) []int {
	var numbers []int
	for howMany > 0 {
		numbers = append(numbers, randNum())
		howMany--
	}
	return numbers
}

func main() {
	rand.Seed(time.Now().UnixNano()) //improve the randomicity of the random number generator
	var howMany int
	fmt.Println("Quantos numeros deseja criar?")
	fmt.Scan(&howMany)
	if howMany == 1 {
		fmt.Println(randNum())
	} else {
		fmt.Println(createTheArrayOfRandNum(howMany))
	}

}
