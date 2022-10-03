// En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
// Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error:
//  el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que
//  “salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var salary int = 100_000

	// EXERCISE 1
	if salary < 150_000 {
		message := myCustomError.Error(myCustomError{"error: el salario ingresado no alcanza el mínimo imponible"})
		println(message)
		os.Exit(1)
	}

	//EXERCISE 2
	err2 := HasToPayTaxExercise2(salary)
	if err2 != nil {
		println(err2.Error())
		os.Exit(1)
	}

	//EXERCISE 3
	err3 := HasToPayTaxExercise3(salary)
	if err3 != nil {
		println(err3.Error())
		os.Exit(1)
	}

	println("Debe pagar impuesto.")
}

type myCustomError struct {
	ErrorMessage string
}

func (customErr myCustomError) Error() string {
	return customErr.ErrorMessage
}

func HasToPayTaxExercise2(salary int) error {
	if salary < 150_000 {
		return errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	}
	return nil
}

func HasToPayTaxExercise3(salary int) error {
	if salary < 150_000 {
		return fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
	}
	return nil
}
