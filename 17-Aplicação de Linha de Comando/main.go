package main

import (
	"fmt"
	"linha-de-comando/app" //linha de comando nome do módulo e o app aplcação que foi criada
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de Partida")

	aplicacao := app.Gerar()
	if erro:= aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
