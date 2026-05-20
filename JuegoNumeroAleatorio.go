package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var numevaluado int
	var contadorintentos int
	var rango int
	contadorintentos = 0
	fmt.Println("El juego trata de adivinar el numero aleatorio")
	fmt.Println("El rango empieza con 1, ingrese un numero que sea el otro extremo ")
	fmt.Scanln(&rango)
	numerobuscado := rand.Intn(rango) + 1
	for {
		fmt.Println("Ingresa el numero ")
		fmt.Scanln(&numevaluado)
		if numevaluado > numerobuscado {
			fmt.Println("El número a descubrir es menor")
			contadorintentos++
		} else if numevaluado < numerobuscado {

			fmt.Println("El número a descubrir es mayor")
			contadorintentos++
		} else {
			contadorintentos++
			fmt.Println("Haz encontrado el numero buscado ", numerobuscado)
			fmt.Println("En el intento ", contadorintentos)
			break
		}

	}

}
