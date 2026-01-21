package main

import "fmt"

// Exemplo básico de slice em Go
func main() {

	// Criando um slice literal
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println("Slice:", numbers)
	fmt.Println("Tamanho:", len(numbers))
	fmt.Println("Capacidade:", cap(numbers))
}
