package main

import "fmt"

func main() {
	modulo := 5 % 2
	switch modulo {
	case 0:
		fmt.Println("Es Par")
	default:
		fmt.Println("Es impar")
	}
}
