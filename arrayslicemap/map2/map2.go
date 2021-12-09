package main

import "fmt"

func main() {
	employeeNSalary := map[string]float64{
		"John Joe":     10000.50,
		"Xamara Zinc":  15555.85,
		"Jason Cory":   500.99,
		"Dorime Ameno": 800.80, // perceba que a vírgula aqui é necessária se a chave de fechamento for colocada na nova linha. Seria possível fechar com "500.99}" sem dar erro.
	}
	employeeNSalary["Joseph Birita"] = 1400.42

	// tentar excluir um elemento inexistente não gera erro, nem impacta no map:
	// fmt.Println("----------------------------------------")
	// fmt.Println("Antes da exclusão de campo inexistente:")
	// fmt.Println(employeeNSalary)
	// fmt.Println("----------------------------------------")
	// delete(employeeNSalary, "123456879")
	// fmt.Println("Depois da exclusão de campo inexistente:")
	// fmt.Println(employeeNSalary)
	// fmt.Println("----------------------------------------")

	for name, salary := range employeeNSalary {
		fmt.Printf("%s Recebe %.2f\n", name, salary)
	}
}
