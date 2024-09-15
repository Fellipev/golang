package main

import "fmt"

func main() {
	slice := []int{10, 11, 12, 13}
	fmt.Println(slice)

	slice = append(slice, 14)
	fmt.Println(slice)

	arr := [...]int{1, 2, 3, 4}

	slice2 := arr[1:3] //Pega fatia do array, retornando 2 e 3
	fmt.Println(slice2)
}
