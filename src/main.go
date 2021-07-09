package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var textReverse string
	text = strings.ToLower(text)

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func main() {
	slice := []string{"hola", "que", "hace"}

	//Iterar key value, osea indice y el valor
	for key, valor := range slice {
		fmt.Println(key, " | ", valor)
	}

	//Iterar solo el valor (Nota: para ignorar alguna variable usar _)
	for _, valor := range slice {
		fmt.Println(valor)
	}

	isPalindromo("Amor a roma")

}
