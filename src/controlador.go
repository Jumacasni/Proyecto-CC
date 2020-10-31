package main

import (
	"src/terremoto"
)

type Controlador struct {  
	terremotos []terremoto.Terremoto
}

func NewControlador() *Controlador{
	// Inicializa terremotos vacíos
	controlador := Controlador{terremotos: []terremoto.Terremoto{}}

	return &controlador
}

func addTerremoto(t terremoto.Terremoto){

}

func getTerremoto() terremoto.Terremoto{
	
}