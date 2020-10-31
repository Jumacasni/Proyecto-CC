package usuario

type Usuario struct {
	nombre string
	apellidos string
	email string
	telefono string
}

func NewUsuario(nombre string, apellidos string, email string, telefono string) *Usuario{
	usuario := Usuario{nombre: nombre, apellidos: apellidos, email: email, telefono: telefono}

	return &usuario
}

func getEmail() string{

}

func getTelefono() string{

}