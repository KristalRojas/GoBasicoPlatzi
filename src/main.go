package main

import "fmt"

func main() {
	//Defer
	//Defer es el keyword que va a ejecutar la ultima funcion antes de que todo muera (referido al main o
	// equivalente), sirve para usar close.Connection en bd, la buena practica es no mas de un defer por funcion
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//Continue y Break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		//Continue
		//Cuando usarlo? es para cuando quieras que siga el for
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}
		//Break
		//Sirve para romper el ciclo de for
		if i == 8 {
			break
		}
	}
}
