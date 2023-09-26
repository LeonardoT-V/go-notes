package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number1 int
var number2 int
var description string
var err error

func InputKeyboardNumber() {
	fmt.Println("Ingrese numero 1: ")
	scan := bufio.NewScanner(os.Stdin)
	if scan.Scan() {
		number1, err = strconv.Atoi(scan.Text())
		if err != nil {
			panic("el dato ingresado es incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese numero 2: ")
	if scan.Scan() {
		number2, err = strconv.Atoi(scan.Text())
		if err != nil {
			panic("el dato ingresado es incorrecto" + err.Error())
		}
	}
	fmt.Println("Ingrese description: ")
	if scan.Scan() {
		description = scan.Text()
	}
	fmt.Println(description, number1*number2)
}
