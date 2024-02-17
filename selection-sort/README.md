# Selection Sort

O algoritmo de ordenação por seleção é um método simples de ordenação que funciona selecionando repetidamente o menor elemento da lista não ordenada e movendo-o para a posição correta. Ele é semelhante a como você selecionaria cartas de baralho, onde você escolhe a menor carta e a coloca na posição correta.

## Explicação

1. **Iteração sobre a lista não ordenada**: O algoritmo itera sobre toda a lista não ordenada, começando pelo primeiro elemento.

2. **Seleção do menor elemento**: Em cada iteração, o algoritmo seleciona o menor elemento da parte não ordenada da lista.

3. **Troca com o elemento atual**: Após encontrar o menor elemento, o algoritmo troca esse elemento com o elemento na posição atual da iteração.

4. **Avanço para a próxima posição**: O algoritmo avança para a próxima posição na lista não ordenada e repete os passos acima até que toda a lista esteja ordenada.

## Complexidade do algoritmo

- **Melhor caso**: O(n^2) - quando a lista está em ordem inversa.
- **Caso médio**: O(n^2) - quando a lista está em ordem aleatória.
- **Pior caso**: O(n^2) - quando a lista está em ordem inversa.

## Diferença entre Selection Sort e Insertion Sort

Embora tanto o algoritmo de ordenação por seleção quanto o algoritmo de ordenação por inserção sejam algoritmos simples de ordenação com complexidade quadrática, existem algumas diferenças entre eles:

1. **Critério de seleção**:

   - O algoritmo de seleção seleciona o menor elemento da lista não ordenada e o move para a posição correta.
   - O algoritmo de inserção seleciona um elemento da lista e o insere na posição correta, deslocando os elementos maiores para a direita.

2. **Número de trocas**:

   - O algoritmo de seleção realiza trocas apenas uma vez em cada iteração, após encontrar o menor elemento.
   - O algoritmo de inserção pode realizar várias trocas para inserir um elemento na posição correta.

3. **Estabilidade**:
   - O algoritmo de seleção não é estável, o que significa que a ordem relativa de elementos iguais pode ser alterada durante a ordenação.
   - O algoritmo de inserção é estável, o que significa que a ordem relativa de elementos iguais permanece a mesma após a ordenação.

## Conclusão

Embora o algoritmo de ordenação por seleção seja simples de entender e implementar, ele pode não ser tão eficiente quanto outros algoritmos de ordenação, especialmente para conjuntos de dados muito grandes. No entanto, é útil para pequenas listas ou como parte de algoritmos mais complexos.
