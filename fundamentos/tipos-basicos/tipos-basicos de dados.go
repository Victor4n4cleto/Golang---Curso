package main

// tipo de dados básicos
import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	// números inteiros
	fmt.Println(1, 2, 1500)

	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))
	// reflect.TypeOf informa o tipo de dado selecionado

	// sem sinal (só positivos)... uint8 = 8bits; uint16 = 2 bytes; uint32= 4bytes; uint64 = 8bytes

	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de 1l é", reflect.TypeOf(i1))

	var i2 rune = 'a' // rune representa um mapeamento da tabela unicode (int32)

	fmt.Println("O rune é", reflect.TypeOf(i2))

	// números reais (float32, float64)
	var x float32 = 49.99
	// var x = float32(49.99) - outra possibilidade
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49,99 é", reflect.TypeOf(49.99)) // float64

	//boolean
	bo := true
	fmt.Println("O tipo da variável é", reflect.TypeOf(bo))
	fmt.Println(!bo) // negação

	// string
	s1 := "Olá mundo"
	fmt.Println("O tamanho da string é", len(s1))

	// string com multiplas linhas utilizando craze
	s2 := `Olá
	mundo
	baum ?`
	fmt.Println("O tamanho da string é", len(s2))

	// char = int32
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
}
