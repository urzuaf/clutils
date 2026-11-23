package geo

// Provincia representa una de las 56 provincias oficiales de Chile según SUBDERE e INE.
type Provincia struct {
	Code       string `json:"code"`        // Código numérico oficial INE de 3 dígitos (ej: "131", "161")
	Name       string `json:"name"`        // Nombre oficial de la provincia
	RegionCode string `json:"region_code"` // Código numérico oficial INE de la región a la que pertenece (ej: "13", "16")
}

// provincias contiene el listado maestro inmutable de las 56 provincias oficiales de Chile.
var provincias = []Provincia{
	{
		Code:       "011",
		Name:       "Iquique",
		RegionCode: "01",
	},
	{
		Code:       "014",
		Name:       "Tamarugal",
		RegionCode: "01",
	},
	{
		Code:       "021",
		Name:       "Antofagasta",
		RegionCode: "02",
	},
	{
		Code:       "022",
		Name:       "El Loa",
		RegionCode: "02",
	},
	{
		Code:       "023",
		Name:       "Tocopilla",
		RegionCode: "02",
	},
	{
		Code:       "031",
		Name:       "Copiapó",
		RegionCode: "03",
	},
	{
		Code:       "032",
		Name:       "Chañaral",
		RegionCode: "03",
	},
	{
		Code:       "033",
		Name:       "Huasco",
		RegionCode: "03",
	},
	{
		Code:       "041",
		Name:       "Elqui",
		RegionCode: "04",
	},
	{
		Code:       "042",
		Name:       "Choapa",
		RegionCode: "04",
	},
	{
		Code:       "043",
		Name:       "Limarí",
		RegionCode: "04",
	},
	{
		Code:       "051",
		Name:       "Valparaíso",
		RegionCode: "05",
	},
	{
		Code:       "052",
		Name:       "Isla de Pascua",
		RegionCode: "05",
	},
	{
		Code:       "053",
		Name:       "Los Andes",
		RegionCode: "05",
	},
	{
		Code:       "054",
		Name:       "Petorca",
		RegionCode: "05",
	},
	{
		Code:       "055",
		Name:       "Quillota",
		RegionCode: "05",
	},
	{
		Code:       "056",
		Name:       "San Antonio",
		RegionCode: "05",
	},
	{
		Code:       "057",
		Name:       "San Felipe de Aconcagua",
		RegionCode: "05",
	},
	{
		Code:       "058",
		Name:       "Marga Marga",
		RegionCode: "05",
	},
	{
		Code:       "061",
		Name:       "Cachapoal",
		RegionCode: "06",
	},
	{
		Code:       "062",
		Name:       "Cardenal Caro",
		RegionCode: "06",
	},
	{
		Code:       "063",
		Name:       "Colchagua",
		RegionCode: "06",
	},
	{
		Code:       "071",
		Name:       "Talca",
		RegionCode: "07",
	},
	{
		Code:       "072",
		Name:       "Cauquenes",
		RegionCode: "07",
	},
	{
		Code:       "073",
		Name:       "Curicó",
		RegionCode: "07",
	},
	{
		Code:       "074",
		Name:       "Linares",
		RegionCode: "07",
	},
	{
		Code:       "081",
		Name:       "Concepción",
		RegionCode: "08",
	},
	{
		Code:       "082",
		Name:       "Arauco",
		RegionCode: "08",
	},
	{
		Code:       "083",
		Name:       "Biobío",
		RegionCode: "08",
	},
	{
		Code:       "091",
		Name:       "Cautín",
		RegionCode: "09",
	},
	{
		Code:       "092",
		Name:       "Malleco",
		RegionCode: "09",
	},
	{
		Code:       "101",
		Name:       "Llanquihue",
		RegionCode: "10",
	},
	{
		Code:       "102",
		Name:       "Chiloé",
		RegionCode: "10",
	},
	{
		Code:       "103",
		Name:       "Osorno",
		RegionCode: "10",
	},
	{
		Code:       "104",
		Name:       "Palena",
		RegionCode: "10",
	},
	{
		Code:       "111",
		Name:       "Coyhaique",
		RegionCode: "11",
	},
	{
		Code:       "112",
		Name:       "Aysén",
		RegionCode: "11",
	},
	{
		Code:       "113",
		Name:       "Capitán Prat",
		RegionCode: "11",
	},
	{
		Code:       "114",
		Name:       "General Carrera",
		RegionCode: "11",
	},
	{
		Code:       "121",
		Name:       "Magallanes",
		RegionCode: "12",
	},
	{
		Code:       "122",
		Name:       "Antártica Chilena",
		RegionCode: "12",
	},
	{
		Code:       "123",
		Name:       "Tierra del Fuego",
		RegionCode: "12",
	},
	{
		Code:       "124",
		Name:       "Última Esperanza",
		RegionCode: "12",
	},
	{
		Code:       "131",
		Name:       "Santiago",
		RegionCode: "13",
	},
	{
		Code:       "132",
		Name:       "Cordillera",
		RegionCode: "13",
	},
	{
		Code:       "133",
		Name:       "Chacabuco",
		RegionCode: "13",
	},
	{
		Code:       "134",
		Name:       "Maipo",
		RegionCode: "13",
	},
	{
		Code:       "135",
		Name:       "Melipilla",
		RegionCode: "13",
	},
	{
		Code:       "136",
		Name:       "Talagante",
		RegionCode: "13",
	},
	{
		Code:       "141",
		Name:       "Valdivia",
		RegionCode: "14",
	},
	{
		Code:       "142",
		Name:       "Ranco",
		RegionCode: "14",
	},
	{
		Code:       "151",
		Name:       "Arica",
		RegionCode: "15",
	},
	{
		Code:       "152",
		Name:       "Parinacota",
		RegionCode: "15",
	},
	{
		Code:       "161",
		Name:       "Diguillín",
		RegionCode: "16",
	},
	{
		Code:       "162",
		Name:       "Itata",
		RegionCode: "16",
	},
	{
		Code:       "163",
		Name:       "Punilla",
		RegionCode: "16",
	},
}

// GetProvincias retorna una copia de las 56 provincias oficiales de Chile.
func GetProvincias() []Provincia {
	result := make([]Provincia, len(provincias))
	copy(result, provincias)
	return result
}

// GetProvinciasByRegion retorna las provincias pertenecientes a una región específica.
// Acepta código de región (ej: "13", "16"), código ISO (ej: "CL-RM", "CL-NB") o nombre de región.
// Si la región no existe o no tiene provincias, retorna un slice vacío.
func GetProvinciasByRegion(regionQuery string) []Provincia {
	reg, ok := FindRegion(regionQuery)
	if !ok || reg == nil {
		return []Provincia{}
	}

	var result []Provincia
	for i := range provincias {
		if provincias[i].RegionCode == reg.Code {
			result = append(result, provincias[i])
		}
	}
	return result
}

// FindProvincia busca una provincia por su código INE de 3 dígitos (ej: "131", "161")
// o por su nombre oficial (tolerante a mayúsculas, tildes y eñes).
// Retorna un puntero a una copia de Provincia y true si fue encontrada, o nil y false si no existe coincidencia.
func FindProvincia(query string) (*Provincia, bool) {
	return searchEntity(
		query,
		provincias,
		func(p Provincia) []string {
			return []string{p.Code}
		},
		func(p Provincia) string {
			return p.Name
		},
	)
}

// SearchProvincias busca y retorna todas las provincias que coincidan con el término de búsqueda,
// ordenadas por relevancia (coincidencias exactas primero, luego prefijos y finalmente subcadenas).
// Ideal para cajas de autocompletado y filtros de búsqueda interactivos en UI.
func SearchProvincias(query string) []Provincia {
	return searchAllEntities(
		query,
		provincias,
		func(p Provincia) []string {
			return []string{p.Code}
		},
		func(p Provincia) string {
			return p.Name
		},
	)
}

// GetRegionByProvincia resuelve la región oficial a la que pertenece una provincia dada.
// Acepta código de provincia (ej: "131", "161") o nombre de provincia (ej: "diguillin", "santiago").
// Retorna un puntero a Region y true si fue resuelta, o nil y false si la provincia no existe.
func GetRegionByProvincia(provinciaQuery string) (*Region, bool) {
	prov, ok := FindProvincia(provinciaQuery)
	if !ok || prov == nil {
		return nil, false
	}
	return FindRegion(prov.RegionCode)
}
