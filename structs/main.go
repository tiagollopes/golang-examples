package main

import "fmt"

// Exemplo 1: Struct simples
type Pessoa struct {
    Nome   string
    Idade  int
    Altura float64
    Ativo  bool
}

// Exemplo 2: Struct com struct aninhada (embedding)
type Endereco struct {
    Rua    string
    Numero int
    Cidade string
}

type Cliente struct {
    Pessoa
    Endereco
    Saldo float64
}

// Exemplo 3: Método (receiver) associado à struct
func (p Pessoa) Saudacao() string {
    return fmt.Sprintf("Olá, meu nome é %s e tenho %d anos.", p.Nome, p.Idade)
}

func (c Cliente) Apresentar() {
    fmt.Printf("Cliente: %s\n", c.Nome)
    fmt.Printf("Idade: %d anos\n", c.Idade)
    fmt.Printf("Endereço: %s, %d - %s\n", c.Rua, c.Numero, c.Cidade)
    fmt.Printf("Saldo atual: R$ %.2f\n", c.Saldo)
}

func main() {
    // Usando struct simples
    p1 := Pessoa{
        Nome:   "Tiago",
        Idade:  30,
        Altura: 1.78,
        Ativo:  true,
    }

    fmt.Println("--- Exemplo 1 ---")
    fmt.Printf("%+v\n", p1)
    fmt.Println(p1.Saudacao())

    // Usando struct com embedding
    c1 := Cliente{
        Pessoa: Pessoa{
            Nome:  "Maria",
            Idade: 25,
        },
        Endereco: Endereco{
            Rua:    "Av. Brasil",
            Numero: 123,
            Cidade: "Santos",
        },
        Saldo: 450.75,
    }

    fmt.Println("\n--- Exemplo 2 + Método ---")
    c1.Apresentar()
}
