package rut

import (
	"errors"
	"strconv"
	"strings"
)

// FormatOptions define las opciones de formato visual para un RUT.
type FormatOptions struct {
	Dots   bool // Si incluye puntos separadores de miles (por defecto true)
	Hyphen bool // Si incluye guión antes del dígito verificador (por defecto true)
}

// DefaultFormatOptions retorna la configuración de formato estándar chilena (con puntos y guión).
func DefaultFormatOptions() FormatOptions {
	return FormatOptions{
		Dots:   true,
		Hyphen: true,
	}
}

// Format convierte una cadena de RUT en una representación visual estándar.
// Si no se especifican opciones, aplica el formato formal estándar chileno (ej. "18.123.458-K").
//
// Ejemplos:
//
//	Format("18123458K") // "18.123.458-K"
//	Format("18123458K", FormatOptions{Dots: false, Hyphen: true}) // "18123458-K"
//	Format("18123458K", FormatOptions{Dots: true, Hyphen: false}) // "18.123.458K"
func Format(inputRut string, opts ...FormatOptions) string {
	cleaned := Clean(inputRut)
	if len(cleaned) < 2 {
		return cleaned
	}

	config := DefaultFormatOptions()
	if len(opts) > 0 {
		config = opts[0]
	}

	lastIndex := len(cleaned) - 1
	body := cleaned[:lastIndex]
	dv := string(cleaned[lastIndex])

	var b strings.Builder
	b.Grow(len(cleaned) + 4)

	bodyLen := len(body)
	for i := range bodyLen {
		if config.Dots && i > 0 && (bodyLen-i)%3 == 0 {
			b.WriteByte('.')
		}
		b.WriteByte(body[i])
	}

	if config.Hyphen {
		b.WriteByte('-')
	}

	b.WriteString(dv)

	return b.String()
}

// RutParts representa la descomposición de un RUT en cuerpo y dígito verificador.
type RutParts struct {
	Body string `json:"body"` // Cuerpo numérico (preserva ceros a la izquierda)
	Dv   string `json:"dv"`   // Dígito verificador normalizado en mayúscula ("0"-"9" o "K")
}

// BodyInt convierte el cuerpo del RUT a un entero. Retorna error si el cuerpo contiene caracteres no numéricos.
func (p RutParts) BodyInt() (int, error) {
	return strconv.Atoi(p.Body)
}

// Split divide un RUT en su cuerpo numérico y dígito verificador.
// Acepta entradas en cualquier formato habitual (con o sin puntos, guiones o espacios).
// Retorna error si la cadena contiene menos de 2 caracteres.
//
// Ejemplos:
//
//	parts, err := Split("18.123.458-K") // RutParts{Body: "18123458", Dv: "K"}, nil
//	parts, err := Split("09.123.456-4") // RutParts{Body: "09123456", Dv: "4"}, nil
//	parts, err := Split("1-9")          // RutParts{Body: "1", Dv: "9"}, nil
func Split(inputRut string) (RutParts, error) {
	cleaned := Clean(inputRut)
	if len(cleaned) < 2 {
		return RutParts{}, errors.New("el RUT no contiene cuerpo y dígito verificador suficientes")
	}

	lastIndex := len(cleaned) - 1
	body := cleaned[:lastIndex]
	dv := string(cleaned[lastIndex])

	return RutParts{
		Body: body,
		Dv:   dv,
	}, nil
}

// MaskOptions define las opciones para ofuscar un RUT por privacidad.
type MaskOptions struct {
	MaskChar     rune // Carácter de reemplazo (por defecto '*')
	VisibleCount int  // Cantidad de caracteres visibles al final (por defecto 4: ej. "458-K")
	Formatted    bool // Si conserva el formato formal con puntos y guión (por defecto true)
}

// DefaultMaskOptions retorna la configuración por defecto de privacidad.
func DefaultMaskOptions() MaskOptions {
	return MaskOptions{
		MaskChar:     '*',
		VisibleCount: 4,
		Formatted:    true,
	}
}

// Mask ofusca los caracteres iniciales de un RUT para proteger la privacidad del titular
// según la Ley 19.628 de Protección de la Vida Privada / Datos Personales,
// manteniendo visibles los últimos N caracteres.
//
// Ejemplos:
//
//	Mask("18123458K") // "**.***.458-K"
//	Mask("18.123.458-k", MaskOptions{Formatted: false}) // "*****458K"
func Mask(inputRut string, opts ...MaskOptions) string {
	cleaned := Clean(inputRut)
	if len(cleaned) < 2 {
		return cleaned
	}

	config := DefaultMaskOptions()
	if len(opts) > 0 {
		userOpts := opts[0]
		if userOpts.MaskChar != 0 {
			config.MaskChar = userOpts.MaskChar
		}
		if userOpts.VisibleCount > 0 {
			config.VisibleCount = userOpts.VisibleCount
		}
		config.Formatted = userOpts.Formatted
	}

	if !config.Formatted {
		if config.VisibleCount >= len(cleaned) {
			return cleaned
		}
		maskLen := max(len(cleaned)-config.VisibleCount, 0)
		var b strings.Builder
		b.Grow(len(cleaned))
		for range maskLen {
			b.WriteRune(config.MaskChar)
		}
		b.WriteString(cleaned[maskLen:])
		return b.String()
	}

	formatted := Format(cleaned)
	if config.VisibleCount >= len(cleaned) {
		return formatted
	}

	totalAlphanumeric := len(cleaned)
	threshold := totalAlphanumeric - config.VisibleCount

	var b strings.Builder
	b.Grow(len(formatted))

	seenAlphanumeric := 0
	for _, char := range formatted {
		if char == '.' || char == '-' {
			b.WriteRune(char)
			continue
		}

		if seenAlphanumeric < threshold {
			b.WriteRune(config.MaskChar)
		} else {
			b.WriteRune(char)
		}
		seenAlphanumeric++
	}

	return b.String()
}
