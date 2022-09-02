package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá Mundo!", canal)

	fmt.Println("Depois da função escrever começar a ser executada")
    // opção 01
	// for {
	// 	mensagem, aberto := <-canal //aguardando um valor no canal
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }
    // opção 02
	for mensagem := range canal {
		fmt.Println(mensagem)
	}
	fmt.Println("Fim do Programa!")

}

func escrever(texto string, canal chan string) {
	// time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		canal <- texto

		time.Sleep(time.Second)
	}
	close(canal)
}
