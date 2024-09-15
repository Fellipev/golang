package main

import "fmt"

func main() {
	var v1 int = 10
	var v2 int = v1

	fmt.Println(v1, v2) //10 10

	v2++

	fmt.Println(v1, v2) //10 11

	//Ponteiro

	var v3 int
	var ponteiro *int //ponteiro criado

	v3 = 100
	ponteiro = &v3

	fmt.Println(v3, ponteiro) //100 0xc0001a...

	v3 = 150

	fmt.Println(v3, *ponteiro) //150 150
}
