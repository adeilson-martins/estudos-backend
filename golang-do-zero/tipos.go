package main

import "fmt"

func main() {
	// bool - (verdadeiro ou falso)
	fmt.Printf("Type: %T - Value: %v\n", true, true)

	// string - (sequência de bytes)
	fmt.Printf("Type: %T - Value: %v\n", "Adeilson", "Adeilson")
	fmt.Printf("Type: %T - Value: %v\n", "11", "11")

	// int - (inteiro)
	fmt.Printf("Type: %T - Value: %v\n", 1, 1)

	// float - (float64/float32 - decimal)
	fmt.Printf("Type: %T - Value: %v\n", 1.28, 1.28)
}
