package main

import (
	"modulo/auxiliar"
	"fmt"
	"github.com/badoux/checkmail"

)	

func main() {

	fmt.Println("Escrevendo modo main")
	auxiliar.Escrever()
    // validação de e-mail 
	erro :=checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)
}