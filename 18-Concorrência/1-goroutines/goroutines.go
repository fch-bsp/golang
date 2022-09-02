package main

import (
	"fmt"
	"time"
)

func main() {
	// concorrência e paralelismo
	go escrever("Olá mundo...") // chamando goroutine para executar as suas funções
	escrever("Programando em Go!")


}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}