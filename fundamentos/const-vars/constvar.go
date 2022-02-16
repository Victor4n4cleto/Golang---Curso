package main

// posso referenciar pacotes colocando uma letra antes de seus nomes na importação
import (
	"fmt"
	"math"
)

func main() {
	// jeito padrão de declarar uma constante ou variável:
	const PI float64 = 3.1415
	var raio = 3.1

	// segue a forma reduzida para criar uma variável e já executada-lá
	// forma reduzida de declarar uma variável
	area := PI * math.Pow(raio, 2)

	fmt.Println("A área da circunferencia é", area)

	// variáveis devem ser declaradas e usada. Não pode declarar uma variável e nao usa-lá !

	// forma pratica de declarar varias variáveis ou constantes:

	const (
		name = "Maria"
		age  = 18
	)

	var (
		monthSignatureOff = "Fev"
		valueSignture     = 55
	)

	fmt.Println("O preço mensal é de", valueSignture, "vecimento é:", monthSignatureOff, "!")

	// pode também ser declarado variáveis de forma extensa apenas separando e colocando em ordem o nome da variável e seu valor após o sinal de =, exemplo abaixo

	var product, price, shelfLife = "leite", 5.50, "Março"

	fmt.Println("O", product, "esta custando apenas:", price, "devido ao vencimento que será no mês de:", shelfLife)
}
