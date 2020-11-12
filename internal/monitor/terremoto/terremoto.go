package terremoto

import (  
	"time"
	"./tipomagnitud"
)

type Terremoto struct {  
	idTerremoto string
	fechahora time.Time
	latitud float64
	longitud float64
	profundidad int
	magnitud float64
	tipoMagnitud tipomagnitud.TipoMagnitud
	localizacion string
}

func NewTerremoto(idTerremoto string, fechahora time.Time, latitud float64, longitud float64, profundidad int, magnitud float64, tipoMagnitud tipomagnitud.TipoMagnitud, localizacion string) *Terremoto{
	terremoto := Terremoto{idTerremoto: idTerremoto, fechahora: fechahora, latitud: latitud, longitud: longitud, profundidad: profundidad, magnitud: magnitud, tipoMagnitud: tipoMagnitud, localizacion: localizacion}

	return &terremoto
}