package main

import (
	"fmt"
)

func main() {

	var a [4][4]int32 = llenadoVector()
	fmt.Println(a)
	b := diagonalDifference(a)
	fmt.Println(b)

}

func diagonalDifference(arr [4][4]int32) int32 {
	var sum, mus, r int32

	for i := 0; i < len(arr); i++ {
		sum += arr[i][i]

	}
	f := len(arr)
	for j := 0; j < len(arr); j++ {
		f--
		mus += arr[j][f]

	}

	if (sum - mus) > 0 {
		r = (sum - mus)

	} else {
		r = (sum - mus) * -1
	}

	return r

}

func llenadoVector() [4][4]int32 {
	var v [4][4]int32
	for f := 0; f < 4; f++ {
		for c := 0; c < 4; c++ {
			fmt.Printf("Ingrese componente %d , %d \n", f, c)
			fmt.Scan(&v[f][c])
		}
	}
	return v

}
