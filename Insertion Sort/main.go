package main

import (
	"fmt"
)

func main() {

	lista := []int{5, 4, 3, 2, 1} //lista a ser ordenada
	fmt.Println(sort(lista))      //exibe a lista ordenada
}

// função que recebe um arry de tamanho indefinido do tipo inteiro e
// retorna um arry de tamanho indefinido do tipo inteiro
func sort(lista []int) []int {

	//loop que vai interar sobre todos os elemento do array a partir da segunda posição
	for j := 1; j < len(lista); j++ {
		chave := lista[j] // armazena o elemento atual da lista

		i := j - 1 // será usado para acessar o elemento anterior em relação a chave

		//{5, 4, 3, 2, 1}
		// enquanto não chegar a posição 0 e o elemento anterior for maior que o próximo
		//5 é maior que quatro e i é maior que zero
		for i >= 0 && lista[i] > chave {
			// o próximo elemento é substituido pelo anterior
			lista[i+1] = lista[i] // neste caso o quatro é substituido pelo 5
			i--                   // decrementa i em 1
		}
		//quando elemento anterior não for maior que o próximo, então ele permanece na mesma posição
		lista[i+1] = chave
	}

	//retorna a lista ordenada
	return lista
}
