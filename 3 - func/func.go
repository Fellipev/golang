package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func calculos(n1 int, n2 int) (int, int) {

	soma := n1 + n2
	sub := n1 - n2

	return soma, sub
}

func main() {
	//func podem ser um tipo de dado tb, e alem disso podem conter mais de um retorno

	soma := somar(1, 2)
	fmt.Println(soma)

	cSoma, cSub := calculos(10, 20)
	fmt.Println(cSoma, cSub)

}
