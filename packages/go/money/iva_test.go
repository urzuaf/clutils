package money_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/urzuaf/clutils/packages/go/money"
)

type IVACalculationCase struct {
	InputAmount   int64  `json:"input_amount"`
	IsBruto       bool   `json:"is_bruto"`
	ExpectedNeto  int64  `json:"expected_neto"`
	ExpectedIVA   int64  `json:"expected_iva"`
	ExpectedBruto int64  `json:"expected_bruto"`
	Description   string `json:"description"`
}

type IVATestSuite struct {
	IVACalculation []IVACalculationCase `json:"iva_calculation"`
}

func loadIVATestSuite(t *testing.T) IVATestSuite {
	t.Helper()

	path := filepath.Join("..", "..", "..", "specs", "tests", "money_formatting.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("error leyendo specs/tests/money_formatting.json: %v", err)
	}

	var suite IVATestSuite
	if err := json.Unmarshal(data, &suite); err != nil {
		t.Fatalf("error parseando JSON de pruebas IVA: %v", err)
	}

	return suite
}

func TestCalculateIVA(t *testing.T) {
	suite := loadIVATestSuite(t)

	for _, tc := range suite.IVACalculation {
		t.Run(tc.Description, func(t *testing.T) {
			got := money.CalculateIVA(tc.InputAmount, tc.IsBruto)

			if got.Neto != tc.ExpectedNeto {
				t.Errorf("CalculateIVA(%d, isBruto: %v).Neto = %d; esperaba %d",
					tc.InputAmount, tc.IsBruto, got.Neto, tc.ExpectedNeto)
			}

			if got.IVA != tc.ExpectedIVA {
				t.Errorf("CalculateIVA(%d, isBruto: %v).IVA = %d; esperaba %d",
					tc.InputAmount, tc.IsBruto, got.IVA, tc.ExpectedIVA)
			}

			if got.Bruto != tc.ExpectedBruto {
				t.Errorf("CalculateIVA(%d, isBruto: %v).Bruto = %d; esperaba %d",
					tc.InputAmount, tc.IsBruto, got.Bruto, tc.ExpectedBruto)
			}

			// Invariante sagrado del SII: Neto + IVA SIEMPRE debe ser exactamente igual a Bruto
			if got.Neto+got.IVA != got.Bruto {
				t.Errorf("Fallo de cuadratura SII: %d + %d != %d", got.Neto, got.IVA, got.Bruto)
			}
		})
	}
}

func TestCalculateIVACustomRate(t *testing.T) {
	// Prueba con tasa personalizada (ej. 10%)
	got := money.CalculateIVA(1000, false, money.IVAOptions{Rate: 0.10})
	if got.Neto != 1000 || got.IVA != 100 || got.Bruto != 1100 {
		t.Errorf("CalculateIVA con 10%% fallo: %+v", got)
	}

	gotBruto := money.CalculateIVA(1100, true, money.IVAOptions{Rate: 0.10})
	if gotBruto.Neto != 1000 || gotBruto.IVA != 100 || gotBruto.Bruto != 1100 {
		t.Errorf("CalculateIVA con 10%% Bruto fallo: %+v", gotBruto)
	}
}
