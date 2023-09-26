package funciones

import "fmt"

func tablaClousure(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}

func CallClousureTabla() {
	tablaOf := 2
	table := tablaClousure(tablaOf)
	fmt.Printf("table: %v\n", table())
	fmt.Printf("table: %v\n", table())
	fmt.Printf("table: %v\n", table())
	fmt.Printf("table: %v\n", table())
	for i := 0; i < 10; i++ {
		fmt.Println(table())
	}
}
