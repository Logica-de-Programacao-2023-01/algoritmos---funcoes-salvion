package main

import (
	"errors"
	"fmt"
)

// Função que aplica a função f em cada elemento do slice e retorna um novo slice com os resultados
func aplicarFuncao(slice []int, f func(int) int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	resultado := make([]int, len(slice))
	for i, v := range slice {
		resultado[i] = f(v)
	}

	return resultado, nil
}

func main() {
	// Exemplo de função para dobrar um número
	dobrar := func(num int) int {
		return num * 2
	}

	// Slice de inteiros de exemplo
	numeros := []int{1, 2, 3, 4, 5}

	// Chamada da função aplicarFuncao com o slice de números e a função dobrar
	resultado, err := aplicarFuncao(numeros, dobrar)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	// Imprimir o resultado
	fmt.Println(resultado) // Output: [2 4 6 8 10]
}
