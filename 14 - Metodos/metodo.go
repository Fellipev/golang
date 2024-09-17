package main

import "fmt"

type usuario struct {
	nome  string
	idade int8
}

func (u usuario) toString() {
	fmt.Printf("nome: %s", u.nome)
	fmt.Printf("\nidade: %d", u.idade)
}

func (u usuario) setIdade(idade int8) bool {
	if idade > 0 {
		u.idade = idade
		fmt.Printf("\nIdade alterada para %d", u.idade)
		return true
	}
	fmt.Print("idade invalida.")
	return false
}

func main() {
	u := usuario{"Fellipe", 21}
	u.toString()
	u.setIdade(10)
}
