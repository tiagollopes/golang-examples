package main

import (
	"errors"
	"fmt"
)

// Dividir retorna o resultado e um erro caso o divisor seja zero
func Dividir(a, b float64) (float64, error) {
	if b == 0 {
		// Criando um erro básico
		return 0, errors.New("não é possível dividir por zero")
	}
	return a / b, nil
}

func main() {
	// Exemplo 1: Sucesso
	resultado, err := Dividir(10, 2)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}

	// Exemplo 2: Falha
	_, err = Dividir(10, 0) // Usamos _ quando não queremos o resultado, apenas o erro
	if err != nil {
		fmt.Printf("Falha esperada: %v\n", err)
	}
}
