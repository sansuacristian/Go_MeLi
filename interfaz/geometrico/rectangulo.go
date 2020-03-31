package geometrico

//Rect ...
type Rect struct {
	Width, Height float64
}

//Area ...
func (r Rect) Area() float64 {
	return r.Width * r.Height
}

//Perim ...
func (r Rect) Perim() float64 {
	return 2*r.Width + 2*r.Height
}
