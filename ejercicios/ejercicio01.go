package ejercicios

import (
	"fmt"
	"strconv"
)

func Ejercicio1(s string) (int, string) {
	if num, _ := strconv.Atoi(s); num >= 100 {
		return num, fmt.Sprintf("El numero %v es mayor a 100", num)
	} else {
		return num, fmt.Sprintf("El numero %v es menor a 100", num)
	}
}
