package main

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	lista := []int{5, 2, 4, 6, 1, 3}

	sort := sort(lista)

	esperado := []int{1, 2, 3, 4, 5, 6}

	if !reflect.DeepEqual(sort, esperado) {
		t.Errorf("Recebido: %v | Esperado: %v", sort, esperado)
	}
}

func TestReverseSort(t *testing.T) {
	lista := []int{1, 2, 3, 4, 5}

	reverse_sort := reverse_sort(lista)

	esperado := []int{5, 4, 3, 2, 1}

	if !reflect.DeepEqual(reverse_sort, esperado) {
		t.Errorf("Recebido: %v | Esperado: %v", reverse_sort, esperado)
	}
}
