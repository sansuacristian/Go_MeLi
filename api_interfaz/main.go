package main

import (
	"Go_MeLi/api_interfaz/internal/controllers/figura"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default() // crear una instancia de gin-lo que sera el servidor (servidor en memoria)
	//r.GET("/rectangulo/area", rectangulo.Area)

	routerRect := r.Group("/rectangulo")
	{
		routerRect.GET("/area", figura.AreaRectangulo)
		routerRect.GET("/perim", figura.PerimRectangulo)

	}

	/*	routerCircle := r.Group("/circulo")
		{
			routerCircle.GET("/area", figura.Geometria.Area)
			routerCircle.GET("/perim", figura.Geometria.Perim)

		}
	*/
	r.Run()
}
