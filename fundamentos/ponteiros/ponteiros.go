package main

import "fmt"

func main() {
	i := 1

	// Go não tem aritmética de ponteiro

	var p *int = nil

	p = &i         // pegando o endereço na memória de i com o comando '&i' e agora p aponta para aquela referência/endereço na memória
	fmt.Println(p) // 0xc***********

	*p++ // o * na frente desreferencia para pegar o valor
	fmt.Println(*p)

	// como a operação *p++ mexeu no exato local da memória da variável 'i', perceba:
	*p++
	i++
	fmt.Printf("Endereço 'p'=> %v;\nendereço '&i' => %v;\nvalor de *p=> %v;\nvalor de i => %v;\n", p, &i, *p, i)

	fmt.Printf("Ou seja 'p == &i'?=> %v;\nE '*p == i'? %v;\n", p == &i, *p == i)
}
