package main

import "fmt"

func main() {

	numero := 20 // criei uma Var que valor de 20

	if numero >= 15 { // se a var for maior que 15 vc faz isso:
		fmt.Println("Maior que 15")
	} else { // caso contrario vc faiz isso
		fmt.Println("Menor ou igual a 10")
	}

	if outroNumero := numero; outroNumero > 1 {
		fmt.Println("Numero é maior que zero, o numero é:", outroNumero)
	}else {
		fmt.Println("Numero é menor que zero, o numero é:", outroNumero)
	}

}
