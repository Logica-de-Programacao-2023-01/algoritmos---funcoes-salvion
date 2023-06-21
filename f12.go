package main

import (
	"errors"
	"fmt"
)

//Escreva uma função que receba um slice de strings como parâmetro e retorne uma string contendo apenas as strings que começam
//com uma letra maiúscula. Caso o slice esteja vazio, retorne um erro.

func maiusc(palavras []string) ([]string, error) {
	if len(palavras) == 0 {
		return nil, errors.New("slice vazio")
	}
	maiusc := make([]string, 0)
	for _, str := range palavras {
		if 64 < int(str[0]) && int(str[0]) < 91 {
			maiusc = append(maiusc, str)
		}
	}
	return maiusc, nil
}

func main() {
	slice := []string{"Alo", "dona maria", "Vigia", "gasolina", "carambola verde", "grande Seu Zé"}
	maiusculas, err := maiusc(slice)
	if err != nil {
		fmt.Println("Erro", err)
	}
	fmt.Println(maiusculas)
}
