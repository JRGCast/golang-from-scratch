package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 15
	s = make([]int, 10, 20)
	// a função make cria um SLICE, podendo criar, internamente, um array. No caso a função é make([]tipo do array, tamanho do SLICE, tamanho do array interno que necessariamente precisa ser >= que o tamanho do slice)
	// veja que a função len (de length) referencia o tamanho do slice, já a função cap (de capacity) referencia a capacidade total do array referenciado
	fmt.Println(s, "length: ", len(s), "capacity: ", cap(s)) // [0 0 0 0 0 0 0 0 0 0] length: 10 capacity: 20
	// vamos testar appendando +10 elementos ao slice:
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println(s, "length: ", len(s), "capacity: ", cap(s)) // [0 0 0 0 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9 10] length:  20 capacity:  20
	// Atenção: o que foi 'preenchido' foi o SLICE, isto é, o ponteiro, mas o array interno continua exatamente igual.
	// Agora um ponto interessante do GO: ao extrapolar o limite do slice, um novo array interno é referenciado, a própria linguagem gerencia isso:

	s = append(s, 11)
	fmt.Println(s, "length: ", len(s), "capacity: ", cap(s)) // [0 0 0 0 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9 10 11] length:  21 capacity:  40
}
