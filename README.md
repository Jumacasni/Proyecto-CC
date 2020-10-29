# Terrake

**Autor:** Juan Manuel Castillo Nievas

## Descripción del proyecto

La descripción de este proyecto se puede consultar en [este enlace](https://github.com/Jumacasni/Terrake/blob/main/docs/descripcion_proyecto.md).

## Arquitectura

Este proyecto se va a desarrollar usando una arquitectura basada en microservicios. Se ha elegido esta arquitectura porque nos permite diferenciar entre distintos microservicios que ofrecen funcionalidades independientes. Por ejemplo, la **la gestión de la base de datos** en la que se van almacenando todos los terremotos registrados forma un microservicio, y la **gestión de consultas de terremotos** forma otro microservicio independiente que sirve para gestionar todas las consultas que realicen los usuarios, comunicándose ambos servicios mediante una API. Esta arquitectura presenta las siguientes ventajas (poniendo como ejemplo los dos microservicios mencionados):
* Si el microservicio de las consultas falla, eso no implica que el microservicio que gestiona la base de datos deje de funcionar. Va a seguir cogiendo datos de terremotos y metiéndolos en la base de datos.
* Implementar una nueva funcionalidad en la gestin de consultas seria mucho más fácil puesto que ambos microservicios no comparten un mismo fichero, si no que cada uno tiene el suyo propio y es mucho más limpio y ordenado incluir nuevas funcionalidades.
* Cada microservicio se escala de forma totalmente independiente de los demás.

En concreto, en este proyecto encontramos los siguientes microservicios:
* **Gestión de la base de datos**
* **Gestión de las consultas de terremotos**
* **Gestión análisis de datos**
* **Gestión de usuarios**
* **Gestión del reporte de terremotos**

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
