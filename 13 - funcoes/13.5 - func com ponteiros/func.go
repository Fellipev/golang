package main

import "fmt"

func numeroInvertido(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 10
	numeroInvertido(&numero)
	fmt.Println(numero)
}
