package main

import (
	"fmt"
	"reflect"
)

func main() {
	// a diferença entre um array e um slice é simplesmente a colocação do tamanho inicial:
	a1 := [3]int{1, 2, 3} // isto é um array
	s1 := []int{1, 2, 3}  // isto é um slice
	fmt.Println(reflect.TypeOf(a1), a1, reflect.TypeOf(s1), s1)

	fmt.Println("-----------------")

	//** Um array pode ser a união de vários slices, mas um slice não é um array! Mas define um pedaço do array, podendo ser inclusive um ponteiro:
	a2 := [9]int{1, 80, 3, 45, 52, 68, 7, 90, 100}
	s2 := a2[2:5] // aqui a primeira posição é o início, incluído no slice, e a última posição é o limite, mas não será incluída no slice, ou seja, seria o mesmo que começa do índice 2(3) e vai até uma posição antes do índice 5(68), ou seja, do índice 2(3) até o índice 4(52)
	s3 := a2[:3] // aqui omitiu-se a posição do início, então começará do início do array

	fmt.Println(a2)
	fmt.Printf("s2 = a2[0:2] => %v\n", s2)
	fmt.Printf("s3 = a2[:3] => %v\n", s3)

	// também é possível fazer um slice de um slice
	s4 := s3[1:3]
	fmt.Printf("s4 = s3[1:3] => %v\n", s4)
}
