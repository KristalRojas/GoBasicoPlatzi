package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	//Instanciacion de un objeto, con ambos parametros
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	//Instanciacion alternativa, al no ingresar todos los parametros del objeto, este quedara en su zero value
	//en este ejemplo: no se ha indicado el valor de year, por lo cual ese parametro queda en cero.
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}
