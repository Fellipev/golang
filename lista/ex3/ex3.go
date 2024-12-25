package main

import "fmt"

func operacao(num1 int, num2 int) int {
	var result int

	if num1 == num2 {
		result = num1 + num2
	} else {
		result = num1 * num2
	}

	return result
}

func main() {
	//Faça um algoritmo que leia dois valores inteiros A e B, se os valores de A e B forem iguais, deverá somar os dois valores,
	//caso contrário devera multiplicar A por B. Ao final de qualquer um dos cálculos deve-se atribuir o resultado a uma variável C e
	//imprimir seu valor na tela

	var a, b, c int

	fmt.Println("Digite o primeiro valor:")
	fmt.Scanln(&a)
	fmt.Println("Digite o segundo valor:")
	fmt.Scanln(&b)

	c = operacao(a, b)
	fmt.Println(c)

}
