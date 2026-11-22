package telecom_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/urzuaf/clutils/packages/go/telecom"
)

type NormalizationTestCase struct {
	Input        string `json:"input"`
	ExpectedE164 string `json:"expected_e164"`
	Valid        bool   `json:"valid"`
	Description  string `json:"description"`
}

type FormattingTestCase struct {
	Input                 string `json:"input"`
	ExpectedInternational string `json:"expected_international"`
	ExpectedNational      string `json:"expected_national"`
	ExpectedE164          string `json:"expected_e164"`
	Description           string `json:"description"`
}

type PhoneTypeTestCase struct {
	Input        string `json:"input"`
	ExpectedType string `json:"expected_type"`
	Description  string `json:"description"`
}

type PhoneTestSuite struct {
	Description   string                  `json:"description"`
	Normalization []NormalizationTestCase `json:"normalization"`
	Formatting    []FormattingTestCase    `json:"formatting"`
	PhoneTypes    []PhoneTypeTestCase     `json:"phone_types"`
}

func loadTestSuite(t *testing.T) PhoneTestSuite {
	t.Helper()

	path := filepath.Join("..", "..", "..", "specs", "tests", "telecom_phone.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("error leyendo specs/tests/telecom_phone.json: %v", err)
	}

	var suite PhoneTestSuite
	if err := json.Unmarshal(data, &suite); err != nil {
		t.Fatalf("error deserializando JSON de pruebas: %v", err)
	}

	return suite
}

func TestNormalizePhone(t *testing.T) {
	suite := loadTestSuite(t)

	for _, tc := range suite.Normalization {
		t.Run(tc.Description, func(t *testing.T) {
			got, err := telecom.NormalizePhone(tc.Input)
			if tc.Valid {
				if err != nil {
					t.Fatalf("NormalizePhone(%q) arrojó error inesperado: %v", tc.Input, err)
				}
				if got != tc.ExpectedE164 {
					t.Errorf("NormalizePhone(%q) = %q; esperaba %q", tc.Input, got, tc.ExpectedE164)
				}
			} else {
				if err == nil {
					t.Errorf("NormalizePhone(%q) = %q; esperaba error", tc.Input, got)
				}
			}
		})
	}
}

func TestValidatePhone(t *testing.T) {
	suite := loadTestSuite(t)

	for _, tc := range suite.Normalization {
		t.Run(tc.Description, func(t *testing.T) {
			isValid := telecom.ValidatePhone(tc.Input)
			if isValid != tc.Valid {
				t.Errorf("ValidatePhone(%q) = %v; esperaba %v", tc.Input, isValid, tc.Valid)
			}
		})
	}
}

func TestFormatPhone(t *testing.T) {
	suite := loadTestSuite(t)

	for _, tc := range suite.Formatting {
		t.Run(tc.Description+" (international)", func(t *testing.T) {
			got, err := telecom.FormatPhone(tc.Input, telecom.FormatInternational)
			if err != nil {
				t.Fatalf("FormatPhone(%q, international) arrojó error: %v", tc.Input, err)
			}
			if got != tc.ExpectedInternational {
				t.Errorf("FormatPhone(%q, international) = %q; esperaba %q", tc.Input, got, tc.ExpectedInternational)
			}
		})

		t.Run(tc.Description+" (default style)", func(t *testing.T) {
			got, err := telecom.FormatPhone(tc.Input)
			if err != nil {
				t.Fatalf("FormatPhone(%q) arrojó error: %v", tc.Input, err)
			}
			if got != tc.ExpectedInternational {
				t.Errorf("FormatPhone(%q) por defecto = %q; esperaba %q", tc.Input, got, tc.ExpectedInternational)
			}
		})

		t.Run(tc.Description+" (national)", func(t *testing.T) {
			got, err := telecom.FormatPhone(tc.Input, telecom.FormatNational)
			if err != nil {
				t.Fatalf("FormatPhone(%q, national) arrojó error: %v", tc.Input, err)
			}
			if got != tc.ExpectedNational {
				t.Errorf("FormatPhone(%q, national) = %q; esperaba %q", tc.Input, got, tc.ExpectedNational)
			}
		})

		t.Run(tc.Description+" (e164)", func(t *testing.T) {
			got, err := telecom.FormatPhone(tc.Input, telecom.FormatE164)
			if err != nil {
				t.Fatalf("FormatPhone(%q, e164) arrojó error: %v", tc.Input, err)
			}
			if got != tc.ExpectedE164 {
				t.Errorf("FormatPhone(%q, e164) = %q; esperaba %q", tc.Input, got, tc.ExpectedE164)
			}
		})
	}
}

func TestGetPhoneType(t *testing.T) {
	suite := loadTestSuite(t)

	for _, tc := range suite.PhoneTypes {
		t.Run(tc.Description, func(t *testing.T) {
			got := telecom.GetPhoneType(tc.Input)
			if string(got) != tc.ExpectedType {
				t.Errorf("GetPhoneType(%q) = %q; esperaba %q", tc.Input, got, tc.ExpectedType)
			}
		})
	}
}

func TestNormalizeMobile(t *testing.T) {
	// Móvil válido
	got, err := telecom.NormalizeMobile("912345678")
	if err != nil || got != "+56912345678" {
		t.Errorf("NormalizeMobile(912345678) = (%q, %v); esperaba (+56912345678, nil)", got, err)
	}

	// Red fija rechazada
	got, err = telecom.NormalizeMobile("221234567")
	if err != telecom.ErrNotMobilePhone {
		t.Errorf("NormalizeMobile(221234567) = (%q, %v); esperaba ErrNotMobilePhone", got, err)
	}

	// Número inválido
	got, err = telecom.NormalizeMobile("12345")
	if err != telecom.ErrInvalidPhone {
		t.Errorf("NormalizeMobile(12345) = (%q, %v); esperaba ErrInvalidPhone", got, err)
	}
}

func TestEmergencyNumbers(t *testing.T) {
	emergencies := []string{"131", "132", "133", "134", "149"}
	for _, em := range emergencies {
		if !telecom.IsEmergencyNumber(em) {
			t.Errorf("IsEmergencyNumber(%q) = false; esperaba true", em)
		}
	}

	if telecom.IsEmergencyNumber("123") {
		t.Errorf("IsEmergencyNumber(123) = true; esperaba false")
	}
}

func TestFormatPhoneInvalid(t *testing.T) {
	_, err := telecom.FormatPhone("invalid")
	if err == nil {
		t.Error("FormatPhone(invalid) esperaba error pero retornó nil")
	}
}
