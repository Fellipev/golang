package main

import "fmt"

func main() {

	var arr1 [5]int
	fmt.Println(arr1)

	arr2 := [3]string{"Index 1", "Index 2", "Index 3"}
	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3, 4}
	fmt.Println(arr3)
}
