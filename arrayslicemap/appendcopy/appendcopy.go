package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// array1 = append(array1, 4, 5, 6) esse código não é possível pois a função append só serve para juntar ao final de um slice

	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1) // [1 2 3] [4 5 6]

	slice2 := make([]int, 2) // foi criado um slice com apenas 2 posições
	copy(slice2, slice1)     // copia do segundo argumento para o primeiro, geralmente de um slice source para um slice destination, mas também é possível copiar de uma string para um slice de bytes

	fmt.Println(slice1, slice2) // veja que por somente conter 2 espaços, o slice 2 é [4 5], então o copy copia, mas não expande o tamanho do slice destino
}
