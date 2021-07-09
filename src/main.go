package main

import "fmt"

func main() {

	//Maps (El map seria una lista de llave valor, en python seria un diccionario)
	//Creacion de map (make se utiliza para crear distintos tipos de variables)
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	//Recorrer map
	for key, value := range m {
		fmt.Println(key, value)
	}

	//Encontrar valor | retorna 14 que es la key de Jose
	v := m["Jose"]
	fmt.Println(v)

	//Encontrar valor que no existe | Este retornara el Zero value del tipo que sea el valor
	//en este caso value es de tipo int, por lo que retornara cero
	v = m["Josep"]
	fmt.Println(v)

	//Para saber si se encuentra la item con la key, agregar una variable seguido de la variable que representa
	//value, esa variable retornara true o false, segun corresponda
	v, ok := m["Josep"]
	fmt.Println(v, ok)

}
