package figura

import (
	"math"

	"github.com/gin-gonic/gin"
)

//Circ ...
type Circ struct {
	Radio float64
}

// Area ...
func Area(c *gin.Context) {
	circulo := Circ{5}
	resultado := circulo.area(c)
	c.JSON(200, gin.H{"resultado area": resultado})

}

// Perim ...
func Perim(c *gin.Context) {
	circulo := Circ{20}
	resultado := circulo.perim(c)
	c.JSON(200, gin.H{"resultado Perimetro": resultado})

}

func (s Circ) area(c *gin.Context) float64 {
	return math.Pi * s.Radio * s.Radio
}

func (s Circ) perim(c *gin.Context) float64 {
	return 2 * math.Pi * s.Radio
}
