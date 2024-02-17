package main

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	lista := []int{5, 2, 4, 6, 1, 3}

	sort := selectionSort(lista)

	esperado := []int{1, 2, 3, 4, 5, 6}

	if !reflect.DeepEqual(sort, esperado) {
		t.Errorf("Recebido: %v | Esperado: %v", sort, esperado)
	}
}
