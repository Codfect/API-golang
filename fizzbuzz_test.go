package main

import (
	"testing"
)

// A palavra “Fizz” deve ser substituída no lugar de múltiplos do número três
func TestFizzBuzz3(t *testing.T) {
	arg := fizzbuzz(3)
	fizz := "Fizz"

	if arg != fizz {
		t.Errorf("fizzbuzz(3) \n arg: %v \n fizz: %v", arg, fizz)
	}
}

// A palavra “Buzz” deve ser substituída no lugar de múltiplos do número cinco
func TestFizzBuzz5(t *testing.T) {
	arg := fizzbuzz(5)
	buzz := "Buzz"

	if arg != buzz {
		t.Errorf("fizzbuzz(5) \n arg: %v \n buzz: %v", arg, buzz)
	}
}

//Para casos de serem multiplos por 3 e por 5
func TestFizzBuzz15(t *testing.T) {
	arg := fizzbuzz(15)
	fizzbuzz := "FizzBuzz"

	if arg != fizzbuzz {
		t.Errorf("fizzbuzz(15) \n arg: %v \n fizzbuzz: %v", arg, fizzbuzz)
	}
}

//Para casos de não serem multiplos por 3 e nem por 5
func TestFizzBuzz4(t *testing.T) {
	arg := fizzbuzz(4)
	unmultiple := "4"

	if arg != unmultiple {
		t.Errorf("fizzbuzz(4) \n arg: %v \n unmultiple: %v", arg, unmultiple)
	}
}
