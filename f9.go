package main

import (
	"errors"
	"fmt"
	"strings"
)

//Crie uma função que receba uma string como parâmetro e retorne um novo slice com todas
//as palavras contidas na string. Considere que as palavras são separadas por espaços em branco.
//Caso a string seja vazia, retorne um erro.

func palavras(s string) ([]string, error) {
	if s == "" {
		return nil, errors.New("string vazia")
	}
	palavras := strings.Split(s, " ")
	return palavras, nil
}

func main() {
	frase := "Crie uma função que receba uma string como parâmetro e retorne um novo slice com todas as palavras contidas na string. Considere que as palavras são separadas por espaços em branco. Caso a string seja vazia, retorne um erro."
	resultado, err := palavras(frase)
	if err != nil {
		fmt.Println("Erro:", nil)
		return
	}
	fmt.Println(resultado)
}
