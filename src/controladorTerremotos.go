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