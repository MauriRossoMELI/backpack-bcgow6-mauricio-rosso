package main

//import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	println("La edad de Benjamin es: ", employees["Benjamin"], " años.")
	mayoresDe21 := 0
	for _, value := range employees {
		if value > 21 {
			mayoresDe21++
		}
	}
	println("Cantidad de empleados mayores de 21 años: ", mayoresDe21)
	employees["Federico"] = 25
	println("Empleado Federico (25) agregado correctamente.")
	delete(employees, "Pedro")
	println("Empleado Pedro eliminado correctamente.")
}
