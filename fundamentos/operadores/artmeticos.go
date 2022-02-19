package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	// o que cada operador faz:
	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Módulo =>", a%b) // resto da divisão

	// bitwise - ele pega o numero inteiro e trabalha com ele na sua forma de bit
	fmt.Println("E =>", a&b) // esse operador faz operações bit a bit

	fmt.Println("OU =>", a|b)

	fmt.Println("Xor =>", a^b)

	// outros tipos de operações pacote math... funciona com float

	c := 5.2
	d := 7.1

	fmt.Println("Maior =>", math.Max(c, d))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Exponenciação =>", math.Pow(c, d))

}
