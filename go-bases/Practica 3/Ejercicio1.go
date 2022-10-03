package main

import "fmt"

func main() {
	result := calcularImpuesto(51000)
	fmt.Println("El impuesto es de $", result)
}

func calcularImpuesto(salario float64) float64 {
	var descuento float64
	if salario > 50000 {
		if salario > 150000 {
			descuento = (27 * salario) / 100
			return salario - descuento
		} else {
			descuento = (17 * salario) / 100
			return salario - descuento
		}
	} else {
		return 0
	}
}
