# Terrake

La configuración correcta del entorno de Git y Github en mi máquina local se puede encontrar [aquí](https://github.com/Jumacasni/Terrake/blob/main/docs/configuracion_git.md).

## Descripción del proyecto

La descripción de este proyecto se puede consultar en [este enlace](https://github.com/Jumacasni/Terrake/blob/main/docs/descripcion_proyecto.md).


## Arquitectura

Este proyecto se va a desarrollar usando una arquitectura basada en microservicios. En esta aplicación se distinguen varias funcionalidades que pueden ser diseñadas e implementadas de forma independiente. Al contrario que una arquitectura monolítica, esta arquitectura presenta las siguientes ventajas:
* Cada funcionalidad se separa en un **microservicio independiente**. Si un microservicio falla, no implica que todos los demás microservicios dejen de ofrecer su funcionalidad.
* **Implementar** nuevas características o **mejorar** las ya existentes es más fácil. En una arquitectura monolítica se tendría todo el código en un mismo fichero y resultaría complejo realizar una modificación, puesto que ese fichero va a ir creciendo en cuanto a código.
* La **comunicación** entre cada microservicio se realiza mediante un sistema de paso de mensajes de forma muy sencilla.
* Cada microservicio se **escala** de forma totalmente independiente de los demás.

Este proyecto está formado por los siguientes microservicios:
1. **Microservicio para acceder al [catálogo de terremotos del IGN](https://www.ign.es/web/ign/portal/sis-catalogo-terremotos) y almacenar esos datos:** se deben introducir unos datos en el formulario para obtener un archivo CSV con todos los terremotos que cumplen dichos datos. Para ello, simplemente se realiza una petición **POST** con los datos que se quieran, produciendo como resultado el contenido del archivo CSV. Para este proyecto se va a usar una base de datos **NoSQL**, ya que no todos los datos que se van a almacenar tienen la misma estructura, que supondría un gran desperdicio de espacio en la base de datos. Como el lenguaje que se va a usar es **Node.js**, la [base de datos que mejor se lleva con este lenguaje](https://www.quora.com/What-is-the-best-Node-js-database) es **MongoDB**. Esta base de datos se puede usar de manera muy sencilla a través del paquete **mongoose**, además de contar con una [documentación](https://docs.mongodb.com/) de gran calidad y una [comunidad](https://developer.mongodb.com/community/forums/) bastante buena.
2. **Microservicio para el análisis de datos:** tal y como se ha dicho anteriormente, resulta interesante analizar las estadísticas semanales de los terremotos ocurridos en determinadas zonas, además de analizar las zonas que han tenido mayor actividad sísmica o el terremoto de mayor magnitud que ha ocurrido. Este análisis de datos se hace de acuerdo a toda la información guardada en la base de datos.
3. **Microservicio para la publicación de información en Twitter:** se publicarán tweets cada vez que ocurra un terremoto, ofreciendo toda la información relevante acerca de este. También se publicarán las estadísticas semanales, tal y como se ha descrito en el microservicio anterior. Para este microservicio se va a usar [twit API](https://github.com/ttezel/twit).
4. **Microservicio para la publicación de información en Telegram e interacción con los usuarios:** este microservicio ofrece la misma publicación de información que se hace en Twitter, con la adición de una posible interacción entre el usuario y el bot. Los usuarios pueden hacer uso del bot para filtrar los datos de acuerdo a sus intereses. Los usuarios también podrán reportar un terremoto en caso de que hayan notado uno. Para este microservicio se va a usar [node-telegram-bot](https://github.com/yagop/node-telegram-bot-api), ya que anteriormente he usado esta misma API pero en Python.
5. **Microservicio para el registro (logs):** se necesita un microservicio para administrar los logs de los distintos microservicios. Para ello se va a usar [Logstash](https://www.elastic.co/es/logstash).

La comunicación entre los distintos microservicios se realizará a través de [Express](https://expressjs.com/es/).

Esta arquitectura se describe en la siguiente imagen.

<img src="https://github.com/Jumacasni/Terrake/blob/main/docs/img/arquitectura.png" width="60%" height="60%">

Se tienen en cuenta las siguientes consideraciones:
* Cada vez que se detecte un nuevo terremoto, se debe almacenar en la base de datos y publicar dicha información tanto en Twitter como en Telegram.
* Se hará un análisis semanal cuyo resultado se publicará tanto en Twitter como en Telegram.
* Se desarrollará una API REST para la comunicación entre los distintos microservicios y la API Gateway.

### Lenguaje

Tal y como se ha mencionado en la explicación de los microservicios, este proyecto se va a realizar en **Node.js**.

### Bases de datos

- **Terremotos:** contiene todos los terremotos que van apareciendo en el catálogo de terremotos del IGN.
- **Terremotos reportados:** contiene todos los terremotos que los usuarios de Telegram van reportando.
- **Usuarios:** usuarios de Telegram que han iniciado el bot.

## Planificación del proyecto

La planificación del proyecto se puede ver en [este enlace](https://github.com/Jumacasni/Terrake/projects/1).

## Clases creadas

Se han creado las siguientes clases:
- [terremoto.js](https://github.com/Jumacasni/Terrake/blob/main/src/terremoto.js)
- [terremo_reportado.js](https://github.com/Jumacasni/Terrake/blob/main/src/terremoto_reportado.js)
- [usuario.js](https://github.com/Jumacasni/Terrake/blob/main/src/usuario.js)

## Licencia

El código de este repositorio está bajo la licencia [GPLv3](./LICENSE).
