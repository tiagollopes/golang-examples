package main

import "fmt"

// Este exemplo demonstra:
// - uso de if / else
// - operadores lógicos (&&, ||, !)
// - operadores de comparação
// - condições comuns do dia a dia

func main() {

	idade := 20
	ehEstudante := true
	temIngresso := false

	fmt.Println("=== If simples ===")
	if idade >= 18 {
		fmt.Println("Adulto")
	}

	fmt.Println("\n=== if / else ===")
	if idade < 18 {
		fmt.Println("Menor de idade")
	} else {
		fmt.Println("Adulto")
	}

	fmt.Println("\n=== if / else if / else ===")
	if idade < 13 {
		fmt.Println("Criança")
	} else if idade < 18 {
		fmt.Println("Adolescente")
	} else {
		fmt.Println("Adulto")
	}

	fmt.Println("\n=== Operador lógico AND (&&) ===")
	if idade >= 18 && temIngresso {
		fmt.Println("Pode entrar no evento")
	} else {
		fmt.Println("Não pode entrar no evento")
	}

	fmt.Println("\n=== Operador lógico OR (||) ===")
	if ehEstudante || idade < 18 {
		fmt.Println("Tem direito a desconto")
	} else {
		fmt.Println("Não tem direito a desconto")
	}

	fmt.Println("\n=== Operador lógico NOT (!) ===")
	if !temIngresso {
		fmt.Println("Usuário não possui ingresso")
	}

	fmt.Println("\n=== Operadores lógicos combinados ===")
	if (idade >= 18 && temIngresso) || ehEstudante {
		fmt.Println("Acesso liberado")
	} else {
		fmt.Println("Acesso negado")
	}
}
