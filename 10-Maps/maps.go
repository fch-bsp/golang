package main

import "fmt"

func main() {
	fmt.Println("Maps:::")

	usuario := map[string]string{
		"nome": "Pedo",
		"sobrenome": "Silva",

	}

	fmt.Println(usuario ["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro nome": "João",
			"segundo nome": "Pedo",
		},

		"curso": {
			"curso 1": "João",
			"segundo nome": "Pedo",
		},


		
	}

	fmt.Println(usuario2)
	

         
}



