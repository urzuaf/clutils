package geo

// Region representa una de las 16 regiones oficiales de Chile según SUBDERE e INE.
type Region struct {
	Code         string `json:"code"`         // Código numérico oficial INE de 2 dígitos (ej: "13", "16", "05")
	IsoCode      string `json:"iso_code"`     // Código estándar internacional ISO 3166-2:CL (ej: "CL-RM", "CL-NB")
	Name         string `json:"name"`         // Nombre oficial completo de la región
	Abbreviation string `json:"abbreviation"` // Abreviatura oficial (ej: "RM", "NUBLE", "VALPO")
	Capital      string `json:"capital"`      // Ciudad capital regional
	Order        int    `json:"order"`        // Posición geográfica canónica de Norte a Sur (1 a 16)
}

// regiones contiene el listado maestro inmutable de las 16 regiones oficiales de Chile ordenadas de Norte a Sur.
var regiones = []Region{
	{
		Code:         "15",
		IsoCode:      "CL-AP",
		Name:         "Arica y Parinacota",
		Abbreviation: "AyP",
		Capital:      "Arica",
		Order:        1,
	},
	{
		Code:         "01",
		IsoCode:      "CL-TA",
		Name:         "Tarapacá",
		Abbreviation: "TPCA",
		Capital:      "Iquique",
		Order:        2,
	},
	{
		Code:         "02",
		IsoCode:      "CL-AN",
		Name:         "Antofagasta",
		Abbreviation: "ANTOF",
		Capital:      "Antofagasta",
		Order:        3,
	},
	{
		Code:         "03",
		IsoCode:      "CL-AT",
		Name:         "Atacama",
		Abbreviation: "ATCMA",
		Capital:      "Copiapó",
		Order:        4,
	},
	{
		Code:         "04",
		IsoCode:      "CL-CO",
		Name:         "Coquimbo",
		Abbreviation: "COQ",
		Capital:      "La Serena",
		Order:        5,
	},
	{
		Code:         "05",
		IsoCode:      "CL-VS",
		Name:         "Valparaíso",
		Abbreviation: "VALPO",
		Capital:      "Valparaíso",
		Order:        6,
	},
	{
		Code:         "13",
		IsoCode:      "CL-RM",
		Name:         "Metropolitana de Santiago",
		Abbreviation: "RM",
		Capital:      "Santiago",
		Order:        7,
	},
	{
		Code:         "06",
		IsoCode:      "CL-LI",
		Name:         "Libertador General Bernardo O'Higgins",
		Abbreviation: "LGBO",
		Capital:      "Rancagua",
		Order:        8,
	},
	{
		Code:         "07",
		IsoCode:      "CL-ML",
		Name:         "Maule",
		Abbreviation: "MAULE",
		Capital:      "Talca",
		Order:        9,
	},
	{
		Code:         "16",
		IsoCode:      "CL-NB",
		Name:         "Ñuble",
		Abbreviation: "NUBLE",
		Capital:      "Chillán",
		Order:        10,
	},
	{
		Code:         "08",
		IsoCode:      "CL-BI",
		Name:         "Biobío",
		Abbreviation: "BBIO",
		Capital:      "Concepción",
		Order:        11,
	},
	{
		Code:         "09",
		IsoCode:      "CL-AR",
		Name:         "La Araucanía",
		Abbreviation: "ARAUC",
		Capital:      "Temuco",
		Order:        12,
	},
	{
		Code:         "14",
		IsoCode:      "CL-LR",
		Name:         "Los Ríos",
		Abbreviation: "RIOS",
		Capital:      "Valdivia",
		Order:        13,
	},
	{
		Code:         "10",
		IsoCode:      "CL-LL",
		Name:         "Los Lagos",
		Abbreviation: "LAGOS",
		Capital:      "Puerto Montt",
		Order:        14,
	},
	{
		Code:         "11",
		IsoCode:      "CL-AI",
		Name:         "Aysén del General Carlos Ibáñez del Campo",
		Abbreviation: "AYSEN",
		Capital:      "Coyhaique",
		Order:        15,
	},
	{
		Code:         "12",
		IsoCode:      "CL-MA",
		Name:         "Magallanes y de la Antártica Chilena",
		Abbreviation: "MAG",
		Capital:      "Punta Arenas",
		Order:        16,
	},
}

// regionCodeToRoman mapea el código oficial de región a sus números romanos históricos.
var regionCodeToRoman = map[string]string{
	"01": "I",
	"02": "II",
	"03": "III",
	"04": "IV",
	"05": "V",
	"06": "VI",
	"07": "VII",
	"08": "VIII",
	"09": "IX",
	"10": "X",
	"11": "XI",
	"12": "XII",
	"13": "RM",
	"14": "XIV",
	"15": "XV",
	"16": "XVI",
}

// GetRegiones retorna una copia de las 16 regiones oficiales de Chile ordenadas geográficamente de Norte a Sur.
func GetRegiones() []Region {
	result := make([]Region, len(regiones))
	copy(result, regiones)
	return result
}

// FindRegion busca una región por su código INE (ej: "13", "16", "1"), código ISO 3166-2 (ej: "CL-RM", "cl-nb"),
// abreviatura oficial (ej: "RM", "NUBLE"), número romano ("VIII", "XVI") o nombre (tolerante a tildes, eñes y mayúsculas).
// Retorna un puntero a una copia de Region y true si fue encontrada, o nil y false si no existe coincidencia.
func FindRegion(query string) (*Region, bool) {
	return searchEntity(
		query,
		regiones,
		func(r Region) []string {
			idents := []string{r.Code, r.IsoCode, r.Abbreviation}
			if roman, ok := regionCodeToRoman[r.Code]; ok {
				idents = append(idents, roman)
			}
			return idents
		},
		func(r Region) string {
			return r.Name
		},
	)
}
