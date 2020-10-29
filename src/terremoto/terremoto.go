package terremoto

import (  
	"time"
	"src/terremoto/tipomagnitud"
)

type Terremoto struct {  
	idTerremoto string
	fechahora time.Time
	latitud float64
	longitud float64
	profundidad int
	magnitud floafds
	tipoMagnitud tipomagnitud.TipoMagnitud
	localizacion string
}

func NewTerremoto(idTerremoto string, fechahora time.Time, latitud float64, longitud float64, profundidad int, magnitud float64, tipoMagnitud tipomagnitud.TipoMagnitud, localizacion string) *Terremoto{
	terremoto := Terremoto{idTerremoto: idTerremoto, fechahora: fechahora, latitud: latitud, longitud: longitud, profundidad: profundidad, magnitud: magnitud, tipoMagnitud: tipoMagnitud, localizacion: localizacion}

	return &terremoto
}

func getFechaHora() int{

}

func getLatitud() float64{

}

func getLongitud() float64{

}

func getProfundidad() int{

}

func getMagnitud() float64{

}

func getTipoMagnitud() tipomagnitud.TipoMagnitud{

}

func getLocalizacion() string{
	
}