package main

import "fmt"

func checarNum(num int16) string {

	var txt string

	if num%2 == 0 {
		txt = "O numero e par "
	} else {
		txt = "O numero e impar "
	}

	if num >= 0 {
		txt += "e positivo"
	} else {
		txt += "e negativo"
	}

	return txt
}

func main() {
	//Faça um algoritmo para receber um número qualquer e imprimir na tela se o número é par ou ímpar, positivo ou negativo

	var num int16

	fmt.Println("Digite um numero:")
	fmt.Scanln(&num)

	response := checarNum(num)
	fmt.Println(response)

}
