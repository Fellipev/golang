package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	for i := uint(1); i <= 10; i++ {
		fmt.Println(fibonacci(i))
	}
}
