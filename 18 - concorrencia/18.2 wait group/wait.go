package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup

	waitGroup.Add(2) // Estabelecendo a quantidade de gorotines

	go func() { //como tem o 'go' na frente, roda concorrente
		escrever("Ola mundo!")
		waitGroup.Done() //-1
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() //-1
	}()

	waitGroup.Wait() //Esperar a contagem chegar em 0

}

func escrever(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
