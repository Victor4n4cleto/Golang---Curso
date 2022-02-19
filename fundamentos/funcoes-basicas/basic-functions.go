package main

// apenas uma leve introdução explicativa
import "fmt"

// uma função é um bloco de codigos que executa uma atividade
func somar(a int, b int) int {
	return a + b
}

func imprimir(valor int) {
	fmt.Println(valor)
}

func main() {
	resultado := somar(3, 4)
	imprimir(resultado)
}
