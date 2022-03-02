package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Repovado com nota", nota)
	}
}
func main() {
	imprimirResultado(7.9)
	imprimirResultado(6.8)
}
