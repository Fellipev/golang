package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Ola mundo!") // goroutine
	escrever("Programando em Go!")

	//Apesar do looping infinito, ele vai executar as 2 funcoes pois ele executa a primeira em goroutine
	//o que faz ele nao para o programa enquanto a outra esta sendo executada, ele continua lendo o programa.
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
