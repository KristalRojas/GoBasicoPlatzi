package main

import (
	pk "curso_golang_platzi/src/mypackage"
	"fmt"
)

func main() {
	//se llama a Car public, el car private aparecera como que no existe
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	pk.PrintMessage("Hola Platzi")

	//Aqui ya no aparece la funcion private (printMessage1), notar que empieza con minusculas
}
