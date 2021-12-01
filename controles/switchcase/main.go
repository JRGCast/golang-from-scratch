package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Switch 1 (switch simples)")
	fmt.Println(switch1(9.8))
	fmt.Println(switch1(7.8))
	fmt.Println(switch1(5.8))
	fmt.Println(switch1(3.8))
	fmt.Println(switch1(2.8))
	fmt.Println(switch1(-1))
	fmt.Println(switch1(19.8))
	fmt.Println("------------------------")
	fmt.Println("Switch 2 (switch sem 'parâmetro')")
	switch2()
	fmt.Println("------------------------")
	fmt.Println("Switch 3 (switch avançado [trabalhando com tipos])")
	fmt.Println(switch3(2.3))
	fmt.Println(switch3(2))
	fmt.Println(switch3("Teste string"))
	fmt.Println(switch3(func() {}))
	fmt.Println(switch3(time.Now()))
}
