# Terrake

**Autor:** Juan Manuel Castillo Nievas

## Descripción del proyecto

La descripción de este proyecto se puede consultar en [este enlace](https://github.com/Jumacasni/Terrake/blob/main/docs/descripcion_proyecto.md).

## Arquitectura

Al principio podría pensarse una arquitectura monolítica, pero una vez que se han estudiado las distintas funcionalidades que integra el producto completo, se ha visto que pueden diferenciarse claramente cuatro microservicios y que cada uno realiza una tarea distinta:
* **Catálogo de terremotos**
* **Cuentas de usuario**
* **Reporte de terremotos**
* **Envío de notificaciones**

Cada microservicio realiza una funcionalidad que no depende de otro, pero que uniendo todos estos resulta en el producto total, con lo cual se ha elegido una **arquitectura basada en microservicios**. Como desarrollador de este proyecto, esta arquitectura me permite una mayor facilidad en el avance de cada microservicio, pues cada uno tiene su propio código y esto hace que resulten más sencillos y ordenados los cambios que puedan surgir en cada uno. Lo que más me beneficia de esta arquitectura es que el software va a ser fácilmente personalizable y escalable cuando se despliegue en la nube, al contrario que ocurriría si se hubiera optado por la arquitectura monolítica.

### Herramientas

* **Lenguaje:** en un principio se pensó *Node.js* por la experiencia que tengo en este lenguaje, pero finalmente se va a realizar en **Go** ya que nunca he tenido la oportunidad de aprenderlo y este proyecto es el momento para hacerlo.

## Planificación del proyecto

El desarrollo de este proyecto se va a dividr en las siguientes fases:

### Primera fase

En esta primera fase se van a construir algunas funcionalidades de la gestión de consultas de terremotos. Para ello, se creará la interfaz de la entidad principal **terremoto**. Durante esta fase se van a construir las funcionaliades básicas que contribuyen a la gestión de consultas mediante las siguientes historias de usuario:

* [[HU1] Consulta general de terremotos en los últimos X días/meses/años](https://github.com/Jumacasni/Terrake/issues/29)
* [[HU2] Consulta de terremotos por filtros](https://github.com/Jumacasni/Terrake/issues/30)
* [[HU4] Consulta de un análisis de datos](https://github.com/Jumacasni/Terrake/issues/32)

### Segunda fase

En esta segunda fase se pretende distinguir entre usuarios anónimos (no registrados) y registrados. Para ello, se crear la interfaz de la entidad **usuario**. Un usuario registrado puede recibir notificaciones de terremotos y además eliminar reportes de terremotos que haya hecho por error o por cualquier otro motivo. En concreto, en esta fase se crearán las funcionalidades que recogen las siguientes historias de usuario:

* [[HU7] Registrar usuario](https://github.com/Jumacasni/Terrake/issues/35)
* [[HU10] Activar notificaciones de terremotos](https://github.com/Jumacasni/Terrake/issues/38)
* [[HU11] Restringir las notificaciones recibidas](https://github.com/Jumacasni/Terrake/issues/39)

### Tercera fase

En esta tercera fase se completarán las funcionalidades adicionales que faltaban en la segunda fase. Se pretende añadir estas funcionalidades una vez funcione todo lo creado en la fase anterior, puesto que lo que hacen es complementar a las ya creadas. Esta fase recoge las siguientes historias de usuario:

* [[HU8] Modificar cuenta de usuario](https://github.com/Jumacasni/Terrake/issues/36)
* [[HU9] Eliminar cuenta de usuario](https://github.com/Jumacasni/Terrake/issues/37)
* [[HU12] Desactivar las notificaciones](https://github.com/Jumacasni/Terrake/issues/40)

### Cuarta fase

En esta última fase de desarrollo se creará el microservicio que lleva la gestión de los reportes de terremotos. Para ello, se creará la entidad **reporte** y podrá hacerla tanto un usuario anónimo como uno registrado. Además, se añadirá una funcionalidad a la gestión de consultas que permite ver en un mapa el epicentro de un terremoto concreto. Esta fase final dará como resultado el **producto completo** con todas las funcionalidades. Las historias de usuario de esta fase son las siguientes:

* [[HU5] Reportar que se ha sentido un terremoto](https://github.com/Jumacasni/Terrake/issues/33)
* [[HU6] Eliminar el reporte de un terremoto](https://github.com/Jumacasni/Terrake/issues/34)
* [[HU3] Visualización de un terremoto en un mapa](https://github.com/Jumacasni/Terrake/issues/31)

## Clases creadas

Se han creado las siguientes clases:
- [terremoto.go](https://github.com/Jumacasni/Terrake/blob/main/src/terremoto/terremoto.go): se crea la clase **terremoto**, enlazada a [[HU1]](https://github.com/Jumacasni/Terrake/issues/29)
- [tipomagnitud.go](https://github.com/Jumacasni/Terrake/blob/main/src/terremoto/tipomagnitud/tipomagnitud.go): contiene el *enum* con los tipos de magnitudes que puede tener un terremoto, creado junto con la clase **terremoto** en [[HU1]](https://github.com/Jumacasni/Terrake/issues/29)
- [controlador.go](https://github.com/Jumacasni/Terrake/blob/main/src/controlador.go): se crea la clase **controlador** para crear o destruir terremotos, enlazada a [[HU1]](https://github.com/Jumacasni/Terrake/issues/29)

### Comprobación de la sintaxis

Para la comprobación de que una clase está sintácticamente correcta se ha utilizado el comando ``gofmt -e <fichero.go>``, cuya salida indica los errores sintácticos encontrados.

## Licencia

El código de este repositorio está bajo la licencia [GPLv3](./LICENSE).
