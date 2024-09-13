package main

import "fmt"

func main() {
	//Go é uma linguagem fortemente tipada e caso alguma var seja declarada e não usada o código da erro.

	var v1 string = "Var 1" //primeira forma de declara uma var
	v2 := "Var 2"           //segunda forma

	var ( //terceira forma
		v3 string = "Var 3"
		v4 string = "Var 4"
	)

	v5, v6 := "Var 5", "Var 6" //quarta forma

	fmt.Println(v1, v2, v3, v4, v5, v6)

}
