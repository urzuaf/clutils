package money

import (
	"strings"
)

// ToWordsOptions configura las opciones para la conversion de montos a palabras.
type ToWordsOptions struct {
	// NoCurrency si es true, omite la denominacion monetaria ("peso", "pesos", "de pesos").
	NoCurrency bool
}

var unitsWords = [...]string{
	"", "un", "dos", "tres", "cuatro", "cinco", "seis", "siete", "ocho", "nueve",
	"diez", "once", "doce", "trece", "catorce", "quince", "dieciséis", "diecisiete", "dieciocho", "diecinueve",
	"veinte", "veintiún", "veintidós", "veintitrés", "veinticuatro", "veinticinco", "veintiséis", "veintisiete", "veintiocho", "veintinueve",
}

var tensWords = [...]string{
	"", "", "", "treinta", "cuarenta", "cincuenta", "sesenta", "setenta", "ochenta", "noventa",
}

var hundredsWords = [...]string{
	"", "ciento", "doscientos", "trescientos", "cuatrocientos", "quinientos", "seiscientos", "setecientos", "ochocientos", "novecientos",
}

// convertGroup convierte un entero entre 1 y 999 a su representacion textual en espanol.
func convertGroup(n int) string {
	if n == 0 {
		return ""
	}
	if n == 100 {
		return "cien"
	}

	var parts []string
	c := n / 100
	d := (n % 100) / 10
	u := n % 10

	if c > 0 {
		parts = append(parts, hundredsWords[c])
	}

	du := n % 100
	if du > 0 {
		if du < 30 {
			parts = append(parts, unitsWords[du])
		} else {
			if u == 0 {
				parts = append(parts, tensWords[d])
			} else {
				parts = append(parts, tensWords[d]+" y "+unitsWords[u])
			}
		}
	}

	return strings.Join(parts, " ")
}

// numberToWords convierte un uint64 a palabras en espanol (escala larga).
func numberToWords(val uint64) string {
	if val == 0 {
		return "cero"
	}
	if val == 1 {
		return "un"
	}

	billones := val / 1_000_000_000_000
	remBillones := val % 1_000_000_000_000

	millonesTotal := remBillones / 1_000_000
	remMillones := remBillones % 1_000_000

	miles := remMillones / 1_000
	unidades := remMillones % 1_000

	var parts []string

	// 1. Billones (10^12)
	if billones > 0 {
		if billones == 1 {
			parts = append(parts, "un billón")
		} else {
			parts = append(parts, convertGroup(int(billones)), "billones")
		}
	}

	// 2. Millones (hasta 999.999 millones)
	if millonesTotal > 0 {
		if millonesTotal == 1 {
			parts = append(parts, "un millón")
		} else {
			milesDeMillones := millonesTotal / 1_000
			restoMillones := millonesTotal % 1_000

			var millParts []string
			if milesDeMillones > 0 {
				if milesDeMillones == 1 {
					millParts = append(millParts, "mil")
				} else {
					millParts = append(millParts, convertGroup(int(milesDeMillones)), "mil")
				}
			}
			if restoMillones > 0 {
				millParts = append(millParts, convertGroup(int(restoMillones)))
			}
			parts = append(parts, strings.Join(millParts, " ")+" millones")
		}
	}

	// 3. Miles
	if miles > 0 {
		if miles == 1 {
			parts = append(parts, "mil")
		} else {
			parts = append(parts, convertGroup(int(miles)), "mil")
		}
	}

	// 4. Unidades
	if unidades > 0 {
		parts = append(parts, convertGroup(int(unidades)))
	}

	return strings.Join(parts, " ")
}

// ToWords convierte un monto entero en pesos chilenos a su representacion textual
// formal en espanol, respetando las normas gramaticales, ortograficas y bancarias chilenas
// (util para emision de cheques, pagares y contratos legales).
func ToWords(amount int64, opts ...ToWordsOptions) string {
	noCurrency := false
	if len(opts) > 0 {
		noCurrency = opts[0].NoCurrency
	}

	isNegative := amount < 0
	var absVal uint64
	if isNegative {
		absVal = uint64(-(amount + 1)) + 1
	} else {
		absVal = uint64(amount)
	}

	text := numberToWords(absVal)

	if noCurrency {
		if isNegative {
			return "menos " + text
		}
		return text
	}

	var currencySuffix string
	if absVal == 1 {
		currencySuffix = "peso"
	} else if absVal >= 1_000_000 && absVal%1_000_000 == 0 {
		currencySuffix = "de pesos"
	} else {
		currencySuffix = "pesos"
	}

	if isNegative {
		return "menos " + text + " " + currencySuffix
	}
	return text + " " + currencySuffix
}
