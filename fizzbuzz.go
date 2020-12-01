package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Escre um número inteiro:")
	var inteiro int
	fmt.Scanf("%d", &inteiro)

	for i := 1; i <= inteiro; i++ {
		fmt.Println(fizzbuzz(i))
	}
}

func fizzbuzz(n int) string {
	var resultado string
	var str = strconv.Itoa(n)
	//Fizz = 3 * n
	if n%3 == 0 {
		resultado = "Fizz"
	}

	if n%5 == 0 {
		resultado += "Buzz"
	}

	if len(resultado) == 0 {
		resultado = str
	}

	return resultado
}
