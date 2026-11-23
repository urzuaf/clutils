package date

import (
	"fmt"
	"strings"
	"time"
)

// DateFormatStyle define el estilo visual para el formateo de fechas.
type DateFormatStyle string

const (
	// DateFormatNumeric representa la fecha en formato numerico estandar chileno con barras ("18/09/2026").
	DateFormatNumeric DateFormatStyle = "numeric"

	// DateFormatNumericHyphen representa la fecha en formato numerico con guiones ("18-09-2026").
	DateFormatNumericHyphen DateFormatStyle = "numeric_hyphen"

	// DateFormatFormal representa la fecha formal en espanol ("18 de septiembre de 2026").
	DateFormatFormal DateFormatStyle = "formal"

	// DateFormatFull representa la fecha formal con dia de la semana ("Viernes, 18 de septiembre de 2026").
	DateFormatFull DateFormatStyle = "full"
)

var spanishMonths = [...]string{
	"",
	"enero",
	"febrero",
	"marzo",
	"abril",
	"mayo",
	"junio",
	"julio",
	"agosto",
	"septiembre",
	"octubre",
	"noviembre",
	"diciembre",
}

var spanishWeekdays = [...]string{
	"Domingo",
	"Lunes",
	"Martes",
	"Miércoles",
	"Jueves",
	"Viernes",
	"Sábado",
}

// isLeapYear determina si un ano es bisiesto segun el calendario gregoriano.
func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// daysInMonth retorna el numero de dias que tiene un mes en un ano determinado.
func daysInMonth(year, month int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if isLeapYear(year) {
			return 29
		}
		return 28
	default:
		return 0
	}
}

// FormatDate formatea una fecha segun los estandares y convenciones locales chilenas.
//
// Estilos disponibles:
//   - "numeric" (por defecto): "18/09/2026"
//   - "numeric_hyphen": "18-09-2026"
//   - "formal": "18 de septiembre de 2026"
//   - "full": "Viernes, 18 de septiembre de 2026"
func FormatDate(t time.Time, style ...DateFormatStyle) string {
	st := DateFormatNumeric
	if len(style) > 0 && style[0] != "" {
		st = style[0]
	}

	day := t.Day()
	month := int(t.Month())
	year := t.Year()

	switch st {
	case DateFormatNumericHyphen:
		return fmt.Sprintf("%02d-%02d-%04d", day, month, year)
	case DateFormatFormal:
		return fmt.Sprintf("%d de %s de %04d", day, spanishMonths[month], year)
	case DateFormatFull:
		weekday := spanishWeekdays[t.Weekday()]
		return fmt.Sprintf("%s, %d de %s de %04d", weekday, day, spanishMonths[month], year)
	default: // DateFormatNumeric
		return fmt.Sprintf("%02d/%02d/%04d", day, month, year)
	}
}

// parseStrictDigits convierte una cadena compuesta exclusivamente por dígitos ASCII '0'-'9' a int.
func parseStrictDigits(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	val := 0
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch < '0' || ch > '9' {
			return 0, false
		}
		val = val*10 + int(ch-'0')
	}
	return val, true
}

// ParseDate interpreta una cadena de texto en formato chileno estandar ("DD/MM/YYYY" o "DD-MM-YYYY"),
// garantizando que el primer segmento sea el dia y el segundo el mes (evitando inversion anglosajona MM/DD/YYYY).
// Valida estrictamente dias por mes, anos bisiestos y rechaza formatos ISO o malformados.
func ParseDate(text string) (time.Time, error) {
	trimmed := strings.TrimSpace(text)
	if len(trimmed) == 0 {
		return time.Time{}, fmt.Errorf("el texto de fecha está vacío")
	}

	var sep rune
	if strings.ContainsRune(trimmed, '/') {
		sep = '/'
	} else if strings.ContainsRune(trimmed, '-') {
		sep = '-'
	} else {
		return time.Time{}, fmt.Errorf("formato no reconocido: se requiere separador '/' o '-'")
	}

	parts := strings.Split(trimmed, string(sep))
	if len(parts) != 3 {
		return time.Time{}, fmt.Errorf("formato de fecha invalido, se esperaba DD%cMM%cYYYY", sep, sep)
	}

	dayStr, monthStr, yearStr := parts[0], parts[1], parts[2]
	if len(dayStr) < 1 || len(dayStr) > 2 || len(monthStr) < 1 || len(monthStr) > 2 {
		return time.Time{}, fmt.Errorf("componentes de día o mes inválidos")
	}

	// El ano debe tener 4 digitos
	if len(yearStr) != 4 {
		return time.Time{}, fmt.Errorf("el año debe tener 4 dígitos: %q", yearStr)
	}

	day, okDay := parseStrictDigits(dayStr)
	month, okMonth := parseStrictDigits(monthStr)
	year, okYear := parseStrictDigits(yearStr)

	if !okDay || !okMonth || !okYear {
		return time.Time{}, fmt.Errorf("componentes numéricos inválidos")
	}

	if year <= 0 {
		return time.Time{}, fmt.Errorf("año inválido: %d", year)
	}

	if month < 1 || month > 12 {
		return time.Time{}, fmt.Errorf("mes fuera de rango: %d (debe ser 1 a 12)", month)
	}

	maxDays := daysInMonth(year, month)
	if day < 1 || day > maxDays {
		return time.Time{}, fmt.Errorf("día fuera de rango: %d para mes %d del año %d", day, month, year)
	}

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil
}

// IsWeekend retorna true si la fecha corresponde a sabado o domingo.
func IsWeekend(t time.Time) bool {
	wd := t.Weekday()
	return wd == time.Saturday || wd == time.Sunday
}
