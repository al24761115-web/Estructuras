// Jesus Alejandro Hernandez HErnandez 24761115
package main

import "fmt"

const APROBADO = 6.0
const EXCELENTE = 9.0
const ASISTENCIA_MINIMA = 80

func main() {
	var n int
	fmt.Print("¿Cuántos estudiantes tienes? ")
	fmt.Scan(&n)

	nombres := make([]string, n)
	calificaciones := make([]float64, n)
	asistencias := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Nombre del estudiante %d: ", i+1)
		fmt.Scan(&nombres[i])
		fmt.Printf("Calificación de %s (0-10): ", nombres[i])
		fmt.Scan(&calificaciones[i])
		fmt.Printf("Asistencia de %s (0-100): ", nombres[i])
		fmt.Scan(&asistencias[i])
	}

	fmt.Println("\n Estadisticas ")
	var suma float64
	for i, nombre := range nombres {
		cal := calificaciones[i]
		asis := asistencias[i]
		suma += cal

		fmt.Printf("\n%s: %.1f → ", nombre, cal)

		if cal >= EXCELENTE {
			if asis >= ASISTENCIA_MINIMA {
				fmt.Println("Excelente")
			} else {
				fmt.Println("Excelente con baja asistencia")
			}
		} else if cal >= APROBADO {
			if asis >= ASISTENCIA_MINIMA {
				fmt.Println("Aprobado")
			} else {
				fmt.Println("Aprobado con baja asistencia")
			}
		} else {
			fmt.Println("Reprobado")
		}

		switch {
		case asis >= 90:
			fmt.Printf("  Asistencia: %d%% (Excelente)\n", asis)
		case asis >= ASISTENCIA_MINIMA:
			fmt.Printf("  Asistencia: %d%% (Buena)\n", asis)
		default:
			fmt.Printf("  Asistencia: %d%% (Baja)\n", asis)
		}
	}

	fmt.Println("\n ESTADÍSTICAS FINALES ")
	mejor := calificaciones[0]
	i := 1
	for i < n {
		if calificaciones[i] > mejor {
			mejor = calificaciones[i]
		}
		i++
	}

	promedio := suma / float64(n)
	fmt.Printf("Promedio general: %.1f\n", promedio)
	fmt.Printf("Mejor calificación: %.1f\n", mejor)
}
