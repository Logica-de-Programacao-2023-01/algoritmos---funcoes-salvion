package main

import "fmt"

// Crie uma função que receba um slice de strings e retorne a concatenação de todas as strings.
func juntar(f []string) string {
	juncao := ""
	for _, elemento := range f {
		juncao += " " + elemento
	}
	return juncao
}

func main() {
	frases := []string{"Bom dia", "flor do dia", "o sol", "já nasceu na fazendinha"}
	fmt.Print(juntar(frases))

}
