package mypackage

import "fmt"

//Para declarar un tipo de dato publico, se debe poner la primera letra del tipo en mayuscula, se debe agregar
//encima del tipo un comentario, que diga que es tipo publico, si el tipo comienza por minuscula, esta es private

//CarPublic Car con acceso publico (Solo la primera parte es necesaria(?))
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

func PrintMessage(message string) {
	fmt.Println(message)
}
func printMessage1(message string) {
	fmt.Println(message)
}
