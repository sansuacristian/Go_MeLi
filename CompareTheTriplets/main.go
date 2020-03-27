package main

import (
	"fmt"
)

// Complete the compareTriplets function below.
func compareTriplets(a [3]int32, b [3]int32) [2]int32 {
	var alice, bot int32
	for i := 0; i < len(a); i++ {
		switch {
		case a[i] > b[i]:
			alice++
		case a[i] < b[i]:
			bot++
		default:

		}

	}
	return [2]int32{alice, bot}
}

func arrayllenado() [3]int32 {
	var a [3]int32
	for i := 0; i < 3; i++ {
		fmt.Scanln(&a[i])

	}
	return a
}

func main() {

	//al := [3]int32{3, 3, 3}
	//bo := [3]int32{1, 1, 1}
	fmt.Println("digite las calificaciones de Alice: ")
	var x [3]int32 = arrayllenado()
	fmt.Println("digite las calificaciones de Bob: ")
	y := arrayllenado()
	fmt.Println(x)
	fmt.Println(y)
	f := compareTriplets(x, y)
	fmt.Println("el resultado es: ", f)
}
