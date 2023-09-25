package variables

import (
	"fmt"
	"time"
)

type actualizdoType any

var Nombre string
var Estado bool
var Sueldo float32
var Actualizado actualizdoType

func ShowRestVar() {
	/* Se pueden llamar las funciones dentro de los mismos paquetes
	ShowInteger() */
	fmt.Printf("")
	Nombre = "Pedro"

	Estado = true
	Sueldo = 125.99
	Actualizado = time.Now().Format(time.DateOnly)
	fmt.Printf("Nombre: %v\nSueldo: %v\nActualizado: %v\n, Estado: %v", Nombre, Sueldo, Actualizado, Estado)
}
