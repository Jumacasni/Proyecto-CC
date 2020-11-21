## Elección del contenedor base

Para la elección del contenedor base se han creado 4 contenedores docker usando las siguientes imágenes base:
* Imagen oficial de ``golang``: [golang-latest](https://hub.docker.com/_/golang)

```
FROM golang:latest
```

* ``Alpine`` como SO que ya tiene la última versión 1.15 de ``golang``: [golang:1.15-alpine](https://hub.docker.com/_/golang)

Adicionalmente hay que instalar ``curl`` y ``bash`` para poder descargar [tusk](https://github.com/rliebz/tusk) y ``build-base`` que instala ``gcc`` y que [es necesario](https://stackoverflow.com/questions/59471545/when-trying-to-build-docker-image-i-get-gcc-executable-file-not-found-in-p) para poder usar ``golang``.

```
FROM golang:1.15-alpine
RUN apk --no-cache add curl \
	&& apk add --no-cache bash \
	&& apk add --no-cache build-base
```

* Solo ``Alpine``: [alpine](https://hub.docker.com/_/alpine)

Además de instalar ``curl``, ``bash`` y ``build-base`` como anteriormente, hay que instalar ``golang``.

```
FROM alpine:latest
RUN apk add --no-cache go \
	&& apk --no-cache add curl \
	&& apk add --no-cache bash \
	&& apk add --no-cache build-base
```

* ``Ubuntu`` como SO base que ya tiene ``golang`` instalado: [partylab/ubuntu-golang](https://hub.docker.com/r/partlab/ubuntu-golang)

```
FROM partlab/ubuntu-golang
```

A continuación, se ha usado ``docker images`` para comprobar el tamaño de cada una y el resultado ha sido el siguiente:

<img src="https://github.com/Jumacasni/Terrake/blob/hito3/docs/img/dockerfiles.png" width="60%" height="60%">

Tal y como se ve, la imagen base ``golang:1.15-alpine`` es la que menos pesa con un tamaño de 507MB, con lo cual esta es la imagen elegida.