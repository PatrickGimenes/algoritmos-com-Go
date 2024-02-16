# Insertion Sort

O algoritmo de ordenação por inserção é um dos algoritmos simples de ordenação que funciona bem em pequenos conjuntos de dados. Ele funciona da mesma maneira que muitas pessoas naturalmente ordenam cartas de baralho em suas mãos: uma carta por vez.

## Explicação

O algoritmo de ordenação por inserção é baseado em comparar elementos adjacentes e trocá-los até que a lista esteja completamente ordenada. Aqui está uma visão geral do funcionamento do algoritmo:

1. Iteração sobre a lista: O algoritmo itera sobre todos os elementos da lista, começando do segundo elemento.

2. Comparação e inserção: Para cada elemento, ele compara o elemento atual com os elementos à sua esquerda e o insere na posição correta.

3. Mover elementos: Se o elemento atual for menor do que o elemento à sua esquerda, ele é movido para a esquerda até encontrar a posição correta.

4. Inserção: Finalmente, o elemento é inserido na posição correta, e o algoritmo avança para o próximo elemento.

Caso queira ver o algoritmo em funcionamente, veja esse [vídeo](https://youtu.be/8oJS1BMKE64?si=oQGcooR2xIovLyej).

## Complexidade do algoritmo

- Melhor caso: O(n) - quando a lista já está ordenada.
- Caso médio: O(n^2) - quando a lista está em ordem aleatória.
- Pior caso: O(n^2) - quando a lista está em ordem inversa.

## Conclusão

Embora o algoritmo de ordenação por inserção seja simples e fácil de implementar, ele não é tão eficiente quanto outros algoritmos de ordenação para conjuntos de dados muito grandes devido à sua complexidade quadrática.

## Sobre o Go

Achei muito interessante o fato do Go não possui While, comumente encontrada em outras linguagens de programação. Em vez disso, o Go utiliza uma abordagem simplificada, onde é possível reproduzir o comportamento de um loop while usando a estrutura for.

Exemplo em C:

```c
#include <stdio.h>

int main() {
    int contador = 0;

    // Loop enquanto o contador for menor que 5
    while (contador < 5) {
        printf("O contador é: %d\n", contador);
        contador++; // Incrementa o contador
    }

    return 0;
}
```

Exemplo em Go:

```go
package main

import "fmt"

func main() {

	contador := 0
    // Loop enquanto o contador for menor que 5
	for contador <= 5 {
		fmt.Println("O contador é:", contador)
		contador++ // Incrementa o contador
	}
}
```

No exemplo acima, podemos ver que, embora o Go não tenha uma construção while, é possível obter o mesmo efeito usando um for com uma expressão de condição semelhante à do while. Essa abordagem simplifica a linguagem e mantém sua sintaxe clara e concisa.
