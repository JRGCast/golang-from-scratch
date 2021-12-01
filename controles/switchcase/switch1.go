package main

func switch1(g float64) string {
	grade := int(g)

	// veja que o switch não tem uma condição:
	switch grade {
	case 10:
		fallthrough // palavra reservada para 'continuar olhando os outros cases', é o comportamento padrão do switch em outras linguagens
	case 9:
		return "A" // em GO não é necessário colocar o 'break'
	case 8, 7: // meio que uma mistura do Python com JS, não?
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}
