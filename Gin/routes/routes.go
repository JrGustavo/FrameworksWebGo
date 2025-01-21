package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Usuario struct {
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

var usuarios []Usuario

func SetupRoutes(r *gin.Engine) {

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title":   "Hola universo",
			"Heading": "Hola mundo",
			"Message": "Hola a todos desde la aplicación de Gin & plantillas HTML",
		})
	})

	r.Static("/static", "./static")
}
