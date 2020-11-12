# Terrake

**Autor:** Juan Manuel Castillo Nievas

* [Descripción del proyecto](https://github.com/Jumacasni/Terrake/blob/main/docs/descripcion_proyecto.md)
* [Arquitectura](https://github.com/Jumacasni/Terrake/blob/main/docs/arquitectura.md)
* [Planificación](https://github.com/Jumacasni/Terrake/blob/main/docs/planificacion.md)

## Gestor de tareas

La justificación del gestor de tareas se encuentra [aquí](https://github.com/Jumacasni/Terrake/blob/main/docs/gestion_tareas.md).

## Biblioteca de aserciones

La justificación de la biblioteca de aseriones se encuentra [aquí](https://github.com/Jumacasni/Terrake/blob/main/docs/biblioteca_aserciones.md).

## Marco de prueba

La justificación del marco de prueba se encuentra [aquí](https://github.com/Jumacasni/Terrake/blob/main/docs/marco_prueba.md).

## Instrucciones

Para ejecutar los tests unitarios simplemente hacer:
```shell
$ tusk test
```

## Estructura del directorio y estándares seguidos

De acuerdo a estos [estándares de Go](https://vsupalov.com/go-folder-structure/), en la carpeta ``internal/`` están los paquetes que son específicos de este proyecto. Además, los archivos de *unit test* están dentro del paquete donde está el código que se va a testear. La estructura del directorio está ahora mismo así:

```
├── internal
│   ├── monitor
│   │   ├── monitor.go
│   │   ├── monitor_test.go
│   │   └── terremoto
│   │       ├── terremoto.go
│   │       └── tipomagnitud
│   │           └── tipoMagnitud.go
│   └── notificaciones
│       ├── notificaciones.go
│       └── notificaciones_test.go
```

## Clases avanzadas

**No ha habido comentarios con respecto a esta rúbrica, con lo cual se ha dejado tal y como estaba**. Se puede consultar en [este enlace](https://github.com/Jumacasni/Terrake/blob/main/docs/clases_avanzadas_hito_2.md)

## Licencia

El código de este repositorio está bajo la licencia [GPLv3](./LICENSE).
