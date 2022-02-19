package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("String:", "banana" == "Banana") // se a primeira palavra é igual a segunda
	fmt.Println("!=", 3 != 5)                    // se a primeira é diferente da segunda
	fmt.Println("<", 3 < 6)                      // menor
	fmt.Println(">", 5 > 8)                      // maior
	fmt.Println("<=", 6 <= 9)                    // menor ou igual
	fmt.Println(">=", 8 >= 10)                   // maior ou igual

	d1 := time.Unix(0, 0) // nesse caso esta sendo considerado o valor
	d2 := time.Unix(0, 0) // nesse caso esta sendo considerado o valor

	fmt.Println("Datas", d1 == d2)
}
