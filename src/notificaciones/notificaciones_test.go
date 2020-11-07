package notificaciones

import (
	"testing"
	"os"
)

func TestMain(m *testing.M){
	// Setup: se añade test@test.com a los emails registrados
	InitEmail("test@test.com")

	exitVal := m.Run()
	os.Exit(exitVal)
}

// Test para comprobar que existe un email. Pasa el test.
func TestEmailExistsOk(t *testing.T){
	err := EmailExists("test@test.com")
	if err != nil{
		t.Error(err)
	}
}

// Test para comprobar que existe un email. Falla el test.
func TestEmailExistsFail(t *testing.T){
	err := EmailExists("prueba@prueba.com")
	if err != nil{
		t.Error(err)
	}
}

// Test que añade un email. Pasa el test.
func TestAddEmailOk(t *testing.T){
	// Añade el email
	err := AddEmail("prueba@prueba.com")
	if err != nil{
		t.Error(err)
	}

	// Comprueba que el email se ha añadido
	err = EmailExists("prueba@prueba.com")
	if err != nil{
		t.Error(err)
	}
}

// Test que añade un email que ya está registrado. Falla el test.
func TestAddEmailFail(t *testing.T){
	err := AddEmail("prueba@prueba.com")
	if err != nil{
		t.Error(err)
	}
}

// Test que modifica un email. Pasa el test.
func TestModifyEmail(t *testing.T){
	err := ModifyEmail("test@test.com", "nuevotest@nuevotest.com")

	if err != nil{
		t.Error(err)
	}

	// Comprueba que el antiguo email ya no existe y le nuevo sí
	_, oldExists := db.emails["test@test.com"]
	_, newExists := db.emails["nuevotest@nuevotest.com"]

	if (oldExists || !newExists){
		t.Error("El email no se ha modificado correctamente")
	}
}

// Test para comprobar si las notificaciones están activadas. Pasa el test
func TestEmailActivatedOk(t *testing.T){
	err := EmailActivated("test@test.com")
	if err != nil{
		t.Error(err)
	}
}

// Test para desactivar las notificaciones de "prueba@prueba.com". Pasa el test
func TestDeactivateEmailOk(t *testing.T){
	err := DeactivateEmail("prueba@prueba.com")
	if err != nil{
		t.Error(err)
	}
}

// Test que comprueba si las notificaciones de "prueba@prueba.com" están activadas. Falla el test
func TestEmailActivatedFail(t *testing.T){
	err := EmailActivated("prueba@prueba.com")
	if err != nil{
		t.Error(err)
	}
}

// Test para desactivar las notificaciones. Falla el test porque ya están desactivadas.
func TestDeactivateEmailFail(t *testing.T){
	err := DeactivateEmail("prueba@prueba.com")
	if err != nil{
		t.Error(err)
	}
}