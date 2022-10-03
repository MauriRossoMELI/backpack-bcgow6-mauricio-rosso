package main

import "os"

func main() {
	_, err := os.OpenFile("./customers.txt", os.O_RDONLY, 0644)

	defer func() { //Se ejecuta cuando termina la funcion en donde esta declarada.
		println("Ejecución finalizada")
	}()

	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado.")
	}
}
