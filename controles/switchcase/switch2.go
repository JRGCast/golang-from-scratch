package main

import (
	"fmt"
	"time"
)

func switch2() {
	t := time.Now()

	switch { // funciona como se fosse uma espécie de 'while(true)', mas a execução funciona como um 'find', isto é, ficará procurando o primeiro case/condição que retorne 'true'
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")
	}
}
