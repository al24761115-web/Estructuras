package main

import "fmt"

// potencia calcula base^exp de forma recursiva.
// Puedes asumir que exp >= 0.
func potencia(base, exp int) int {
	// TODO: caso base
	// Cualquier número elevado a 0 es 1
	if exp == 0 {
		return 1
		// TODO: caso recursivo
		// base^exp = base * potencia(base, exp-1)

	}
	return base * potencia(base, exp-1)

}

func main() {
	fmt.Println(potencia(2, 0))  // 1
	fmt.Println(potencia(2, 1))  // 2
	fmt.Println(potencia(2, 8))  // 256
	fmt.Println(potencia(3, 4))  // 81
	fmt.Println(potencia(5, 3))  // 125
	fmt.Println(potencia(10, 0)) // 1
}
