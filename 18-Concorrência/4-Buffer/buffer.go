package main

import (
	"fmt"
)

func main() {
	canal := make(chan string, 2) // canal com buffer

	canal <- "Olá Mundo!" // enviando um valor para o canal
	canal <- "Programando em Go!"

	mensagem := <-canal  // declarando infomrção do canal na var "mensagem"
	mensagem2 := <-canal
	
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
