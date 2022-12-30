package main

import "testing"

func TestFizzBuzz(t *testing.T){
	got := fizzbuzz(3)
	want := "Fizz"

	if got != want {
		t.Errorf("fizzbuzz(3) \n got %v \n want: \n%v", got, want)
	}
}