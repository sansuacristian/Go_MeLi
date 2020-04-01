package figura

//Rect ...
type Rect struct {
	Ancho float64 `json:"Ancho"`
	Alto  float64 `json:"Alto"`
}

//Area ...
func (r Rect) Area() float64 { // con el context no podria retornar el valor al cliente
	return r.Ancho * r.Alto
}

//Perim ...
func (r Rect) Perim() float64 {
	return 2*r.Ancho + 2*r.Alto
}
