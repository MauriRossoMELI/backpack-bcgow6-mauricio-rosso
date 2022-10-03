package main

import "fmt"

func main() {
	numeroMes := 4
	switch numeroMes {
	case 1:
		fmt.Println("Enero")
	case 2:
		fmt.Println("Febero")
	case 3:
		fmt.Println("Marzo")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Mayo")
	case 6:
		fmt.Println("Junio")
	case 7:
		fmt.Println("Julio")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Septiembre")
	case 10:
		fmt.Println("Octubre")
	case 11:
		fmt.Println("Noviembre")
	case 12:
		fmt.Println("Diciembre")
	}
}

//Respuesta punto 2:
//Este ejercicio se puede resolver mediante un gran número de sentencias "If else" de forma anidadas.
//Elijo el método del switch ya que es: más eficiente, más legible y más práctico al momento de programar.
