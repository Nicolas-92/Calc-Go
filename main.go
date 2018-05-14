package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	operando1, err := strconv.Atoi(os.Args[1])
	mostrarError(err)
	operando2, err := strconv.Atoi(os.Args[2])
	mostrarError(err)

	if operando1 < 0 || operando2 < 0 {
		fmt.Println("Debe ingresar un numero diferente")
	} else {
		switch os.Args[3] {
		case "suma":
			fmt.Printf("La suma es %d", operando1 + operando2)

		case "resta":
			fmt.Printf("La resta es %d", operando1 - operando2)

		case "multiplicacion":
			fmt.Printf("La multiplicacion es %d", operando1 * operando2)

		case "division":
			if operando2 == 0 {
				fmt.Printf("El denominador no puede ser cero")
				break
			} else {
				fmt.Printf("La division es %e", (float32 (operando1) / float32 (operando2)))
			}
		default:
			fmt.Printf("Error")
		}
	}
}

func mostrarError(e error) {
	if e != nil {
		panic(e)
	}
}


/*
	Autores:
		Dri Nicolas
		Pinto Rodrigo


*/