package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup //WaitGruop acumula un conjunto de goroutines y las va liberando de a poco

	fmt.Println("Hello")
	wg.Add(1) //Se agrega una goroutine

	go say("world", &wg) //Keyword go, quiere decir que esa funcino correra de forma concurrente

	wg.Wait() //Aqui se le dice que espere, hasta que las goroutine hayan sido disparadas

	go func(text string) {
		fmt.Println(text)
	}("Adios")

	time.Sleep(time.Second * 1)
}
