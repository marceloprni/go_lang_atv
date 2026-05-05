package main

// model e um alias para o pacote model, para evitar de ficar escrevendo model.Endereco toda hora, so Endereco ja funciona.
import (
	"fmt"
	model "golangestudo/model"
	"time"
)

func main() {
	fmt.Println("Hello, VR1!")

	endereco1 := model.Endereco{
		Rua:    "Rua A",
		Numero: 123,
		Cidade: "Cidade X",
	}

	pessoa := model.Pessoa{
		Nome:           "João",
		Endereco:       endereco1,
		DataNascimento: time.Date(1989, 12, 3, 0, 0, 0, 0, time.Local),
	}

	automovelMoto := model.Moto{
		Cilindradas: 160,
	}

	automovelCarro := model.Carro{
		QuantidadeDePortas:   4,
		Potencia:             150,
		PossuiArCondicionado: true,
	}

	//fmt.Println(endereco1)
	fmt.Println(pessoa)
	//idade := model.CalculaIdade(pessoa)
	//idadeAtual0 := pessoa.IdadeAtual()
	//pessoa.CalculaIdade()
	//fmt.Printf("A idade de %s é %d anos.\n", pessoa.Nome, idade)
	//fmt.Printf("A idade atual de %s é %d anos.\n", pessoa.Nome, idadeAtual0)
	fmt.Printf("A moto tem %d cilindradas.\n", automovelMoto.Cilindradas)
	fmt.Printf("O carro tem %d portas, %d de potencia e possui ar condicionado? %t.\n", automovelCarro.QuantidadeDePortas, automovelCarro.Potencia, automovelCarro.PossuiArCondicionado)
}
