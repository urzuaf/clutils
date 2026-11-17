package rut_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/urzuaf/clutils/packages/go/rut"
)

type FormattingTestCase struct {
	Input            string `json:"input"`
	ExpectedDefault  string `json:"expected_default"`
	ExpectedNoDots   string `json:"expected_no_dots"`
	ExpectedNoHyphen string `json:"expected_no_hyphen"`
	ExpectedPlain    string `json:"expected_plain"`
	Description      string `json:"description"`
}

type FormattingTestSuite struct {
	Description string               `json:"description"`
	Cases       []FormattingTestCase `json:"cases"`
}

func loadFormattingTestCases(t *testing.T) []FormattingTestCase {
	t.Helper()

	path := filepath.Join("..", "..", "..", "specs", "tests", "rut_formatting.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("error leyendo specs/tests/rut_formatting.json: %v", err)
	}

	var suite FormattingTestSuite
	if err := json.Unmarshal(data, &suite); err != nil {
		t.Fatalf("error parseando JSON de pruebas de formateo: %v", err)
	}

	return suite.Cases
}

func TestFormat(t *testing.T) {
	cases := loadFormattingTestCases(t)

	for _, tc := range cases {
		t.Run(tc.Description+" (default)", func(t *testing.T) {
			got := rut.Format(tc.Input)
			if got != tc.ExpectedDefault {
				t.Errorf("Format(%q) = %q; esperaba %q", tc.Input, got, tc.ExpectedDefault)
			}
		})

		t.Run(tc.Description+" (no dots)", func(t *testing.T) {
			got := rut.Format(tc.Input, rut.FormatOptions{Dots: false, Hyphen: true})
			if got != tc.ExpectedNoDots {
				t.Errorf("Format(%q, no dots) = %q; esperaba %q", tc.Input, got, tc.ExpectedNoDots)
			}
		})

		t.Run(tc.Description+" (no hyphen)", func(t *testing.T) {
			got := rut.Format(tc.Input, rut.FormatOptions{Dots: true, Hyphen: false})
			if got != tc.ExpectedNoHyphen {
				t.Errorf("Format(%q, no hyphen) = %q; esperaba %q", tc.Input, got, tc.ExpectedNoHyphen)
			}
		})

		t.Run(tc.Description+" (plain)", func(t *testing.T) {
			got := rut.Format(tc.Input, rut.FormatOptions{Dots: false, Hyphen: false})
			if got != tc.ExpectedPlain {
				t.Errorf("Format(%q, plain) = %q; esperaba %q", tc.Input, got, tc.ExpectedPlain)
			}
		})
	}
}

type SplitTestCase struct {
	Input        string  `json:"input"`
	ExpectedBody *string `json:"expected_body"`
	ExpectedDv   *string `json:"expected_dv"`
	ExpectedInt  *int    `json:"expected_int"`
	IsError      bool    `json:"is_error"`
	Description  string  `json:"description"`
}

type SplitTestSuite struct {
	Description string          `json:"description"`
	Cases       []SplitTestCase `json:"cases"`
}

func loadSplitTestCases(t *testing.T) []SplitTestCase {
	t.Helper()

	path := filepath.Join("..", "..", "..", "specs", "tests", "rut_split.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("error leyendo specs/tests/rut_split.json: %v", err)
	}

	var suite SplitTestSuite
	if err := json.Unmarshal(data, &suite); err != nil {
		t.Fatalf("error parseando JSON de pruebas de split: %v", err)
	}

	return suite.Cases
}

func TestSplit(t *testing.T) {
	cases := loadSplitTestCases(t)

	for _, tc := range cases {
		t.Run(tc.Description, func(t *testing.T) {
			parts, err := rut.Split(tc.Input)
			if tc.IsError {
				if err == nil {
					t.Errorf("Split(%q) esperaba error, pero retornó nil", tc.Input)
				}
				return
			}

			if err != nil {
				t.Fatalf("Split(%q) falló inesperadamente: %v", tc.Input, err)
			}

			if tc.ExpectedBody != nil && parts.Body != *tc.ExpectedBody {
				t.Errorf("Split(%q).Body = %q; esperaba %q", tc.Input, parts.Body, *tc.ExpectedBody)
			}
			if tc.ExpectedDv != nil && parts.Dv != *tc.ExpectedDv {
				t.Errorf("Split(%q).Dv = %q; esperaba %q", tc.Input, parts.Dv, *tc.ExpectedDv)
			}
			if tc.ExpectedInt != nil {
				num, err := parts.BodyInt()
				if err != nil {
					t.Fatalf("parts.BodyInt() falló: %v", err)
				}
				if num != *tc.ExpectedInt {
					t.Errorf("parts.BodyInt() = %d; esperaba %d", num, *tc.ExpectedInt)
				}
			}
		})
	}
}

type MaskTestCase struct {
	Input           string `json:"input"`
	ExpectedDefault string `json:"expected_default"`
	ExpectedPlain   string `json:"expected_plain"`
	Description     string `json:"description"`
}

type MaskTestSuite struct {
	Description string         `json:"description"`
	Cases       []MaskTestCase `json:"cases"`
}

func loadMaskTestCases(t *testing.T) []MaskTestCase {
	t.Helper()

	path := filepath.Join("..", "..", "..", "specs", "tests", "rut_mask.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("error leyendo specs/tests/rut_mask.json: %v", err)
	}

	var suite MaskTestSuite
	if err := json.Unmarshal(data, &suite); err != nil {
		t.Fatalf("error parseando JSON de pruebas de mask: %v", err)
	}

	return suite.Cases
}

func TestMask(t *testing.T) {
	cases := loadMaskTestCases(t)

	for _, tc := range cases {
		t.Run(tc.Description+" (default)", func(t *testing.T) {
			got := rut.Mask(tc.Input)
			if got != tc.ExpectedDefault {
				t.Errorf("Mask(%q) = %q; esperaba %q", tc.Input, got, tc.ExpectedDefault)
			}
		})

		t.Run(tc.Description+" (plain)", func(t *testing.T) {
			got := rut.Mask(tc.Input, rut.MaskOptions{Formatted: false})
			if got != tc.ExpectedPlain {
				t.Errorf("Mask(%q, plain) = %q; esperaba %q", tc.Input, got, tc.ExpectedPlain)
			}
		})
	}

	// Casos adicionales con opciones personalizadas
	t.Run("con caracter personalizado X", func(t *testing.T) {
		got := rut.Mask("18123458K", rut.MaskOptions{MaskChar: 'X', VisibleCount: 4, Formatted: true})
		expected := "XX.XXX.458-K"
		if got != expected {
			t.Errorf("Mask con X = %q; esperaba %q", got, expected)
		}
	})

	t.Run("con solo 2 caracteres visibles", func(t *testing.T) {
		got := rut.Mask("18123458K", rut.MaskOptions{VisibleCount: 2, Formatted: true})
		expected := "**.***.**8-K"
		if got != expected {
			t.Errorf("Mask con 2 visibles = %q; esperaba %q", got, expected)
		}
	})
}
