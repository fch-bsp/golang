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
	
    // criando daos de "PESSOAS"
	p1 := pessoa{"Fernando |", "Horas |", 39, "|", 179}
	fmt.Println(p1)
	p2 := pessoa{"João", "Mario", 45, "|", 189}
	fmt.Println(p2)
	// criando daos de "ESTUDANTE"
	e1 := estudante{p1, "TI", "Uninove"}
	fmt.Println(e1)
	e2 := estudante{p2,"Engenharia", "USB"}
	fmt.Println(e2.altura,e2.nome)
	

}
