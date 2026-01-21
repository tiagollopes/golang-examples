package main

import "fmt"

// Criando slices com make
func main() {

	// make(tipo, tamanho, capacidade)
	numbers := make([]int, 3, 5)

	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30

	fmt.Println("Slice:", numbers)
	fmt.Println("Tamanho:", len(numbers))
	fmt.Println("Capacidade:", cap(numbers))

	// Adicionando elementos
	numbers = append(numbers, 40, 50)

	fmt.Println("Após append:", numbers)
	fmt.Println("Nova capacidade:", cap(numbers))
}
