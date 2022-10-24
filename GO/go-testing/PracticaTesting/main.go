package main

import (
	"fmt"

	"github.com/backpack-bcgow6-mauricio-rosso/go-testing/pkg/calculadora"
	"github.com/backpack-bcgow6-mauricio-rosso/go-testing/pkg/ordenamiento"
)

func main() {
	a, b := 10, 5
	resta := calculadora.Restar(a, b)
	fmt.Println(resta)
	lista := []int{1, 4, 3}
	ordenamiento.OrdenarAscendente(lista)
	fmt.Println(resta)
}
