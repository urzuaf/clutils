package rut_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/urzuaf/clutils/packages/go/rut"
)

type TestCase struct {
	Input         string  `json:"input"`
	ExpectedClean string  `json:"expected_clean"`
	ExpectedBody  *string `json:"expected_body"`
	ExpectedDv    *string `json:"expected_dv"`
	IsValid       bool    `json:"is_valid"`
	Description   string  `json:"description"`
}

type TestSuite struct {
	Description string     `json:"description"`
	Cases       []TestCase `json:"cases"`
}

func loadTestCases(t *testing.T) []TestCase {
	t.Helper()

	path := filepath.Join("..", "..", "..", "specs", "tests", "rut_validation.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("error leyendo specs/tests/rut_validation.json: %v", err)
	}

	var suite TestSuite
	if err := json.Unmarshal(data, &suite); err != nil {
		t.Fatalf("error parseando JSON de pruebas: %v", err)
	}

	return suite.Cases
}

func TestClean(t *testing.T) {
	cases := loadTestCases(t)

	for _, tc := range cases {
		t.Run(tc.Description, func(t *testing.T) {
			got := rut.Clean(tc.Input)
			if got != tc.ExpectedClean {
				t.Errorf("Clean(%q) = %q; esperaba %q", tc.Input, got, tc.ExpectedClean)
			}
		})
	}
}

func TestCalculateDv(t *testing.T) {
	cases := loadTestCases(t)

	for _, tc := range cases {
		if tc.ExpectedBody == nil || tc.ExpectedDv == nil {
			continue
		}

		body, err := strconv.Atoi(*tc.ExpectedBody)
		if err != nil {
			continue
		}

		t.Run(tc.Description, func(t *testing.T) {
			got := rut.CalculateDv(body)
			if got != *tc.ExpectedDv {
				t.Errorf("CalculateDv(%d) = %q; esperaba %q", body, got, *tc.ExpectedDv)
			}
		})
	}

	// Casos de borde adicionales
	t.Run("cuerpo negativo o cero", func(t *testing.T) {
		if got := rut.CalculateDv(0); got != "" {
			t.Errorf("CalculateDv(0) = %q; esperaba \"\"", got)
		}
		if got := rut.CalculateDv(-123); got != "" {
			t.Errorf("CalculateDv(-123) = %q; esperaba \"\"", got)
		}
	})
}

func TestValidate(t *testing.T) {
	cases := loadTestCases(t)

	for _, tc := range cases {
		t.Run(tc.Description, func(t *testing.T) {
			got := rut.Validate(tc.Input)
			if got != tc.IsValid {
				t.Errorf("Validate(%q) = %t; esperaba %t", tc.Input, got, tc.IsValid)
			}
		})
	}
}

func TestValidateDetailed(t *testing.T) {
	tests := []struct {
		input        string
		expectedCode rut.ValidationErrorCode
	}{
		{"12.345.678-5", rut.ErrCodeNone},
		{"", rut.ErrCodeEmpty},
		{"   ", rut.ErrCodeEmpty},
		{"12--345-678-5", rut.ErrCodeMalformedFormat},
		{"12..345..678-5", rut.ErrCodeMalformedFormat},
		{"1.2.3.4.5-K", rut.ErrCodeMalformedFormat},
		{"12-345-678-5", rut.ErrCodeMalformedFormat},
		{"1", rut.ErrCodeInvalidLength},
		{"K", rut.ErrCodeInvalidLength},
		{"12345678901-5", rut.ErrCodeInvalidLength},
		{"18.12A.458-K", rut.ErrCodeInvalidCharacters},
		{"abcdefgh-K", rut.ErrCodeInvalidCharacters},
		{"12.345.678-0", rut.ErrCodeInvalidDv},
		{"18.123.458-1", rut.ErrCodeInvalidDv},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			err := rut.ValidateDetailed(tt.input)
			if tt.expectedCode == rut.ErrCodeNone {
				if err != nil {
					t.Errorf("ValidateDetailed(%q) esperaba nil, obtuvo %v (código %d)", tt.input, err, err.Code)
				}
			} else {
				if err == nil {
					t.Fatalf("ValidateDetailed(%q) esperaba error con código %d, obtuvo nil", tt.input, tt.expectedCode)
				}
				if err.Code != tt.expectedCode {
					t.Errorf("ValidateDetailed(%q) código = %d; esperaba %d (msg: %s)", tt.input, err.Code, tt.expectedCode, err.Message)
				}
			}
		})
	}
}
