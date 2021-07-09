package main

import (
	pp "curso_golang_platzi/src/pcpackage"
	"fmt"
)

// type pc struct {
// 	ram   int
// 	disk  int
// 	brand string
// }

// //Seguido del func se agrega el struct (pc) entre parentesis, asi se asocia que esa funcion es de esa struct
// func (myPc pc) ping() {
// 	fmt.Println(myPc.brand, "Pong")
// }

// func (myPc *pc) duplicateRam() {
// 	myPc.ram = myPc.ram * 2
// }

func main() {

	// “&” accede a la dirección del espacio de memoria de la variable.
	// “*” accede al valor que contiene ese espacio de memoria, dado el nombre de una variable o una dirección especifica.
	a := 50
	//Se le pasa a B la direccion del espacio en mememoria de la variable A
	b := &a

	//Valor original de A
	println(a)

	//Imprime el espacio en memoria que usa A
	println(b)

	//Imprime el valor original de A llamando a B, se le agrega el *  para acceder al valor que
	//contiene el espacio en memoria
	println(*b)

	//Al ser punteros de A, si A cambia su valor, todos los punteros tambien cambian, y viceversa
	*b = 100
	fmt.Println(a)

	var myPc pp.Pc

	myPc.Ram = 16
	myPc.Disk = 200
	myPc.Brand = "Asus"

	fmt.Println(myPc)

	//Se llama a la instancia de pc y luego se llama la funcion
	myPc.Ping()

	fmt.Println(myPc)
	myPc.DuplicateRam()

	fmt.Println(myPc)
	myPc.DuplicateRam()

	fmt.Println(myPc)

	//Reto: transcribir todo esto pero en un paquete aparte y luego llamar las funciones desde aca
}
