package main

import (
	"src/terremoto"
)

type Notificaciones struct {
	// Emails registrados en el sistema
	emails map[string]string
}

func NewNotificaciones() *Notificaciones{
	// Inicializa los emails
	notificaciones := Notificaciones{emails: make(map[string]string)
	return &notificaciones
}

// Añade un email
func addEmail(email string){

}

// Modifica un email
func modifyEmail(email string){

}

// Elimina un email
func deleteEmail(email string){

}