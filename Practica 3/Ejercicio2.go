package main

import (
	"errors"
	"fmt"
)

func main() {
	promedio, err := calcularPromedio(9, 9, 7)
	if err != nil {
		println(0, err.Error())
	} else {
		fmt.Printf("%f", promedio)
		fmt.Println()
	}
}

func calcularPromedio(calificaciones ...int) (promedio float64, err error) {
	suma := 0
	for _, calif := range calificaciones {
		if calif < 0 {
			err = errors.New("Error: todas las calificaciones deben ser mayores a 0.")
			return
		}
		suma += calif
	}
	promedio = float64(suma) / float64(len(calificaciones))
	return promedio, nil
}
