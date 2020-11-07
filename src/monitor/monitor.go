package monitor

import (
	"./terremoto"
	"time"
)

type Monitor struct {
	// URL a la que se va a hacer la petición
	url string  
}

var (
	monitor Monitor
)

func init() {
	InitURL("https://www.ign.es/web/ign/portal/sis-catalogo-terremotos?p_p_id=IGNSISCatalogoTerremotos_WAR_IGNSISCatalogoTerremotosportlet&p_p_lifecycle=2&p_p_state=normal&p_p_mode=view&p_p_cacheability=cacheLevelPage&p_p_col_id=column-1&p_p_col_count=1&_IGNSISCatalogoTerremotos_WAR_IGNSISCatalogoTerremotosportlet_jspPage=%2Fjsp%2Fterremoto.jsp")
}

func InitURL(url string){
	monitor.url = url
}

// Interfaz para la función mock csvToTerremotos que parsea el input y lo convierte a terremotos
type GetTerremotos interface{
	csvToTerremotos(string) []terremoto.Terremoto
}

// Consultar de terremotos, devuelve una lista de terremotos
func consultarTerremoto(fechaInicio time.Time, fechaFinal time.Time, latitudMinima float64, latitudMaxima float64, longitudMinima float64, longitudMaxima float64, profundidadMinima int, profundidadMaxima int, magnitudMinima float64, magnitudMaxima float64, getTerremotos GetTerremotos) (error, []terremoto.Terremoto){
	// Inicializa el array de terremotos a vacío
	var terremotos []terremoto.Terremoto

	terremotos = getTerremotos.csvToTerremotos("")
	return nil, terremotos
}