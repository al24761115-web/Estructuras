//Jesus Alejandro Hernandez Hernandez 24761115

package main

import (
	"fmt"
	"sort"
)

// Producto representa un artículo del catálogo de la cafetería.
type Producto struct {
	Nombre      string
	Precio      float64
	Categoria   string
	Descripcion string
}

// catalogo contiene todos los productos disponibles.
// No modifiques esta variable.
var catalogo = []Producto{
	{Nombre: "Agua mineral", Precio: 20.00, Categoria: "Bebida", Descripcion: "Botella 600ml sin gas"},
	{Nombre: "Té verde", Precio: 30.00, Categoria: "Bebida", Descripcion: "Infusión caliente"},
	{Nombre: "Café Americano", Precio: 35.00, Categoria: "Bebida", Descripcion: "Café negro doble shot"},
	{Nombre: "Jugo de naranja", Precio: 45.00, Categoria: "Bebida", Descripcion: "Natural exprimido 355ml"},
	{Nombre: "Pastel de chocolate", Precio: 55.00, Categoria: "Postre", Descripcion: "Rebanada individual 120g"},
	{Nombre: "Burrito de res", Precio: 75.00, Categoria: "Comida", Descripcion: "Tortilla, carne, frijoles"},
	{Nombre: "Sandwich de pollo", Precio: 85.00, Categoria: "Comida", Descripcion: "Pan integral, pollo, verduras"},
	{Nombre: "Ensalada César", Precio: 95.00, Categoria: "Comida", Descripcion: "Lechuga romana, crutones"},
	{Nombre: "Pizza personal", Precio: 110.00, Categoria: "Comida", Descripcion: "4 rebanadas, queso y jitomate"},
	{Nombre: "Combo del día", Precio: 130.00, Categoria: "Combo", Descripcion: "Plato fuerte + bebida + postre"},
}

// encontrarCombinaciones recibe el catálogo y un presupuesto y
// devuelve todas las combinaciones de productos distintos que
// se pueden comprar sin exceder ese monto.
// Cada producto puede aparecer solo una vez por combinación.
func encontrarCombinaciones(productos []Producto, presupuesto float64) [][]Producto {
	// TODO: implementa el algoritmo aquí.
	// Sugerencia: usa backtracking o itera con dos ciclos anidados.
	// Recuerda que para evitar combinaciones duplicadas el índice
	// de inicio en cada nivel de recursión debe avanzar.
	var resultado [][]Producto

	var backtrack func(inicio int, actual []Producto, gastoActual float64)
	backtrack = func(inicio int, actual []Producto, gastoActual float64) {
		if len(actual) > 0 {
			copia := make([]Producto, len(actual))
			copy(copia, actual)
			resultado = append(resultado, copia)
		}

		for i := inicio; i < len(productos); i++ {
			p := productos[i]
			if gastoActual+p.Precio <= presupuesto {
				actual = append(actual, p)
				backtrack(i+1, actual, gastoActual+p.Precio)
				actual = actual[:len(actual)-1]
			}
		}
	}

	backtrack(0, []Producto{}, 0)
	return resultado

}

// imprimirResultados muestra en consola el resumen de combinaciones.
func imprimirResultados(combis [][]Producto, presupuesto float64) {
	// TODO: implementa la presentación de resultados.

	if len(combis) == 0 {
		fmt.Println("Con ese presupuesto no alcanza para ningun combo")
		return

	}

	fmt.Printf(" Tiendita Gourmet COmbos para $%.2f\n ", presupuesto)

	//   - Total de combinaciones encontradas

	fmt.Printf("Estan todos estos combos: %d\n\n", len(combis))

	//   - Agrupación por cantidad de productos
	grupos := make(map[int][][]Producto)
	for _, c := range combis {
		n := len(c)
		grupos[n] = append(grupos[n], c)

	}

	llaves := make([]int, 0, len(grupos))
	for k := range grupos {
		llaves = append(llaves, k)
	}
	sort.Ints(llaves)

	var mayorValor float64
	var mejorCombo []Producto
	for _, c := range combis {
		var total float64
		for _, p := range c {
			total += p.Precio
		}
		if total > mayorValor {
			mayorValor = total
			mejorCombo = c
		}
	}
	for _, n := range llaves {
		plural := "producto"
		if n > 1 {
			plural = "productos"
		}
		fmt.Printf("┌─ %d %s (%d combinaciones)\n", n, plural, len(grupos[n]))

		for idx, combo := range grupos[n] {
			var total float64
			for _, p := range combo {
				total += p.Precio
			}
			cambio := presupuesto - total

			fmt.Printf("│  [%d] ", idx+1)
			for i, p := range combo {
				if i > 0 {
					fmt.Print(" + ")
				}
				fmt.Printf("%s ($%.0f)", p.Nombre, p.Precio)
			}
			fmt.Printf("\n│      Por todo fue: $%.2f  |  Te regreso: $%.2f\n", total, cambio)
		}
		fmt.Println("│")
	}

	fmt.Printf(" La combinacion mas cara ($%.2f):\n", mayorValor)
	for _, p := range mejorCombo {
		fmt.Printf("• %-25s $%.2f   [%s]\n", p.Nombre, p.Precio, p.Categoria)
	}
	fmt.Print("Cambio: $%.2f\n", presupuesto-mayorValor)
}

func main() {

	var presupuesto float64
	fmt.Print("Dame tu presupuesto: $")
	// TODO: leer el valor ingresado
	for {
		_, err := fmt.Scan(&presupuesto)
		if err == nil && presupuesto > 0 {
			break
		}
		fmt.Print("Solo aceptamos dinero, porfavor dame tu presupuesto real: $")
		fmt.Scanln()
	}

	// TODO: llamar a encontrarCombinaciones
	combos := encontrarCombinaciones(catalogo, presupuesto)

	// TODO: llamar a imprimirResultados
	imprimirResultados(combos, presupuesto)

}
