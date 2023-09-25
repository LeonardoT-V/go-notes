package variables

import "fmt"

/* Las funciones publicas deben empezar con mayusculas */
func ShowInteger() {
	/* al iniciar una variable con su tipo de datos esta se le asignara con el valor mas bajo respectivo a su tipo asignado */
	var intCommon int

	/* := declara el tipo de dato */
	intOf16 := int16(10)
	intOf32 := int32(10)
	intOf64 := int64(10)

	fmt.Printf("common: %v, 16: %v, 32: %v, 64: %v", intCommon, intOf16, intOf32, intOf64)
}

/* esto no se encuentra publicamente */
func hola() {
	fmt.Println("hola")
}
