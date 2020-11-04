# Terrake

**Autor:** Juan Manuel Castillo Nievas

## Descripción del proyecto

La descripción de este proyecto se puede consultar en [este enlace](https://github.com/Jumacasni/Terrake/blob/main/docs/descripcion_proyecto.md).

## Arquitectura

Al principio podría pensarse una arquitectura monolítica, de forma que agruparíamos toda la funcionalidad en una misma base de código, pero una vez que se han estudiado las distintas funcionalidades que integra el producto completo, se ha visto que pueden diferenciarse claramente cuatro microservicios y que cada uno de ellos realiza una tarea distinta:
* **Monitorización:** se encarga de realizar las consultas de los usuarios acerca de los terremotos. Las consultas se harán a través de una petición POST a la base de datos del IGN.
* **Notificaciones:** se encarga de notificar a los usuarios a través de sus correos electrónicos de los terremotos que ocurren.

Cada microservicio realiza una funcionalidad que no depende de otro pero que uniendo todos estos resulta en el producto total, con lo cual se ha elegido finalmente una **arquitectura basada en microservicios**. Como desarrollador de este proyecto, esta arquitectura me permite una mayor facilidad en el avance de cada microservicio, pues cada uno tiene su propio código y esto hace que resulten más sencillos y ordenados los cambios que puedan surgir en cada uno. Lo que más me beneficia de esta arquitectura es que el software va a ser fácilmente personalizable y escalable cuando se despliegue en la nube, al contrario que ocurriría si se hubiera optado por la arquitectura monolítica.

### Herramientas

* **Lenguaje:** en un principio se pensó *Node.js* por la experiencia que tengo en este lenguaje, pero finalmente se va a realizar en **Go** ya que nunca he tenido la oportunidad de aprenderlo y este proyecto es el momento para hacerlo.

## Planificación del proyecto

El desarrollo de este proyecto se va a dividir en las siguientes fases:

### *Milestone* 1, primera fase: [**consultas**](https://github.com/Jumacasni/Terrake/milestone/7)

En esta primera fase se va a desarrollar este *milestone* que da como resultado el producto mínimo viable en el que los usuarios pueden realizar consultas de terremotos. La historia de usuario que cubre este *milestone* es:

* [[HU1] Consultar terremotos](https://github.com/Jumacasni/Terrake/issues/70)
  * [Acceder al catálogo de terremotos](https://github.com/Jumacasni/Terrake/issues/74)

### *Milestone* 2, segunda fase: [**notificaciones**](https://github.com/Jumacasni/Terrake/milestone/6)

Una vez que los usuarios realizan consultas de terremotos, los usuarios pueden indicar su correo electrónico para recibir notificaciones cada vez que ocurra un terremoto. Del mismo modo, podrán modificar este correo electrónico o bien eliminarlo para desactivar las notificaciones. Las historias de usuario que se han creado son las siguientes:

* [[HU] Activar notificaciones](https://github.com/Jumacasni/Terrake/issues/71)
* [[HU] Modificar el correo electrónico de las notificaciones](https://github.com/Jumacasni/Terrake/issues/72)
* [[HU] Desactivar notificaciones](https://github.com/Jumacasni/Terrake/issues/73)

## Clases creadas

* En la [[HU1] Consultar terremotos](https://github.com/Jumacasni/Terrake/issues/70) se han avanzado las siguientes clases:
  * [terremoto.go](https://github.com/Jumacasni/Terrake/blob/main/src/terremoto/terremoto.go): se crea la clase **terremoto**.
  * [tipoMagnitud.go](https://github.com/Jumacasni/Terrake/blob/main/src/terremoto/tipomagnitud/tipoMagnitud.go): contiene el *enum* con los tipos de magnitudes que puede tener un terremoto. Necesario para el atributo *tipoMagnitud* de la clase **terremoto**.
  * [monitor.go](https://github.com/Jumacasni/Terrake/blob/main/src/controladorTerremotos.go): se crea la clase **monitor** que es la que permite consultar terremotos a través de la URL del IGN. Contiene la función que permite a un usuario consultar terremotos.
  
La estructura del directorio ``src`` está así actualmente:

```
src
├── controladorTerremotos.go
├── terremoto
│   ├── terremoto.go
│   └── tipomagnitud
│       └── tipoMagnitud.go
└── usuario
    └── usuario.go

```

### Comprobación de la sintaxis

Para la comprobación de que una clase está sintácticamente correcta se ha utilizado el comando ``gofmt -e <fichero.go>``, cuya salida indica los errores sintácticos encontrados.

## Licencia

El código de este repositorio está bajo la licencia [GPLv3](./LICENSE).
