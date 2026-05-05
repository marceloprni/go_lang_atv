package compras

import "time"

type Compra struct {
	Mercado string
	Data    time.Time
	Itens   []ItensDaCompra
}

type ItensDaCompra struct {
	Nome string
}
