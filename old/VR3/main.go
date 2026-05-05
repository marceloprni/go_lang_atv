package main

import (
	"fmt"
	"time"
)

func main() {

	var itens1 []itensDaCompra

	itens1 = append(itens1, itensDaCompra{nome: "Arroz"})
	itens1 = append(itens1, itensDaCompra{nome: "Feijão"})
	itens1 = append(itens1, itensDaCompra{nome: "Macarrão"})

	compra1 := compra{
		mercado: "Mercado do Zé",
		data:    time.Now(),
		itens:   itens1,
	}

	fmt.Printf("Compra: %+v\n", compra1)

}

// go run main.go compras.go para rodar tudo
