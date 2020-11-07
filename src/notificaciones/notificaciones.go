package notificaciones

import (
	"fmt"
)

// Emails registrados en el sistema
// test@test.com = true -> el email tiene activadas las notificaciones
// test@test.com = false -> el email no tiene activadas las notificaciones
type Notificaciones struct {
	emails map[string]bool
}

var (
	db Notificaciones
)

// Inicializa la db con un email
func init() {
	InitEmail("prueba@prueba.com")
}

func InitEmail(email string){
	db.emails = make(map[string]bool)
	db.emails[email] = true
}

// Comprueba si el email existe en la db
func EmailExists(email string) error{
	_, exists := db.emails[email]
	if !exists{
		return fmt.Errorf("El email '%s' no existe", email)
	}
	return nil
}

// Añade un email
func AddEmail(email string) error{
	// Busca si el email ya existe
	found := EmailExists(email)
	if found == nil{
		return fmt.Errorf("El email '%s' ya está registrado", email)
	}

	// Inserta el email en la db
	db.emails[email] = true
	return nil
}

// Modifica un email
func modifyEmail(email string){

}

// Elimina un email
func deleteEmail(email string){

}