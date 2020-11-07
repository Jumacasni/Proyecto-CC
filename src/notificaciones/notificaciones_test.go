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

/***** MOCK PARA COMPROBAR QUE EXISTE EMAIL *****/
var emailExistsMock func(email string) bool

type checkEmailMock struct{}

func (u checkEmailMock) emailExists(email string) bool{
	return emailExistsMock(email)
}

/***** TEST ADD EMAIL *****/
// Test que añade un email que no está registrado
func TestAddEmailOk(t *testing.T){
	// Función mock, el email no existe
	checkEmail := checkEmailMock{}
	emailExistsMock = func(email string) bool{
		return false
	}

	// El length del map antes de meter el nuevo email
	prevLen := len(db.emails)
	// Añade el email
	err := AddEmail("prueba@prueba.com", checkEmail)
	if err != nil{
		t.Error(err)
	}
	// El length del map después de meter el nuevo email
	afterLen := len(db.emails)
	
	// El length del map debe ser mayor que cero para pasar el test
	if (afterLen == 0){
		t.Error("El length del map es 0, no se ha metido el email correctamente")
	}
	// El length del map debe incrementar en uno
	if (afterLen != (prevLen+1)){
		t.Error("El email no se ha metido correctamente")
	}
}

// Test que añade un email que ya está registrado
func TestAddEmailFail(t *testing.T){
	// Función mock, el email ya existe
	checkEmail := checkEmailMock{}
	emailExistsMock = func(email string) bool{
		return true
	}

	// El length del map antes AddEmail
	prevLen := len(db.emails)
	// Añade el email
	AddEmail("prueba@prueba.com", checkEmail)
	// El length del map después de AddEmail
	afterLen := len(db.emails)
	
	// El length del map no debe cambiar
	if (afterLen != prevLen){
		t.Error("El length del map ha cambiado")
	}
}

/***** TEST MODIFY EMAIL *****/
// Test que modifica un email. Pasa el test.
func TestModifyEmail(t *testing.T){
	// Función mock, el email existe
	checkEmail := checkEmailMock{}
	emailExistsMock = func(email string) bool{
		return true
	}

	// El length del map antes de modificar
	prevLen := len(db.emails)
	// Modifica el email
	err := ModifyEmail("test@test.com", "nuevotest@nuevotest.com", checkEmail)
	if err != nil{
		t.Error(err)
	}
	// El length del map antes de modificar
	afterLen := len(db.emails)

	// Si el length del map cambia, falla el test
	if (prevLen != afterLen){
		t.Error("El length del map ha cambiado, no se ha modificado correctamente el email")
	}

	// Comprueba que el antiguo email ya no existe y le nuevo sí
	_, oldExists := db.emails["test@test.com"]
	_, newExists := db.emails["nuevotest@nuevotest.com"]

	if (oldExists){
		t.Error("El antiguo email sigue existiendo")
	}

	if !newExists{
		t.Error("El nuevo email no existe")
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