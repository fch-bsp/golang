package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	espaco    string
	altura    uint8
}

type estudante struct {
	pessoa    // INFORMANDO STRUCT "pessoa"
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Dados dos Funcionários: ")

	p1 := pessoa{"Fernando |", "Horas |", 39, "|", 179}
	p2 := pessoa{"João", "Mario", 45, "|", 189}
	p3 := estudante{p1, "TI", "Uninove"}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

}
