package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 2.4
	y := 2
	// a linguagem pede para ser explicito o termo para conversao de dados
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)

	fmt.Println(notaFinal)

	// cuidado quando for caucatenar strings com inteiros
	// pois pode haver interpretação duvidosa, exemplo:
	// fmt.Println("Teste " + string(97))

	// inteiro para string

	// maneira correta de transformar um inteiro em string
	fmt.Println("Teste " + strconv.Itoa(123))

	// convertendo string para inteiro
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}
