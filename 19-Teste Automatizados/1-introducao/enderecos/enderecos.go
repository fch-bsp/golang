package enderecos

import "strings"

// INFORMAR UMA ENDEREÇO E APLICAÇÃO TI RETORNA O QUE É TIPO ENDEREÇO OU TIPO RUA

//TipoDeEndreco DE ENDEREÇO VERIFICA SE TEM UM ENDEREÇO VALIDO E AVISA.
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"} //declaro os tipos
	// CRIAÇÃO DE SLICE
	enderecoEmLetraMinuscula := strings.ToLower(endereco) //metodo que recebe letra M e transforma m
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false //interando com tipo
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo Inválido"

}
