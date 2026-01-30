package main

import "fmt"

// Exemplo de map com slice como valor
// map[chave][]valor
func main() {

	// Criando um map onde a chave é o nome da pessoa
	// e o valor é um slice de notas
	notas := make(map[string][]int)

	// Adicionando valores
	notas["Joao"] = []int{7, 8, 9}
	notas["Maria"] = []int{10, 9, 8}

	// Adicionando mais notas usando append
	notas["Joao"] = append(notas["Joao"], 10)

	fmt.Println("Map com slices:")
	fmt.Println(notas)

	// Percorrendo o map e os slices
	fmt.Println("\nNotas dos alunos:")
	for nome, listaNotas := range notas {
		fmt.Println("Aluno:", nome)

		for i, nota := range listaNotas {
			fmt.Printf("  Prova %d: %d\n", i+1, nota)
		}
	}

	// Calculando a média de cada aluno
	fmt.Println("\nMédia dos alunos:")
	for nome, listaNotas := range notas {
		soma := 0

		for _, nota := range listaNotas {
			soma += nota
		}

		media := float64(soma) / float64(len(listaNotas))
		fmt.Printf("%s -> média: %.2f\n", nome, media)
	}
}
