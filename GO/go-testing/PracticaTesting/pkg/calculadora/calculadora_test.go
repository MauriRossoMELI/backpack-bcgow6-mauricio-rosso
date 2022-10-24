package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 10
	num2 := 2
	resultadoEsperado := 8

	// Se ejecuta el test
	resultado := Restar(num1, num2)

	// Se validan los resultados
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")

}

//AYUDA MEMORIA
// func AddBad(t *testing.T) {
// 	// Arrange
// 	num1 := 0
// 	num2 := 5
// 	errorEsperado := fmt.Sprintf("a no puede ser:%d", num1)
// 	// Act
// 	_, err := Add(num1, num2)
// 	// Assert
// 	assert.NotNil(t, err)
// 	assert.ErrorContains(t, err, errorEsperado)
// }

// func TestAddGood(t *testing.T) {
// 	// Arrange
// 	num1 := 10
// 	num2 := 5
// 	esperado := 15
// 	add()
// 	// Act
// 	resultado, err := Add(num1, num2)
// 	// Assert
// 	assert.Equal(t, esperado, resultado, "El numero resultado: %d, es distinto del esperado: %d ", resultado, esperado)
// 	assert.Nil(t, err)
// 	//if resultado != esperado {
// 	//	t.Errorf("El numero resultado: %d, es distinto del esperado: %d ", resultado, esperado)
// 	//}
// }
