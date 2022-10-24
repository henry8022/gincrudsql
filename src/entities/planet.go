package entities

import (
	"fmt"
)

type Planet struct {
	Id             int64
	Clima          string
	Fecha_creacion string
	Diametro       string
	Fecha_edicion  string
	Pelicula       string

	Gravedad        string
	Nombre          string
	Periodo_orbital string
	Poblacion       string
	Residentes      string

	Periodo_de_rotacion string
	Agua_superficial    string
	Terreno             string
	Url                 string
}

func (planet Planet) ToString() string {
	return fmt.Sprintf("id: %d\nclima: %s\nfecha_creacion: %s\ndiametro: %s\nfecha_edicion: %s\npelicula: %s\ngravedad: %s\nnombre: %s\nperiodo_orbital: %s\npoblacion: %s\nresidente: %s\nperiodo_de_rotacion: %s\nagua_superficial: %s\nterreno: %s\nurl: %s", planet.Id, planet.Clima, planet.Fecha_creacion, planet.Diametro, planet.Fecha_edicion, planet.Pelicula, planet.Gravedad, planet.Nombre, planet.Periodo_orbital, planet.Poblacion, planet.Residentes, planet.Periodo_de_rotacion, planet.Agua_superficial, planet.Terreno, planet.Url)

}
