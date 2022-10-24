package models

import (
	"database/sql"

	"mod.go/src/entities"
)

type PlanetModel struct {
	Db *sql.DB
}

//Busqueda por Id
func (planetModel PlanetModel) FindAll() (Planet []entities.Planet, err error) {
	rows, err := planetModel.Db.Query("select * from planet")
	if err != nil {
		return nil, err
	} else {
		var planets []entities.Planet
		for rows.Next() {
			var id int64
			var clima string
			var fecha_creacion string
			var diametro string
			var fecha_edicion string
			var pelicula string

			var gravedad string
			var nombre string
			var periodo_orbital string
			var poblacion string
			var residentes string

			var periodo_de_rotacion string
			var agua_superficial string
			var terreno string
			var url string

			err2 := rows.Scan(&id, &clima, &fecha_creacion, &diametro, &fecha_edicion, &pelicula, &gravedad, &nombre,
				&periodo_orbital, &poblacion, &residentes, &periodo_de_rotacion, &agua_superficial, &terreno, &url)
			if err2 != nil {
				return nil, err2
			} else {
				planet := entities.Planet{
					Id:              id,
					Clima:           clima,
					Fecha_creacion:  fecha_creacion,
					Diametro:        diametro,
					Fecha_edicion:   fecha_edicion,
					Pelicula:        pelicula,
					Gravedad:        gravedad,
					Nombre:          nombre,
					Periodo_orbital: periodo_de_rotacion,
					Poblacion:       poblacion,
					Residentes:      residentes,

					Periodo_de_rotacion: periodo_de_rotacion,
					Agua_superficial:    agua_superficial,
					Terreno:             terreno,
					Url:                 url,
				}
				planets = append(planets, planet)
			}
		}
		return planets, nil
	}
}

//Busqueda por id
func (planetModel PlanetModel) Find(id int) (entities.Planet, error) {
	rows, err := planetModel.Db.Query("select * from planet where id = ?", id)
	if err != nil {
		return entities.Planet{}, err
	} else {
		var planet entities.Planet
		for rows.Next() {
			var id int64
			var clima string
			var fecha_creacion string
			var diametro string
			var fecha_edicion string
			var pelicula string

			var gravedad string
			var nombre string
			var periodo_orbital string
			var poblacion string
			var residentes string

			var periodo_de_rotacion string
			var agua_superficial string
			var terreno string
			var url string

			err2 := rows.Scan(&id, &clima, &fecha_creacion, &diametro, &fecha_edicion, &pelicula, &gravedad, &nombre,
				&periodo_orbital, &poblacion, &residentes, &periodo_de_rotacion, &agua_superficial, &terreno, &url)
			if err2 != nil {
				return entities.Planet{}, err2
			} else {
				planet = entities.Planet{
					Id:              id,
					Clima:           clima,
					Fecha_creacion:  fecha_creacion,
					Diametro:        diametro,
					Fecha_edicion:   fecha_edicion,
					Pelicula:        pelicula,
					Gravedad:        gravedad,
					Nombre:          nombre,
					Periodo_orbital: periodo_de_rotacion,
					Poblacion:       poblacion,
					Residentes:      residentes,

					Periodo_de_rotacion: periodo_de_rotacion,
					Agua_superficial:    agua_superficial,
					Terreno:             terreno,
					Url:                 url,
				}
			}
		}
		return planet, nil
	}
}

//Registro de datos
func (planetModel PlanetModel) Create(planet *entities.Planet) error {
	row, err := planetModel.Db.Query("insert into Planet values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) select convert(bigint, SCOPE_IDENTITY());",
		planet.Id, planet.Clima, planet.Fecha_creacion, planet.Diametro, planet.Fecha_edicion, planet.Pelicula, planet.Gravedad, planet.Nombre,
		planet.Periodo_orbital, planet.Poblacion, planet.Residentes, planet.Periodo_de_rotacion, planet.Agua_superficial, planet.Terreno, planet.Url)
	if err != nil {
		return err
	} else {
		var newId int64
		row.Next()
		row.Scan(&newId)
		planet.Id = newId
		return nil

	}
}

//Actualizar datos
func (planetModel PlanetModel) Update(planet *entities.Planet) (int64, error) {
	result, err := planetModel.Db.Exec("update planet set clima = ?, fecha_creacion = ?, diametro = ?, fecha_edicion = ?, pelicula = ?, gravedad = ?, nombre = ?, periodo_orbital = ?, poblacion = ?, residentes = ?, periodo_de_rotacion = ?, agua_superficial = ?, terreno = ?, url = ? where id = ?",
		planet.Clima, planet.Fecha_creacion, planet.Diametro, planet.Fecha_edicion, planet.Pelicula, planet.Gravedad, planet.Nombre,
		planet.Periodo_orbital, planet.Poblacion, planet.Residentes, planet.Periodo_de_rotacion, planet.Agua_superficial, planet.Terreno, planet.Url, planet.Id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

//Eliminar datos
func (planetModel PlanetModel) Delete(id int) (int64, error) {
	result, err := planetModel.Db.Exec("delete from planet where id = ?", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
