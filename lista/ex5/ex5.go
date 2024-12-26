package main

import "fmt"

func qtdSalarioMin(salario float32, salarioMinimo float32) int16 {

	qtd := salario / salarioMinimo

	return int16(qtd)
}

func main() {
	//aça um algoritmo que leia o valor do salário mínimo e o valor do salário de um usuário, calcule quantos salários mínimos esse
	//usuário ganha e imprima na tela o resultado. (Base para o Salário mínimo R$ 1.293,20)

	var sMinino float32 = 1.293
	var salario float32

	fmt.Println("Digite seu salario")
	fmt.Scanln(&salario)

	resultado := qtdSalarioMin(salario, sMinino)
	fmt.Println("O salario informado, eh equivalente a", resultado, "salarios minimos.")

}
