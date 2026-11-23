package money_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/urzuaf/clutils/packages/go/money"
)

type UTMFormattingCase struct {
	Amount           int64  `json:"amount"`
	ExpectedDefault  string `json:"expected_default"`
	ExpectedNoSymbol string `json:"expected_no_symbol"`
	Description      string `json:"description"`
}

type UTMTestSuite struct {
	UTMFormatting []UTMFormattingCase `json:"utm_formatting"`
}

func loadUTMTestSuite(t *testing.T) UTMTestSuite {
	t.Helper()

	path := filepath.Join("..", "..", "..", "specs", "tests", "money_formatting.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("error leyendo specs/tests/money_formatting.json: %v", err)
	}

	var suite UTMTestSuite
	if err := json.Unmarshal(data, &suite); err != nil {
		t.Fatalf("error parseando JSON de pruebas UTM: %v", err)
	}

	return suite
}

func TestFormatUTM(t *testing.T) {
	suite := loadUTMTestSuite(t)

	for _, tc := range suite.UTMFormatting {
		t.Run(tc.Description+" (default)", func(t *testing.T) {
			got := money.FormatUTM(tc.Amount)
			if got != tc.ExpectedDefault {
				t.Errorf("FormatUTM(%d) = %q; esperaba %q", tc.Amount, got, tc.ExpectedDefault)
			}
		})

		t.Run(tc.Description+" (no symbol)", func(t *testing.T) {
			got := money.FormatUTM(tc.Amount, money.UTMFormatOptions{NoSymbol: true})
			if got != tc.ExpectedNoSymbol {
				t.Errorf("FormatUTM(%d, NoSymbol: true) = %q; esperaba %q", tc.Amount, got, tc.ExpectedNoSymbol)
			}
		})
	}
}
