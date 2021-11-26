package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	fmt.Println("Operador ==", "Banana == Banana ?", "Banana" == "Banana")
	fmt.Println("Operador !=", "Banana != Açaí?", "Banana" != "Açaí")
	fmt.Println("Operador <", "3 < 2?", 3 < 2)
	fmt.Println("Operador >", "3 > 2?", 3 > 2)
	fmt.Println("Operadores >= e <= já sabemos como funciona")

	// em go, temos a função time. O VSCode já explica bem o funcionamento (ano, mês, dia, hora, minuto, segundo, nanosegundo, localidade).
	t1 := time.Date(2020, 12, 24, 15, 20, 40, 0, time.UTC)
	fmt.Println(t1)

	// O time.Unix mede os segundos passados, iniciando em 31-12-1969, às 21h.
	data1 := time.Unix(0, 0)
	data2 := time.Unix(0, 0)

	// perceba que diferente de outras linguagens, aqui ele não faz referência a onde o objeto foi alocado na memória, mas sim a sua 'essência'
	fmt.Println(data1 == data2, reflect.TypeOf(data1), reflect.TypeOf(data2))

	// veja que isso acontece também em structs
	type Pessoa struct {
		Nome string
	}
	obj1 := Pessoa{"João"}
	obj2 := Pessoa{"João"}
	fmt.Println("Pessoas:", obj1 == obj2) 

	// é possível utilizar a função .Equal() nativa, porém a equalização no Golang é complexa
	fmt.Println(data1.Equal(data2))
}
