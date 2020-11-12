## Marco de prueba

Ya que este proyecto se va a desarrollar en **Go** y además se va a usar el paquete **Testing** para testear, se va a usar el marco de prueba que proporcionarna el mismo lenguaje: ``go test``. Este comando ejecutará los archivos cuyo nombre sea ``<nombre>_test.go``.

Tal y como se ha comentado en la sección de justificación del gestor de tareas, si simplemente usamos ``go test`` aparecerán *warnings* si tenemos un archivo ``.go`` que no tenga un test. La flexibilidad de este marco de pruebas permite usar ``go test <archivo>`` para ejecutar solamente el archivo deseado.

De acuerdo a los [estándares de **Go**](https://vsupalov.com/go-folder-structure/), lo normal es que los *unit tests* estén en la misma carpeta del código que van a testear.