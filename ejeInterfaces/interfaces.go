package ejeinterfaces

import (
	"fmt"

	"github.com/leonardoT-V/go-notes/interfaces"
)

func HumWalking(hu interfaces.Humano) {
	hu.Thinking()
	fmt.Printf("Soy un %v y estoy respirando", hu.Genrer())
}
