package main

import "fmt"

// Este exemplo demonstra:
// - uso do break dentro do for
// - uso do continue dentro do for
// - uso de if para controlar o fluxo
// - diferença entre break e continue

func main() {

	fmt.Println("=== Exemplo de break ===")
	for i := 1; i <= 10; i++ {

		if i == 5 {
			fmt.Println("Encontrou o valor 5, parando o loop")
			break
		}

		fmt.Println("Valor de i:", i)
	}

	fmt.Println("\n=== Exemplo de continue ===")
	for i := 1; i <= 10; i++ {

		if i%2 == 0 {
			fmt.Println("Pulando número par:", i)
			continue
		}

		fmt.Println("Número ímpar:", i)
	}

	fmt.Println("\n=== Exemplo combinando if, break e continue ===")
	for i := 1; i <= 10; i++ {

		if i == 3 {
			fmt.Println("Pulando o número 3")
			continue
		}

		if i == 8 {
			fmt.Println("Encontrou o número 8, encerrando o loop")
			break
		}

		fmt.Println("Processando número:", i)
	}
}
