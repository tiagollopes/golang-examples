package main

import "fmt"

// Constantes simples
const language = "Go"
const version = 1.22

// Constantes agrupadas
const (
	pi      = 3.14159
	maxUser = 100
	active  = true
)

// Constantes com iota
const (
	sunday = iota
	monday
	tuesday
	wednesday
	thursday
	friday
	saturday
)

func main() {
	fmt.Println("Language:", language)
	fmt.Println("Version:", version)

	fmt.Println("PI:", pi)
	fmt.Println("Max users:", maxUser)
	fmt.Println("Active:", active)

	fmt.Println("Days of week:")
	fmt.Println(sunday, monday, tuesday, wednesday, thursday, friday, saturday)
}