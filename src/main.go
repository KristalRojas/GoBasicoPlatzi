package main

import "fmt"

type pc struct {
	ram   int
	brand string
	disk  int
}

//Stringer
func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d DB Disco y es una %s", myPc.ram, myPc.disk, myPc.brand)
}

func main() {
	myPc := pc{ram: 16, brand: "Asus", disk: 100}

	fmt.Println(myPc)
}
