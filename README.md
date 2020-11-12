# Terrake

**Autor:** Juan Manuel Castillo Nievas

* [Descripción del proyecto](https://github.com/Jumacasni/Terrake/blob/main/docs/descripcion_proyecto.md)
* [Arquitectura](https://github.com/Jumacasni/Terrake/blob/main/docs/arquitectura.md)
* [Planificación](https://github.com/Jumacasni/Terrake/blob/main/docs/planificacion.md)

## Gestor de tareas

La justificación del gestor de tareas se encuentra [aquí](https://github.com/Jumacasni/Terrake/blob/main/docs/gestion_tareas.md).

## Biblioteca de aserciones

**Go** tiene su propia biblioteca de aserciones [testing](https://golang.org/pkg/testing/). Sin embargo, **Go** ha pensado en los programadores que vienen de otros lenguajes de programación y que siempre han usado *assertions*. De este modo, proporcionan la biblioteca [assert](https://godoc.org/github.com/stretchr/testify/assert).

En este proyecto se va a usar su propia biblioteca **testing**, ya que en lugar de utilizar funciones específicas de aserciones, lo que hace es devolver un error cuando falla el test, y esto simplifica las pruebas.

## Marco de prueba

Se analizan dos posibles marcos de prueba:
* [go test](https://golang.org/pkg/testing/): viene por defecto en **Go**. No requiere librerías externas.
* [godoc](https://github.com/cucumber/godog): creada para aplicar *Behavior Driven Development* en **Go**.

Se va a optar por ```go test``` y se va a desarrollar usando *Test Driven Development*. Es cierto que BDD ayuda a entender mejor los tests y, cuando se trata de un proyecto grande en el que hay clientes que no entienden mucho de programación, puede mejorar la comunicación desarrollador-cliente, pero este proyecto es propio y no va a entregarse a nadie que no entienda de programación. Además, con TDD se consigue una mejor calidad de código, que es lo prioritario para mí.

## Instrucciones

Para ejecutar los tests unitarios simplemente hacer:
```shell
$ make test
```

## Clases avanzadas

Se han avanzado las siguientes HU:
* **Milestone: [Notificaciones](https://github.com/Jumacasni/Terrake/milestone/6)** - [notificaciones.go](https://github.com/Jumacasni/Terrake/blob/main/src/notificaciones/notificaciones.go) y [notificaciones_test.go](https://github.com/Jumacasni/Terrake/blob/main/src/notificaciones/notificaciones_test.go):
  * [[HU] Registrar email](https://github.com/Jumacasni/Terrake/issues/71): se ha creado la función [``AddEmail``](https://github.com/Jumacasni/Terrake/blob/9cda240270382e4a0157332ea6261b7af7be60b8/src/notificaciones/notificaciones.go#L53) que registra un email de un usuario. Se han hechos dos tests: [``TestAddEmailOk``](https://github.com/Jumacasni/Terrake/blob/9cda240270382e4a0157332ea6261b7af7be60b8/src/notificaciones/notificaciones_test.go#L36) que testea cuando se añade un email que aún no está registrado, y [``TestAddEmailFail``](https://github.com/Jumacasni/Terrake/blob/9cda240270382e4a0157332ea6261b7af7be60b8/src/notificaciones/notificaciones_test.go#L64), que testea cuando se intenta registrar un email que ya está registrado. Se ha usado mocking para la función que comprueba la existencia de un email (``emailExistsMock``).
  * [[HU] Modificar correo electrónico de notificaciones](https://github.com/Jumacasni/Terrake/issues/72): se ha creado la función [``ModifyEmail``](https://github.com/Jumacasni/Terrake/blob/9cda240270382e4a0157332ea6261b7af7be60b8/src/notificaciones/notificaciones.go#L66) que modifica el email de un usuario. Se ha hecho un test: [``TestModifyEmail``](https://github.com/Jumacasni/Terrake/blob/9cda240270382e4a0157332ea6261b7af7be60b8/src/notificaciones/notificaciones_test.go#L86), que testea que un email se modifica correctamente. Se ha usado mocking para la función que comprueba la existencia de un email (``emailExistsMock``).
  * [[HU] Dejar de notificar por correo electrónico](https://github.com/Jumacasni/Terrake/issues/73): se ha creado la función [``DeactivateEmail``](https://github.com/Jumacasni/Terrake/blob/9cda240270382e4a0157332ea6261b7af7be60b8/src/notificaciones/notificaciones.go#L87) que desactiva las notificaciones de un usuario. Se han creado dos tests: [``TestDeactivateEmailOk``](https://github.com/Jumacasni/Terrake/blob/9cda240270382e4a0157332ea6261b7af7be60b8/src/notificaciones/notificaciones_test.go#L140) que testea que se desactivan las notificaciones de un usuario que tiene las notificaciones activadas, y [``TestDeactivateEmailFail``](https://github.com/Jumacasni/Terrake/blob/9cda240270382e4a0157332ea6261b7af7be60b8/src/notificaciones/notificaciones_test.go#L174) que testea cuando un usuario intenta desactivar las notificaciones que ya tiene desactivadas. Se ha usado mocking para la función que comprueba la existencia de un email (``emailExistsMock``) y que comprueba si un usuario tiene las notificaciones activadas (``emailActivatedMock``).
  * [[HU] Notificar por correo electrónico](https://github.com/Jumacasni/Terrake/issues/84): se ha creado la función [``SendEmail``](https://github.com/Jumacasni/Terrake/blob/9cda240270382e4a0157332ea6261b7af7be60b8/src/notificaciones/notificaciones.go#L131) que envía un correo al usuario indicado con el mensaje indicado. Se ha creado el test [``TestSendEmailOk``](https://github.com/Jumacasni/Terrake/blob/9cda240270382e4a0157332ea6261b7af7be60b8/src/notificaciones/notificaciones_test.go#L225) que testea que un email se envía con el mensaje correcto. Se ha usado mocking para la función que envía un correo, ``mockSend``.
* **Milestone: [Consultas](https://github.com/Jumacasni/Terrake/milestone/7)** - [monitor.go](https://github.com/Jumacasni/Terrake/blob/main/src/monitor/monitor.go) y [monitor_test.go](https://github.com/Jumacasni/Terrake/blob/main/src/monitor/monitor_test.go):
  * [[HU] Consultar terremotos](https://github.com/Jumacasni/Terrake/issues/70): se ha creado la función [``consultarTerremoto``](https://github.com/Jumacasni/Terrake/blob/9cda240270382e4a0157332ea6261b7af7be60b8/src/monitor/monitor.go#L31) que consulta los terremotos que cumplen los criterios indicados. Se ha creado el test [``TestConsultarTerremotos``](https://github.com/Jumacasni/Terrake/blob/9cda240270382e4a0157332ea6261b7af7be60b8/src/monitor/monitor_test.go#L27) que testea que se devuelve un array de terremotos. Se ha usado mocking para la función que crea el array de terremotos a partir de un string ([``csvToTerremotosMock``](https://github.com/Jumacasni/Terrake/blob/9cda240270382e4a0157332ea6261b7af7be60b8/src/monitor/monitor_test.go#L27)).

## Licencia

El código de este repositorio está bajo la licencia [GPLv3](./LICENSE).
