package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()                    // crear una instancia de gin-lo que sera el servidor (servidor en memoria)
	r.GET("/ping", func(c *gin.Context) { // metodo http llamada Get (traer) recibe dos parametros un string y una funcion
		c.JSON(200, gin.H{ //funcion JSON // retorna status code
			"message": "pong", // esto es lo que debe responder el api
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080") levanta el servidor
}
