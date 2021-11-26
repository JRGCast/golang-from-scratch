package main

import (
	f "fmt"
	"math"
	r "reflect"
)

func main() {
	// números inteiros
	f.Println(1, 2, 1000)
	// tipo genérico é int, que será a inferência padrão
	f.Println("Literal inteiro é", r.TypeOf(36000))

	// tipos sem sinal (somente positivos) : uint8 uint16 uint32 uint64 (o 'u' significa unsinght), cada número final está mais relacionado ao byte do número.

	const byt byte = 255 // o tipo byte é um alias para uint8, que vai de 0 a 255
	maxUin16 := math.MaxInt16
	maxUin32 := math.MaxInt32
	maxUin64 := math.MaxInt64

	f.Println("tipo de byt, que ", byt, "é", r.TypeOf(byt))
	f.Println("tipo de uint16", maxUin16, "é", r.TypeOf(maxUin16))
	f.Println("tipo de uint32", maxUin32, "é", r.TypeOf(maxUin32))
	f.Println("tipo de uint64", maxUin64, "é", r.TypeOf(maxUin64))

	// o tipo 'rune' é um alias para o int32, que retorna o unicode do que foi atribuído. Por exemplo o unicode do char 'a' é 97, 'b' é 98. Porém não existe o tipo 'chat' em Go
	const runetype1 rune = 'a'
	const runetype2 rune = 'b'

	f.Println("tipo de runetype1", runetype1, "é", r.TypeOf(runetype1))
	f.Println("tipo de runetype2", runetype2, "é", r.TypeOf(runetype2))

	// ainda existe os tipos chamados números reais (float32 e float64)
	var n32 float32 = 32
	f.Println("tipo de n32", n32, "é", r.TypeOf(n32))
	// geralmente números com pontos flutuantes será float64
	var n64 = 64.99
	f.Println("tipo de n64", n64, "é", r.TypeOf(n64))
	// mas é possível 'burlar' colocando tipo
	var burlan64 float32 = 65.99
	f.Println("tipo de burlan64", burlan64, "por conta do tipo colocado, é", r.TypeOf(burlan64), "porém, naturalmente o tipo seria", r.TypeOf(65.99))

	// de resto há os conhecidos bool, podendo ser negado
	boleano := true
	f.Println("tipo de boleano", boleano, "assim como seu oposto", !boleano, "é", r.TypeOf(boleano))

	// e string, que no go é exclusivamente por aspas duplas, podendo concatenar sem problemas.
	stringman := "sou uma string"
	f.Println("tipo de stringman", stringman, "é", r.TypeOf(stringman))
	// aqui existe o backtick (aspas) com efeito idêntico ao JS
	thebacktick := `Este é
	o poder do backtick`

	f.Println("tipo de thebacktick", thebacktick, "é", r.TypeOf(thebacktick))
	// o length é semelhante ao python
	f.Println(stringman+"O tamanho da string é", len(stringman))

	// por fim, os tipos 'default' de cada um, incluindo o tipo ponteiro:

	var (
		a int
		b float64
		c bool
		d string
		e *int
	)
  // em Go, null é nil
	f.Printf("tipo default de: a=%v, b=%v, c=%v, d=%q, e=%v", a, b, c, d, e)
	// utilizou-se %q apenas para mostrar a string vazia
}
