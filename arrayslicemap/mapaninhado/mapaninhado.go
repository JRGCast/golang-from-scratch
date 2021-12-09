package main

import "fmt"

func main() {
	// dificilmente se trabalhará com esse tipo de estrutura, já que o struct é uma maneira muito melhor para organizar, mas, saiba que existe tal possibilidade:
	funcsPerLetter := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 1546.79,
			"Guga Pereira":   8456.99,
		},
		"J": {
			"Joseph Pelintro": 10000.56,
			"Jamil Bonfire":   8000.66,
		},
		"P": {
			"Paleontolog Habil": 5000.99,
			"Pleiades Honor":    105069.33,
		},
		"X": {"Xamanic": 8888888.99},
	}
	delete(funcsPerLetter, "X")

	for letra, funcs := range funcsPerLetter {
		fmt.Printf("Funcionários com letra %s:\n", letra)
		for employee, salary := range funcs {
			fmt.Printf("%s Recebe %f\n", employee, salary)
		}
	}
}
