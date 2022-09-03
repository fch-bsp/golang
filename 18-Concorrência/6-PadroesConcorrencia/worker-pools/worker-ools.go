package main

import "fmt"

func main() {
	fmt.Println("Padrões de Workerpools") // fila de processo onde pega os work para ecexuta-las
	tarefas := make(chan int, 45) // canal para as tareg=fas
	resultados := make(chan int, 45) // canal para o resultado

	go worker(tarefas, resultados) // chamando a função worker 
	go worker(tarefas, resultados) // chamando a função worker
	go worker(tarefas, resultados) // chamando a função worker
	go worker(tarefas, resultados) // chamando a função worker

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultados := <- resultados
		fmt.Println(resultados)
	} 
		
		
	



}

func worker(tarefas <-chan int, resultados chan<- int ) { // função worker
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}

}

func fibonacci(posicao int)int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}