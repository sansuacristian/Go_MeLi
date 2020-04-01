package main

import (
	"Go_MeLi/api_interfaz/internal/controllers/figura"
	"strconv"

	"github.com/gin-gonic/gin"
)

//var cir = figura.Circ{5}

//RectA creo la función que recibe el contexto de gin, en esta función utilizo query para pasar los datos
func RectA(c *gin.Context) {
	ancho, _ := strconv.ParseFloat(c.DefaultQuery("ancho", "0"), 64)
	alto, _ := strconv.ParseFloat(c.DefaultQuery("alto", "0"), 64)
	rectangulo := figura.Rect{Ancho: ancho, Alto: alto}
	c.JSON(200, gin.H{"el area del rectangulo": figura.Geometria.Area(rectangulo)})

}

//CircA en esta función utilizo c.param para  pasar los datos desde la URL
func CircA(c *gin.Context) {
	Radio, _ := strconv.ParseFloat(c.Param("Radio"), 64)
	circulo := figura.Circ{Radio: Radio}
	c.JSON(200, gin.H{"el area del circulo": figura.Geometria.Area(circulo)})
}

//CircP ...
func CircP(c *gin.Context) {
	Radio, _ := strconv.ParseFloat(c.Param("Radio"), 64)
	circulo := figura.Circ{Radio: Radio}
	c.JSON(200, gin.H{"el perimetro del circulo": figura.Geometria.Perim(circulo)})
}

//RectP ...
func RectP(c *gin.Context) {
	ancho, _ := strconv.ParseFloat(c.DefaultQuery("ancho", "0"), 64)
	alto, _ := strconv.ParseFloat(c.DefaultQuery("alto", "0"), 64)
	rectangulo := figura.Rect{Ancho: ancho, Alto: alto}
	c.JSON(200, gin.H{"el perimetro del rectangulo": figura.Geometria.Perim(rectangulo)})

}

func main() {

	r := gin.Default() // crear una instancia de gin-lo que sera el servidor (servidor en memoria)
	//r.GET("/rectangulo/area", rectangulo.Area)

	// Grupo de Rutas para rectangulo usando Query
	routerRect := r.Group("/rectangulo")
	{

		routerRect.GET("/perim", RectA)
		routerRect.GET("/area", RectP)

	}
	// Grupo de rutas usando c.Param como parte de la información de la URL
	routerCircle := r.Group("/circulo")
	{
		routerCircle.GET("/area/:Radio", CircA)
		routerCircle.GET("/perim/:Radio", CircP)

	}

	r.Run()
}
