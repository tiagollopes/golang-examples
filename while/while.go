package main

import "fmt"

// ATENÇÃO:
// Go NÃO possui a estrutura "while".
// Em Go, usamos o "for" para representar
// o comportamento de um while.

// Este exemplo demonstra:
// - como simular um while usando for
// - controle de loop com condição
// - uso de break dentro do loop

func main() {

	fmt.Println("=== Simulando while com for ===")

	contador := 1

	// Este for funciona como um while:
	// enquanto a condição for verdadeira,
	// o bloco continuará executando
	for contador <= 5 {
		fmt.Println("Contador:", contador)
		contador++
	}

	fmt.Println("\n=== While com break ===")

	valor := 1

	for {
		if valor > 5 {
			fmt.Println("Condição atendida, saindo do loop")
			break
		}

		fmt.Println("Valor:", valor)
		valor++
	}
}
