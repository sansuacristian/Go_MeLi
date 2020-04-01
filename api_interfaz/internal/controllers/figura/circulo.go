package figura

import (
	"math"
)

//Circ ...
type Circ struct {
	Radio float64 `json:"Radio"`
}

//Area ...
func (c Circ) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

// Perim ...
func (c Circ) Perim() float64 {
	return 2 * math.Pi * c.Radio
}
