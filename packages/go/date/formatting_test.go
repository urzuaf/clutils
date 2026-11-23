package date_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/urzuaf/clutils/packages/go/date"
)

type ParseDateCase struct {
	Input         string `json:"input"`
	ExpectedYear  int    `json:"expected_year"`
	ExpectedMonth int    `json:"expected_month"`
	ExpectedDay   int    `json:"expected_day"`
	Valid         bool   `json:"valid"`
	Description   string `json:"description"`
}

type FormatDateCase struct {
	Year                  int    `json:"year"`
	Month                 int    `json:"month"`
	Day                   int    `json:"day"`
	ExpectedNumeric       string `json:"expected_numeric"`
	ExpectedNumericHyphen string `json:"expected_numeric_hyphen"`
	ExpectedFormal        string `json:"expected_formal"`
	ExpectedFull          string `json:"expected_full"`
	Description           string `json:"description"`
}

type IsWeekendCase struct {
	Year        int    `json:"year"`
	Month       int    `json:"month"`
	Day         int    `json:"day"`
	Expected    bool   `json:"expected"`
	Description string `json:"description"`
}

type DateTestSuite struct {
	Description string           `json:"description"`
	ParseDate   []ParseDateCase  `json:"parse_date"`
	FormatDate  []FormatDateCase `json:"format_date"`
	IsWeekend   []IsWeekendCase  `json:"is_weekend"`
}

func loadDateTestSuite(t *testing.T) DateTestSuite {
	t.Helper()

	path := filepath.Join("..", "..", "..", "specs", "tests", "date_formatting.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("error leyendo specs/tests/date_formatting.json: %v", err)
	}

	var suite DateTestSuite
	if err := json.Unmarshal(data, &suite); err != nil {
		t.Fatalf("error parseando JSON de pruebas date: %v", err)
	}

	return suite
}

func TestParseDate(t *testing.T) {
	suite := loadDateTestSuite(t)

	for _, tc := range suite.ParseDate {
		t.Run(tc.Description, func(t *testing.T) {
			got, err := date.ParseDate(tc.Input)
			if tc.Valid {
				if err != nil {
					t.Fatalf("ParseDate(%q) retorno error inesperado: %v", tc.Input, err)
				}
				if got.Year() != tc.ExpectedYear || int(got.Month()) != tc.ExpectedMonth || got.Day() != tc.ExpectedDay {
					t.Errorf("ParseDate(%q) = (%d, %d, %d); esperaba (%d, %d, %d)",
						tc.Input, got.Year(), got.Month(), got.Day(),
						tc.ExpectedYear, tc.ExpectedMonth, tc.ExpectedDay)
				}
			} else {
				if err == nil {
					t.Errorf("ParseDate(%q) debio retornar error pero devolvio %v", tc.Input, got)
				}
			}
		})
	}
}

func TestFormatDate(t *testing.T) {
	suite := loadDateTestSuite(t)

	for _, tc := range suite.FormatDate {
		tDate := time.Date(tc.Year, time.Month(tc.Month), tc.Day, 12, 0, 0, 0, time.UTC)

		t.Run(tc.Description+" (numeric)", func(t *testing.T) {
			got := date.FormatDate(tDate, date.DateFormatNumeric)
			if got != tc.ExpectedNumeric {
				t.Errorf("FormatDate(numeric) = %q; esperaba %q", got, tc.ExpectedNumeric)
			}
		})

		t.Run(tc.Description+" (numeric default)", func(t *testing.T) {
			got := date.FormatDate(tDate)
			if got != tc.ExpectedNumeric {
				t.Errorf("FormatDate(default) = %q; esperaba %q", got, tc.ExpectedNumeric)
			}
		})

		t.Run(tc.Description+" (numeric hyphen)", func(t *testing.T) {
			got := date.FormatDate(tDate, date.DateFormatNumericHyphen)
			if got != tc.ExpectedNumericHyphen {
				t.Errorf("FormatDate(numeric_hyphen) = %q; esperaba %q", got, tc.ExpectedNumericHyphen)
			}
		})

		t.Run(tc.Description+" (formal)", func(t *testing.T) {
			got := date.FormatDate(tDate, date.DateFormatFormal)
			if got != tc.ExpectedFormal {
				t.Errorf("FormatDate(formal) = %q; esperaba %q", got, tc.ExpectedFormal)
			}
		})

		t.Run(tc.Description+" (full)", func(t *testing.T) {
			got := date.FormatDate(tDate, date.DateFormatFull)
			if got != tc.ExpectedFull {
				t.Errorf("FormatDate(full) = %q; esperaba %q", got, tc.ExpectedFull)
			}
		})
	}
}

func TestIsWeekend(t *testing.T) {
	suite := loadDateTestSuite(t)

	for _, tc := range suite.IsWeekend {
		t.Run(tc.Description, func(t *testing.T) {
			tDate := time.Date(tc.Year, time.Month(tc.Month), tc.Day, 12, 0, 0, 0, time.UTC)
			got := date.IsWeekend(tDate)
			if got != tc.Expected {
				t.Errorf("IsWeekend(%v) = %v; esperaba %v", tDate, got, tc.Expected)
			}
		})
	}
}
