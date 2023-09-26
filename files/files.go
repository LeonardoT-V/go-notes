package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/leonardoT-V/go-notes/ejercicios"
)

const filename string = "./files/data/table.txt"

func SaveTableInText() {
	table := ejercicios.ExcecuteEje2()
	data, err := os.Create(filename)
	if err != nil {
		fmt.Println("error al create File" + err.Error())
		return
	}
	fmt.Fprintln(data, table)
	data.Close()
}

func AddMoreTablesInText() {
	table := ejercicios.ExcecuteEje2()
	if !appendTable(filename, table) {
		fmt.Println("Error en contenido")
	}
}

func appendTable(filename string, table string) bool {
	arch, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("error in append")
		return false
	}
	_, err = arch.WriteString(table)
	if err != nil {
		fmt.Println("error in write")
		return false
	}
	arch.Close()
	return true
}

func ReadData() {
	// este lee el archivo completo
	//archivo,_:= ioutil.ReadFile(filename)
	/* archivo, _ := os.ReadFile(filename)
	fmt.Println(string(archivo)) */

	// esto lee por linea
	archivo, _ := os.Open(filename)
	scan := bufio.NewScanner(archivo)
	fmt.Println("asd" + scan.Text())
	for scan.Scan() {
		reg := scan.Text()
		fmt.Println("> " + reg)
	}
	archivo.Close()
}
