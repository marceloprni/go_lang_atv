package main

import "time"

type compra struct {
	mercado string
	data    time.Time
	itens   []itensDaCompra
}

type itensDaCompra struct {
	nome string
}
