package money

import (
	"strings"
)

// UTMFormatOptions configura las opciones de formateo para Unidades Tributarias Mensuales (UTM).
type UTMFormatOptions struct {
	// Symbol define el prefijo de la unidad (por defecto "UTM").
	Symbol string

	// NoSymbol si es true, omite completamente el simbolo de la unidad.
	NoSymbol bool
}

// FormatUTM formatea un monto entero en Unidades Tributarias Mensuales (UTM)
// utilizando puntos como separador de miles y sin decimales segun la normativa tributaria chilena.
func FormatUTM(amount int64, opts ...UTMFormatOptions) string {
	symbol := "UTM"

	if len(opts) > 0 {
		userOpt := opts[0]
		if userOpt.NoSymbol {
			symbol = ""
		} else if userOpt.Symbol != "" {
			symbol = userOpt.Symbol
		}
	}

	isNegative := amount < 0
	var absVal uint64
	if isNegative {
		absVal = uint64(-(amount + 1)) + 1
	} else {
		absVal = uint64(amount)
	}

	numStr := formatThousands(absVal)

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
