package main

import "fmt"

// Este exemplo demonstra:
// - for simples
// - for com condição (estilo while)
// - for com sub-for (for dentro de for)

func main() {

	fmt.Println("=== For simples ===")
	for i := 1; i <= 5; i++ {
		fmt.Println("Valor de i:", i)
	}

	fmt.Println("\n=== For com condição (tipo while) ===")
	contador := 1
	for contador <= 5 {
		fmt.Println("Contador:", contador)
		contador++
	}

	fmt.Println("\n=== For com sub-for (for dentro de for) ===")
	for linha := 1; linha <= 3; linha++ {
		fmt.Println("Linha:", linha)

		for coluna := 1; coluna <= 4; coluna++ {
			fmt.Println("  Coluna:", coluna)
		}
	}

	fmt.Println("\n=== Exemplo prático: tabuada ===")
	for numero := 1; numero <= 3; numero++ {
		fmt.Println("Tabuada do", numero)

		for multiplicador := 1; multiplicador <= 5; multiplicador++ {
			fmt.Println(numero, "x", multiplicador, "=", numero*multiplicador)
		}
	}
}
