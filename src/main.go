package main

import "fmt"

//Como buena practica, al usar un channel en los parametros de una funcion, se debe indicar si el parametro sera de entrada o de salida
//Ej: channel sin indicacion: "func say(text string, c chan string) {"
//Ej: channel indica que sera de entrada: "func say(text string, c chan<- string) {"
//Ej: channel indica que sera de salida: "func say(text string, c <-chan string) {"
func say(text string, c chan<- string) {
	c <- text //c<-, es utiliza para almacenar un valor dentro del channel, este caso almacena el valor de text
}

func main() {
	//Creacion de un Channel, es buena practica señalar la cantidad de datos que almacenara el channel
	//en este caso, el 1 quiere decir que almacenara hasta 1 dato, en caso de no asignar un numero, el valor sera dinamico
	c := make(chan string, 1)

	fmt.Println("Hello")

	go say("Bye", c) //recordar que go, indica goroutine

	fmt.Println(<-c) //Al usar <-c, se esta extrayendo el valor guardado del channel
}
