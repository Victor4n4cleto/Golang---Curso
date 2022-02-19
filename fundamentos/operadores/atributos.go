package main

import "fmt"

func main() {

	var b byte = 3
	fmt.Println(b)

	i := 3 // inferência de tipo
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 3 // i = i / 2
	i *= 3 // i = i * 2
	i %= 2 // i = i % 2

	fmt.Println(i)

	// atribuindo o valor da variável e executando
	x, y := 1, 2

	// invertendo os valores
	x, y = y, x
}
