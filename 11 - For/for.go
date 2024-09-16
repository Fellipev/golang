package main

import (
	"fmt"
)

func main() {
	i := 0

	for i < 10 {
		i++
		// fmt.Println(i)
		// time.Sleep(time.Second)
	}

	for j := 0; j < 10; j++ {
		// fmt.Println(j)
	}

	nomes := [3]string{"Joao", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	user := map[string]string{
		"nome":      "Fellipe",
		"sobrenome": "Reis",
	}

	for chave, valor := range user {
		fmt.Println(chave, valor)
	}
}
