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
	x := 10
	y := 50

	//Suma
	result := x + y
	println("Suma: ", result)

	//Resta
	result = y - x
	println("Resta: ", result)

	//Multiplicacion
	result = x * y
	println("Multiplicacion: ", result)

	//Division
	result = y / x
	println("Division: ", result)

	//Modulo
	result = y % x
	println("Modulo: ", result)

	// Incremental
	x++
	println("Incremental: ", x)

	// Decremental
	x--
	println("Decremental: ", x)

	//RETOS
	//RECTANGULO, TRAPECIO Y DE UN CIRCULO

	//Area Rectangulo
	const largo = 2
	const ancho = 4

	resultArea := largo * ancho

	println("Area Rectangulo: ", resultArea)

	//Area Trapecio
	baseBaja := 10
	baseAlta := 4

	resultArea = ((baseBaja + baseAlta) * altura) / 2
	println("Area Trapecio: ", resultArea)

	//Area Circulo
	var radio = 5

	var resultCircle float64
	resultCircle = (pi * (float64(radio) * float64(radio)))

	println("Area Circulo: ", resultCircle)
}
