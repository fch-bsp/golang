package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero uint16

}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Fernando"
	u.idade = 41
	fmt.Println(u)


	u15 := endereco{"Lago das rosas", 0}

	// OPÇÃO 2 REFERENCIA DE TIPO

	u2 := usuario{"Fernando", 41, u15}
	fmt.Println(u2, u15)

	// OPÇÃO 3 REFERENCIA DE TIPO

	u3 := usuario{nome: "Fernando"}
	fmt.Println(u3)

	// OPÇÃO 4 REFERENCIA DE TIPO

	u4 := usuario{idade: 41}
	fmt.Println(u4)

}
