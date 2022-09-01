package main

import (
	"fmt"
	"time"
)
// 1 tipo de FOR
func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("incrementando...")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	// 2 TIPO DE fOR
	for j := 0; j <10; j++ {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	// fmt.Println(j)

}
