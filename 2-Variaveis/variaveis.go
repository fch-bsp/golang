package main

import "fmt"


// criação de uma função main e criação de uma variavél com nome de "variavel1" do tipo string 
func main() {
	var variavel1 string = "variavel1" // tipo 1 de criação de variavel
	variavel2 := "variavel2" // tipo 2 de criação de variavel inferencia de tipo 
	
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// criação de variavel multiplas com os dois tipo 

	// tipo 1

	var (
		variavel3 string = "alguma coisa"
		variavel4 string = "alguma coisa"
	)
    // tipo 2 

	variavel5, variavel6 := "alguma coisa", "alguma coisa"

	fmt.Println(variavel3, variavel4)
	fmt.Println(variavel5, variavel6)
}