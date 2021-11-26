package main

import "fmt"

// infelizmente não colocaram operador ternário em GO =(
func gradeOK(grade float64) string {
	if grade >= 6 {
		return "Aprovado"
	} else {
		return "Reprovado"
	}
}

func main() {
	fmt.Println(gradeOK(6.5))
}
