package main

import "fmt"

func main() {
	idade := map[string]int{}
	idade["Adeilson"] = 31
	idade["Ana Clara"] = 19
	fmt.Println(idade)
	fmt.Println(idade["Adeilson"])
	fmt.Println(idade["Ana Clara"])

	anoNasc := map[string]int{
		"Adeilson": 1994,
		"Ana Clara": 2007,
	}

	fmt.Println(anoNasc)
	fmt.Println(anoNasc["Adeilson"])
	fmt.Println(anoNasc["Ana Clara"])
	anoNasc["GolangDoZero"] = 2026
	fmt.Println(anoNasc)
}

/* 2 - Maps: Heterogêneos
pode misturar tipos
estrutura chave - valor
[key] = value
chave tem um tipo, e o valor pode ter outro
map[k]v -> k = chave, v = valor

map[string]int
{ "Adeilson": 31, "Ana Clara": 19}
map[string]string
{ "Adeilson": "Martins", "Ana Clara": "Jardim" }
*/
