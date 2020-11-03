package main

import (
	"src/terremoto"
	"src/usuario"
)

type ControladorTerremotos struct {  
	terremotos map[string]terremoto.Terremoto
	usuarios map[string]usuario.Usuario
}

func NewControladorTerremotos() *ControladorTerremotos{
	// Inicializa terremotos vacíos
	controlador := ControladorTerremotos{terremotos: make(map[string]terremoto.Terremoto), usuarios: make(map[string]usuario.Usuario)}

	return &controlador
}

// Como desarrollador tengo que permitir que se añada un terremoto
func addTerremoto(idTerremoto string, fechahora time.Time, latitud float64, longitud float64, profundidad int, magnitud float64, tipoMagnitud tipomagnitud.TipoMagnitud, localizacion string){

}

// Un usuario consulta un terremoto con su identificador
func searchTerremoto(string id){

}

// Un usuario se registra en el sistema añadiendo su nombre, apellidos, email y teléfono
func addUsuario(nombre string, apellidos string, email string, telefono string){

}