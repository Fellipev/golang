package main

import "fmt"

func recuperar() {
	if r := recover(); r != nil {
		fmt.Println("Tentativa de recuperar a execucao.")
	}
}

func media(n1 int, n2 int) bool {
	defer recuperar()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("MEDIA E 6 AAAAAAAAA")
}

func main() {
	resultado := media(6, 6)
	fmt.Println(resultado)
}
