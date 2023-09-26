package main

import (
	"fmt"

	"github.com/leonardoT-V/go-notes/promesas"
)

func main() {
	// go no espera a que termine la ejecucion de una promesa
	go promesas.DelayShowName("Leonardo")

	fmt.Println("jiji")
	var x string
	fmt.Scanln(&x)
}
