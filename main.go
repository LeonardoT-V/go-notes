package main

import (
	inter "github.com/leonardoT-V/go-notes/ejeInterfaces"
	"github.com/leonardoT-V/go-notes/modelos"
)

func main() {
	Pedro := new(modelos.Hombre)

	inter.HumWalking(Pedro)
}
