package main

import "fmt"

//Forma de burlar a tipagem da linguagem, recomendado em fazer isso só em casos muito específicos.
func f(param interface{}) {
	fmt.Println(param)
}

func main() {
	f(1)
	f("OI")
	f('A')
	f(3.9)
	f(false)

	mapa := map[interface{}]interface{}{
		1:      "String",
		false:  0.8,
		"Bool": true,
	}

	fmt.Println(mapa)
}
