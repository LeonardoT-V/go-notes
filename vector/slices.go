package vector

import "fmt"

var tablaSlice []int = []int{2, 4, 5}
var arreglo [5]int = [5]int{1, 2, 3, 4, 5}

func ShowSlices() {
	fmt.Println(tablaSlice)

	/* slices
	n: -> desde el numero al final
	:n -> desde el inicio hasta el numero
	n:n -> desde el numero definido
	*/
	porcion := arreglo[3:]
	fmt.Print(porcion)
}

func Capacidad() {
	el := make([]int, 5, 20)

	fmt.Println(el)
	fmt.Printf("Tamaño: %v\nCapacidad: %v\nValores: %v\n", len(el), cap(el), el)

	elAppend := make([]int, 0)

	for i := 0; i < 100; i++ {
		elAppend = append(elAppend, i)
	}
	fmt.Printf("Tamaño: %v\nCapacidad: %v\nValores: %v", len(elAppend), cap(elAppend), elAppend)
}
