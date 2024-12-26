package main

import "fmt"

func reajuste(num int) int {

	reajuste := (5 * num) / 100
	reajuste += num

	return reajuste
}

func main() {
	//Faça um algoritmo que leia um valor qualquer e imprima na tela com um reajuste de 5%
	var num int

	fmt.Println("Digite um valor:")
	fmt.Scanln(&num)

	novoNum := reajuste(num)
	fmt.Println("O numero (", num, ") com reajuste de 5%:", novoNum)

}
