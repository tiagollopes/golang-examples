// range/main.go
// Exemplos básicos da cláusula 'range' em Go
// Útil para iterar em slices, arrays, maps, strings e channels
// Execute com: go run main.go

package main

import "fmt"

func main() {
	fmt.Println("=== Exemplos de range em Go ===")

	// ─────────────────────────────────────────────
	// 1. Range em SLICE (o mais comum)
	// ─────────────────────────────────────────────
	frutas := []string{"maçã", "banana", "laranja", "uva"}

	fmt.Println("1. Apenas índice:")
	for i := range frutas {
		fmt.Printf("índice %d\n", i)
	}

	fmt.Println("\n2. Índice + valor:")
	for i, fruta := range frutas {
		fmt.Printf("índice %d → %s\n", i, fruta)
	}

	fmt.Println("\n3. Só o valor (ignora índice):")
	for _, fruta := range frutas {
		fmt.Println(fruta)
	}

	// ─────────────────────────────────────────────
	// 2. Range em ARRAY (muito parecido com slice)
	// ─────────────────────────────────────────────
	numeros := [5]int{10, 20, 30, 40, 50}

	fmt.Println("\n4. Range em array:")
	for i, num := range numeros {
		fmt.Printf("posição %d = %d\n", i, num)
	}

	// ─────────────────────────────────────────────
	// 3. Range em MAP (muito usado!)
	// ─────────────────────────────────────────────
	idades := map[string]int{
		"Ana":    25,
		"Bruno":  32,
		"Clara":  19,
		"Daniel": 45,
	}

	fmt.Println("\n5. Range em map (ordem não garantida):")
	for nome, idade := range idades {
		fmt.Printf("%s tem %d anos\n", nome, idade)
	}

	// Só as chaves
	fmt.Println("\n6. Só as chaves do map:")
	for nome := range idades {
		fmt.Println(nome)
	}

	// ─────────────────────────────────────────────
	// 4. Range em STRING (itera por RUNE / caractere)
	// ─────────────────────────────────────────────
	palavra := "Golang é top!"

	fmt.Println("\n7. Range em string (índice + rune):")
	for i, r := range palavra {
		fmt.Printf("posição %2d → %q (rune %d)\n", i, r, r)
	}

	// Dica: perceba que "ç" e emojis ocupam mais de 1 byte
	emoji := "Olá 😊 mundo"
	fmt.Println("\n8. String com emoji:")
	for i, r := range emoji {
		fmt.Printf("pos %2d → %q\n", i, r)
	}

	// ─────────────────────────────────────────────
	// 5. Range em canal (channels) – básico
	// ─────────────────────────────────────────────
	fmt.Println("\n9. Range em channel (até fechar):")
	ch := make(chan string, 4)
	ch <- "café"
	ch <- "água"
	ch <- "suco"
	ch <- "chá"
	close(ch) // ← obrigatório para o range terminar

	for bebida := range ch {
		fmt.Println("Bebida:", bebida)
	}

	fmt.Println("\nFim dos exemplos básicos de range!")
}
