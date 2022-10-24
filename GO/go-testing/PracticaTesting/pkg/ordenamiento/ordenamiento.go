package ordenamiento

import "sort"

func OrdenarAscendente(lista []int) []int {
	sort.Ints(lista)
	return lista
}
