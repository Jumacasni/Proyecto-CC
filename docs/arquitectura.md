## Arquitectura

Se ha elegido una **arquitectura basada en microservicios** para así para poder manejar las distintas funcionalidades de forma independiente en un servicio diferente. Por ejemplo, en este proyecto se necesita acceder al catálogo de terremotos mediante peticiones HTTP POST, con lo cual esta arquitectura nos permite realizar esta funcionalidad en un servicio independiente, sin afectar a las demás funcionalidades que se vayan a implementar. Lo que más me beneficia de esta arquitectura es que el software va a ser fácilmente personalizable y escalable cuando se despliegue en la nube, al contrario que ocurriría si se hubiera optado por la arquitectura monolítica.

### Herramientas

* **Lenguaje:** en un principio se pensó *Node.js* por la experiencia que tengo en este lenguaje, pero finalmente se va a realizar en **Go** ya que nunca he tenido la oportunidad de aprenderlo y este proyecto es el momento para hacerlo.