package main

import "fmt"

func main() {

	fmt.Println("Aula de Ponteiro::")

	// PONTEIRO É UMA REFERENCIA DE MEMÓRIA

	var variavel int
	var ponteiro  *int

	variavel = 100
	ponteiro = &variavel

	fmt.Println(variavel, *ponteiro)// "*" É MESMO QUE DESREFERẼNCIAÇÃO
	variavel = 150
	fmt.Println(variavel, *ponteiro)



}