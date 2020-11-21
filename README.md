# Terrake

**Autor:** Juan Manuel Castillo Nievas

**Nota: este proyecto aún está en desarrollo**

## Descripción y planificación del proyecto

* [Descripción del proyecto](https://github.com/Jumacasni/Terrake/blob/main/docs/descripcion_proyecto.md)
* [Arquitectura](https://github.com/Jumacasni/Terrake/blob/main/docs/arquitectura.md)
* [Planificación](https://github.com/Jumacasni/Terrake/blob/main/docs/planificacion.md)

## Tests

* **r1** - [Gestor de tareas](https://github.com/Jumacasni/Terrake/blob/main/docs/gestion_tareas.md)
* **r2** - [Biblioteca de aserciones](https://github.com/Jumacasni/Terrake/blob/main/docs/biblioteca_aserciones.md)
* **r3** - [Marco de prueba](https://github.com/Jumacasni/Terrake/blob/main/docs/marco_prueba.md)

## Uso de contenedores

* **r1** - [Elección del contenedor base](https://github.com/Jumacasni/Terrake/blob/main/docs/eleccion_contenedor_base.md)
* **r2** - [Dockerfile](https://github.com/Jumacasni/Terrake/blob/main/docs/dockerfile.md)
* **r3** - [Subida a Docker Hub y actualización automática](https://github.com/Jumacasni/Terrake/blob/main/docs/dockerhub.md)
* **r4** - [GitHub Container Registry](https://github.com/Jumacasni/Terrake/blob/main/docs/githubcontainerregistry.md)

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

## Licencia

El código de este repositorio está bajo la licencia [GPLv3](./LICENSE).
