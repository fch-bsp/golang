package main

import "fmt"

// opção 1 do Switch
func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda Feira"
	case 3:
		return "Terça Feira"
	case 4:
		return "Quarta Feira"
	case 5:
		return "Quinta Feira"
	case 6:
		return "Sexta Feira"
	case 7:
		return "Sabado"
	default:
		return "Número Inválido"
	}
}	
// opção 2 do Switch
func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda Feira"
	case numero == 3:
		return "Terça Feira"
	case numero == 4:
		return "Quarta Feira"
	case numero == 5:
		return "Quinta Feira"
	case numero == 6:
		return "Sexta Feira"
	case numero == 7:
		return "Sabado"
	default:
		return "Número Inválido"
	}

}

func main() {

	dia := diaDaSemana(70)
	fmt.Println(dia)

	dia2 := diaDaSemana2(5)
	fmt.Println(dia2)

}
