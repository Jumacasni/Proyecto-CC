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
	err := EmailActivated("prueba@prueba.com")
	if err != nil{
		t.Error(err)
	}
}

/***** MOCK PARA COMPROBAR ACTIVACIÓN DE NOTIFICACIONES *****/
var emailActivatedMock func(email string) bool

type checkEmailActivatedMock struct{}

func (u checkEmailActivatedMock) emailActivated(email string) bool{
	return emailActivatedMock(email)
}

/***** TEST DEACTIVATE EMAIL *****/
// Test para desactivar las notificaciones de "prueba@prueba.com" (están activadas)
func TestDeactivateEmailOk(t *testing.T){
	// Función mock, el email existe
	checkEmail := checkEmailMock{}
	emailExistsMock = func(email string) bool{
		return true
	}
	checkEmailActivated := checkEmailActivatedMock{}
	// Función mock, el email tiene las notificaciones activadas
	emailActivatedMock = func(email string) bool{
		return true
	}

	// Email de prueba
	emailPrueba := "prueba@prueba.com"
	// Valor de las notificaciones antes de desactivarlas
	prevValue := db.emails[emailPrueba]
	err := DeactivateEmail(emailPrueba, checkEmail, checkEmailActivated)
	if err != nil{
		t.Error(err)
	}
	// Valor de las notificaciones después de desactivarlas
	afterValue := db.emails[emailPrueba]

	// El valor del email tiene que ser distinto del valor antes de DeactivateEmail
	if (db.emails[emailPrueba] == prevValue){
		t.Error("Las notificaciones siguen activadas")
	}
	// El valor del email tiene que ser igual al valor después de DeactivateEmail
	if (db.emails[emailPrueba] != afterValue){
		t.Error("Las notificaciones no se han desactivado")
	}
}

// Test para desactivar las notificaciones de "prueba@prueba.com" (están desactivadas)
func TestDeactivateEmailFail(t *testing.T){
	// Función mock, el email existe
	checkEmail := checkEmailMock{}
	emailExistsMock = func(email string) bool{
		return true
	}
	checkEmailActivated := checkEmailActivatedMock{}
	// Función mock, el email tiene las notificaciones desactivadas
	emailActivatedMock = func(email string) bool{
		return false
	}

	// Email de prueba
	emailPrueba := "prueba@prueba.com"
	// Valor de las notificaciones antes de desactivarlas
	prevValue := db.emails[emailPrueba]
	DeactivateEmail(emailPrueba, checkEmail, checkEmailActivated)
	// Valor de las notificaciones después de desactivarlas
	afterValue := db.emails[emailPrueba]

	// El valor de antes y después de la función DeactivateEmail tiene que ser el mismo
	if (prevValue != afterValue){
		t.Error("El email tenía las notificaciones desactivadas y se han activado")
	}
	// El valor del email tiene que ser igual al valor después de DeactivateEmail
	if (afterValue != false){
		t.Error("Las notificaciones se han activado")
	}
}