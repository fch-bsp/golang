package main

import "fmt"

func main() {
	// ARITIMÉTICOS
	soma := 1 + 2
	subtração := 1 - 2
	divisao := 10 / 4
	multplicacao := 10 * 5
	restoDivisao := 10 % 2

	fmt.Println(soma, subtração, divisao, multplicacao, restoDivisao)

	var numero1 int32 = 10
	var numero2 int32 = 25
	soma1 := numero1 + numero2

	fmt.Println(soma1)

	// FIM DOS ARITIMÉTICOS


	// ATRIBUIÇÃO

	var variavel1 string = "String"
	variavel2 := "String"

	fmt.Println(variavel1, variavel2)

	// FIM DOS OPERADORES DE ATRIBUIÇÃO


	//OPERADORES RELACIONAIS

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(2 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(2 != 2)

	//FIM DOS RELACIONAIS


	// OPERADORES LOGICOS
	verdadeiro, falso := true, false

	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// FIM DOS OPERADORES LÓGICOS


	// OPERADORES UNIÁRIOS
	numero3 := 10
	numero3 ++
	numero3 += 15  // numero = numero + 15
	fmt.Println(numero3)

	numero4 := 10
	numero4 -- 
	numero4 -= 5
	fmt.Println(numero4)

	numero5 := 10
	numero5 *= 2
	fmt.Println(numero5)

	numero6 := 10
	numero6 /= 2
	fmt.Println(numero6)

	numero7 := 10
	numero7 %= 4
	fmt.Println(numero7)
	//FIM DOS OPERADORES UNIDARIO

	// OPERADOR TENÁRIO

	var texto string
	if numero > 5 {
		texto = "Maior que 5 "

	} else {
		texto = "Menor que 5 "
	} 

	fmt.Println(texto)



}