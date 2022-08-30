package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Criação de um Arrays::")

	var array1 [5]string
	array1[0] = "P1"
	array1[1] = "P2"
	array1[2] = "P3"
	array1[3] = "P4"
	array1[4] = "P5"
	fmt.Println(array1)

	// segunda opção
	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	fmt.Println("Criação de Slices::")

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)
	// importanto um b devolver um tipo

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array2))
	fmt.Println(reflect.TypeOf(array1))

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	//Arrays Internos
	fmt.Println("-----------------")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) // capacidade






}
