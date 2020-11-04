package main

import (
	"src/terremoto"
	"time"
)

type Monitor struct {
	// URL a la que se va a hacer la petición
	url string  
}

func NewMonitor() *Monitor{
	// Inicializa la URL
	monitor := Monitor{url: "https://www.ign.es/web/ign/portal/sis-catalogo-terremotos?p_p_id=IGNSISCatalogoTerremotos_WAR_IGNSISCatalogoTerremotosportlet&p_p_lifecycle=2&p_p_state=normal&p_p_mode=view&p_p_cacheability=cacheLevelPage&p_p_col_id=column-1&p_p_col_count=1&_IGNSISCatalogoTerremotos_WAR_IGNSISCatalogoTerremotosportlet_jspPage=%2Fjsp%2Fterremoto.jsp"}

	return &monitor
}

// Consultar de terremotos, devuelve una lista de terremotos
func consultarTerremoto(fechaInicio time.Date, fechaFinal time.Date, latitudMinima float, latitudMaxima float, longitudMinima float, longitudMaxima float, profundidadMinima int, profundidadMaxima int, magnitudMinima float, magnitudMaxima float) []terremoto.Terremoto{

}