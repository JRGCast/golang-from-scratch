package main

import "fmt"

// em golang, se ambos os argumentos forem iguais, é possível escrever arg1, arg2 bool ao invés de arg1 tipo, arg2, e o retorno, se mais de um, deve ser colocado entre parênteses após os argumentos. É Possível colocar o retorno como valores nomeados, mas, por hora, deixemos assim:
func shopping(job1, job2 bool) (bool, bool, bool) {
	buyGreatTV := job1 && job2
	buyMedTV := job1 != job2 // infelizmente em GO não existe o operador conhecido como 'ou exclusivo' então a comparação feita foi com '!='
	buyIceCream := job1 || job2
	return buyGreatTV, buyMedTV, buyIceCream
}

func main() {
	tv50, tv32, iceCream := shopping(false, false)
	fmt.Printf("Tv grande? %t, Tv média? %t, Sorvete? %t, Tristeza? %t", tv50, tv32, iceCream, !iceCream)
}
