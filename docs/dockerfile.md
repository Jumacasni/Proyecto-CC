## Dockerfile

* En una primera versión del Dockerfile se ha usado el siguiente script:

```
FROM golang:1.15-alpine
LABEL maintainer="Juan Manuel <jumacasni@correo.ugr.es>"

RUN apk --no-cache add curl
RUN apk add --no-cache bash
RUN apk add --no-cache build-base
RUN curl -sL https://git.io/tusk | bash -s -- -b /usr/local/bin latest

COPY internal/ /terrake/internal/
COPY tusk.yml /terrake

WORKDIR /terrake

CMD ["tusk", "test"]
```

En esta primera versión se han utilizado cuatro órdenes ``RUN`` para depurar. Una vez se haya depurado, se unirán en una misma orden. También se ha copiado el contenido de ``internal/`` que contiene los archivos ``.go`` y el archivo ``tusk.yml`` que lanza los tests. Se obtienen **11 capas**:

<kbd><img src="https://github.com/Jumacasni/Terrake/blob/main/docs/img/dockerfile-v1.png" width="80%" height="80%"></kbd>

---

* En una segunda versión del Dockerfile se ha usado el siguiente script:

```
FROM golang:1.15-alpine
LABEL maintainer="Juan Manuel <jumacasni@correo.ugr.es>"

RUN apk --no-cache add curl \
	&& apk add --no-cache bash \
	&& apk add --no-cache build-base \
	&& curl -sL https://git.io/tusk | bash -s -- -b /usr/local/bin latest

COPY internal/ /terrake/internal/
COPY tusk.yml /terrake

WORKDIR /terrake

CMD ["tusk", "test"]
```

En esta segunda versión se han juntado las cuatro órdenes ``RUN`` en una misma orden. El número de capas es **8**:

<kbd><img src="https://github.com/Jumacasni/Terrake/blob/main/docs/img/dockerfile-v2.png" width="80%" height="80%"></kbd>

---

* En una tercera versión del dockerfile se ha usado el siguiente script:

```
FROM golang:1.15-alpine
LABEL maintainer="Juan Manuel <jumacasni@correo.ugr.es>"

RUN apk --no-cache add curl \
	&& apk add --no-cache bash \
	&& apk add --no-cache build-base \
	&& curl -sL https://git.io/tusk | bash -s -- -b /usr/local/bin latest

COPY internal/ tusk.yml /terrake/

WORKDIR /terrake

CMD ["tusk", "test"]
```

En esta última versión se han juntado las dos órdenes ``COPY`` en una misma orden. Esto implica que se debe modificar levemente el fichero ``tusk.yml`` eliminando la orden ``cd internal/`` que tenía originalmente, ya que no es necesaria puesto que está al mismo nivel que las carpetas ``monitor/`` y ``notificaciones/``. El fichero ``tusk.yml`` se puede consultar en [este enlace](https://github.com/Jumacasni/Terrake/blob/main/tusk.yml).

De esta forma se ahorra una capa, teniendo en total **7 capas**:

<kbd><img src="https://github.com/Jumacasni/Terrake/blob/main/docs/img/dockerfile-v3.png" width="80%" height="80%"></kbd>

---

### Conclusiones

Se ha usado ``docker images`` para comprobar el tamaño que queda en cada dockerfile:

<kbd><img src="https://github.com/Jumacasni/Terrake/blob/main/docs/img/dockerfiles-v*.png" width="80%" height="80%"></kbd>

Aunque es cierto que el tamaño no ha variado, al seguir las [buenas prácticas](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/) de Dockerfile se ha conseguido pasar de **11 capas** a **7 capas**.
