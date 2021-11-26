package main

import "fmt"

// repare a semelhança com o TypeScript
// quanto há explicitado que a função retorna algo, o return é obrigatório
func sum(x int, y int) int {
	return x + y
}

// já quando não há retorno explicitado, não é necessário o return
func printing(value int) {
	fmt.Println(value)
}

// Se as funções estiverem num mesmo 'escopo' de diretório, é possível utilizar as funções em outro arquivo. Vamos copiar a função main() abaixo para outroarquivo.go

// func main() {
// 	result := sum(1, 2)
// 	printing(result)
// }
