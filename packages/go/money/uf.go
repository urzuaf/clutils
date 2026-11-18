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

// FormatUF formatea un valor en Unidades de Fomento (UF) segun la convencion oficial chilena,
// utilizando puntos para separador de miles y coma para decimales.
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

	isNegative := amount < 0
	absAmount := math.Abs(amount)

	// Formatear a string con la cantidad exacta de decimales y redondeo estandar
	formattedFloat := strconv.FormatFloat(absAmount, 'f', decimals, 64)

	var integerPartStr string
	var decimalPartStr string

	if before, after, ok := strings.Cut(formattedFloat, "."); ok {
		integerPartStr = before
		decimalPartStr = after
	} else {
		integerPartStr = formattedFloat
	}

	intVal, _ := strconv.ParseUint(integerPartStr, 10, 64)
	formattedInt := formatThousands(intVal)

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
