// Package money proporciona utilidades para operaciones monetarias en Chile,
// incluyendo formateo y parseo de CLP, calculos de IVA (normativa SII),
// UF, UTM y conversion a palabras segun la convencion local.
package money

import (
	"fmt"
	"strconv"
	"strings"
)

// NegativeFormat define el estilo de representacion para montos negativos.
type NegativeFormat string

const (
	// NegativeFormatMinusSign representa montos negativos con signo menos (ej: "-$1.500").
	NegativeFormatMinusSign NegativeFormat = "-sign"

	// NegativeFormatParentheses representa montos negativos en formato contable entre parentesis (ej: "($1.500)").
	NegativeFormatParentheses NegativeFormat = "parentheses"
)

// CLPFormatOptions configura las opciones de formateo para pesos chilenos.
type CLPFormatOptions struct {
	// Symbol define el simbolo o prefijo de la moneda (por defecto "$").
	Symbol string

	// NoSymbol si es true, omite completamente el simbolo de la moneda.
	NoSymbol bool

	// Space indica si se debe incluir un espacio entre el simbolo y el monto (ej: "$ 1.500").
	Space bool

	// NegativeFormat define el estilo para numeros negativos ("-sign" por defecto o "parentheses").
	NegativeFormat NegativeFormat
}

// formatThousands formatea un entero positivo con separadores de miles '.' de forma lineal y sin regex.
func formatThousands(val uint64) string {
	s := strconv.FormatUint(val, 10)
	n := len(s)
	if n <= 3 {
		return s
	}

	dots := (n - 1) / 3
	res := make([]byte, n+dots)
	j := len(res) - 1
	count := 0

	for i := n - 1; i >= 0; i-- {
		res[j] = s[i]
		j--
		count++
		if count == 3 && i > 0 {
			res[j] = '.'
			j--
			count = 0
		}
	}

	return string(res)
}

// FormatCLP formatea un monto entero a pesos chilenos (CLP) segun la convencion local,
// utilizando puntos como separador de miles y sin centavos.
func FormatCLP(amount int64, opts ...CLPFormatOptions) string {
	opt := CLPFormatOptions{
		Symbol:         "$",
		NegativeFormat: NegativeFormatMinusSign,
	}

	if len(opts) > 0 {
		userOpt := opts[0]
		if userOpt.NoSymbol {
			opt.Symbol = ""
		} else if userOpt.Symbol != "" {
			opt.Symbol = userOpt.Symbol
		}

		opt.Space = userOpt.Space
		if userOpt.NegativeFormat != "" {
			opt.NegativeFormat = userOpt.NegativeFormat
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

	spaceStr := ""
	if opt.Space && opt.Symbol != "" {
		spaceStr = " "
	}

	var sb strings.Builder
	if !isNegative {
		sb.WriteString(opt.Symbol)
		sb.WriteString(spaceStr)
		sb.WriteString(numStr)
		return sb.String()
	}

	if opt.NegativeFormat == NegativeFormatParentheses {
		sb.WriteByte('(')
		sb.WriteString(opt.Symbol)
		sb.WriteString(spaceStr)
		sb.WriteString(numStr)
		sb.WriteByte(')')
	} else {
		sb.WriteByte('-')
		sb.WriteString(opt.Symbol)
		sb.WriteString(spaceStr)
		sb.WriteString(numStr)
	}

	return sb.String()
}

// ParseCLP extrae y convierte cualquier representacion en texto de pesos chilenos
// a un numero entero limpio (int64), reconociendo signos negativos estandar y contables.
func ParseCLP(amount string) (int64, error) {
	trimmed := strings.TrimSpace(amount)
	if len(trimmed) == 0 {
		return 0, fmt.Errorf("el texto está vacío")
	}

	// Normalizar removiendo "CLP" (prefijo o sufijo) si existe
	clean := trimmed
	cleanUpper := strings.ToUpper(clean)
	if strings.HasPrefix(cleanUpper, "CLP") {
		clean = strings.TrimSpace(clean[3:])
	} else if strings.HasSuffix(cleanUpper, "CLP") {
		clean = strings.TrimSpace(clean[:len(clean)-3])
	}

	var (
		isNegative  bool
		hasParen    bool
		minusCount  int
		openParens  int
		closeParens int
		digits      []byte
	)

	for i := 0; i < len(clean); i++ {
		ch := clean[i]
		switch {
		case ch >= '0' && ch <= '9':
			if closeParens > 0 {
				return 0, fmt.Errorf("dígito posterior al paréntesis de cierre")
			}
			digits = append(digits, ch)
		case ch == '.':
			// Separador de miles valido
			continue
		case ch == '$' || ch == ' ' || ch == '\t':
			// Simbolo o espacio valido
			continue
		case ch == '-':
			if len(digits) > 0 {
				return 0, fmt.Errorf("signo menos no puede ir después de dígitos")
			}
			minusCount++
			isNegative = true
		case ch == '(':
			if len(digits) > 0 {
				return 0, fmt.Errorf("paréntesis de apertura no puede ir después de dígitos")
			}
			openParens++
		case ch == ')':
			if openParens == 0 {
				return 0, fmt.Errorf("paréntesis de cierre sin apertura previa")
			}
			closeParens++
		default:
			return 0, fmt.Errorf("carácter no válido en monto CLP: %q", ch)
		}
	}

	// Validaciones de consistencia de signo
	if openParens > 0 || closeParens > 0 {
		if openParens != 1 || closeParens != 1 {
			return 0, fmt.Errorf("paréntesis contables desbalanceados")
		}
		if minusCount > 0 {
			return 0, fmt.Errorf("formato negativo ambiguo (combina '-' y paréntesis)")
		}
		hasParen = true
	}

	if minusCount > 1 {
		return 0, fmt.Errorf("múltiples signos menos inválidos")
	}

	if len(digits) == 0 {
		return 0, fmt.Errorf("el texto no contiene dígitos numéricos válidos")
	}

	val, err := strconv.ParseInt(string(digits), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("desborde o error numérico: %w", err)
	}

	if isNegative || hasParen {
		val = -val
	}

	return val, nil
}
