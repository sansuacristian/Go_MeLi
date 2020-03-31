package main

import (
	"Go_MeLi/interfaz/geometrico"
	"fmt"
)

func main() {

	r := geometrico.Rect{Width: 3, Height: 4}
	c := geometrico.Circle{Radius: 5}
	measure(r)
	measure(c)
}

func measure(g geometrico.Geometry) {
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Perim())
}
