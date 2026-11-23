package geo_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/urzuaf/clutils/packages/go/geo"
)

type findComunaTestCase struct {
	Query              string  `json:"query"`
	ExpectedCut        *string `json:"expected_cut"`
	ExpectedName       *string `json:"expected_name"`
	ExpectedRegionCode *string `json:"expected_region_code"`
	Description        string  `json:"description"`
}

type getRegionByComunaTestCase struct {
	ComunaQuery        string  `json:"comuna_query"`
	ExpectedRegionCode *string `json:"expected_region_code"`
	ExpectedRegionName *string `json:"expected_region_name"`
	Description        string  `json:"description"`
}

type fullGeoTestSuite struct {
	Counts struct {
		ExpectedRegions   int `json:"expected_regions"`
		ExpectedProvinces int `json:"expected_provinces"`
		ExpectedCommunes  int `json:"expected_communes"`
	} `json:"counts"`
	FindComuna         []findComunaTestCase        `json:"find_comuna"`
	GetRegionByComuna  []getRegionByComunaTestCase `json:"get_region_by_comuna"`
}

func loadFullGeoTestSuite(t *testing.T) fullGeoTestSuite {
	t.Helper()

	path := filepath.Join("..", "..", "..", "specs", "tests", "geo_dpa.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("error leyendo specs/tests/geo_dpa.json: %v", err)
	}

	var suite fullGeoTestSuite
	if err := json.Unmarshal(data, &suite); err != nil {
		t.Fatalf("error deserializando JSON de pruebas: %v", err)
	}

	return suite
}

func TestGetComunas(t *testing.T) {
	suite := loadFullGeoTestSuite(t)

	comunas := geo.GetComunas()
	if len(comunas) != suite.Counts.ExpectedCommunes {
		t.Fatalf("GetComunas() retornó %d comunas; esperaba %d", len(comunas), suite.Counts.ExpectedCommunes)
	}

	// Inmutabilidad
	first := geo.GetComunas()
	first[0].Name = "Modificado"
	second := geo.GetComunas()
	if second[0].Name == "Modificado" {
		t.Error("GetComunas() no protegió la inmutabilidad interna")
	}
}

func TestGetComunasByRegion(t *testing.T) {
	// 1. Ñuble (16) -> 21 comunas
	nubleComunas := geo.GetComunasByRegion("16")
	if len(nubleComunas) != 21 {
		t.Fatalf("GetComunasByRegion(16) retornó %d comunas; esperaba 21", len(nubleComunas))
	}

	// 2. Ñuble por código ISO "CL-NB" -> 21 comunas
	nubleComunasIso := geo.GetComunasByRegion("CL-NB")
	if len(nubleComunasIso) != 21 {
		t.Errorf("GetComunasByRegion(CL-NB) retornó %d comunas; esperaba 21", len(nubleComunasIso))
	}

	// 3. Región Metropolitana (13) -> 52 comunas
	rmComunas := geo.GetComunasByRegion("13")
	if len(rmComunas) != 52 {
		t.Fatalf("GetComunasByRegion(13) retornó %d comunas; esperaba 52", len(rmComunas))
	}

	// 4. Arica y Parinacota (15) -> 4 comunas
	aypComunas := geo.GetComunasByRegion("15")
	if len(aypComunas) != 4 {
		t.Errorf("GetComunasByRegion(15) retornó %d comunas; esperaba 4", len(aypComunas))
	}

	// 5. Región inválida "99" -> []
	invalidComunas := geo.GetComunasByRegion("99")
	if len(invalidComunas) != 0 {
		t.Errorf("GetComunasByRegion(99) retornó %d comunas; esperaba 0", len(invalidComunas))
	}
}

func TestGetComunasByProvincia(t *testing.T) {
	// Provincia de Santiago (131) -> 32 comunas
	santiagoComunas := geo.GetComunasByProvincia("131")
	if len(santiagoComunas) != 32 {
		t.Fatalf("GetComunasByProvincia(131) retornó %d comunas; esperaba 32", len(santiagoComunas))
	}

	// Por nombre "diguillin" -> 9 comunas
	diguillinComunas := geo.GetComunasByProvincia("diguillin")
	if len(diguillinComunas) != 9 {
		t.Fatalf("GetComunasByProvincia(diguillin) retornó %d comunas; esperaba 9", len(diguillinComunas))
	}

	// Provincia inexistente "999" -> []
	invalid := geo.GetComunasByProvincia("999")
	if len(invalid) != 0 {
		t.Errorf("GetComunasByProvincia(999) retornó %d comunas; esperaba 0", len(invalid))
	}
}

func TestFindComuna(t *testing.T) {
	suite := loadFullGeoTestSuite(t)

	for _, tc := range suite.FindComuna {
		t.Run(tc.Description, func(t *testing.T) {
			comuna, ok := geo.FindComuna(tc.Query)
			if tc.ExpectedCut != nil {
				if !ok || comuna == nil {
					t.Fatalf("FindComuna(%q) no encontró la comuna esperada", tc.Query)
				}
				if comuna.Cut != *tc.ExpectedCut {
					t.Errorf("FindComuna(%q) Cut = %q; esperaba %q", tc.Query, comuna.Cut, *tc.ExpectedCut)
				}
				if tc.ExpectedName != nil && comuna.Name != *tc.ExpectedName {
					t.Errorf("FindComuna(%q) Name = %q; esperaba %q", tc.Query, comuna.Name, *tc.ExpectedName)
				}
				if tc.ExpectedRegionCode != nil && comuna.RegionCode != *tc.ExpectedRegionCode {
					t.Errorf("FindComuna(%q) RegionCode = %q; esperaba %q", tc.Query, comuna.RegionCode, *tc.ExpectedRegionCode)
				}
			} else {
				if ok || comuna != nil {
					t.Errorf("FindComuna(%q) retornó %+v; esperaba nil/false", tc.Query, comuna)
				}
			}
		})
	}
}

func TestSearchComunas(t *testing.T) {
	// Búsqueda "San" -> más de 30 comunas (San Bernardo, San Miguel, San Joaquín, etc.)
	sanResults := geo.SearchComunas("San")
	if len(sanResults) < 30 {
		t.Errorf("SearchComunas(San) retornó %d comunas; esperaba más de 30", len(sanResults))
	}

	// Búsqueda "vina" -> Viña del Mar debe estar incluida
	vinaResults := geo.SearchComunas("vina")
	if len(vinaResults) == 0 || vinaResults[0].Name != "Viña del Mar" {
		t.Errorf("SearchComunas(vina) = %+v; esperaba Viña del Mar como primer resultado", vinaResults)
	}

	// Búsqueda por CUT exacto "13120" -> 1 resultado Ñuñoa
	nunoaResults := geo.SearchComunas("13120")
	if len(nunoaResults) != 1 || nunoaResults[0].Name != "Ñuñoa" {
		t.Errorf("SearchComunas(13120) = %+v; esperaba [Ñuñoa]", nunoaResults)
	}

	// Búsqueda vacía -> 0
	emptyResults := geo.SearchComunas("   ")
	if len(emptyResults) != 0 {
		t.Errorf("SearchComunas('   ') retornó %d resultados; esperaba 0", len(emptyResults))
	}

	// Búsqueda sin coincidencias
	noneResults := geo.SearchComunas("comuna_inexistente_xyz")
	if len(noneResults) != 0 {
		t.Errorf("SearchComunas(inexistente) retornó %d resultados; esperaba 0", len(noneResults))
	}
}

func TestGetRegionByComuna(t *testing.T) {
	suite := loadFullGeoTestSuite(t)

	for _, tc := range suite.GetRegionByComuna {
		t.Run(tc.Description, func(t *testing.T) {
			reg, ok := geo.GetRegionByComuna(tc.ComunaQuery)
			if tc.ExpectedRegionCode != nil {
				if !ok || reg == nil {
					t.Fatalf("GetRegionByComuna(%q) no resolvió la región", tc.ComunaQuery)
				}
				if reg.Code != *tc.ExpectedRegionCode {
					t.Errorf("GetRegionByComuna(%q) Region Code = %q; esperaba %q", tc.ComunaQuery, reg.Code, *tc.ExpectedRegionCode)
				}
				if tc.ExpectedRegionName != nil && reg.Name != *tc.ExpectedRegionName {
					t.Errorf("GetRegionByComuna(%q) Region Name = %q; esperaba %q", tc.ComunaQuery, reg.Name, *tc.ExpectedRegionName)
				}
			} else {
				if ok || reg != nil {
					t.Errorf("GetRegionByComuna(%q) retornó %+v; esperaba nil/false", tc.ComunaQuery, reg)
				}
			}
		})
	}
}

func TestGetProvinciaByComuna(t *testing.T) {
	// Ñuñoa (13120) -> Provincia de Santiago (131)
	prov, ok := geo.GetProvinciaByComuna("Ñuñoa")
	if !ok || prov == nil || prov.Code != "131" || prov.Name != "Santiago" {
		t.Errorf("GetProvinciaByComuna(Ñuñoa) = %+v, %v; esperaba Provincia Santiago 131", prov, ok)
	}

	// Chillán (16101) -> Provincia de Diguillín (161)
	prov, ok = geo.GetProvinciaByComuna("16101")
	if !ok || prov == nil || prov.Code != "161" || prov.Name != "Diguillín" {
		t.Errorf("GetProvinciaByComuna(16101) = %+v, %v; esperaba Provincia Diguillín 161", prov, ok)
	}

	// Comuna inexistente
	prov, ok = geo.GetProvinciaByComuna("inexistente")
	if ok || prov != nil {
		t.Errorf("GetProvinciaByComuna(inexistente) = %+v, %v; esperaba nil, false", prov, ok)
	}
}
