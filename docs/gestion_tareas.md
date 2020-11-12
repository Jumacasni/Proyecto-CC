## Gestor de tareas

Se han explorado tres alternativas, y para cada alternativa se ha probado a realizar la tarea que ejecuta los test. Las opciones son:
* [tusk](https://github.com/rliebz/tusk)
* [task](https://github.com/go-task/task)
* [robo](https://github.com/tj/robo)

### Tusk

Se debe crear un archivo llamado ``tusk.yml`` que tiene la siguiente sintaxis:
```
tasks:
  test:
    usage: test notificaciones and monitor
    run: cd internal/; go test ./monitor; go test ./notificaciones
```

Se puede ver la descripción y uso de la tarea con ``tusk <tarea> --help``:
<img src="https://github.com/Jumacasni/Terrake/blob/main/docs/img/tusk1.png" width="55%" height="55%">

Al ejecutar ``tusk test`` se muestra la siguiente salida:
<img src="https://github.com/Jumacasni/Terrake/blob/main/docs/img/tusk2.png" width="55%" height="55%">

El comando ejecutado podría simplificarse a ``cd internal; go test ./...``, pero de esta manera mostraría *warnings* de que hay ficheros que no tienen test, y estos ficheros realmente no los necesitan:
<img src="https://github.com/Jumacasni/Terrake/blob/main/docs/img/tusk3.png" width="55%" height="55%">

### Task
Se debe crear un archivo llamado ``Taskfile.yml`` que tiene la siguiente sintaxis:
```
version: '3'

tasks:
  test:
    summary: test notificaciones and monitor
    cmds:
      - cd internal/; go test ./monitor; go test ./notificaciones
```

Con el comando ``task <tarea> --summary`` se puede ver una descripción de la tarea:
<img src="https://github.com/Jumacasni/Terrake/blob/main/docs/img/task2.png" width="55%" height="55%">

Al ejecutar ``task test`` se muestra la siguiente salida:
<img src="https://github.com/Jumacasni/Terrake/blob/main/docs/img/task1.png" width="55%" height="55%">

### Robo
Se debe crear un archivo llamado ``robo.yml`` que tiene la siguiente sintaxis:
```
test:
  summary: test notificaciones and monitor
  command: cd internal/; go test ./monitor; go test ./notificaciones
```

Al ejecutar ``robo help <tarea>`` se muestra una descripción y uso de la tarea:
<img src="https://github.com/Jumacasni/Terrake/blob/main/docs/img/robo1.png" width="55%" height="55%">

Al ejecutar ``robo test`` se muestra la siguiente salida:
<img src="https://github.com/Jumacasni/Terrake/blob/main/docs/img/robo2.png" width="55%" height="55%">

### Conclusiones

Se puede ver que **robo** es el que tiene la sintaxis más simple, pero al ejecutar el test no se muestra nada más allá de lo que hace el comando. En cambio, **task** y **tusk** muestran en la primera línea lo que se ejecuta, con lo cual la decisión va a estar entre estos dos porque esto ayuda visualmente a saber qué se está haciendo.

Finalmente se ha elegido **tusk** porque al ejecutar la tarea se marca de manera mucho más visual los comandos que se ejecutan y porque me ha resultado mucho más fácil seguir su documentación que la de **task**, que no la encontraba muy bien redactada.