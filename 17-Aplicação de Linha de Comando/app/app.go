package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retonar aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Buscar IPS e Nomes de Servidor na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},

	}		
	

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca de IPS de endereços na Internet",
			Flags: 	flags,
			Action: buscarIps,
		},
		{
			Name: "servidores",
			Usage: "BUscar o nome dos srevidores na internet",
			Flags: flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) { // criado a função de buscar IPS
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func buscarServidores(c *cli.Context) { // criado a função de buscar servidores
	host := c.String("host")

	servidores, erro := net.LookupNS(host) //name server
	if erro != nil {
		log.Fatal(erro)

	} 
	
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}

}

// 1 comando para buscar servidores "go run main.go ip --host amazon.com.br"
// 2 comando para buscar servidores "go run main.go servidores --host amazon.com.br"

// arquivo cobilado:

//opção 1: ./linha-de-comando ip --host amazon.com.br

//opção 2: ./linha-de-comando servidores --host amazon.com.br




