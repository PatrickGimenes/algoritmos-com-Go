package main

import "fmt"

func main() {

	lista := []int{5, 2, 4, 6, 1, 3}

	fmt.Println(selectionSort(lista))

}

func selectionSort(lista []int) []int {

	tamanho := len(lista) //armazena o tamanho total da lista

	for i := 0; i < tamanho-1; i++ { //itera por todos elementos da lista a partir do primeiro elemento

		menor := i                         //armazena o índice do menor numero
		for j := i + 1; j < tamanho; j++ { // itera por todos elementos da lista a partir do segundo elemento
			if lista[j] < lista[menor] { //verifica qual é o menor elemento
				menor = j //armazena o índice do menor elemento
			}
		}

		if menor != i {
			lista[i], lista[menor] = lista[menor], lista[i] //troca os elementos de posição
		}
	}

	return lista
}
