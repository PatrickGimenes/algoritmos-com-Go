package main

import "fmt"

func main() {

	lista := []int{45, 84, 78, 78, 75, 15}

	valor := 15

	fmt.Println(linearSearch(lista, valor))

}

func linearSearch(lista []int, v int) int {
	index := 0
	for i := 0; i < len(lista); i++ {

		if lista[i] == v {
			index = i
		} else {
			index = -1
		}
	}
	return index
}
