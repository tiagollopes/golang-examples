package main

import "fmt"

func main() {

	// Criando um map com make
	// map[chave]valor
	idades := make(map[string]int)

	idades["Joao"] = 25
	idades["Maria"] = 30
	idades["Pedro"] = 20

	fmt.Println("Map:", idades)

	// Acessando um valor pela chave
	fmt.Println("\nIdade da Maria:", idades["Maria"])

	// Verificando se a chave existe
	idade, existe := idades["Ana"]
	if existe {
		fmt.Println("Idade da Ana:", idade)
	} else {
		fmt.Println("Ana não encontrada no map")
	}

	// Atualizando valor
	idades["Joao"] = 26
	fmt.Println("\nJoao atualizado:", idades["Joao"])

	// Removendo chave
	delete(idades, "Pedro")
	fmt.Println("\nMap após delete:", idades)

	// Percorrendo o map
	fmt.Println("\nPercorrendo o map:")
	for nome, idade := range idades {
		fmt.Printf("%s tem %d anos\n", nome, idade)
	}
}
