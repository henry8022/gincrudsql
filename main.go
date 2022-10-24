package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Planet struct {
	ID                  string `json:"id"`
	Clima               string `json:"clima"`
	Fecha_creacion      string `json:"fecha_creacion"`
	Diametro            string `json:"diametro"`
	Fecha_edicion       string `json:"fecha_edicion"`
	Pelicula            string `json:"pelicula"`
	Gravedad            string `json:"gravedad"`
	Nombre              string `json:"nombre"`
	Periodo_orbital     string `json:"periodo_orbital"`
	Poblacion           string `json:"poblacion"`
	Residentes          string `json:"residentes"`
	Periodo_de_rotacion string `json:"period_de_rotacion"`
	Agua_superficial    string `json:"agua_superficial"`
	Terreno             string `json:"terreno"`
	Url                 string `json:"url"`
}

var Planets []Planet //nil

func main() {
	r := gin.Default()

	planetRoutes := r.Group("/planets")
	{
		planetRoutes.GET("/mostrar", GetPlanet)
		planetRoutes.POST("/guardar", CreatePlanet)
		planetRoutes.PUT("/actualizar/:id", EditPlanet)
		planetRoutes.DELETE("/eliminar/:id", DeletePlanet)

	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err.Error())
	}
}
func GetPlanet(c *gin.Context) {
	c.JSON(200, Planets)

}

func CreatePlanet(c *gin.Context) {
	var reqBody Planet
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return
	}
	reqBody.ID = uuid.New().String()

	Planets = append(Planets, reqBody)

	c.JSON(200, gin.H{
		"error": false,
	})
}

func EditPlanet(c *gin.Context) {
	id := c.Param("id")
	var reqBody Planet
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return
	}

	for i, u := range Planets {
		if u.ID == id {
			Planets[i].Nombre = reqBody.Nombre
			Planets[i].Clima = reqBody.Clima
			Planets[i].Fecha_creacion = reqBody.Fecha_creacion
			Planets[i].Diametro = reqBody.Diametro
			Planets[i].Fecha_edicion = reqBody.Fecha_edicion
			Planets[i].Pelicula = reqBody.Pelicula
			Planets[i].Gravedad = reqBody.Gravedad
			Planets[i].Nombre = reqBody.Nombre
			Planets[i].Periodo_orbital = reqBody.Periodo_orbital
			Planets[i].Poblacion = reqBody.Poblacion
			Planets[i].Residentes = reqBody.Residentes
			Planets[i].Periodo_de_rotacion = reqBody.Periodo_de_rotacion
			Planets[i].Agua_superficial = reqBody.Agua_superficial
			Planets[i].Terreno = reqBody.Terreno
			Planets[i].Url = reqBody.Url

			c.JSON(200, gin.H{
				"error": false,
			})
			return
		}
	}

	c.JSON(404, gin.H{
		"error":   false,
		"message": "id incorrecto de planeta",
	})
}

func DeletePlanet(c *gin.Context) {
	id := c.Param("id")

	for i, u := range Planets {
		if u.ID == id {
			Planets = append(Planets[:i], Planets[i+1:]...)

			c.JSON(200, gin.H{
				"error": false,
			})
			return
		}
	}

	c.JSON(404, gin.H{
		"error":   false,
		"message": "Invalid user id",
	})
}
