package middleware

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func OperatorMiddleware() {
	fmt.Println("inicio")
	result := operacionMiddleware(sumar)(2, 3)
	fmt.Println(result)
	resultRes := operacionMiddleware(restar)(2, 3)
	fmt.Println(resultRes)
	resultMu := operacionMiddleware(multiplicar)(2, 3)
	fmt.Println(resultMu)
}

func operacionMiddleware(f func(int, int) int) func(int, int) int {
	return func(i1, i2 int) int {
		fmt.Println("Inicio de operacion")
		return f(i1, i2)
	}
}
