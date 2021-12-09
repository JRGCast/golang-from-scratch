package main

import "fmt"

func main() {
	// var aproved1 map[int]string maps NECESSARIAMENTE devem ser inicializados
	aproved1 := make(map[int]string) // a função make serve também para inicializar objetos do tipo slice, map ou chan

	// a estrutura map se inicia com [tipoChave]tipoValor, logo, é [chave]valor:
	aproved1[123456789] = "Maria"
	aproved1[987654321] = "Pedro"
	aproved1[915618596] = "Joe"

	// também, como já vimos, pode ser inicializada com valores já determinados
	aproved2 := map[string]int{"John": 123456789, "Xangrila": 45678910, "XABLAU": 12131415}

	// o map se assemelha ao objeto JS:
	fmt.Println(aproved1)
	fmt.Println(aproved2)

	// veja que no laço for ele dividirá por chave,valor:
	for cpf, nome := range aproved1 {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	for nome, num := range aproved2 {
		fmt.Printf("%s (CPF: %d)\n", nome, num)
	}

	// semelhante ao objeto, só é possível acessar o valor pela chave:
	fmt.Println(aproved1[987654321]) // Pedro
	fmt.Println(aproved2["XABLAU"])  // 12131415

	// para remover uma chave, utilizamos a função delete(map, chave):
	fmt.Println("-------------------")
	fmt.Println("ANTES DA EXCLUSÃO:")
	fmt.Println(aproved1)
	fmt.Println(aproved2)
	fmt.Println("-------------------")

	delete(aproved1, 915618596)
	delete(aproved2, "XABLAU")
	fmt.Println("DEPOIS DA EXCLUSÃO:")
	fmt.Println(aproved1)
	fmt.Println(aproved2)
	fmt.Println("-------------------")

	// assim como no objeto, se passar um novo valor para uma chave existente, ele será sobrescrito:
	aproved1[123456789] = "SOBRESCREVEU MARIA!"
	aproved2["XANGRILA"] = 10000000000000000
	fmt.Println("DEPOIS DE SOBRESCREVER:")
	fmt.Println(aproved1)
	fmt.Println(aproved2)
}
