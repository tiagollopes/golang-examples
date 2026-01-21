package main

import "fmt"

// Acessando elementos de um slice por índice
func main() {

	names := []string{"Ana", "Bruno", "Carla"}

	fmt.Println("Primeiro elemento:", names[0])
	fmt.Println("Segundo elemento:", names[1])
	fmt.Println("Último elemento:", names[len(names)-1])

	// Alterando valor pelo índice
	names[1] = "Daniel"

	fmt.Println("Slice atualizado:", names)
}
