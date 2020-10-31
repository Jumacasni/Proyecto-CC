package main

import (
	"src/terremoto"
)

type ControladorTerremotos struct {  
	terremotos []terremoto.Terremoto
}

func NewControladorTerremotos() *ControladorTerremotos{
	// Inicializa terremotos vacíos
	controlador := ControladorTerremotos{terremotos: []terremoto.Terremoto{}}

	return &controlador
}

func addTerremoto(t terremoto.Terremoto){

}

func getTerremoto() terremoto.Terremoto{
	
}