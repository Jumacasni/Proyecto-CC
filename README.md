# Terrake

La configuración correcta del entorno de Git y Github en mi máquina local se puede encontrar [aquí](https://github.com/Jumacasni/Proyecto-CC/blob/main/docs/configuracion_git.md).

## Descripción del proyecto

Este proyecto tiene el nombre de Terrake. Se pretende hacer una monitorización de la actividad sísmica que tiene lugar en España. En este [enlace](http://www.ign.es/web/resources/docs/IGNCnig/SIS-Tablas_estadisticas_PIberica.pdf) se pueden encontrar unas estadísticas del número de terremotos que han tenido lugar en los últimos 40 años. Dado que la frecuencia de actividad sísmica es muy alta en España debido al contacto de la placa africana y euroasiática que tiene lugar en el sur de la península, el mar de Alborán y el norte de Marruecos, resulta interesante monitorizar esta actividad para mantener en alerta a los habitantes en caso de terremotos muy sucesivos que puedan desembocar en un terremoto de magnitud mayor.

El Instituto Geográfico Nacional [(IGN)](https://www.ign.es/web/ign/portal/sis-catalogo-terremotos) proporciona un catálogo de terremotos en formato CSV en el que se puede consultar la actividad sísmica, pudiendo analizar distintos datos como fecha y hora, latitud y longitud, profundidad, magnitud y localización.

Se realizará un análisis de datos semanal en el que se analicen, entre otras cosas, las zonas más afectadas y los terremotos de mayor magnitud. Este análisis de datos, junto con la información actualizada de cada terremoto, se publicará vía **Twitter** y **Telegram**:
  1. **Twitter**: se usa para publicar la información acerca de cada terremoto ocurrido y las estadísticas semanales. Su finalidad es ofrecer la información más actualizada a los usuarios.
  2. **Telegram**: además de ofrecer la información de cada terremoto y las estadísticas semanales, los usuarios podrán usar este bot para filtrar los terremotos utilizando distintos filtros tales como la fecha y hora, profunidad, magnitud y localización.
  
La elección de estas dos aplicaciones se debe a que son aplicaciones muy populares actualmente, teniendo un gran número de usuarios activos que pueden consultar esta información de forma muy sencilla y sin necesidad de instalar aplicaciones externas.
