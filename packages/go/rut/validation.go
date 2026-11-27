package rut

import (
	"strconv"
	"strings"
)

// ValidationErrorCode representa el código numérico del error de validación de un RUT.
type ValidationErrorCode int

const (
	ErrCodeNone              ValidationErrorCode = iota // 0: Sin error (Válido)
	ErrCodeEmpty                                        // 1: El RUT está vacío o contiene solo espacios
	ErrCodeMalformedFormat                              // 2: Puntos o guiones mal ubicados o duplicados
	ErrCodeInvalidLength                                // 3: Longitud del RUT inválida (fuera del rango permitido)
	ErrCodeInvalidCharacters                            // 4: El cuerpo contiene letras o caracteres no numéricos
	ErrCodeInvalidDv                                    // 5: El dígito verificador no coincide con el Módulo 11
)

// ValidationError representa un error detallado de validación de RUT.
// Contiene un código numérico para fácil serialización/i18n y un mensaje legible.
type ValidationError struct {
	Code    ValidationErrorCode `json:"code"`
	Message string              `json:"message"`
}

// Error implementa la interfaz estándar error de Go.
func (e *ValidationError) Error() string {
	return e.Message
}

// Clean sanitiza una cadena de RUT eliminando puntos, guiones y espacios en blanco,
// y normalizando el dígito verificador 'k' a mayúscula 'K'.
//
// Ejemplo:
//
//	Clean(" 18.123.458-k ") // "18123458K"
func Clean(inputRut string) string {
	var cleaned strings.Builder
	cleaned.Grow(len(inputRut))

	for _, char := range inputRut {
		if char == '.' || char == '-' || char == ' ' || char == '\t' || char == '\n' || char == '\r' {
			continue
		}

		if char == 'k' {
			cleaned.WriteRune('K')
			continue
		}

		cleaned.WriteRune(char)
	}

	return cleaned.String()
}

// CalculateDv calcula el Dígito Verificador oficial para un cuerpo numérico
// aplicando el algoritmo de Módulo 11 con ponderadores cíclicos [2..7].
//
// Retorna "0"-"9" o "K". Si el cuerpo es menor o igual a cero, retorna cadena vacía "".
//
// Ejemplo:
//
//	CalculateDv(18123458) // "K"
//	CalculateDv(18123466) // "0"
//	CalculateDv(12345678) // "5"
func CalculateDv(body int) string {
	if body <= 0 {
		return ""
	}

	sum := 0
	multiplier := 2

	for n := body; n > 0; n /= 10 {
		digit := n % 10
		sum += digit * multiplier

		multiplier++
		if multiplier > 7 {
			multiplier = 2
		}
	}

	dv := 11 - (sum % 11)

	switch dv {
	case 11:
		return "0"
	case 10:
		return "K"
	default:
		return strconv.Itoa(dv)
	}
}

// ValidateDetailed valida un RUT y retorna nil si es válido,
// o un puntero a ValidationError con el código numérico y motivo exacto del fallo.
//
// Ejemplo:
//
//	if err := ValidateDetailed("18.12A.458-K"); err != nil {
//	    switch err.Code {
//	    case ErrCodeInvalidCharacters:
//	        fmt.Println("No uses letras en el cuerpo")
//	    }
//	}
func ValidateDetailed(inputRut string) *ValidationError {
	trimmed := strings.TrimSpace(inputRut)
	if len(trimmed) == 0 {
		return &ValidationError{
			Code:    ErrCodeEmpty,
			Message: "el RUT no puede estar vacío",
		}
	}

	// 1. Control de separadores repetidos o malformados
	if strings.Contains(inputRut, "--") || strings.Contains(inputRut, "..") {
		return &ValidationError{
			Code:    ErrCodeMalformedFormat,
			Message: "el RUT contiene puntos o guiones repetidos",
		}
	}

	if strings.Count(inputRut, "-") > 1 || strings.Count(inputRut, ".") > 2 {
		return &ValidationError{
			Code:    ErrCodeMalformedFormat,
			Message: "el RUT excede la cantidad máxima de puntos o guiones",
		}
	}

	var bodyWithDots string
	if before, after, ok := strings.Cut(trimmed, "-"); ok {
		if strings.Contains(after, ".") {
			return &ValidationError{
				Code:    ErrCodeMalformedFormat,
				Message: "no se permiten puntos después del guión del RUT",
			}
		}
		afterHyphen := []rune(Clean(after))
		if len(afterHyphen) != 1 {
			return &ValidationError{
				Code:    ErrCodeMalformedFormat,
				Message: "el guión del RUT debe separar únicamente el dígito verificador",
			}
		}
		bodyWithDots = stripWhitespace(before)
	} else {
		noSpace := stripWhitespace(trimmed)
		if strings.Contains(noSpace, ".") {
			noSpaceRunes := []rune(noSpace)
			if len(noSpaceRunes) < 2 || noSpaceRunes[len(noSpaceRunes)-1] == '.' {
				return &ValidationError{
					Code:    ErrCodeMalformedFormat,
					Message: "ubicación inválida de puntos separadores en el RUT",
				}
			}
			bodyWithDots = string(noSpaceRunes[:len(noSpaceRunes)-1])
		}
	}

	if strings.Contains(bodyWithDots, ".") {
		groups := strings.Split(bodyWithDots, ".")
		if len([]rune(groups[0])) < 1 || len([]rune(groups[0])) > 3 {
			return &ValidationError{
				Code:    ErrCodeMalformedFormat,
				Message: "ubicación inválida de puntos separadores en el RUT",
			}
		}
		for i := 1; i < len(groups); i++ {
			if len([]rune(groups[i])) != 3 {
				return &ValidationError{
					Code:    ErrCodeMalformedFormat,
					Message: "ubicación inválida de puntos separadores en el RUT",
				}
			}
		}
	}

	cleaned := Clean(inputRut)
	cleanedRunes := []rune(cleaned)

	// 2. Control de longitud (por runas, no por bytes UTF-8)
	if len(cleanedRunes) < 2 || len(cleanedRunes) > 10 {
		return &ValidationError{
			Code:    ErrCodeInvalidLength,
			Message: "la longitud del RUT es inválida",
		}
	}

	lastIndex := len(cleanedRunes) - 1
	body := string(cleanedRunes[:lastIndex])
	dv := string(cleanedRunes[lastIndex])

	// 3. Control de cuerpo numérico (estrictamente dígitos ASCII '0'-'9', rechaza '+' o '-')
	for i := 0; i < len(body); i++ {
		if body[i] < '0' || body[i] > '9' {
			return &ValidationError{
				Code:    ErrCodeInvalidCharacters,
				Message: "el cuerpo del RUT contiene caracteres no numéricos",
			}
		}
	}

	numBody, err := strconv.Atoi(body)
	if err != nil {
		return &ValidationError{
			Code:    ErrCodeInvalidCharacters,
			Message: "el cuerpo del RUT contiene caracteres no numéricos",
		}
	}

	// 4. Control de Dígito Verificador Módulo 11
	if CalculateDv(numBody) != dv {
		return &ValidationError{
			Code:    ErrCodeInvalidDv,
			Message: "el dígito verificador no coincide",
		}
	}

	return nil
}

func stripWhitespace(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		if r == ' ' || r == '\t' || r == '\n' || r == '\r' {
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

// Validate verifica si una cadena representa un RUT matemáticamente válido según la normativa chilena.
// Retorna true si es válido, o false si está malformado, vacío o su DV es incorrecto.
//
// Si requieres conocer el motivo específico del fallo o su código de error, utiliza ValidateDetailed.
//
// Ejemplo:
//
//	Validate("12.345.678-5") // true
//	Validate("18.123.458-k") // true
//	Validate("12.345.678-0") // false
func Validate(inputRut string) bool {
	return ValidateDetailed(inputRut) == nil
}
