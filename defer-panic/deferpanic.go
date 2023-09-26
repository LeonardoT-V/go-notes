package deferpanic

import "fmt"

func DeferExample() {
	defer fmt.Println("Mensaje final 1")
	fmt.Println("Primer mensage")
	defer fmt.Println("Mensaje final")
	fmt.Println("Segundo mensage")
}
