package main

import "fmt"

func main() {
	//DECLARACION DE CONSTANTES
	const pi float64 = 3.14
	const pi2 = 3.1415
	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	//DECLARACION DE VARIABLES ENTERAS
	base := 12
	var altura int = 14
	var area int
	println("base: ", base)
	println("altura: ", altura)
	println("area: ", area)

	//ZERO VALUES
	var a int
	var b float64
	var c string
	var d bool

	println(a, b, c, d)

	//Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	println("Area cuadrado es: ", areaCuadrado)
}
