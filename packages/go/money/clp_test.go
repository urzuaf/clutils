package money_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/urzuaf/clutils/packages/go/money"
)

type CLPFormattingCase struct {
	Amount               int64  `json:"amount"`
	ExpectedDefault      string `json:"expected_default"`
	ExpectedSpace        string `json:"expected_space"`
	ExpectedSymbolCLP    string `json:"expected_symbol_clp"`
	ExpectedNoSymbol     string `json:"expected_no_symbol"`
	ExpectedParentheses  string `json:"expected_parentheses"`
	Description          string `json:"description"`
}

type CLPParsingCase struct {
	Input       string `json:"input"`
	Expected    int64  `json:"expected"`
	Valid       bool   `json:"valid"`
	Description string `json:"description"`
}

type MoneyTestSuite struct {
	Description   string              `json:"description"`
	CLPFormatting []CLPFormattingCase `json:"clp_formatting"`
	CLPParsing    []CLPParsingCase    `json:"clp_parsing"`
}

func loadMoneyTestSuite(t *testing.T) MoneyTestSuite {
	t.Helper()

	path := filepath.Join("..", "..", "..", "specs", "tests", "money_formatting.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("error leyendo specs/tests/money_formatting.json: %v", err)
	}

	var suite MoneyTestSuite
	if err := json.Unmarshal(data, &suite); err != nil {
		t.Fatalf("error parseando JSON de pruebas de money: %v", err)
	}

	return suite
}

func TestFormatCLP(t *testing.T) {
	suite := loadMoneyTestSuite(t)

	for _, tc := range suite.CLPFormatting {
		t.Run(tc.Description+" (default)", func(t *testing.T) {
			got := money.FormatCLP(tc.Amount)
			if got != tc.ExpectedDefault {
				t.Errorf("FormatCLP(%d) = %q; esperaba %q", tc.Amount, got, tc.ExpectedDefault)
			}
		})

		t.Run(tc.Description+" (space)", func(t *testing.T) {
			got := money.FormatCLP(tc.Amount, money.CLPFormatOptions{Space: true})
			if got != tc.ExpectedSpace {
				t.Errorf("FormatCLP(%d, Space: true) = %q; esperaba %q", tc.Amount, got, tc.ExpectedSpace)
			}
		})

		t.Run(tc.Description+" (symbol CLP)", func(t *testing.T) {
			got := money.FormatCLP(tc.Amount, money.CLPFormatOptions{Symbol: "CLP", Space: true})
			if got != tc.ExpectedSymbolCLP {
				t.Errorf("FormatCLP(%d, Symbol: CLP, Space: true) = %q; esperaba %q", tc.Amount, got, tc.ExpectedSymbolCLP)
			}
		})

		t.Run(tc.Description+" (no symbol)", func(t *testing.T) {
			got := money.FormatCLP(tc.Amount, money.CLPFormatOptions{NoSymbol: true})
			if got != tc.ExpectedNoSymbol {
				t.Errorf("FormatCLP(%d, NoSymbol: true) = %q; esperaba %q", tc.Amount, got, tc.ExpectedNoSymbol)
			}
		})

		if tc.ExpectedParentheses != "" {
			t.Run(tc.Description+" (parentheses)", func(t *testing.T) {
				got := money.FormatCLP(tc.Amount, money.CLPFormatOptions{
					NegativeFormat: money.NegativeFormatParentheses,
				})
				if got != tc.ExpectedParentheses {
					t.Errorf("FormatCLP(%d, NegativeFormat: parentheses) = %q; esperaba %q", tc.Amount, got, tc.ExpectedParentheses)
				}
			})
		}
	}
}

func TestParseCLP(t *testing.T) {
	suite := loadMoneyTestSuite(t)

	for _, tc := range suite.CLPParsing {
		t.Run(tc.Description, func(t *testing.T) {
			got, err := money.ParseCLP(tc.Input)
			if tc.Valid {
				if err != nil {
					t.Fatalf("ParseCLP(%q) retorno error inesperado: %v", tc.Input, err)
				}
				if got != tc.Expected {
					t.Errorf("ParseCLP(%q) = %d; esperaba %d", tc.Input, got, tc.Expected)
				}
			} else {
				if err == nil {
					t.Errorf("ParseCLP(%q) = %d, esperaba error", tc.Input, got)
				}
			}
		})
	}
}
