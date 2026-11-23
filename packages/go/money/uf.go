package money

import (
	"math"
	"strconv"
	"strings"
)

// UFFormatOptions configura las opciones de formateo para Unidades de Fomento (UF).
type UFFormatOptions struct {
	// Decimals indica la cantidad de decimales a formatear (por defecto 2).
	Decimals int

	// Symbol define el prefijo de la unidad (por defecto "UF").
	Symbol string

	// NoSymbol si es true, omite completamente el simbolo de la unidad.
	NoSymbol bool
}

// isAllZeros verifica si una cadena de dígitos contiene únicamente '0'.
func isAllZeros(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != '0' {
			return false
		}
	}
	return true
}

// FormatUF formatea un valor en Unidades de Fomento (UF) segun la convencion oficial chilena,
// utilizando puntos para separador de miles, coma para decimales y redondeo estandar aritmetico
// (round half away from zero) consistente con TypeScript y la normativa local.
func FormatUF(amount float64, opts ...UFFormatOptions) string {
	decimals := 2
	symbol := "UF"

	if len(opts) > 0 {
		userOpt := opts[0]
		if userOpt.Decimals > 0 {
			decimals = userOpt.Decimals
		}
		if userOpt.NoSymbol {
			symbol = ""
		} else if userOpt.Symbol != "" {
			symbol = userOpt.Symbol
		}
	}

	if math.IsNaN(amount) || math.IsInf(amount, 0) {
		amount = 0
	}

	isNegative := amount < 0
	absAmount := math.Abs(amount)

	var integerPartStr string
	var decimalPartStr string

	// Redondeo aritmético determinista (half away from zero) para paridad exacta con TS
	factor := math.Pow10(decimals)
	scaled := absAmount * factor
	if decimals <= 12 && !math.IsInf(scaled, 0) && scaled <= 9e15 {
		roundedInt := uint64(math.Round(scaled))
		rawDigits := strconv.FormatUint(roundedInt, 10)
		if len(rawDigits) <= decimals {
			integerPartStr = "0"
			decimalPartStr = strings.Repeat("0", decimals-len(rawDigits)) + rawDigits
		} else {
			splitPos := len(rawDigits) - decimals
			integerPartStr = rawDigits[:splitPos]
			decimalPartStr = rawDigits[splitPos:]
		}
	} else {
		formattedFloat := strconv.FormatFloat(absAmount, 'f', decimals, 64)
		if before, after, ok := strings.Cut(formattedFloat, "."); ok {
			integerPartStr = before
			decimalPartStr = after
		} else {
			integerPartStr = formattedFloat
		}
	}

	// Evitar "-0,00" cuando un valor negativo pequeño redondea a cero
	if isNegative && isAllZeros(integerPartStr) && isAllZeros(decimalPartStr) {
		isNegative = false
	}

	formattedInt := formatThousandsStr(integerPartStr)

	var numberBuilder strings.Builder
	numberBuilder.WriteString(formattedInt)
	if decimals > 0 && len(decimalPartStr) > 0 {
		numberBuilder.WriteByte(',')
		numberBuilder.WriteString(decimalPartStr)
	}

	numStr := numberBuilder.String()

	var sb strings.Builder
	if symbol != "" {
		sb.WriteString(symbol)
		sb.WriteByte(' ')
	}

	if isNegative {
		sb.WriteByte('-')
	}

	sb.WriteString(numStr)
	return sb.String()
}
