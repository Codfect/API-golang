package main

import "fmt"

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
	//Fizz = 3 * n
	if n%3 == 0 {
		resultado = "Fizz"
	}

	return resultado
}
