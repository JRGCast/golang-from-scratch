package main

import (
	formatted "fmt" // o 'alias' do go é simplesmente colocá-lo antes do import
	"math"

	// se você quiser importar um pacote e não usá-lo, coloque '_' na frente, semelhante ao que fazemos no JS, por exemplo
	_ "reflect"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) é inferido pelo compilador

	// forma reduzida de criar uma variável
	// ATENÇÃO colocar os dois pontos antes do '='significa dizer que a variável (no caso area) ainda não existe, então atribui-se um valor e já inicializa numa única construção. Dê preferência a essa forma
	area := PI * math.Pow(raio, 2)

	// é possível declarar variáveis 'em cascata':

	const (
		variable1 = 1
		variable2 = 2
	)

	var (
		a = "a"
		b = "b"
	)

	// também é possível fazer a atribuição igual no Python
	const c, d = "c", "d"

	// assim como semelhante ao TypeScript
	const tipobool1 bool = false
	const tipostring1 string = "sou uma string"
	var tiponum1 int = 10

	// e é possível atribuir diversas variáveis na mesma linha. Façamos do jeito recomendado
	e, f, g := 2, false, "string"

	// ainda, vale ressaltar que go é uma linguagem fortemente tipada, logo, os tipos das variáveis serão 'imutáveis' ao longo do código, isto é, uma variável 'boolean' sempre será 'boolean'
  // ou seja, e = "string" daria erro, porém:
	e = 20 // é possível

		// o Go não vai rodar se não utilizar todas as variáveis
	formatted.Println("Área da circunferência é", area)
	formatted.Println(a, b, c, d, e, f, g, variable1, variable2, tipobool1, tipostring1, tiponum1)
}
