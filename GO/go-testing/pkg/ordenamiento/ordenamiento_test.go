package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	lista := []int{3, 4, 1}
	resultadoEsperado := []int{1, 3, 4}

	// Se ejecuta el test
	resultado := OrdenarAscendente(lista)

	// Se validan los resultados
	assert.Equal(t, resultadoEsperado, resultado, "Debe estar ordenado ascendente")

}
