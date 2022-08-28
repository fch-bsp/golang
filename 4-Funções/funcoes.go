package main

import "fmt"

func somar(n1 int8, n2 int8) int8 { //criação de uma função do tipo inteiro que retorna inteiro 
	return n1 + n2 
}

func subtrair(n3 int8, n4 int8) int8 {
	return n3 - n4
}

func multplicar(n5 int8, n6 int8) int8 {
	return n5 * n6
}

func divisao(n7 int8, n8 int8) int8 {
	return n7 / n8
}

func calculosMatematicos(n1, n2 int8) (int8, int8, int8, int8) { //criação de duas variavéis com 4 retorno
	soma := n1 + n2
	subtracao := n1 - n2
	multiplicacao := n1 * n2 
	divisao := n1 /n2
	return soma, subtracao, multiplicacao, divisao
}

func main() {
	// somar funções do app
	soma := somar(10, 20) // criação de uma variavel "soma" recebe somar da func anterior com 10 e 20
	fmt.Println(soma)
    // subtração
	sub := subtrair(50,10)
	fmt.Println(sub)
	// multplicação
	mult := multplicar(10,78)
	fmt.Println(mult)
	// divisão
	div := divisao(69, 81)
	fmt.Println(div)

	total1, total2, total3, _ := calculosMatematicos(50, 23)
	fmt.Println(total1, total2,total3)
}