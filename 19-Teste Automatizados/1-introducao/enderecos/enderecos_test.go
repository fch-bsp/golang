package enderecos

import "testing"

// Teste de unidade para chamar eh= go test

type cenariosDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTeste := []cenariosDeTeste{
		{"Rua Lago das Rosas", "Rua"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Avenida 23 de Maio", "Avenida"},
		{"Esquina treze", "Tipo Inválido"},
		{"Travessa boa menina", "Tipo Inválido"},
		{"Estrada das coisas", "Estrada"},
		{"", "Tipo Invalido"},
		
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido: %s é diferente do esperado %s",
				tipoDeEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}

}
