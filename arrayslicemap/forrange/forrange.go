package main

import "fmt"

func main() {
	numbers := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(numbers)

	for index, number := range numbers { // em se tratando de arrays, é possível pegar o índice e os números desta forma (que já mistura a atribuição com vírgulas, operação for e range numa única coisa)
		fmt.Printf("index %d=> %d\n", index, number)
	}

	fmt.Println("--------------------------")

	// ATENÇÃO, o primeiro parâmetro sempre será o índice, não o elemento!
	words := [...]string{"alo", "bebe", "cece", "dede", "errei"}
	// assim, esse tipo de laço 'for' só retorna o índice, não as palavras
	for notwords := range words {
		fmt.Printf("É O ÍNDEX, não o elemento: %d\n", notwords)
	}
	fmt.Println("--------------------------")
	// por isso é possível ignorar o índice:
	for _, thewords := range words {
		fmt.Printf("Agora sim é o elemento do array: %s\n", thewords)
	}
}
