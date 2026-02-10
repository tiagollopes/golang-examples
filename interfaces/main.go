package main

import "fmt"

// Notificador é a nossa interface.
// Qualquer coisa que tenha o método Enviar(string) será um Notificador.
type Notificador interface {
	Enviar(mensagem string)
}

// Estrutura para Email
type Email struct {
	Endereco string
}

func (e Email) Enviar(msg string) {
	fmt.Printf("Enviando Email para %s: %s\n", e.Endereco, msg)
}

// Estrutura para WhatsApp
type WhatsApp struct {
	Numero string
}

func (w WhatsApp) Enviar(msg string) {
	fmt.Printf("Enviando Zap para %s: %s\n", w.Numero, msg)
}

// Esta função não quer saber SE é email ou zap, ela só quer um Notificador
func dispararAlerta(n Notificador, mensagem string) {
	n.Enviar(mensagem)
}

func main() {
	meuEmail := Email{Endereco: "tiago@exemplo.com"}
	meuZap := WhatsApp{Numero: "13999999999"}

	// Podemos passar ambos para a mesma função!
	fmt.Println("Sistema de Alertas:")
	dispararAlerta(meuEmail, "O servidor Okay-OS iniciou!")
	dispararAlerta(meuZap, "Alerta: Uso de CPU elevado!")
}
