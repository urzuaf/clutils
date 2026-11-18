package money_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/urzuaf/clutils/packages/go/money"
)

type UFFormattingCase struct {
	Amount            float64 `json:"amount"`
	Decimals          int     `json:"decimals"`
	ExpectedDefault   string  `json:"expected_default"`
	ExpectedNoSymbol  string  `json:"expected_no_symbol"`
	Description       string  `json:"description"`
}

type UFTestSuite struct {
	UFFormatting []UFFormattingCase `json:"uf_formatting"`
}

func loadUFTestSuite(t *testing.T) UFTestSuite {
	t.Helper()

	path := filepath.Join("..", "..", "..", "specs", "tests", "money_formatting.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("error leyendo specs/tests/money_formatting.json: %v", err)
	}

	var suite UFTestSuite
	if err := json.Unmarshal(data, &suite); err != nil {
		t.Fatalf("error parseando JSON de pruebas UF: %v", err)
	}

	return suite
}

func TestFormatUF(t *testing.T) {
	suite := loadUFTestSuite(t)

	for _, tc := range suite.UFFormatting {
		t.Run(tc.Description+" (default)", func(t *testing.T) {
			got := money.FormatUF(tc.Amount, money.UFFormatOptions{
				Decimals: tc.Decimals,
			})
			if got != tc.ExpectedDefault {
				t.Errorf("FormatUF(%f, Decimals: %d) = %q; esperaba %q", tc.Amount, tc.Decimals, got, tc.ExpectedDefault)
			}
		})

		t.Run(tc.Description+" (no symbol)", func(t *testing.T) {
			got := money.FormatUF(tc.Amount, money.UFFormatOptions{
				Decimals: tc.Decimals,
				NoSymbol: true,
			})
			if got != tc.ExpectedNoSymbol {
				t.Errorf("FormatUF(%f, NoSymbol: true) = %q; esperaba %q", tc.Amount, got, tc.ExpectedNoSymbol)
			}
		})
	}
}
