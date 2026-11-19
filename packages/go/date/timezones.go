// Package date proporciona utilidades canonicas para la gestion de fecha y hora
// segun la normativa y husos horarios oficiales de Chile (Continental, Magallanes e Insular).
package date

import (
	"sync"
	"time"
	_ "time/tzdata" // Garantiza base IANA embebida incluso en contenedores scratch o entornos minimos
)

const (
	// TzSantiago es el identificador IANA oficial para Chile Continental ("America/Santiago").
	TzSantiago = "America/Santiago"

	// TzPuntaArenas es el identificador IANA oficial para la Region de Magallanes y de la Antartica Chilena ("America/Punta_Arenas").
	TzPuntaArenas = "America/Punta_Arenas"

	// TzEaster es el identificador IANA oficial para Chile Insular Occidental / Rapa Nui ("Pacific/Easter").
	TzEaster = "Pacific/Easter"
)

var (
	locSantiago       *time.Location
	locPuntaArenas    *time.Location
	locEaster         *time.Location
	initLocationsOnce sync.Once
)

func initLocations() {
	initLocationsOnce.Do(func() {
		var err error
		locSantiago, err = time.LoadLocation(TzSantiago)
		if err != nil {
			locSantiago = time.FixedZone("CLT", -4*3600)
		}

		locPuntaArenas, err = time.LoadLocation(TzPuntaArenas)
		if err != nil {
			locPuntaArenas = time.FixedZone("CLST", -3*3600)
		}

		locEaster, err = time.LoadLocation(TzEaster)
		if err != nil {
			locEaster = time.FixedZone("EAST", -6*3600)
		}
	})
}

// LocationSantiago retorna el *time.Location oficial de Chile Continental.
func LocationSantiago() *time.Location {
	initLocations()
	return locSantiago
}

// LocationPuntaArenas retorna el *time.Location oficial de la Region de Magallanes (UTC-3 permanente).
func LocationPuntaArenas() *time.Location {
	initLocations()
	return locPuntaArenas
}

// LocationEaster retorna el *time.Location oficial de Rapa Nui / Isla de Pascua.
func LocationEaster() *time.Location {
	initLocations()
	return locEaster
}

// NowChile retorna la fecha y hora actual en la zona de Chile Continental ("America/Santiago"),
// aplicando automaticamente la compensacion por horario de invierno (UTC-4) o verano (UTC-3).
func NowChile() time.Time {
	return time.Now().In(LocationSantiago())
}

// NowMagallanes retorna la fecha y hora actual en la Region de Magallanes y de la Antartica Chilena,
// la cual se mantiene permanentemente en UTC-3 por decreto supremo.
func NowMagallanes() time.Time {
	return time.Now().In(LocationPuntaArenas())
}

// NowEasterIsland retorna la fecha y hora actual en Rapa Nui ("Pacific/Easter").
func NowEasterIsland() time.Time {
	return time.Now().In(LocationEaster())
}

// ToChileTime convierte cualquier time.Time a la zona oficial de Chile Continental ("America/Santiago").
func ToChileTime(t time.Time) time.Time {
	return t.In(LocationSantiago())
}

// ToMagallanesTime convierte cualquier time.Time a la zona oficial de Magallanes ("America/Punta_Arenas").
func ToMagallanesTime(t time.Time) time.Time {
	return t.In(LocationPuntaArenas())
}

// ToEasterIslandTime convierte cualquier time.Time a la zona oficial de Rapa Nui ("Pacific/Easter").
func ToEasterIslandTime(t time.Time) time.Time {
	return t.In(LocationEaster())
}
