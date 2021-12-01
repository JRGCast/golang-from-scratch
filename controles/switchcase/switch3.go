package main

func switch3(i interface{}) string {
	switch i.(type) { // sintaxe utilizada para esse tipo de switch
	case int:
		return "Inteiro"
	case float32, float64:
		return "Real"
	case func():
		return "Função"
	default:
		return "Não sei"
	}
}
