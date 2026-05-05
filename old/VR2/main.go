package main

import (
	"fmt"
	"goexec1/compras"
	"time"
)

func main() {
	compra := compras.Compra{
		Mercado: "Mercado X",
		Data:    time.Now(),
		Itens: []compras.ItensDaCompra{
			{Nome: "Arroz"},
			{Nome: "Feijão"},
			{Nome: "Macarrão"},
		},
	}

	fmt.Printf("Compra realizada no mercado: %s, na data: %s\n", compra.Mercado, compra.Data.Format("02/01/2006"))
	fmt.Println("Itens da compra:")
	for _, item := range compra.Itens {
		fmt.Printf("- %s\n", item.Nome)
	}
}
