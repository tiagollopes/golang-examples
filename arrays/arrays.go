package main

import "fmt"

func main() {
    // Para leigos: explicamos que um Array/Slice é como uma gaveta com várias divisórias

    // Criando uma lista (slice) de strings
    itens := []string{"Arroz", "Feijão", "Azeite", "Café"}

    fmt.Println("--- Minha Lista de Compras ---")

    // Adicionando um item novo na lista
    itens = append(itens, "Açúcar")

    // Mostrando cada item da lista e sua posição (índice)
    for i, item := range itens {
        fmt.Printf("Item %d: %s\n", i+1, item)
    }

    fmt.Println("------------------------------")
    fmt.Printf("Total de itens na lista: %d\n", len(itens))
}
