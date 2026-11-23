package geo_test

import (
	"testing"

	"github.com/urzuaf/clutils/packages/go/geo"
)

func TestGetProvincias(t *testing.T) {
	provincias := geo.GetProvincias()
	if len(provincias) != 56 {
		t.Fatalf("GetProvincias() retornó %d provincias; esperaba exactamente 56", len(provincias))
	}

	// Inmutabilidad
	firstCall := geo.GetProvincias()
	firstCall[0].Name = "Modificado"
	secondCall := geo.GetProvincias()
	if secondCall[0].Name == "Modificado" {
		t.Error("GetProvincias() no protegió la inmutabilidad interna")
	}
}

func TestGetProvinciasByRegion(t *testing.T) {
	// 1. Ñuble por código INE "16" -> 3 provincias
	provsNuble := geo.GetProvinciasByRegion("16")
	if len(provsNuble) != 3 {
		t.Fatalf("GetProvinciasByRegion(16) retornó %d provincias; esperaba 3", len(provsNuble))
	}
	nubleNames := map[string]bool{}
	for _, p := range provsNuble {
		nubleNames[p.Name] = true
	}
	for _, expected := range []string{"Diguillín", "Itata", "Punilla"} {
		if !nubleNames[expected] {
			t.Errorf("GetProvinciasByRegion(16) no contiene provincia %q", expected)
		}
	}

	// 2. Ñuble por código ISO "CL-NB"
	provsNubleIso := geo.GetProvinciasByRegion("CL-NB")
	if len(provsNubleIso) != 3 {
		t.Errorf("GetProvinciasByRegion(CL-NB) retornó %d provincias; esperaba 3", len(provsNubleIso))
	}

	// 3. Región Metropolitana "13" -> 6 provincias
	provsRM := geo.GetProvinciasByRegion("13")
	if len(provsRM) != 6 {
		t.Fatalf("GetProvinciasByRegion(13) retornó %d provincias; esperaba 6", len(provsRM))
	}

	// 4. Región inexistente "99" -> []
	provsInvalid := geo.GetProvinciasByRegion("99")
	if len(provsInvalid) != 0 {
		t.Errorf("GetProvinciasByRegion(99) retornó %d provincias; esperaba 0", len(provsInvalid))
	}
}

func TestFindProvincia(t *testing.T) {
	tests := []struct {
		query        string
		expectedCode string
		expectedName string
		valid        bool
	}{
		{"131", "131", "Santiago", true},
		{"161", "161", "Diguillín", true},
		{"Santiago", "131", "Santiago", true},
		{"diguillin", "161", "Diguillín", true},
		{"DIGUILLIN", "161", "Diguillín", true},
		{"isla de pascua", "052", "Isla de Pascua", true},
		{"marga marga", "058", "Marga Marga", true},
		{"cordillera", "132", "Cordillera", true},
		{"provincia_inexistente", "", "", false},
		{"", "", "", false},
	}

	for _, tc := range tests {
		t.Run(tc.query, func(t *testing.T) {
			p, ok := geo.FindProvincia(tc.query)
			if tc.valid {
				if !ok || p == nil {
					t.Fatalf("FindProvincia(%q) no encontró la provincia", tc.query)
				}
				if p.Code != tc.expectedCode {
					t.Errorf("FindProvincia(%q) Code = %q; esperaba %q", tc.query, p.Code, tc.expectedCode)
				}
				if p.Name != tc.expectedName {
					t.Errorf("FindProvincia(%q) Name = %q; esperaba %q", tc.query, p.Name, tc.expectedName)
				}
			} else {
				if ok || p != nil {
					t.Errorf("FindProvincia(%q) retornó %+v; esperaba nil/false", tc.query, p)
				}
			}
		})
	}
}

func TestGetRegionByProvincia(t *testing.T) {
	// Diguillín -> Ñuble (16)
	reg, ok := geo.GetRegionByProvincia("Diguillín")
	if !ok || reg == nil || reg.Code != "16" {
		t.Errorf("GetRegionByProvincia(Diguillín) = %+v, %v; esperaba Región 16", reg, ok)
	}

	// Por código de provincia "131" (Santiago) -> RM (13)
	reg, ok = geo.GetRegionByProvincia("131")
	if !ok || reg == nil || reg.Code != "13" {
		t.Errorf("GetRegionByProvincia(131) = %+v, %v; esperaba Región 13", reg, ok)
	}

	// Sin tilde "diguillin"
	reg, ok = geo.GetRegionByProvincia("diguillin")
	if !ok || reg == nil || reg.Code != "16" {
		t.Errorf("GetRegionByProvincia(diguillin) = %+v, %v; esperaba Región 16", reg, ok)
	}

	// Provincia inexistente
	reg, ok = geo.GetRegionByProvincia("inexistente")
	if ok || reg != nil {
		t.Errorf("GetRegionByProvincia(inexistente) = %+v, %v; esperaba nil, false", reg, ok)
	}
}

func TestSearchProvincias(t *testing.T) {
	// Búsqueda "San" -> San Antonio, San Felipe de Aconcagua, Santiago (3 provincias)
	results := geo.SearchProvincias("San")
	if len(results) < 3 {
		t.Errorf("SearchProvincias(San) retornó %d provincias; esperaba al menos 3", len(results))
	}

	// Búsqueda "Val" -> Valparaíso, Valdivia (2 provincias)
	resultsVal := geo.SearchProvincias("Val")
	if len(resultsVal) != 2 {
		t.Errorf("SearchProvincias(Val) retornó %d provincias; esperaba 2 (Valparaíso y Valdivia)", len(resultsVal))
	}

	// Búsqueda sin tilde "diguillin" -> 1 resultado Diguillín
	resultsDig := geo.SearchProvincias("diguillin")
	if len(resultsDig) != 1 || resultsDig[0].Name != "Diguillín" {
		t.Errorf("SearchProvincias(diguillin) = %+v; esperaba [Diguillín]", resultsDig)
	}

	// Búsqueda sin coincidencias
	resultsNone := geo.SearchProvincias("inexistente")
	if len(resultsNone) != 0 {
		t.Errorf("SearchProvincias(inexistente) retornó %d resultados; esperaba 0", len(resultsNone))
	}

	// Búsqueda vacía
	resultsEmpty := geo.SearchProvincias("   ")
	if len(resultsEmpty) != 0 {
		t.Errorf("SearchProvincias('   ') retornó %d resultados; esperaba 0", len(resultsEmpty))
	}
}
