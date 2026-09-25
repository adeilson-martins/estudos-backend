package main

import "fmt"

func main() {
	// Variáveis
	var nome string = "Adeilson" // string - (sequência de bytes)
	var idade int = 31           // int - (inteiro)
	var altura float64 = 1.70    // float - (float64/float32 - decimal)
	var ativo bool = true        // bool - (verdadeiro ou falso)

	// Declaração implícita com := (forma curta)
	sobrenome := "Martins"
	peso := 65.51

	// Constantes
	const linguagem = "Golang"

	//Exibindo as Variáveis
	fmt.Println(nome)
	fmt.Println(idade)
	fmt.Println(altura)
	fmt.Println(ativo)
	fmt.Println(sobrenome)
	fmt.Println(peso)
	fmt.Println(linguagem)
}
