package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	operando1, err := strconv.Atoi(os.Args[1])

	operando2, er := strconv.Atoi(os.Args[2])

	if err != nil || er != nil {
		fmt.Println("entrada invalida")
	} else {
		switch os.Args[3] {
		case "suma":
			suma := operando1 + operando2
			fmt.Printf("La suma es %d", suma)

		case "resta":
			resta := operando1 - operando2
			fmt.Printf("La resta es %d", resta)

		case "multiplicacion":
			multi := operando1 * operando2
			fmt.Printf("La multiplicacion es %d", multi)

		case "division":
			if operando2 == 0 {
				fmt.Printf("El denominador no puede ser cero")
				break
			} else {
				div := (float32(operando1) / float32(operando2))
				fmt.Printf("La division es %f", div)
			}
		default:
			fmt.Printf("El operador no es reconocido")
		}
	}
}

/*
	Autores:
		Dri Nicolas
		Pinto Rodrigo


*/
