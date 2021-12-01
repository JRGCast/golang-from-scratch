package main

import "fmt"

func main() {
	// em GO, o array é homogêneo (mesmo tipo) e estático (tamanho fixo). Apesar de comum em outras linguagens, go não modifica o tamanho dos arrays originais 'naturalmente' (é possível fazer com slices)

	// há algumas formas de se inicializar um array:
	// forma mais utilizada variável := [length]tipo{valor1, valor2, ..., valorN}:
	arrNum := [5]int{1, 2, 3, 4, 5}
	// se não colocar nada, lembra-te dos tipos 'default' de cada um.
	arrNum2 := [5]int{}
	arrNum3 := [...]int{}
	fmt.Println(arrNum)  // [1 2 3 4 5]
	fmt.Println(arrNum2) // [0 0 0 0]
	fmt.Println(arrNum3) // []
	arrNum2[4] = 5
	fmt.Println(arrNum2) // [0 0 0 5]

	var grades [3]float64
	fmt.Println(grades) // [0 0 0]

	grades[0], grades[1], grades[2] = 4.4, 7.8, 9.5
	fmt.Println(grades) // [4.4 7.8 9.5]

	// vamos calcular a média utilizando o laço for simples:

	total := 0.0
	for i := 0; i < len(grades); i++ { // aqui o len é idêntico ao Python
		total += grades[i]
	}
	average := total / float64(len(grades)) // necessário adequar a length de grades (que é um int) ao total (que é um float64)

	fmt.Println(average)                  // média não 'filtrada' no ponto flutuante
	fmt.Printf("Média = %.2f\n", average) // lembrando o "toFixed" de GO
}
