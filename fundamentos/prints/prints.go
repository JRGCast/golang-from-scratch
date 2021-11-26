package main

import f "fmt"

func main() {
	f.Print("Veja que eu estou ")
	f.Print("na mesma linha.")

	f.Println(" Aqui ainda estou naquela linha") // a partir do primeiro Println, começa a quebrar linhas
	f.Println("Esta aqui porém não")
	f.Println("E aqui já quebra linha")

	// formas de concatenar strings:
	pi := 3.141516

	// 1 - criando uma variável que 'stringfy' o valor
	piString := f.Sprint(pi)
	f.Println("Sprint + Println: O valor de pi é " + piString)
	// 2 - adicionando como segundo parâmetro de Println
	f.Println("Println puro: O valor de pi é", pi)

	// na função Printf é impresso na mesma linha, o % indica onde será inserido o argumento, com a consoante seguinte indicando seu tipo ('%s' para string, '%f' para 'float', '%t' para booleano, '%v' para valor genérico, etc.)
	f.Printf("Printf comum: O valor de pi é %f\n", pi)
	f.Printf("Printf comum: %f é o valor de pi\n", pi)
	// também há um 'toFixed' embutido no Printf
	f.Printf("Prinft com 'toFixed(2)': O valor de pi é %.2f", pi)
}
