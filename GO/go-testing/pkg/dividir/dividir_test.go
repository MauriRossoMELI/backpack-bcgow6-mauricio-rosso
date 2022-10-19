package dividir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 10
	num2 := 2
	resultadoEsperado := 5
	// Se ejecuta el test
	resultado, err := Dividir(num1, num2)

	// Se validan los resultados
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
	assert.Nil(t, err)
}

func TestZeroDenominator(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 10
	num2 := 0
	errorEsperado := "el denominador no puede ser 0"

	// Se ejecuta el test
	_, err := Dividir(num1, num2)

	// Se validan los resultados
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, errorEsperado)
}
