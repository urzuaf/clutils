package rut_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/urzuaf/clutils/packages/go/rut"
)

func TestGenerate(t *testing.T) {
	t.Run("genera 100 RUTs validos consecutivos por defecto", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			generated := rut.Generate()
			if !rut.Validate(generated) {
				t.Fatalf("Generate() produjo un RUT matemáticamente inválido: %q", generated)
			}
		}
	})

	t.Run("respeta opciones de formato plano", func(t *testing.T) {
		generated := rut.Generate(rut.FormatOptions{Dots: false, Hyphen: false})
		if !rut.Validate(generated) {
			t.Fatalf("Generate() plano produjo un RUT inválido: %q", generated)
		}
		// No debe contener puntos ni guiones
		for _, c := range generated {
			if c == '.' || c == '-' {
				t.Errorf("Generate() plano contiene separadores: %q", generated)
				break
			}
		}
	})

	t.Run("GeneratePerson produce RUTs en rango y validos", func(t *testing.T) {
		for i := 0; i < 50; i++ {
			generated := rut.GeneratePerson()
			if !rut.Validate(generated) {
				t.Fatalf("GeneratePerson() produjo un RUT inválido: %q", generated)
			}
			parts, err := rut.Split(generated)
			if err != nil {
				t.Fatalf("Split fallo en RUT generado: %v", err)
			}
			bodyInt, err := parts.BodyInt()
			if err != nil {
				t.Fatalf("BodyInt fallo: %v", err)
			}
			if bodyInt < 5_000_000 || bodyInt > 28_000_000 {
				t.Errorf("RUT %q fuera del rango esperado de personas (5M-28M)", generated)
			}
		}
	})

	t.Run("GenerateBusiness produce RUTs en rango y validos", func(t *testing.T) {
		for i := 0; i < 50; i++ {
			generated := rut.GenerateBusiness()
			if !rut.Validate(generated) {
				t.Fatalf("GenerateBusiness() produjo un RUT inválido: %q", generated)
			}
			parts, err := rut.Split(generated)
			if err != nil {
				t.Fatalf("Split fallo: %v", err)
			}
			bodyInt, err := parts.BodyInt()
			if err != nil {
				t.Fatalf("BodyInt fallo: %v", err)
			}
			if bodyInt < 60_000_000 || bodyInt > 90_000_000 {
				t.Errorf("RUT %q fuera del rango esperado de empresas (60M-90M)", generated)
			}
		}
	})

	t.Run("GenerateBetween es seguro con limites invertidos o negativos", func(t *testing.T) {
		// Límites invertidos
		inverted := rut.GenerateBetween(80_000_000, 20_000_000)
		if !rut.Validate(inverted) {
			t.Errorf("GenerateBetween invertido produjo RUT inválido: %q", inverted)
		}

		// Negativos
		negative := rut.GenerateBetween(-100, -50)
		if !rut.Validate(negative) {
			t.Errorf("GenerateBetween negativos produjo RUT inválido: %q", negative)
		}
	})

	t.Run("constantes en Go coinciden 1:1 con data/rut/ranges.json", func(t *testing.T) {
		type RangeSpec struct {
			Min int `json:"min"`
			Max int `json:"max"`
		}
		type Schema struct {
			Ranges struct {
				Person   RangeSpec `json:"person"`
				Business RangeSpec `json:"business"`
				General  RangeSpec `json:"general"`
			} `json:"ranges"`
		}

		path := filepath.Join("..", "..", "..", "data", "rut", "ranges.json")
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("error leyendo data/rut/ranges.json: %v", err)
		}

		var schema Schema
		if err := json.Unmarshal(data, &schema); err != nil {
			t.Fatalf("error parseando JSON de rangos: %v", err)
		}

		if rut.DefaultPersonMin != schema.Ranges.Person.Min || rut.DefaultPersonMax != schema.Ranges.Person.Max {
			t.Errorf("Paridad de Person rota: Go=(%d-%d) vs JSON=(%d-%d)",
				rut.DefaultPersonMin, rut.DefaultPersonMax, schema.Ranges.Person.Min, schema.Ranges.Person.Max)
		}

		if rut.DefaultBusinessMin != schema.Ranges.Business.Min || rut.DefaultBusinessMax != schema.Ranges.Business.Max {
			t.Errorf("Paridad de Business rota: Go=(%d-%d) vs JSON=(%d-%d)",
				rut.DefaultBusinessMin, rut.DefaultBusinessMax, schema.Ranges.Business.Min, schema.Ranges.Business.Max)
		}

		if rut.DefaultGeneralMin != schema.Ranges.General.Min || rut.DefaultGeneralMax != schema.Ranges.General.Max {
			t.Errorf("Paridad de General rota: Go=(%d-%d) vs JSON=(%d-%d)",
				rut.DefaultGeneralMin, rut.DefaultGeneralMax, schema.Ranges.General.Min, schema.Ranges.General.Max)
		}
	})
}
