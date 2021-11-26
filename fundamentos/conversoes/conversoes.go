package main

import (
	f "fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	// por se tratar de float64 (x) e int (y), o Golang não permite imprimir diretamente f.Println(x / y), por isso é necessário conversão. A maneira mais verbosa seria:

	f.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota) // int() em um numero float arredondará para baixo, independente da casa decimal
	f.Println(notaFinal, nota)

	// ATENÇÃO: string() em um número converterá para o símbolo da tabela ASCII:
	const onetwothree = 123
	asciiOfa := 97
	f.Println("string(123) dá nisso: " + string(onetwothree))
	f.Println("string(97) retornará a string 'a': " + string(asciiOfa))

	// para isso é necessário utilizar o método específico de conversão strconv, com o .Itoa (inteiro para string) ou seu inverso .Atoi:
	number_to_string := 123
	f.Println("Itoa(123) dá o que queremos: " + strconv.Itoa(number_to_string))

	// atenção que o método .Atoi retorna dois valores, a conversão e eventual erro (quanto a string não é um número, por isso a lógica necessária é):
	string_to_number_ok := "555"
	string_to_number_notok := "abc"
	num, err := strconv.Atoi(string_to_number_ok)
	// como não daria erro, também é possível fazer o num, _, ignorando o erro.

	num2, err2 := strconv.Atoi(string_to_number_notok)
	f.Println(`Atoi em "555" dá o que queremos: num =>`, num, "err =>", err)
	f.Println(`Atoi em "555" inclusive já torna possível realizar operações matemáticas de imediato: 555-554 =`, num-554)
	f.Println(`Atoi em "avc" dá erro: num =>`, num2, "err =>", err2)

	// também é possível converter booleanos, de maneira semelhante ao Atoi:

	// vale lembrar que "1" é true e "0" e demais números são false
	bol, errbol := strconv.ParseBool("true")
	bol2, errbol2 := strconv.ParseBool("truee")
	if bol {
		f.Println(`converteu a string "true" para:`, bol, "sem erros:", errbol)
	}
	if bol2 {
		f.Println(`porém não é possível converter "truee"`)
	} else {
		f.Printf(`porém não é possível converter "truee", já que dá erro: (tentando resulta em) %t, (o erro) %v`, bol2, errbol2)
	}
}
