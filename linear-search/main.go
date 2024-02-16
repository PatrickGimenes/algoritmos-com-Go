package main

import "fmt"

func main() {

	lista := []int{45, 84, 78, 78, 75, 15}

	valor := 78

	fmt.Println(linearSearch(lista, valor))

}

func linearSearch(lista []int, v int) int {
	index := -1
	for i := 0; i < len(lista); i++ { //itera sobre todos os elementos da lista

		if lista[i] == v { //verifica se v é igual ao elemento da lista[i]
			index = i //se for, armazena o index
		}
	}
	return index
}
