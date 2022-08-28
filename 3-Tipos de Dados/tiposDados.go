package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int16 = 100 // criação de uma variavél "numero" do tipo inteiro que recebe o valor 100
	fmt.Println(numero)

	var numero2 uint32 = 1000000
	fmt.Println(numero2)

	

	// alias
	// INT32 = rune
	// BYTE = UINT8

	var numero3 rune = 123456
	fmt.Println(numero3)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 3210000000.54
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.45
	fmt.Println(numeroReal3)



	// NUMEROS REAIS 

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	//char é um numero da tabela ask 

	char := "B"
	fmt.Println(char)


	// strings valor "0"

	texto := 5
	fmt.Println(texto)

	// booleano

	var booleano1 bool = true
	fmt.Println(booleano1)

	// erro 

	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Erro Interno")  // var / tipo / pacote
	fmt.Println(erro2)


}
