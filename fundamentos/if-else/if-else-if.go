package main

import "fmt"

func notaParaConceito(N float64) string {
	if N >= 9 && N <= 10 {
		return "A"
	} else if N >= 7 && N < 9 {
		return "B"
	} else if N >= 5 && N < 8 {
		return "C"
	} else if N >= 3 && N < 5 {
		return "D"
	} else {
		return "Reprovado"
	}
}

func main() {

	fmt.Println(notaParaConceito(8.9))
	fmt.Println(notaParaConceito(7.8))
	fmt.Println(notaParaConceito(6.7))
}
