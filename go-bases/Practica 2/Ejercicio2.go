package main

//import "fmt"

func main() {
	edad := 23
	es_empleado := true
	antiguedad := 2
	sueldo := 100000

	if edad > 22 {
		if es_empleado {
			if antiguedad > 1 {
				println("¡Prestamo otorgado!")
				if sueldo > 100000 {
					println("No deberás pagar intereses.")
				} else {
					println("Deberás pagar intereses.")
				}
			} else {
				println("Préstamo rechazado: No tienes la suficiente antiguedad.")
			}
		} else {
			println("Préstamo rechazado: No eres empleado.")
		}
	} else {
		println("Préstamo rechazado: No tienes más de 22 años.")
	}
}
