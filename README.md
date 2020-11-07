# Terrake

**Autor:** Juan Manuel Castillo Nievas

* [Descripción del proyecto](https://github.com/Jumacasni/Terrake/blob/main/docs/descripcion_proyecto.md)
* [Arquitectura](https://github.com/Jumacasni/Terrake/blob/main/docs/arquitectura.md)
* [Planificación](https://github.com/Jumacasni/Terrake/blob/main/docs/planificacion.md)

## Gestor de tareas

Al usar **Go** como lenguaje de programación, una característica de este lenguaje es que tiene un gestor de tareas implícito. A través del comando ```go <comando> [argumentos]``` se pueden ejecutar diferentes tareas como compilar e instalar paquetes y dependencias o ejecutar tests, entre otras. Todos los comandos que se pueden ejecutar con **Go** se encuentran en [esta sección](https://golang.org/cmd/go/) de su página web.

Una alternativa muy potente al gestor implícito de Go es [realize](https://github.com/oxequa/realize). Añade la posibilidad de manejar diferentes proyectos a la misma vez y la posibilidad de ver los errores y logs en una web de forma más limpia.

Como ahora mismo solo se tiene un proyecto y no considero necesario ver los errores y logs en la web, el **gestor de tareas implícito** es más que suficiente para las tareas de este proyecto, con lo cual se ha decidido usar este.

## Biblioteca de aserciones

**Go** tiene su propia biblioteca de aserciones [testing](https://golang.org/pkg/testing/). Sin embargo, **Go** ha pensado en los programadores que vienen de otros lenguajes de programación y que siempre han usado *assertions*. De este modo, proporcionan la biblioteca [assert](https://godoc.org/github.com/stretchr/testify/assert).

En este proyecto se va a usar su propia biblioteca **testing**, ya que en lugar de utilizar funciones específicas de aserciones, lo que hace es devolver un error cuando falla el test, y esto simplifica las pruebas.

## Marco de prueba

Se analizan dos posibles marcos de prueba:
* [go test](https://golang.org/pkg/testing/): viene por defecto en **Go**. No requiere librerías externas, con lo cual se tiene tanto la biblioteca de aserciones como la tarea para ejecutar los tests..
* [godoc](https://github.com/cucumber/godog): creada para aplicar *Behavior Driven Development* en **Go**.

Se va a optar por ```go test``` y se va a desarrollar usando *Test Driven Development*. Es cierto que BDD ayuda a entender mejor los tests y, cuando se trata de un proyecto grande en el que hay clientes que no entienden mucho de programación, puede mejorar la comunicación desarrollador-cliente, pero este proyecto es propio y no va a entregarse a nadie que no entienda de programación. Además, con TDD se consigue una mejor calidad de código, que es lo prioritario para mí.

## Licencia

El código de este repositorio está bajo la licencia [GPLv3](./LICENSE).
