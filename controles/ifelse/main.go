package main

import "fmt"

func main() {
	fmt.Println("ifelse:")
	printResult(7.5)
	printResult(5.6)
	fmt.Println("---------------")
	fmt.Println("ifelseif:")
	fmt.Println(conceptGrade(9.5))
	fmt.Println(conceptGrade(8.5))
	fmt.Println(conceptGrade(7.5))
	fmt.Println(conceptGrade(6.5))
	fmt.Println(conceptGrade(5.9))
	fmt.Println("---------------")
	fmt.Println("ifinit:")
	// é possível inicializar um bloco if de maneira semelhante ao laço for, isto é, a variável inicial que fazemos (for(let i = 0[...])), podemos fazer no if:
	if init := randomNum(); init > 5 { // também podemos fazer isso no switch
		fmt.Printf("%d é acima de 5!\n", init)
	} else {
		fmt.Printf("%d é abaixo de 5!\n", init)
	}
	// perceba que a variável init não existe fora do escopo do if:
	// fmt.Printf("%d é abaixo de 5!\n", init) dá erro de 'init undeclared name'
}
