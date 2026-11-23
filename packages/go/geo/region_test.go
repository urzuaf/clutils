package geo_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/urzuaf/clutils/packages/go/geo"
)

type findRegionTestCase struct {
	Query        string  `json:"query"`
	ExpectedCode *string `json:"expected_code"`
	ExpectedIso  *string `json:"expected_iso"`
	ExpectedName *string `json:"expected_name"`
	Description  string  `json:"description"`
}

type geoTestSuite struct {
	Counts struct {
		ExpectedRegions int `json:"expected_regions"`
	} `json:"counts"`
	FindRegion []findRegionTestCase `json:"find_region"`
}

func loadGeoTestSuite(t *testing.T) geoTestSuite {
	t.Helper()

	path := filepath.Join("..", "..", "..", "specs", "tests", "geo_dpa.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("error leyendo specs/tests/geo_dpa.json: %v", err)
	}

	var suite geoTestSuite
	if err := json.Unmarshal(data, &suite); err != nil {
		t.Fatalf("error deserializando JSON de pruebas: %v", err)
	}

	return suite
}

func TestGetRegiones(t *testing.T) {
	suite := loadGeoTestSuite(t)

	regiones := geo.GetRegiones()

	if len(regiones) != suite.Counts.ExpectedRegions {
		t.Fatalf("GetRegiones() retornó %d regiones; esperaba %d", len(regiones), suite.Counts.ExpectedRegions)
	}

	// Verificar orden de Norte a Sur
	for i, reg := range regiones {
		expectedOrder := i + 1
		if reg.Order != expectedOrder {
			t.Errorf("Region %q tiene Order=%d; esperaba %d", reg.Name, reg.Order, expectedOrder)
		}
	}

	// Verificar primera región (Arica y Parinacota)
	if regiones[0].Code != "15" || regiones[0].IsoCode != "CL-AP" || regiones[0].Capital != "Arica" {
		t.Errorf("Primera región inesperada: %+v", regiones[0])
	}

	// Verificar Región Metropolitana
	if regiones[6].Code != "13" || regiones[6].IsoCode != "CL-RM" || regiones[6].Capital != "Santiago" {
		t.Errorf("Región Metropolitana inesperada: %+v", regiones[6])
	}

	// Verificar Región de Ñuble
	if regiones[9].Code != "16" || regiones[9].IsoCode != "CL-NB" || regiones[9].Capital != "Chillán" {
		t.Errorf("Región de Ñuble inesperada: %+v", regiones[9])
	}

	// Verificar última región (Magallanes)
	if regiones[15].Code != "12" || regiones[15].IsoCode != "CL-MA" || regiones[15].Capital != "Punta Arenas" {
		t.Errorf("Última región inesperada: %+v", regiones[15])
	}
}

func TestGetRegionesImmutability(t *testing.T) {
	firstCall := geo.GetRegiones()
	firstCall[0].Name = "Modificado"

	secondCall := geo.GetRegiones()
	if secondCall[0].Name == "Modificado" {
		t.Error("GetRegiones() no protegió la inmutabilidad interna; la modificación del slice afectó llamadas posteriores")
	}
}

func TestFindRegion(t *testing.T) {
	suite := loadGeoTestSuite(t)

	for _, tc := range suite.FindRegion {
		t.Run(tc.Description, func(t *testing.T) {
			reg, ok := geo.FindRegion(tc.Query)
			if tc.ExpectedCode != nil {
				if !ok || reg == nil {
					t.Fatalf("FindRegion(%q) no encontró la región esperada", tc.Query)
				}
				if reg.Code != *tc.ExpectedCode {
					t.Errorf("FindRegion(%q) Code = %q; esperaba %q", tc.Query, reg.Code, *tc.ExpectedCode)
				}
				if tc.ExpectedIso != nil && reg.IsoCode != *tc.ExpectedIso {
					t.Errorf("FindRegion(%q) IsoCode = %q; esperaba %q", tc.Query, reg.IsoCode, *tc.ExpectedIso)
				}
				if tc.ExpectedName != nil && reg.Name != *tc.ExpectedName {
					t.Errorf("FindRegion(%q) Name = %q; esperaba %q", tc.Query, reg.Name, *tc.ExpectedName)
				}
			} else {
				if ok || reg != nil {
					t.Errorf("FindRegion(%q) retornó %+v; esperaba nil/false", tc.Query, reg)
				}
			}
		})
	}
}

func TestFindRegionRomanAndPadded(t *testing.T) {
	// Probar número romano VIII -> Biobío (08)
	reg, ok := geo.FindRegion("VIII")
	if !ok || reg.Code != "08" {
		t.Errorf("FindRegion(VIII) = %+v, %v; esperaba Code 08", reg, ok)
	}

	// Probar dígito simple "1" -> Tarapacá (01)
	reg, ok = geo.FindRegion("1")
	if !ok || reg.Code != "01" {
		t.Errorf("FindRegion(1) = %+v, %v; esperaba Code 01", reg, ok)
	}

	// Cadena vacía
	reg, ok = geo.FindRegion("   ")
	if ok || reg != nil {
		t.Errorf("FindRegion('   ') = %+v, %v; esperaba nil, false", reg, ok)
	}
}
