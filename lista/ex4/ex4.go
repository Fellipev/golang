package main

import "fmt"

func main() {
	//Faça um algoritmo que receba um número inteiro e imprima na tela o seu antecessor e o seu sucessor.

	var num int

	fmt.Println("Digite um numero:")
	fmt.Scanln(&num)

	fmt.Println("O Antecessor de", num, "é", (num - 1))
	fmt.Println("O Sucessor de", num, "é", (num + 1))
}
