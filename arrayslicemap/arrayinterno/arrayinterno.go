package main

import "fmt"

func main() {

	// aqui vamos comprovar que o array interno apontado é o mesmo:
	slic1 := make([]int, 10, 20)
	slic2 := append(slic1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)     // no limite do array
	slic3 := append(slic1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11) // ultrapassando o limite

	fmt.Println(slic1, slic2)                                            // [0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9 10]
	fmt.Println(slic1, "length: ", len(slic1), "capacity: ", cap(slic1)) // [0 0 0 0 0 0 0 0 0 0] length:  10 capacity:  20
	fmt.Println(slic2, "length: ", len(slic2), "capacity: ", cap(slic2)) // [0 0 0 0 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9 10] length:  20 capacity:  20
	fmt.Println(slic3, "length: ", len(slic3), "capacity: ", cap(slic3)) // 	[0 0 0 0 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9 10 11] length:  21 capacity:  40
	fmt.Println("-----------------------------------------------------")
	// alteremos apenas o array interno do slic1:
	slic1[1] = 2

	// mas perceba que, como apontam para o mesmo array interno, também foi alterado no slic2
	fmt.Println(slic1, slic2)                                            // [0 2 0 0 0 0 0 0 0 0] [0 2 0 0 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9 10]
	fmt.Println(slic1, "length: ", len(slic1), "capacity: ", cap(slic1)) // [0 2 0 0 0 0 0 0 0 0] length:  10 capacity:  20
	fmt.Println(slic2, "length: ", len(slic2), "capacity: ", cap(slic2)) // [0 2 0 0 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9 10 11] length:  21 capacity:  40
	fmt.Println("-----------------------------------------------------")

	fmt.Println(slic1, slic2)
	fmt.Println(slic1, "length: ", len(slic1), "capacity: ", cap(slic1)) // [0 2 0 0 0 0 0 0 0 0] length:  10 capacity:  20
	fmt.Println(slic2, "length: ", len(slic2), "capacity: ", cap(slic2)) // [0 2 0 0 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9 10] length:  20 capacity:  20
	fmt.Println(slic3, "length: ", len(slic3), "capacity: ", cap(slic3)) // [0 0 0 0 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9 10 11] length:  21 capacity:  40
	slic3 = append(slic1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	fmt.Println(slic3, "length: ", len(slic3), "capacity: ", cap(slic3)) // [0 2 0 0 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9 10 11] length:  21 capacity:  40
	fmt.Println("-----------------------------------------------------")
}
