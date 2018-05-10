package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

		operando1, err := strconv.Atoi(os.Args[1])
		fmt.Println(err)
		operando2, err := strconv.Atoi(os.Args[2])
		fmt.Println(err)

		switch os.Args[3] {
		case "s":
			//Suma
			fmt.Print("La suma es: ")
			fmt.Println(operando1 + operando2)
		case "r":
			//Resta
			fmt.Print("La resta es: ")
			fmt.Println(operando1 - operando2)
		case "m":
			//Multiplicacion
			fmt.Print("La multiplicacion es: ")
			fmt.Println(operando1 * operando2)
		case "d":
			//Division
			fmt.Print("La division es: ")
			fmt.Println(float64(operando1) / float64(operando2))
		default:
			fmt.Println("Error")
		}

	}
// func almacen(int a, int b)
// 	historial, err := os.OpenFile("historial.txt", os.O_APPEND, 0777)
// 	showError(err)

// 	escribir, err := historial.WriteString(fmt.Println(operando1 + operando2))
// 	fmt.Println(escribir)
// 	showError(err)

// 	historial.Close()
	
// 	texto, errorDeFichero := ioutil.ReadFile("historial.txt")
// 	showError(errorDeFichero)

// 	fmt.Println(string(texto))
// }

// func showError(e error) {
// 	if e != nil {
// 		panic(e)
// 	}

