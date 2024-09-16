package main

import "fmt"

func main() {

	usuario := map[string]string{
		"nome":      "Fellipe",
		"sobrenome": "Reis",
	}

	fmt.Println(usuario)

	u2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Fellipe",
			"meio":     "Vieira",
			"ultimo":   "Reis",
		},
		"curso": {
			"nome":   "ADS",
			"status": "concluido",
		},
	}

	u2["trabalho"] = map[string]string{
		"cargo":   "Analista de sistemas web jr.",
		"empresa": "Unisa",
	}

	delete(u2, "trabalho")

	fmt.Println(u2["nome"]["primeiro"])
}
