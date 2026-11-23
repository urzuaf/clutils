package geo

import (
	"strings"
	"unicode"
)

// normalizeSearch limpia y normaliza una cadena para búsquedas insensibles
// a mayúsculas, tildes, eñes, puntuación y espacios redundantes.
func normalizeSearch(s string) string {
	trimmed := strings.TrimSpace(s)
	if trimmed == "" {
		return ""
	}

	var b strings.Builder
	b.Grow(len(trimmed))

	for _, r := range trimmed {
		// Omitir marcas diacríticas combinantes Unicode (NFD, ej. 'n' + '\u0303')
		if r >= 0x0300 && r <= 0x036F {
			continue
		}
		lr := unicode.ToLower(r)
		switch lr {
		case 'á', 'à', 'ä', 'â':
			b.WriteRune('a')
		case 'é', 'è', 'ë', 'ê':
			b.WriteRune('e')
		case 'í', 'ì', 'ï', 'î':
			b.WriteRune('i')
		case 'ó', 'ò', 'ö', 'ô':
			b.WriteRune('o')
		case 'ú', 'ù', 'ü', 'û':
			b.WriteRune('u')
		case 'ñ':
			b.WriteRune('n')
		case '\'', '.', '-', '_', '/':
			continue
		default:
			b.WriteRune(lr)
		}
	}

	words := strings.Fields(b.String())
	return strings.Join(words, " ")
}

// searchEntity es el motor genérico de búsqueda territorial para retornar un único resultado:
// 1. Coincidencia por identificadores exactos (códigos INE, ISOs, abreviaturas, CUTs).
// 2. Coincidencia exacta por nombre normalizado (sin tildes, eñes ni mayúsculas).
// 3. Coincidencia parcial contenida (subcadena) en el nombre normalizado.
func searchEntity[T any](
	query string,
	items []T,
	getIdents func(T) []string,
	getName func(T) string,
) (*T, bool) {
	trimmed := strings.TrimSpace(query)
	if trimmed == "" {
		return nil, false
	}

	normQuery := normalizeSearch(query)

	padded := trimmed
	if len(trimmed) == 1 && unicode.IsDigit(rune(trimmed[0])) {
		padded = "0" + trimmed
	}

	// Paso 1: Coincidencia por identificadores exactos
	for i := range items {
		for _, id := range getIdents(items[i]) {
			if id == "" {
				continue
			}
			if strings.EqualFold(id, trimmed) || strings.EqualFold(id, padded) || normalizeSearch(id) == normQuery {
				res := items[i]
				return &res, true
			}
		}
	}

	if normQuery == "" {
		return nil, false
	}

	// Paso 2: Coincidencia exacta por nombre normalizado (ej: "nunoa" == "nunoa")
	for i := range items {
		if normalizeSearch(getName(items[i])) == normQuery {
			res := items[i]
			return &res, true
		}
	}

	// Paso 3: Coincidencia por prefijo del nombre normalizado (coherente con searchAllEntities[0])
	for i := range items {
		if strings.HasPrefix(normalizeSearch(getName(items[i])), normQuery) {
			res := items[i]
			return &res, true
		}
	}

	// Paso 4: Coincidencia parcial contenida (ej: "metropolitana" en "metropolitana de santiago")
	for i := range items {
		if strings.Contains(normalizeSearch(getName(items[i])), normQuery) {
			res := items[i]
			return &res, true
		}
	}

	return nil, false
}

// searchAllEntities busca y retorna todas las entidades coincidentes como un subconjunto ([]T),
// ordenando primero coincidencias exactas, luego coincidencias por prefijo y finalmente subcadenas.
func searchAllEntities[T any](
	query string,
	items []T,
	getIdents func(T) []string,
	getName func(T) string,
) []T {
	trimmed := strings.TrimSpace(query)
	if trimmed == "" {
		return []T{}
	}

	normQuery := normalizeSearch(query)

	padded := trimmed
	if len(trimmed) == 1 && unicode.IsDigit(rune(trimmed[0])) {
		padded = "0" + trimmed
	}

	var exactMatches []T
	var prefixMatches []T
	var substringMatches []T
	seen := make(map[int]bool)

	// 1. Identificadores exactos
	for i := range items {
		matched := false
		for _, id := range getIdents(items[i]) {
			if id == "" {
				continue
			}
			if strings.EqualFold(id, trimmed) || strings.EqualFold(id, padded) || normalizeSearch(id) == normQuery {
				matched = true
				break
			}
		}
		if matched && !seen[i] {
			seen[i] = true
			exactMatches = append(exactMatches, items[i])
		}
	}

	if normQuery == "" {
		if exactMatches == nil {
			return []T{}
		}
		return exactMatches
	}

	// 2. Coincidencia exacta por nombre
	for i := range items {
		if seen[i] {
			continue
		}
		if normalizeSearch(getName(items[i])) == normQuery {
			seen[i] = true
			exactMatches = append(exactMatches, items[i])
		}
	}

	// 3. Coincidencia por prefijo del nombre
	for i := range items {
		if seen[i] {
			continue
		}
		if strings.HasPrefix(normalizeSearch(getName(items[i])), normQuery) {
			seen[i] = true
			prefixMatches = append(prefixMatches, items[i])
		}
	}

	// 4. Coincidencia por subcadena dentro del nombre
	for i := range items {
		if seen[i] {
			continue
		}
		if strings.Contains(normalizeSearch(getName(items[i])), normQuery) {
			seen[i] = true
			substringMatches = append(substringMatches, items[i])
		}
	}

	total := len(exactMatches) + len(prefixMatches) + len(substringMatches)
	result := make([]T, 0, total)
	result = append(result, exactMatches...)
	result = append(result, prefixMatches...)
	result = append(result, substringMatches...)
	return result
}
