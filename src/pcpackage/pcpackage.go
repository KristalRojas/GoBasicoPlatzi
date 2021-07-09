package pcpackage

import "fmt"

//PcPublic  <--- Si es necesario el comentario
type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

func (myPc Pc) Ping() {
	fmt.Println(myPc.Brand, "Pong")
}
func (myPc *Pc) DuplicateRam() {
	myPc.Ram = myPc.Ram * 2
}
