package main

import "fmt"

// Escreva uma função que receba um slice de strings como parâmetro e
// retorne uma string com todas as strings concatenadas e separadas por vírgulas.
// Caso o slice esteja vazio, retorne um erro.
func sep(str []string) string {
	if str == nil {
		erro := "erro"
		return erro
	}
	frase := str[0]
	for i, valor := range str {
		if i > 0 {
			frase += ", " + valor
		}
	}
	return frase
}

func main() {
	frases := []string{"Oi", "tudo otimo", "nossa"}
	fmt.Print(sep(frases))
}
