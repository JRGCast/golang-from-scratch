package main

func main() {
	x := 1
	y := 2

	// em go só existe a operação 'pós-fixadas', isto é, o incremento só vem após o operando, ou seja não existe, por exemplo ++x ou --y.

	x++ // forma padrão de x+=1
	y-- // forma padrão de y-=1

	// não é possível fazer, por exemplo, print(x == y--)
}
