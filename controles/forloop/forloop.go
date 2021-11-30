package main

import (
	"fmt"
	"time"
)

func main() {
	// existe apenas o laço 'for' em GO, porém podemos moldá-lo a realizar o mesmo trabalho que os demais existentes em outras linguagens

	// 1- Com o contador fora do laço (imitando um while):
	i := 1
	for i <= 10 {
		fmt.Printf("%d é menor que 10\n", i)
		i += 1
	}

	// 2- O laço 'for' bem conhecido:
	// perceba que a variável 'i' daqui é apenas no escopo do laço for, e não se mistura com o 'i' da linha 12
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d é Par\n", i)
		} else {
			fmt.Printf("%d é Ímpar\n", i)
		}
	}

	// 3- O laço infinito (basta abrir for):
	for {
		fmt.Println("Laço eterno...")
		time.Sleep(time.Second * 2) // função time.Second é simplesmente 1 segundo.
	}

	// veremos foreach e similares no bloco voltado a arrays
}
