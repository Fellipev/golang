package main

import "fmt"

func media(n1 int, n2 int) string {
	defer println("Media calculada, aguarde o resultado")

	media := (n1 + n2) / 2
	if media >= 6 {
		return "APROVADO!"
	}

	return "Reprovado."
}

func main() {

	r := media(7, 2)
	fmt.Println(r)
}
