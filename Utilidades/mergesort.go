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

func sliceOfTen() []int { //YEEEEEEEEEEEEESSS [10]int type exists dont forget it!!!!
	var slice []int
	for i := 0; i < 10; i++ {
		slice = append(slice, randNum())
	}
	return slice
}

func Merge(left, right []int) []int {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}

func MergeSort(slice []int) []int {

	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return Merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

func main() {
	rand.Seed(time.Now().UnixNano()) //improve the randomicity of the random number generator
	theArray := sliceOfTen()
	fmt.Println(theArray)
	fmt.Println(MergeSort(theArray))
}
