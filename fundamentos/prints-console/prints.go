package main

import "fmt"

func main() {
	fmt.Print("Mesma ")
	// esse comando acima ele imprime os dados na mesma linha, ele não quebra linha.

	fmt.Println("Que viagem")
	fmt.Println("menino maluquinho")
	// esse comando acima ele imprime os dados, porém pula/quebra a linha quando imprime os dados

	x := 3.141516
	// a função abaixo converte número em string
	valueStringConvert := fmt.Sprint(x)

	fmt.Println(" o valor de x é " + valueStringConvert)

	// formas de concatenação ( + ou , )
	// simbolo de % quando utilizado na impressão busca o último valor de:
	// uma string declarada "%s"
	/* último float "%f" - colocando . e um número entre o % e o f é impresso
	somente a quantidade de casas da mesma quantidade do número informado */
	// último número inteiro "%n"
	// último booleano "%t"
	// quando utilizado o comando Printf para pular linha é necessário utilizar \n

	a := 2
	b := 2.8888
	c := true
	d := "valeu"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d) // %v é uma forma genérica para chamar dados
}
