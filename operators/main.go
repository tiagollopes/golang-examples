package main

import "fmt"

func main() {
	a := 10
	b := 3

	// ======================
	// Operadores aritméticos
	// ======================
	fmt.Println("Soma:", a+b)
	fmt.Println("Subtração:", a-b)
	fmt.Println("Multiplicação:", a*b)
	fmt.Println("Divisão:", a/b)
	fmt.Println("Resto:", a%b)

	// ======================
	// Operadores relacionais
	// ======================
	fmt.Println("a == b:", a == b)
	fmt.Println("a != b:", a != b)
	fmt.Println("a > b:", a > b)
	fmt.Println("a < b:", a < b)
	fmt.Println("a >= b:", a >= b)
	fmt.Println("a <= b:", a <= b)

	// ======================
	// Operadores lógicos
	// ======================
	fmt.Println("true && false:", true && false)
	fmt.Println("true || false:", true || false)
	fmt.Println("!true:", !true)

	// ======================
	// Operadores de atribuição
	// ======================
	c := 5
	c += 2 // c = c + 2
	fmt.Println("c += 2:", c)

	c -= 1 // c = c - 1
	fmt.Println("c -= 1:", c)

	c *= 3 // c = c * 3
	fmt.Println("c *= 3:", c)

	c /= 2 // c = c / 2
	fmt.Println("c /= 2:", c)

	// ======================
	// Operadores de incremento/decremento
	// ======================
	d := 10
	d++
	fmt.Println("d++:", d)

	d--
	fmt.Println("d--:", d)
}
