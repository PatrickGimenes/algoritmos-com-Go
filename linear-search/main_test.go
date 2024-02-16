package main

import "testing"

func TestSeach(t *testing.T) {

	lista := []int{45, 84, 78, 78, 75, 15}

	valor := 75

	index := linearSearch(lista, valor)

	esperado := 4

	if esperado != index {
		t.Errorf("Recebido: %v | Esperado: %v", index, esperado)
	}

}
