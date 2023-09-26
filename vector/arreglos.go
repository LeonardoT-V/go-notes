package vector

import "fmt"

var tabla [10]int = [10]int{1, 2, 3, 0, 5}
var matrix [10][5]int

func ShowArrays() {
	tabla[1] = 22
	tabla[9] = 10

	fmt.Println(tabla)
	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matrix[7][2] = 9
	fmt.Println(matrix)
}
