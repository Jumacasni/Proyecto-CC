package main

import (
	"src/usuario"
)

type ControladorUsuarios struct {  
	usuarios []usuario.Usuario
}

func NewControladorUsuarios() *ControladorUsuarios{
	// Inicializa terremotos vacíos
	usuarios := ControladorUsuarios{usuarios: []usuario.Usuarios{}}

	return &usuarios
}

func addUsuario(u usuario.Usuario){

}