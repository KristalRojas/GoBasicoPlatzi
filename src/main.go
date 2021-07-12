package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "Mensaje1"
	c <- "Mensaje2"

	fmt.Println(len(c), cap(c)) //len indicara la cantidad de datos que tiene, cap indicar la capacidad maxima de channel

	//Close, es una buena practica cerrar el channel cuando llegue a la capacidad maxima
	//Es necesario cerrar el channel para recorrerlo, asi ya no espera ningun dato
	close(c)

	// c <- "Mensaje3"

	//Range (Recorrido)
	for message := range c {
		fmt.Println(message)
	}

	//Select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("Mensaje1", email1)
	go message("Mensaje2", email2)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1: ", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email1: ", m2)
		}
	}
}
