package main

import "fmt"

func usuario(txt string) (texto string, ano int32) {
	texto = "Texto da func: " + txt
	ano = 2003
	return
}

func main() {

	texto, ano := usuario("AAAA")
	fmt.Println(texto, ano)
}
