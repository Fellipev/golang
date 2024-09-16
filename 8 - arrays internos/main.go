package main

import "fmt"

func main() {

	slice := make([]int8, 2, 3) //tipo de dados do slice, lenght e capacidade
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3) //atingiu a capacidade, quando ele atinge ele pega a capacidade + 1 e dobra, se tornando entao 8
	fmt.Println("- SLICE -")

	fmt.Println(cap(slice)) // 8 de capacidade
}
