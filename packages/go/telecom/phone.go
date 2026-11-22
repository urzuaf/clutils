package telecom

import (
	"errors"
	"fmt"
	"strings"
)

// PhoneType representa la clasificación oficial del número según SUBTEL.
type PhoneType string

const (
	PhoneTypeUnknown    PhoneType = "unknown"
	PhoneTypeMobile     PhoneType = "mobile"
	PhoneTypeLandline   PhoneType = "landline"
	PhoneTypeVoIP       PhoneType = "voip"
	PhoneTypeTollFree   PhoneType = "toll_free"
	PhoneTypeSharedCost PhoneType = "shared_cost"
	PhoneTypeEmergency  PhoneType = "emergency"
)

// PhoneFormatStyle representa los formatos visuales soportados para teléfonos chilenos.
type PhoneFormatStyle string

const (
	FormatInternational PhoneFormatStyle = "international" // Ej: +56 9 1234 5678
	FormatNational      PhoneFormatStyle = "national"      // Ej: 9 1234 5678
	FormatE164          PhoneFormatStyle = "e164"          // Ej: +56912345678
)

var (
	// ErrInvalidPhone indica que el número telefónico no cumple con el formato o prefijos oficiales de SUBTEL.
	ErrInvalidPhone = errors.New("número telefónico chileno inválido")

	// ErrEmptyPhone indica que la entrada está vacía o contiene solo espacios en blanco.
	ErrEmptyPhone = errors.New("el número telefónico no puede estar vacío")

	// ErrNotMobilePhone indica que el número es válido pero no corresponde a un teléfono móvil celular.
	ErrNotMobilePhone = errors.New("el número telefónico no corresponde a una línea móvil")
)

// regionalAreaCodes contiene los códigos de área de 2 dígitos de red fija oficial según SUBTEL.
var regionalAreaCodes = map[string]string{
	"32": "Valparaíso / Isla de Pascua",
	"33": "Quillota",
	"34": "Los Andes / San Felipe",
	"35": "San Antonio",
	"41": "Concepción",
	"42": "Chillán",
	"43": "Los Ángeles",
	"45": "Temuco",
	"51": "La Serena / Coquimbo",
	"52": "Copiapó",
	"53": "Ovalle",
	"55": "Antofagasta / Calama",
	"57": "Iquique",
	"58": "Arica",
	"61": "Punta Arenas",
	"63": "Valdivia",
	"64": "Osorno",
	"65": "Puerto Montt",
	"67": "Coyhaique",
	"71": "Talca",
	"72": "Rancagua / O'Higgins",
	"73": "Linares",
	"75": "Curicó",
}

// emergencyNumbers contiene los números cortos oficiales de emergencia en Chile. (https://es.wikipedia.org/wiki/Anexo:Prefijos_telefónicos_de_Chile)
var emergencyNumbers = map[string]string{
	"130":  "CONAF Incendios Forestales",
	"131":  "SAMU Ambulancia",
	"132":  "Bomberos",
	"133":  "Carabineros",
	"134":  "PDI Policía de Investigaciones",
	"135":  "Fono Drogas",
	"136":  "Cuerpo de Socorro Andino",
	"137":  "Rescate Marítimo Armada",
	"138":  "Rescate Aéreo FACh",
	"139":  "Fono Información Policial",
	"143":  "Fono Acción",
	"147":  "Fono Niños",
	"149":  "Fono Familia",
	"911":  "Central de Carabineros más cercana (solo desde teléfonos móviles)",
	"1404": "Servicio de Rescate",
	"1407": "Fono Emergencia (Mutual de Seguridad)",
	"1411": "Fono Emergencia (Metro de Santiago)",
	"1488": "Fono Denuncia (Metro de Santiago (acoso))",
	"1412": "Fono Drogas (SENDA)",
}

// parsedPhone almacena la información estructurada de un número telefónico procesado.
type parsedPhone struct {
	NationalNumber string
	Type           PhoneType
	AreaCode       string
}

// parsePhone limpia, valida y clasifica cualquier número telefónico chileno.
func parsePhone(phone string) (*parsedPhone, error) {
	trimmed := strings.TrimSpace(phone)
	if trimmed == "" {
		return nil, ErrEmptyPhone
	}

	// Verificar caracteres válidos: dígitos, +, -, espacios, paréntesis, puntos, barras
	for _, r := range trimmed {
		if (r < '0' || r > '9') && r != '+' && r != '-' && r != ' ' && r != '(' && r != ')' && r != '.' && r != '/' {
			return nil, ErrInvalidPhone
		}
	}

	// Extraer solo dígitos
	var digitsBuilder strings.Builder
	digitsBuilder.Grow(len(trimmed))
	for _, r := range trimmed {
		if r >= '0' && r <= '9' {
			digitsBuilder.WriteRune(r)
		}
	}
	digits := digitsBuilder.String()

	if digits == "" {
		return nil, ErrEmptyPhone
	}

	var national string

	// Caso con prefijo de país 56
	if strings.HasPrefix(digits, "56") {
		rest := digits[2:]
		switch len(rest) {
		case 8:
			// +56 seguido de 8 dígitos (móvil legacy sin el 9 inicial)
			national = "9" + rest
		case 9:
			national = rest
		case 10:
			// Líneas 600 de 10 dígitos nacionales
			if strings.HasPrefix(rest, "600") {
				national = rest
			} else {
				return nil, ErrInvalidPhone
			}
		default:
			return nil, ErrInvalidPhone
		}
	} else if strings.HasPrefix(digits, "0") {
		// Casos legacy con cero inicial: 09XXXXXXXX, 02XXXXXXXX, 032XXXXXXX
		rest := digits[1:]
		if len(rest) == 9 {
			national = rest
		} else {
			return nil, ErrInvalidPhone
		}
	} else if len(digits) == 8 {
		// Móvil de 8 dígitos sin el prefijo 9 (ej: 1234 5678)
		national = "9" + digits
	} else if len(digits) == 9 {
		national = digits
	} else if len(digits) == 10 && strings.HasPrefix(digits, "600") {
		// Línea 600 nacional de 10 dígitos (ej: 6006001234)
		national = digits
	} else {
		return nil, ErrInvalidPhone
	}

	// Clasificación y validación de la numeración nacional
	// 1. Móviles: inician con 9 y tienen 9 dígitos
	if strings.HasPrefix(national, "9") && len(national) == 9 {
		return &parsedPhone{
			NationalNumber: national,
			Type:           PhoneTypeMobile,
			AreaCode:       "9",
		}, nil
	}

	// 2. Red Fija Santiago (RM): inicia con 2 y tiene 9 dígitos
	if strings.HasPrefix(national, "2") && len(national) == 9 {
		return &parsedPhone{
			NationalNumber: national,
			Type:           PhoneTypeLandline,
			AreaCode:       "2",
		}, nil
	}

	// 3. Red Fija Regional: código de área de 2 dígitos y 9 dígitos en total
	if len(national) == 9 {
		prefix2 := national[:2]
		if _, ok := regionalAreaCodes[prefix2]; ok {
			return &parsedPhone{
				NationalNumber: national,
				Type:           PhoneTypeLandline,
				AreaCode:       prefix2,
			}, nil
		}
	}

	// 4. Telefonía IP (VoIP): código de área virtual 44 y 9 dígitos en total
	if strings.HasPrefix(national, "44") && len(national) == 9 {
		return &parsedPhone{
			NationalNumber: national,
			Type:           PhoneTypeVoIP,
			AreaCode:       "44",
		}, nil
	}

	// 5. Toll-Free: 800 XXXXXX (9 dígitos)
	if strings.HasPrefix(national, "800") && len(national) == 9 {
		return &parsedPhone{
			NationalNumber: national,
			Type:           PhoneTypeTollFree,
			AreaCode:       "800",
		}, nil
	}

	// 5. Shared-Cost: 600 XXXXXX (9 dígitos) o 600 XXXXXXX (10 dígitos)
	if strings.HasPrefix(national, "600") && (len(national) == 9 || len(national) == 10) {
		return &parsedPhone{
			NationalNumber: national,
			Type:           PhoneTypeSharedCost,
			AreaCode:       "600",
		}, nil
	}

	return nil, ErrInvalidPhone
}

// NormalizePhone normaliza cualquier número telefónico chileno (móvil, red fija o especiales)
// al formato canónico internacional E.164 (+569XXXXXXXX, +56221234567, etc.).
func NormalizePhone(phone string) (string, error) {
	p, err := parsePhone(phone)
	if err != nil {
		return "", err
	}
	return "+56" + p.NationalNumber, nil
}

// NormalizeMobile normaliza específicamente un número móvil celular chileno al estándar E.164 (+569XXXXXXXX).
// Si el número es válido pero corresponde a red fija o servicio especial, retorna ErrNotMobilePhone.
func NormalizeMobile(phone string) (string, error) {
	p, err := parsePhone(phone)
	if err != nil {
		return "", err
	}
	if p.Type != PhoneTypeMobile {
		return "", ErrNotMobilePhone
	}
	return "+56" + p.NationalNumber, nil
}

// ValidatePhone verifica si un número telefónico es válido según el plan de numeración de SUBTEL.
func ValidatePhone(phone string) bool {
	_, err := parsePhone(phone)
	return err == nil
}

// FormatPhone formatea un número telefónico chileno en el estilo visual solicitado.
// Si no se especifica el estilo, se utiliza FormatInternational por defecto.
func FormatPhone(phone string, style ...PhoneFormatStyle) (string, error) {
	p, err := parsePhone(phone)
	if err != nil {
		return "", err
	}

	selectedStyle := FormatInternational
	if len(style) > 0 && style[0] != "" {
		selectedStyle = style[0]
	}

	nat := p.NationalNumber

	switch selectedStyle {
	case FormatE164:
		return "+56" + nat, nil

	case FormatNational:
		switch p.Type {
		case PhoneTypeMobile:
			// 9 1234 5678
			return fmt.Sprintf("%s %s %s", nat[:1], nat[1:5], nat[5:]), nil
		case PhoneTypeLandline:
			if p.AreaCode == "2" {
				// 2 2123 4567
				return fmt.Sprintf("%s %s %s", nat[:1], nat[1:5], nat[5:]), nil
			}
			// 32 212 3456
			return fmt.Sprintf("%s %s %s", nat[:2], nat[2:5], nat[5:]), nil
		case PhoneTypeVoIP:
			// 44 123 4567
			return fmt.Sprintf("%s %s %s", nat[:2], nat[2:5], nat[5:]), nil
		case PhoneTypeTollFree:
			// 800 123 456
			return fmt.Sprintf("%s %s %s", nat[:3], nat[3:6], nat[6:]), nil
		case PhoneTypeSharedCost:
			// 600 600 1234
			return fmt.Sprintf("%s %s %s", nat[:3], nat[3:6], nat[6:]), nil
		default:
			return nat, nil
		}

	case FormatInternational:
		fallthrough
	default:
		switch p.Type {
		case PhoneTypeMobile:
			// +56 9 1234 5678
			return fmt.Sprintf("+56 %s %s %s", nat[:1], nat[1:5], nat[5:]), nil
		case PhoneTypeLandline:
			if p.AreaCode == "2" {
				// +56 2 2123 4567
				return fmt.Sprintf("+56 %s %s %s", nat[:1], nat[1:5], nat[5:]), nil
			}
			// +56 32 212 3456
			return fmt.Sprintf("+56 %s %s %s", nat[:2], nat[2:5], nat[5:]), nil
		case PhoneTypeVoIP:
			// +56 44 123 4567
			return fmt.Sprintf("+56 %s %s %s", nat[:2], nat[2:5], nat[5:]), nil
		case PhoneTypeTollFree:
			// +56 800 123 456
			return fmt.Sprintf("+56 %s %s %s", nat[:3], nat[3:6], nat[6:]), nil
		case PhoneTypeSharedCost:
			// +56 600 600 1234
			return fmt.Sprintf("+56 %s %s %s", nat[:3], nat[3:6], nat[6:]), nil
		default:
			return "+56 " + nat, nil
		}
	}
}

// GetPhoneType retorna la clasificación del número telefónico chileno según SUBTEL.
// Detecta números móviles, red fija, líneas 800, líneas 600 y números cortos de emergencia.
func GetPhoneType(phone string) PhoneType {
	trimmed := strings.TrimSpace(phone)

	// Verificar primero si corresponde a un número corto de emergencia (ej: 133, 131, etc.)
	var cleanDigits strings.Builder
	hasOnlyDigitsOrSeparators := true
	for _, r := range trimmed {
		if r >= '0' && r <= '9' {
			cleanDigits.WriteRune(r)
		} else if r != ' ' && r != '-' {
			hasOnlyDigitsOrSeparators = false
		}
	}

	cleanStr := cleanDigits.String()
	if hasOnlyDigitsOrSeparators && len(cleanStr) == 3 {
		if _, ok := emergencyNumbers[cleanStr]; ok {
			return PhoneTypeEmergency
		}
	}

	p, err := parsePhone(phone)
	if err != nil {
		return PhoneTypeUnknown
	}

	return p.Type
}

// IsEmergencyNumber verifica si una cadena corresponde a un número de emergencia oficial en Chile.
func IsEmergencyNumber(phone string) bool {
	return GetPhoneType(phone) == PhoneTypeEmergency
}
