package main

import (
	"testing"
)

// A palavra “Fizz” deve ser substituída no lugar de múltiplos do número três
func TestFizzBuzz(t *testing.T) {
	param := fizzbuzz(3)
	fizz := "Fizz"

	if param != fizz {
		t.Errorf("fizzbuzz3(3) \n param: %v \n fizz: %v", param, fizz)
	}
}
