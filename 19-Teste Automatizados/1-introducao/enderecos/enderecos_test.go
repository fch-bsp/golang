package enderecos

import "testing"

// Teste de unidade para chamar eh= go test



type cenariosDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {  // parametro para executar test
	t.Parallel() //executa teste em paralelo

	cenariosDeTeste := []cenariosDeTeste{
		{"Rua Lago das Rosas", "Rua"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Avenida 23 de Maio", "Avenida"},
		{"Esquina treze", "Tipo Inválido"},
		{"Travessa boa menina", "Tipo Inválido"},
		{"Estrada das coisas", "Estrada"},
		{"", "Tipo Inválido"},
		
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

// executar sem precisar ir até o dir= go test ./... ou go test ou go test -v  ou 

// go test --cover inf % do test

// gerar um relatório TXT comando: go test --coverprofile nome.txt 

// para ler TXT comando go tool cover --func=nome.txt

// para ler TXT comando go tool cover --html=nome.txt



func TestQaulquer(t *testing.T) {
	t.Parallel()  //executa teste em paralelo
	if 1 > 2 {
		t.Errorf("Teste Quebrou!!")
	}
}
