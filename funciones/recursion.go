package funciones

import "fmt"

func ExponentNumber(n int) {
	if n > 100 {
		return
	}
	fmt.Println(n)
	ExponentNumber(n * 2)
}
