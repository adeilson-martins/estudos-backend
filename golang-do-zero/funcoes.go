package main

import "fmt"

func soma(a, b int) int {
	return a + b
}

func SomaSoma(a, b, c int) int {
	return a + b + c
}

func main() {
	res := soma(1, 2)
	fmt.Println("1+2 =", res)
	res = SomaSoma(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}

/* Função começando com letra minúscula: FUNÇÃO PRIVADA
Ela só pode ser utilizada no próprio pacote.

Função começando com letra maiuscula: FUNÇÃO PÚBLICA
Ela pode ser utilizada fora do próprio pacote.

Como utilizaria ela fora do pacote:
Ex. main.nome da função(tipo da função Ex. int) */
