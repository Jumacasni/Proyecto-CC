package monitor

import (
	"testing"
	"os"
	"time"
	"./terremoto"
)

func TestMain(m *testing.M){
	InitURL("https://ejemplo.com/post")
	exitVal := m.Run()

	os.Exit(exitVal)
}

/***** MOCK PARA CONVERTIR LOS DATOS DE RESPUESTA DEL HTTP POST EN TERREMOTOS *****/
var csvToTerremotosMock func(data string) []terremoto.Terremoto

type convertCsvToTerremotosMock struct{}

func (u convertCsvToTerremotosMock) csvToTerremotos(data string) []terremoto.Terremoto{
	return csvToTerremotosMock(data)
}

/***** TEST PARA CONSULTAR TERREMOTOS *****/
func TestConsultarTerremotos(t *testing.T){
	// Se define la función csvToTerremotoMock que es una función mock que devuelve un terremoto
	csvToTerremotos := convertCsvToTerremotosMock{}
	csvToTerremotosMock = func(data string) []terremoto.Terremoto{
		var terremotos []terremoto.Terremoto
		t := terremoto.NewTerremoto("o",time.Now(),0,0,0,0,"Md","0")
		terremotos = append(terremotos, *t)
		return terremotos
	}

	err, terremotos := consultarTerremoto(time.Now(), time.Now(), 0, 0, 0, 0, 0, 0, 0, 0, csvToTerremotos)

	if err != nil{
		t.Error(err)
	}

	if len(terremotos) != 1 {
		t.Error("No están todos los terremotos")
	}
}