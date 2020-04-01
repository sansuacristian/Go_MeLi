package figura

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

//Rect ...
type Rect struct {
	Ancho, Alto float64
}

//Area ...
func Area(c *gin.Context) { // utilizo c*gin.Context saber el contexto de la solicitud (respuesta-peticion)
	ancho, _ := strconv.ParseFloat(c.DefaultQuery("ancho", "0"), 64) //por query me llegara un valor que  se llama ancho y por defecto seria 0
	alto, _ := strconv.ParseFloat(c.DefaultQuery("alto", "0"), 64)   // llega una info por URL defalultQuery
	rectangulo := Rect{ancho, alto}                                  //  al objeto le damos los paramentros
	result := rectangulo.area(c)
	c.JSON(200, gin.H{"resultado area": result}) //devuelve el resultado para que lo vea desde laweb gracias al context

}

//Perim ...
func Perim(c *gin.Context) { // por qué dice que  variable replicada, acaso no toma la interface?
	rectangulo := Rect{5, 4}
	result := rectangulo.perim(c) // se le pasa al objeto rectangulp
	c.JSON(200, gin.H{"resultado perimetro": result})

}

//Area ...
func (g Geometria) area(c *gin.Context, r Rect) float64 { // con el context no podria retornar el valor al cliente
	return r.Ancho * r.Alto
}

//Perim ...
func (g Geometria) perim(c *gin.Context, r Rect) float64 {
	return 2*r.Ancho + 2*r.Alto
}
