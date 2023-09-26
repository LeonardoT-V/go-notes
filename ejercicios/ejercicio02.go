package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var selectTable bool = true

func ExcecuteEje2() string {

	var formatTable string
	scan := bufio.NewScanner(os.Stdin)

	for selectTable {
		fmt.Print("Tabla de multiplicacion a mostrar: ")
		if scan.Scan() {
			table, err := strconv.Atoi(scan.Text())
			if err != nil {
				fmt.Println("valor incorrecto" + err.Error())
				continue
			}
			for i := 0; i < 12; i++ {
				//fmt.Printf("%d x %d = %d\n", table, i, table*i)
				formatTable += fmt.Sprintf("%d x %d = %d\n", table, i, table*i)
			}
			selectTable = false
		}
	}
	return formatTable
}
