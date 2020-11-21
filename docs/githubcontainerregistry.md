## GitHub Container Registry

Se ha publicado el [contenedor](https://github.com/users/Jumacasni/packages/container/terrake/356186) en **GitHub Container Registry** siguiendo [estas instrucciones](https://codefresh.io/docs/docs/integrations/docker-registries/github-container-registry/#using-the-github-container-registry).

**1.** He creado un *Personal Access Token* (PAT) con los permisos que se encuentran en la siguiente imagen.

<img src="https://github.com/Jumacasni/Terrake/blob/main/docs/img/gcr1.png" width="80%" height="80%">

**2.** He exportado el PAT creado a ``$CR_PAT`` y me he logeado con la url ``ghcr.io``, mi usuario de GitHub y el PAT.

<img src="https://github.com/Jumacasni/Terrake/blob/main/docs/img/gcr2.png" width="80%" height="80%">

**3.** He tageado la imagen que se va a subir con el nombre de ``ghcr.io/jumacasni/terrake:latest``. La imagen se corresponde con la que se creó usando la [tercera versión](https://github.com/Jumacasni/Terrake/blob/hito3/docs/dockerfile.md) del Dockerfile.

<img src="https://github.com/Jumacasni/Terrake/blob/main/docs/img/gcr3.png" width="80%" height="80%">

**4.** Se ha hecho ``push`` de la imagen

<img src="https://github.com/Jumacasni/Terrake/blob/main/docs/img/gcr4.png" width="80%" height="80%">

**5.** Una vez pusheada, en el apartado **Packages** de mi perfil de GitHub aparece el paquete.

<img src="https://github.com/Jumacasni/Terrake/blob/main/docs/img/gcr5.png" width="80%" height="80%">

<img src="https://github.com/Jumacasni/Terrake/blob/main/docs/img/gcr6.png" width="80%" height="80%">

**6.** Se ha linkeado el paquete a este repositorio y se ha hecho público (por defecto estaba privado)

<img src="https://github.com/Jumacasni/Terrake/blob/main/docs/img/gcr7.png" width="80%" height="80%">

<img src="https://github.com/Jumacasni/Terrake/blob/main/docs/img/gcr8.png" width="80%" height="80%">