package main

import "fmt"

func main() {
	// Faça um algoritmo que leia os valores de A, B, C e em seguida imprima na tela a soma entre A e B é mostre se a soma é menor que C.
	var a, b, c int

	fmt.Println("Digite o valor de A:")
	fmt.Scanln(&a)
	fmt.Println("Digite o valor de B:")
	fmt.Scanln(&b)
	fmt.Println("Digite o valor de C:")
	fmt.Scanln(&c)

	soma := a + b
	fmt.Println("A soma entre A e B é:", soma)

	if soma < c {
		fmt.Println("A soma entre A e B é menor que C")
	} else if soma == c {
		fmt.Println("A soma entre A e B é igual a C")
	} else {
		fmt.Println("A soma entre A e B é maior que C")
	}
}
