package main

import "fmt"

// --- EXEMPLO SIMPLES ---
// Função que tenta dobrar um número sem usar ponteiros
func dobrarValor(n int) {
	n = n * 2
}

// Função que dobra o valor REAL usando ponteiros
func dobrarComPonteiro(n *int) {
	*n = *n * 2 // O asterisco aqui acessa o valor no endereço
}

// --- EXEMPLO MEDIANO ---
// Em Go, passar structs grandes por ponteiro evita cópias desnecessárias na memória
type PerfilUsuario struct {
	Nome  string
	Nivel int
}

// Método que promove o usuário. Usamos (u *PerfilUsuario) para alterar o original.
func promoverUsuario(u *PerfilUsuario) {
	u.Nivel++ // Go facilita: não precisamos de (*u).Nivel++, ele entende o atalho
}

func main() {
	// 1. Teste Simples (Ponteiros com Tipos Primitivos)
	numero := 10

	dobrarValor(numero)
	fmt.Println("Após dobrarValor (sem ponteiro):", numero) // Continua 10

	dobrarComPonteiro(&numero) // Passamos o endereço usando &
	fmt.Println("Após dobrarComPonteiro (com &):", numero)  // Agora é 20

	fmt.Println("--------------------------------")

	// 2. Teste Mediano (Ponteiros com Structs)
	player := PerfilUsuario{Nome: "Tupã", Nivel: 1}

	fmt.Printf("Antes da promoção: %+v\n", player)

	// Passando o endereço da struct para a função
	promoverUsuario(&player)

	fmt.Printf("Depois da promoção: %+v\n", player)

	// Dica: O endereço de memória onde o player está guardado:
	fmt.Printf("Endereço de memória do player: %p\n", &player)
}
