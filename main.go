package main

import (
	"fmt"

	"github.com/leonardoT-V/go-notes/control"
	"github.com/leonardoT-V/go-notes/variables"
)

func main() {
	fmt.Println("hola")
	//variables.ShowInteger()
	//variables.ShowRestVar()
	est, text := variables.ConvertToText(10)
	fmt.Println(est, text)
	control.ShowDeviceRunning()
}
