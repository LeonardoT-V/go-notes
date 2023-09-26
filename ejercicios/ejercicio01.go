package ejercicios

import (
	"fmt"
	"strconv"
)

func Ejercicio1(s string) (int, string) {
	num, _ := strconv.Atoi(s)
	var textToSend string
	if num >= 100 {
		textToSend = fmt.Sprintf("El numero %v es mayor a 100", num)
	} else {
		textToSend = fmt.Sprintf("El numero %v es menor a 100", num)
	}

	return num, textToSend
}
