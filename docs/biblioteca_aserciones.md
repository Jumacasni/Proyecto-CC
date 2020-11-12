## Biblioteca de aserciones

Dentro de la biblioteca [**testify**](https://github.com/stretchr/testify) que proporciona **Go**, se encuentra el paquete [**testing**](https://golang.org/pkg/testing/) en el que podemos encontrar dos alternativas para testear:
* Utilizar simplemente el paquete **testing** que proporciona herramientas para realizar tests. Se pueden encontrar funciones como ``t.Error()``, que devuelve un mensaje de error.
* Utilizar el paquete **testing** y además la librería [**assert**](https://godoc.org/github.com/stretchr/testify/assert): una librería externa en la que podemos crear nuestros propios métodos de aserción como por ejemplo ``assertTrue``o ``assertEquals``, entre otros.

### Paquete Testing

Una función test que testea que se ha registrado un email en nuestro sistema es el siguiente:
```
func TestAddEmailOk(t *testing.T){

	prevLen := len(db.emails)
	// Añade el email
	err := AddEmail("prueba@prueba.com", checkEmail)
	if err != nil{
		t.Error(err)
	}
	afterLen := len(db.emails)
	
	if (afterLen == 0){
		t.Error("El length del map es 0, no se ha metido el email correctamente")
	}

	if (afterLen != (prevLen+1)){
		t.Error("El email no se ha metido correctamente")
	}
}
```

En este ejemplo, se lanza un mensaje de error si después de insertar el email:
* El número de emails registrados en el sistema es cero
* El número de emails registrados en el sistema no es el número de emails que había antes de insertar el email + 1

### Paquete Testing + Assert

La misma función test de antes pero usando **assert**, quedaría así:

```
func TestAddEmailOk(t *testing.T){

	prevLen := len(db.emails)
	// Añade el email
	err := AddEmail("prueba@prueba.com", checkEmail)
	if err != nil{
		t.Error(err)
	}
	afterLen := len(db.emails)
	
	assert.Equal(t, afterLen, 0, "El length del map debería ser 0")

	assert.Equal(t, afterLen, prevLen+1, "El length del map debería ser el length de antes + 1")

}
```

Hay que tener en cuenta que se debe importar la librería externa.

### Conclusiones

La diferencia que puedo apreciar entre las dos alternativas es que con **Assert** se testea en una sola línea de código, mientras que usando solo el paquete **Testing** se necesita un **if**. Personalmente, me gusta más el código que queda en la primera alternativa, porque así se separa en líneas de código la condición de error y el mensaje enviado. Usando **Assert** todo va en la misma línea de código y me resulta menos visual.

Finalmente he decidido usar la primera alternativa por la razón que he dicho anteriormente.