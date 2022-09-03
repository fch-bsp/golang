package enderecos

import "strings"

//TipoDeEndreco DE ENDREÇO VERIFICA SE TEM UM ENDEREÇO VALIDO E AVISA
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodivia"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return primeiraPalavraDoEndereco
	}

	return "Tipo Inválido"

}
