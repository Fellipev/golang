package main

import "fmt"

type pessoa struct {
	nome     string
	idade    int8
	altura   float32
	endereco endereco //assim é uma relação de classe
}

type endereco struct {
	logradouro string
	numero     int
}

type estudante struct {
	pessoa    //Herança feita
	curso     string
	faculdade string
}

func main() {
	var u pessoa
	u.nome = "Fellipe"
	u.idade = 21
	u.altura = 1.72

	u2 := pessoa{"Giovanna", 21, 1.68, endereco{"Giovanni Gronchi", 30}}

	e := estudante{u, "ADS", "SENAC"}
	e2 := estudante{u2, "Odontologia", "UNESP"}

	fmt.Println(u, u2)
	fmt.Println("________")
	fmt.Println(e, e2)
}
