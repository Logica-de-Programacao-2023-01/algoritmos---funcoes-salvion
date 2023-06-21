package main

import "fmt"

//Crie uma função que receba um número inteiro como parâmetro e retorne verdadeiro se ele
//for um número primo e falso caso contrário. Caso o número seja negativo, retorne um erro.

func ehPrimo(num int) (bool, error) {
	if num < 0 {
		fmt.Println("Erro", nil)
	}
	divisores := 0
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			divisores++
		}
	}
	if divisores == 2 {
		return true, nil
	} else {
		return false, nil
	}
}

func main() {
	var n int
	fmt.Print("Digite e verifique aqui se o número é primo: ")
	fmt.Scan(&n)
	ver, err := ehPrimo(n)
	if err != nil {
		fmt.Println("Erro", err)
	}
	if ver == true {
		fmt.Println("É primo!")
	} else {
		fmt.Println("Não é primo :(")
	}
}
