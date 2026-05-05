package model

// isso e heranbça, tem que ser letra maiuscula para ser acessado de fora do pacote, se fosse letra
// minuscula, so dentro do pacote model que poderia acessar.
type Automovel struct {
	Ano    int
	Placa  string
	Modelo string
}

type Moto struct {
	Automovel
	Cilindradas int
}

type Carro struct {
	Automovel
	QuantidadeDePortas   int
	Potencia             int
	PossuiArCondicionado bool
}
