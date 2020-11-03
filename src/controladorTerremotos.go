package main

import (
	"src/terremoto"
)

type ControladorTerremotos struct {  
	terremotos []terremoto.Terremoto
	terremotos map[string]terremoto.Terremoto
}

func NewControladorTerremotos() *ControladorTerremotos{
	// Inicializa terremotos vacíos
	controlador := ControladorTerremotos{terremotos: make(map[string]terremoto.Terremoto)}

	return &controlador
}

// Como desarrollador tengo que permitir que se añada un terremoto
func addTerremoto(idTerremoto string, fechahora time.Time, latitud float64, longitud float64, profundidad int, magnitud float64, tipoMagnitud tipomagnitud.TipoMagnitud, localizacion string){

}