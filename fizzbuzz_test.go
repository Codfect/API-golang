package main

import (
	"testing"
)

// A palavra “Fizz” deve ser substituída no lugar de múltiplos do número três
func TestFizzBuzz3(t *testing.T) {
	arg := fizzbuzz(3)
	fizz := "Fizz"

	if arg != fizz {
		t.Errorf("fizzbuzz3(3) \n param: %v \n fizz: %v", arg, fizz)
	}
}

func TestFizzBuzz5(t *testing.T) {
	arg := fizzbuzz(5)
	buzz := "Buzz"

	if arg != buzz {
		t.Errorf("fizzbuzz5(5) \n param: %v \n fizz: %v", arg, buzz)
	}
}
