package main

import (
	"errors"
	"fmt"
	"strconv"
)

//Crie uma função que receba um número inteiro como parâmetro e retorne a soma de todos os
//seus dígitos. Caso o número seja negativo, retorne um erro.

func somaD(numero int) (int, error) {
	if numero < 0 {
		return 0, errors.New("sem numero")
	}
	soma := 0
	strNumero := strconv.Itoa(numero)
	for _, digito := range strNumero {
		digitoInt, _ := strconv.Atoi(string(digito))
		soma += digitoInt
	}
	return soma, nil
}

func main() {
	var n int
	fmt.Print("n: ")
	fmt.Scan(&n)
	soma, err := somaD(n)
	if err != nil {
		fmt.Println("erro", nil)
	}
	fmt.Print(soma)
}
