package main

import "fmt"

func esPar(n int) string {
	if n%2 == 0 {
		return "Es Par"
	}
	return "Es Impar"
}

func loginAdmin(user, pass string) string {
	correctName := "Admin"
	correctPass := "AdminPass"
	if correctName == user && correctPass == pass {
		return "Logeado"
		// return true 	}
	}
	return "Usuario y/o Contraseña incorrectos!"
	// return false
}

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	//With AND
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	}

	//With OR
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Es verdad, OR")
	}

	//hacer funcion que diga si el numero es par o impar
	fmt.Println(esPar(13))
	fmt.Println(esPar(2))

	//hacer funcion de login que retorne true o false
	fmt.Println(loginAdmin("Admin", "AdminPass"))
}
