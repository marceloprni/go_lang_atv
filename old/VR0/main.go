package main

import "fmt"

func main() {
	x := 5
	y := &x // ponteiro
	*y = 10 //* colocar para mudar o valor do ponteiro, quando voce altera o y tambem altera o x
	fmt.Println("Main")
	fmt.Println(x, y)
	fmt.Println(&x, y)
	ImprimirValores(x, *y)

	Ponteiro(&x, y)
	fmt.Println("Main2")
	fmt.Println(x, y)
	fmt.Println(&x, y)
}

func ImprimirValores(x int, y int) {
	fmt.Println("ImprimirValores")
	fmt.Println(x, y)
	fmt.Println(&x, &y)
}

func Ponteiro(x *int, y *int) {
	*x = 15
}
