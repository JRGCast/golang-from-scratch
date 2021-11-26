package main

import "fmt"

func main() {
	// já aprendemos as atribuições possíveis:
	var a byte = 3
	const b = "string"
	c := true
	fmt.Println(a, b, c)

	// e as demais funcionam de maneira similar as outras linguagens
	i := 3
	fmt.Println("i é", i)
	i += 1
	fmt.Println("+=1 =>", i)
	i -= 1
	fmt.Println("-=1 =>", i)
	i *= 2
	fmt.Println("*=2 =>", i)
	i /= 2
	fmt.Println("/=2 =>", i)
	i %= 2
	fmt.Println("%=2 =>", i)

	// também é possível substituir variáveis
	x, y := 3, 80
	fmt.Println("x é", x, "y é", y)
	// invertendo os valores
	fmt.Println("Invertendo os valores...")
	x, y = y, x
	fmt.Println("agora x é",x, "e y é", y)
}
