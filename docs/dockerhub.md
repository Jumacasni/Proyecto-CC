## Subida del contenedor a Docker Hub

Para subir el contenedor creado a Docker Hub, se han seguido [estas instrucciones](https://ropenscilabs.github.io/r-docker-tutorial/04-Dockerhub.html).

1. Me he creado una cuenta en Docker Hub y he hecho login desde la terminal

<img src="https://github.com/Jumacasni/Terrake/blob/main/docs/img/dockerhub1.png" width="80%" height="80%">

2. Se ha renombrado la imagen a ``jumacasni/terrake:latest`` y se ha subido usando ``docker push jumacasni/terrake``.

<img src="https://github.com/Jumacasni/Terrake/blob/main/docs/img/dockerhub2.png" width="80%" height="80%">

El contenedor se encuentra subido en [este enlace](https://hub.docker.com/repository/docker/jumacasni/terrake).

### Actualización automática

Para que el contenedor se construya automáticamente cada vez que se actualice el repositorio se han seguido los siguientes pasos:

1. Se ha vinculado mi cuenta de GitHub a mi cuenta de Docker Hub.

<img src="https://github.com/Jumacasni/Terrake/blob/main/docs/img/dockerhub3.png" width="80%" height="80%">

2. En la pestaña de **Builds** del contenedor subido se ha configurado la actualización automática.

<img src="https://github.com/Jumacasni/Terrake/blob/main/docs/img/dockerhub4.png" width="80%" height="80%">

3. Como resultado, se ha configurado la actualización automática cada vez que se actualice la rama principal del repositorio.

<img src="https://github.com/Jumacasni/Terrake/blob/main/docs/img/dockerhub5.png" width="80%" height="80%">