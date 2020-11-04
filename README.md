# Terrake

**Autor:** Juan Manuel Castillo Nievas

## Descripción del proyecto

La descripción de este proyecto se puede consultar en [este enlace](https://github.com/Jumacasni/Terrake/blob/main/docs/descripcion_proyecto.md).

## Arquitectura

### Herramientas

* **Lenguaje:** en un principio se pensó *Node.js* por la experiencia que tengo en este lenguaje, pero finalmente se va a realizar en **Go** ya que nunca he tenido la oportunidad de aprenderlo y este proyecto es el momento para hacerlo.

## Planificación del proyecto

El desarrollo de este proyecto se va a dividir en las siguientes fases:

### *Milestone* 1, primera fase: [**consultas**](https://github.com/Jumacasni/Terrake/milestone/7)

En esta primera fase se va a desarrollar este *milestone* que da como resultado el producto mínimo viable en el que los usuarios pueden realizar consultas de terremotos. Una consulta es una búsqueda de terremotos de acuerdo a sus características, por ejemplo, *los terremotos ocurridos entre el 01/01/2020 y 31/01/2020*. Dichas características se detallan en la historia de usuario. La historia de usuario que cubre este *milestone* es:

* [[HU] Consultar terremotos](https://github.com/Jumacasni/Terrake/issues/70)
  * [Acceder al catálogo de terremotos](https://github.com/Jumacasni/Terrake/issues/74)

### *Milestone* 2, segunda fase: [**notificaciones**](https://github.com/Jumacasni/Terrake/milestone/6)

Una vez que los usuarios realizan consultas de terremotos, los usuarios pueden indicar su correo electrónico y su nombre para recibir notificaciones cada vez que ocurra un terremoto. Del mismo modo, podrán modificar este correo electrónico o bien eliminarlo para desactivar las notificaciones. Las historias de usuario que se han creado son las siguientes:

* [[HU] Notificar por correo electrónico](https://github.com/Jumacasni/Terrake/issues/71)
* [[HU] Modificar el correo electrónico de las notificaciones](https://github.com/Jumacasni/Terrake/issues/72)
* [[HU] Dejar de notificar por correo electrónico](https://github.com/Jumacasni/Terrake/issues/73)

## Clases creadas

* En la [[HU] Consultar terremotos](https://github.com/Jumacasni/Terrake/issues/70) se han avanzado las siguientes clases:
  * [terremoto.go](https://github.com/Jumacasni/Terrake/blob/main/src/terremoto/terremoto.go): se crea la clase **terremoto**.
  * [tipoMagnitud.go](https://github.com/Jumacasni/Terrake/blob/main/src/terremoto/tipomagnitud/tipoMagnitud.go): contiene el *enum* con los tipos de magnitudes que puede tener un terremoto. Necesario para el atributo *tipoMagnitud* de la clase **terremoto**.
  * [monitor.go](https://github.com/Jumacasni/Terrake/blob/main/src/monitor.go): se crea la clase **monitor** que es la que permite consultar terremotos a través de la URL del IGN. Contiene la función que permite a un usuario consultar terremotos.

* En las [[HU] Notificar por correo electrónico](https://github.com/Jumacasni/Terrake/issues/71), [[HU] Modificar correo electrónico de notificaciones](https://github.com/Jumacasni/Terrake/issues/72) y [[HU] Dejar de notificar por correo electrónico](https://github.com/Jumacasni/Terrake/issues/73) se ha avanzado la siguiente clase:
  * [notificaciones.go](https://github.com/Jumacasni/Terrake/blob/main/src/notificaciones.go): encargada de los correos electrónicos de las notificaciones de los usuarios. Los emails se almacenan en un map cuya clave es el email del usuario y el valor es su nombre.
  
La estructura del directorio ``src`` está así actualmente:

```
├── monitor.go
├── notificaciones.go
├── terremoto
│   ├── terremoto.go
│   └── tipomagnitud
│       └── tipoMagnitud.go
```

### Comprobación de la sintaxis

Para la comprobación de que una clase está sintácticamente correcta se ha utilizado el comando ``gofmt -e <fichero.go>``, cuya salida indica los errores sintácticos encontrados.

## Licencia

El código de este repositorio está bajo la licencia [GPLv3](./LICENSE).
