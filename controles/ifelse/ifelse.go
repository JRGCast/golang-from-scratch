package main

import "fmt"

func printResult(grade float64) {
	if grade >= 7 {
		fmt.Println("Aprovado com a nota", grade)
	} else {
		fmt.Println("Reprovado com a nota", grade)
	}
}
