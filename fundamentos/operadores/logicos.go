package main

import "fmt"

func compras(trabOne, trabTwo bool) (bool, bool, bool) {
	comprarTvGrande := trabOne && trabTwo
	comprarTvPequena := trabOne != trabTwo // ou exclusivo
	comprarSorvete := trabOne || trabTwo

	return comprarTvGrande, comprarTvPequena, comprarSorvete

}

func main() {
	tvGrande, tvPequena, sorvete := compras(true, true)
	fmt.Printf("tvGrande: %t, tvPequena: %t, sorvete: %t, Saudável: %t",
		tvGrande, tvPequena, sorvete, !sorvete)
}
