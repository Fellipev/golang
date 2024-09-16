package main

import "fmt"

func diaSemana(num uint8) string {

	switch num {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sabado"
	default:
		return "Num inválido"
	}
}

func main() {
	dia := diaSemana(uint8(5))
	fmt.Println(dia)

}
