package main

import "fmt"

type usuario struct {
	nome   string
	idade  int8
	altura float32
}

func main() {
	var u usuario
	u.nome = "Fellipe"
	u.idade = 21
	u.altura = 1.72

	u2 := usuario{"Giovanna", 21, 1.68}

	fmt.Println(u, u2)
}
