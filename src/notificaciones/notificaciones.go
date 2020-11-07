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

// Comprueba si el email tiene activadas las notificaciones
func EmailActivated(email string) error{
	value, _ := db.emails[email]
	if !value{
		return fmt.Errorf("El email '%s' no tiene activadas las notificaciones", email)
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
func ModifyEmail(email string, newEmail string) error{
	// Busca si el email existe
	found := EmailExists(email)
	if found != nil{
		return fmt.Errorf("El email '%s' no existe", email)
	}

	// Se añade el nuevo email con el valor que tenía antes
	db.emails[newEmail] = db.emails[email]
	// Se borra el antiguo email
	delete(db.emails, email)

	return nil
}

// Desactiva las notificaciones de un email
func DeactivateEmail(email string) error{
	// Busca si el email ya existe
	found := EmailExists(email)

	if found != nil{
		return fmt.Errorf("El email '%s' no está registrado", email)
	}

	activated := EmailActivated(email)

	if activated != nil{
		return fmt.Errorf("El email '%s' ya tiene las notificaciones desactivadas", email)
	}

	// Desactiva las notificaciones
	db.emails[email] = false

	return nil
}