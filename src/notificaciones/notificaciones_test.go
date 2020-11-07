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