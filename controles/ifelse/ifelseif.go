package main

func conceptGrade(grade float64) string {
	// no caso é melhor utilizar o switch quando há muitos if/else if, como no presente caso
	if grade >= 9 && grade <= 10 {
		return "A"
	} else if grade >= 8 && grade < 9 {
		return "B"
	} else if grade >= 7 && grade < 8 {
		return "C"
	} else if grade >= 6 && grade < 7 {
		return "D"
	} else {
		return "E"
	}
}


