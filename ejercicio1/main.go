package main

import (
	"fmt"
)

func main() {

	var num1, num2, num3 int

	fmt.Println("digite el primer numero: ")
	fmt.Scanln(&num1)
	fmt.Println("digite el segundo numero: ")
	fmt.Scanln(&num2)

	if num2 < num1 || num1 == num2 {

		fmt.Println(" No es posible ejecutar el programa")

	} else {

		for i := num1; i < num2; i++ {

			if i%2 != 0 {

				num3++

			}

		}

		fmt.Printf("El numero de numeros impares que hay entre %d y %d es igual a %d ", num1, num2, num3)
	}
}
