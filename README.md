# Terrake

**Autor:** Juan Manuel Castillo Nievas

## Descripción del proyecto

La descripción de este proyecto se puede consultar en [este enlace](https://github.com/Jumacasni/Terrake/blob/main/docs/descripcion_proyecto.md).

## Arquitectura

Este proyecto se va a desarrollar usando una arquitectura basada en microservicios. En esta aplicación se distinguen varias funcionalidades que pueden ser diseñadas e implementadas de forma independiente. Al contrario que una arquitectura monolítica, esta arquitectura presenta las siguientes ventajas:
* Cada funcionalidad se separa en un **microservicio independiente**. Si un microservicio falla, no implica que todos los demás microservicios dejen de ofrecer su funcionalidad.
* **Implementar** nuevas características o **mejorar** las ya existentes es más fácil. En una arquitectura monolítica se tendría todo el código en un mismo fichero y resultaría complejo realizar una modificación, puesto que ese fichero va a ir creciendo en cuanto a código.
* La **comunicación** entre cada microservicio se realiza mediante un sistema de paso de mensajes de forma muy sencilla.
* Cada microservicio se **escala** de forma totalmente independiente de los demás.

Esta arquitectura encaja muy bien en este proyecto porque se quiere tener distintos servicios independientes que se comuniquen entre ellos. Por ejemplo, la gestión de la base de datos y el análisis de datos son dos funcionalidades que se separan en microservicios independientes. En concreto, en este proyecto encontramos los siguientes microservicios:
* Gestión de la base de datos
* Gestión de las consultas de terremotos
* Gestión del análisis de datos
* Gestión de las notificaciones

### Lenguaje

Tal y como se ha mencionado en la explicación de los microservicios, este proyecto se va a realizar en **Node.js**.

### Bases de datos

- **Terremotos:** contiene todos los terremotos que van apareciendo en el catálogo de terremotos del IGN.
- **Terremotos reportados:** contiene todos los terremotos que los usuarios de Telegram van reportando.
- **Usuarios:** usuarios de Telegram que han iniciado el bot.

## Planificación del proyecto (ROADMAP)

La planificación del proyecto se puede ver en [este enlace](https://github.com/Jumacasni/Terrake/projects/1).

### Historias de Usuario

- [[HU1] Consulta de terremotos a través de Twitter](https://github.com/Jumacasni/Terrake/issues/9)
- [[HU2] Notificación de terremotos a través de Telegram](https://github.com/Jumacasni/Terrake/issues/10)
- [[HU3] Consulta de terremotos a través de Telegram](https://github.com/Jumacasni/Terrake/issues/11)
- [[HU4] Reportar un terremoto a través de Telegram ](https://github.com/Jumacasni/Terrake/issues/12)
- [[HU5] Consultar terremotos reportados recientemente en Telegram](https://github.com/Jumacasni/Terrake/issues/13)
- [[HU6] Visualización gráfica de terremotos a través de Telegram](https://github.com/Jumacasni/Terrake/issues/14)
- [[HU7] Dar de baja notificaciones de nuevos terremotos en Telegram](https://github.com/Jumacasni/Terrake/issues/17)

## Clases creadas

Se han creado las siguientes clases:
- [terremoto.js](https://github.com/Jumacasni/Terrake/blob/main/src/terremoto.js)
- [terremo_reportado.js](https://github.com/Jumacasni/Terrake/blob/main/src/terremoto_reportado.js)
- [usuario.js](https://github.com/Jumacasni/Terrake/blob/main/src/usuario.js)

## Licencia

El código de este repositorio está bajo la licencia [GPLv3](./LICENSE).
