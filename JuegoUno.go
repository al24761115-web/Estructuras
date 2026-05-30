package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type Carta struct {
	Color string
	Valor string
}

type Jugador struct {
	Nombre string
	Cartas []Carta
}

type Partida struct {
	Mazo      []Carta
	Jugadores []Jugador
	Descarte  []Carta
	Turno     int
}

func main() {
	juego := Partida{
		Jugadores: []Jugador{{Nombre: "Alex"}, {Nombre: "Taquito de Pastor"}}, //Es mi gamertag
	}

	colores := []string{"Rojo", "Verde", "Azul", "Amarillo"}
	nums := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, c := range colores {
		for _, v := range nums {
			juego.Mazo = append(juego.Mazo, Carta{c, v})
		}
	}
	rand.Shuffle(len(juego.Mazo), func(i, j int) {
		juego.Mazo[i], juego.Mazo[j] = juego.Mazo[j], juego.Mazo[i]
	})

	for i := 0; i < 5; i++ {
		for j := range juego.Jugadores {
			juego.Jugadores[j].Cartas = append(juego.Jugadores[j].Cartas, juego.Mazo[len(juego.Mazo)-1])
			juego.Mazo = juego.Mazo[:len(juego.Mazo)-1]
		}
	}
	juego.Descarte = append(juego.Descarte, juego.Mazo[len(juego.Mazo)-1])
	juego.Mazo = juego.Mazo[:len(juego.Mazo)-1]

	lector := bufio.NewReader(os.Stdin)

	for {
		p := &juego.Jugadores[juego.Turno]
		tope := juego.Descarte[len(juego.Descarte)-1]

		fmt.Println("Le toca a", p.Nombre)
		fmt.Println("La carta que esta:", tope.Color, tope.Valor)
		fmt.Println("Tu tienes las cartas: ")
		for i, c := range p.Cartas {
			fmt.Printf("%d: %s %s | ", i, c.Color, c.Valor)
		}

		var input string
		var op int
		fmt.Print("\nElegir carta (o R para robar): ")
		fmt.Scan(&input)

		if input == "R" || input == "r" {
			op = -1
		} else {
			fmt.Sscan(input, &op)
		}

		jugadaValida := false

		if op == -1 {
			if len(juego.Mazo) > 0 {
				carta := juego.Mazo[len(juego.Mazo)-1]
				juego.Mazo = juego.Mazo[:len(juego.Mazo)-1]
				p.Cartas = append(p.Cartas, carta)
				fmt.Println("Le robaste una carta.")
				jugadaValida = true
			} else {
				fmt.Println("Ya no hay cartas.")
			}
		} else if op >= 0 && op < len(p.Cartas) {
			c := p.Cartas[op]
			if c.Color == tope.Color || c.Valor == tope.Valor {
				juego.Descarte = append(juego.Descarte, c)
				p.Cartas = append(p.Cartas[:op], p.Cartas[op+1:]...)
				fmt.Println("¡Jugaste:", c.Color, c.Valor, "!")
				jugadaValida = true
			} else {
				fmt.Println("Esta carta no sirve, no tiene ni el color ni el valor que deberia.")
			}
		} else {
			fmt.Println("Esta opcion no es valida.")
		}

		if len(p.Cartas) == 0 {
			perdedor := juego.Jugadores[(juego.Turno+1)%len(juego.Jugadores)]
			fmt.Println(p.Nombre, "ha humillado a", perdedor.Nombre)
			break
		}

		if jugadaValida {
			juego.Turno = (juego.Turno + 1) % len(juego.Jugadores)
		}

		fmt.Println("\nPresiona Enter para el siguiente turno...")
		lector.ReadString('\n')
	}
}
