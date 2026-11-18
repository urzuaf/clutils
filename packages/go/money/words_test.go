package money_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/urzuaf/clutils/packages/go/money"
)

type ToWordsCase struct {
	Amount             int64  `json:"amount"`
	Expected           string `json:"expected"`
	ExpectedNoCurrency string `json:"expected_no_currency"`
	Description        string `json:"description"`
}

type ToWordsTestSuite struct {
	ToWords []ToWordsCase `json:"to_words"`
}

func loadToWordsTestSuite(t *testing.T) ToWordsTestSuite {
	t.Helper()

	path := filepath.Join("..", "..", "..", "specs", "tests", "money_formatting.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("error leyendo specs/tests/money_formatting.json: %v", err)
	}

	var suite ToWordsTestSuite
	if err := json.Unmarshal(data, &suite); err != nil {
		t.Fatalf("error parseando JSON de pruebas toWords: %v", err)
	}

	return suite
}

func TestToWords(t *testing.T) {
	suite := loadToWordsTestSuite(t)

	for _, tc := range suite.ToWords {
		t.Run(tc.Description+" (con moneda)", func(t *testing.T) {
			got := money.ToWords(tc.Amount)
			if got != tc.Expected {
				t.Errorf("ToWords(%d) = %q; esperaba %q", tc.Amount, got, tc.Expected)
			}
		})

		t.Run(tc.Description+" (sin moneda)", func(t *testing.T) {
			got := money.ToWords(tc.Amount, money.ToWordsOptions{NoCurrency: true})
			if got != tc.ExpectedNoCurrency {
				t.Errorf("ToWords(%d, NoCurrency: true) = %q; esperaba %q", tc.Amount, got, tc.ExpectedNoCurrency)
			}
		})
	}
}
