# Busca linear

A busca linear é um método simples de encontrar um valor em uma lista percorrendo-a sequencialmente e comparando cada elemento com o valor desejado. Este método é semelhante a procurar um item em uma lista na vida real, onde você verifica cada item até encontrar o que está procurando.

## Explicação

1. Iteração sobre a lista: O algoritmo percorre cada elemento da lista sequencialmente, começando pelo primeiro elemento.

2. Comparação: Para cada elemento, ele compara o valor do elemento com o valor que estamos procurando.

3. Verificação: Se o elemento atual for igual ao valor desejado, o algoritmo retorna o índice desse elemento na lista.

4. Conclusão: Se o elemento não for encontrado após percorrer toda a lista, o algoritmo retorna -1 para indicar que o valor não está presente na lista.

## Complexidade do algoritmo

- Melhor caso: O(1) - quando o valor é encontrado no primeiro elemento da lista.
- Caso médio: O(n/2) - quando o valor é encontrado na metade da lista.
- Pior caso: O(n) - quando o valor não está presente na lista ou está presente no último elemento.

## Conclusão

Embora a busca linear seja fácil de entender e implementar, ela pode ser ineficiente para listas muito grandes, pois requer percorrer todos os elementos da lista sequencialmente. Para conjuntos de dados não ordenados ou pequenos, a busca linear pode ser uma escolha viável.
