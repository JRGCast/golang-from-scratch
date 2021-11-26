package main

import (
	"fmt"
	"math"
)

func main() {
	const (
		a = 4
		b = 2
	)

	// os operadores simples são nativos da linguagem GO
	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Módulo =>", a%b)

	// também é nativo da linguagem GO operações bitwise, mais conhecidas como 'bit a bit'
	c := 3
	d := 2
	// 11 é o binário para 3 e 10 o binário para 2, logo:
	fmt.Println("E =>", c&d)   // 11 & 10 = 10
	fmt.Println("Ou =>", c|d)  // 11 | 10 = 11
	fmt.Println("Xor =>", c^d) // 11 ^ 10 = 01

	// contudo operações como raiz quadrada ou potenciação não são nativas, necessitando do pacote math
	var (
		e = 3.5
		f = 2.5
	)
	// para usar esse método em números inteiros, é necessário converter int para float64
	fmt.Println("O maior entre números os inteiros", c, "e", d, "=>", math.Max(float64(c), float64(d)))

	fmt.Println("O menor entre números os inteiros", c, "e", d, "=>", math.Min(float64(c), float64(d)))

	// logo é desnecessários se os números já forem float64
	fmt.Println("O maior entre os números com ponto flutuante", e, "e", f, "=>", math.Max(e, f))
	fmt.Println("O menor entre os números com ponto flutuante", e, "e", f, "=>", math.Min(e, f))
}
