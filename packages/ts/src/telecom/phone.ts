/**
 * @file phone.ts
 * @description Normalización E.164, validación, formateo visual y clasificación de teléfonos chilenos (SUBTEL).
 */

/**
 * Clasificación oficial del número telefónico chileno según la SUBTEL.
 */
export type PhoneType =
  | 'unknown'
  | 'mobile'
  | 'landline'
  | 'voip'
  | 'toll_free'
  | 'shared_cost'
  | 'emergency';

/**
 * Formatos visuales soportados para teléfonos chilenos.
 */
export type PhoneFormatStyle = 'international' | 'national' | 'e164';

/**
 * Códigos de área oficiales de 2 dígitos de red fija regional según SUBTEL.
 */
export const REGIONAL_AREA_CODES = new Set([
  '32', // Valparaíso / Isla de Pascua
  '33', // Quillota
  '34', // Los Andes / San Felipe
  '35', // San Antonio
  '41', // Concepción
  '42', // Chillán
  '43', // Los Ángeles
  '45', // Temuco
  '51', // La Serena / Coquimbo
  '52', // Copiapó
  '53', // Ovalle
  '55', // Antofagasta / Calama
  '57', // Iquique
  '58', // Arica
  '61', // Punta Arenas
  '63', // Valdivia
  '64', // Osorno
  '65', // Puerto Montt
  '67', // Coyhaique
  '71', // Talca
  '72', // Rancagua / O'Higgins
  '73', // Linares
  '75', // Curicó
]);

/**
 * Números cortos oficiales de servicios de emergencia en Chile.
 */
export const EMERGENCY_NUMBERS = new Set([
  '131', // SAMU Ambulancia
  '132', // Bomberos
  '133', // Carabineros
  '134', // PDI
  '149', // Fono Familia Carabineros
  '130', // CONAF Incendios Forestales
  '135', // Fono Drogas
  '136', // Cuerpo de Socorro Andino
  '137', // Rescate Marítimo Armada
  '138', // Rescate Aéreo FACh
  '139', // Fono Información Policial
  '147', // Fono Niños
]);

/**
 * Estructura interna de un teléfono chileno procesado.
 */
interface ParsedPhone {
  nationalNumber: string;
  type: PhoneType;
  areaCode: string;
}

/**
 * Limpia, valida y clasifica cualquier número telefónico chileno.
 *
 * @throws {Error} Si el formato es inválido, vacío o contiene caracteres no permitidos.
 */
function parsePhone(phone: string): ParsedPhone {
  if (typeof phone !== 'string') {
    throw new Error('El número telefónico no puede estar vacío');
  }

  const trimmed = phone.trim();
  if (trimmed === '') {
    throw new Error('El número telefónico no puede estar vacío');
  }

  // Verificar caracteres permitidos: dígitos, +, -, espacios, paréntesis, puntos, barras
  for (let i = 0; i < trimmed.length; i++) {
    const ch = trimmed[i];
    if (
      (ch < '0' || ch > '9') &&
      ch !== '+' &&
      ch !== '-' &&
      ch !== ' ' &&
      ch !== '(' &&
      ch !== ')' &&
      ch !== '.' &&
      ch !== '/'
    ) {
      throw new Error('Número telefónico chileno inválido');
    }
  }

  // Extraer dígitos
  let digits = '';
  for (let i = 0; i < trimmed.length; i++) {
    const ch = trimmed[i];
    if (ch >= '0' && ch <= '9') {
      digits += ch;
    }
  }

  if (digits === '') {
    throw new Error('El número telefónico no puede estar vacío');
  }

  let national: string;

  // Caso con prefijo de país 56
  if (digits.startsWith('56')) {
    const rest = digits.slice(2);
    if (rest.length === 8) {
      // +56 seguido de 8 dígitos (móvil legacy sin el 9 inicial)
      national = '9' + rest;
    } else if (rest.length === 9) {
      national = rest;
    } else if (rest.length === 10 && rest.startsWith('600')) {
      // Líneas 600 de 10 dígitos nacionales
      national = rest;
    } else {
      throw new Error('Número telefónico chileno inválido');
    }
  } else if (digits.startsWith('0')) {
    // Casos legacy con cero inicial: 09XXXXXXXX, 02XXXXXXXX, 032XXXXXXX
    const rest = digits.slice(1);
    if (rest.length === 9) {
      national = rest;
    } else {
      throw new Error('Número telefónico chileno inválido');
    }
  } else if (digits.length === 8) {
    // Móvil de 8 dígitos sin el prefijo 9 (ej: 1234 5678)
    national = '9' + digits;
  } else if (digits.length === 9) {
    national = digits;
  } else if (digits.length === 10 && digits.startsWith('600')) {
    // Línea 600 nacional de 10 dígitos (ej: 6006001234)
    national = digits;
  } else {
    throw new Error('Número telefónico chileno inválido');
  }

  // Clasificación y validación de la numeración nacional
  // 1. Móviles: inician con 9 y tienen 9 dígitos
  if (national.startsWith('9') && national.length === 9) {
    return {
      nationalNumber: national,
      type: 'mobile',
      areaCode: '9',
    };
  }

  // 2. Red Fija Santiago (RM): inicia con 2 y tiene 9 dígitos
  if (national.startsWith('2') && national.length === 9) {
    return {
      nationalNumber: national,
      type: 'landline',
      areaCode: '2',
    };
  }

  // 3. Red Fija Regional: código de área de 2 dígitos y 9 dígitos en total
  if (national.length === 9) {
    const prefix2 = national.slice(0, 2);
    if (REGIONAL_AREA_CODES.has(prefix2)) {
      return {
        nationalNumber: national,
        type: 'landline',
        areaCode: prefix2,
      };
    }
  }

  // 4. Telefonía IP (VoIP): código de área virtual 44 y 9 dígitos en total
  if (national.startsWith('44') && national.length === 9) {
    return {
      nationalNumber: national,
      type: 'voip',
      areaCode: '44',
    };
  }

  // 5. Toll-Free: 800 XXXXXX (9 dígitos)
  if (national.startsWith('800') && national.length === 9) {
    return {
      nationalNumber: national,
      type: 'toll_free',
      areaCode: '800',
    };
  }

  // 5. Shared-Cost: 600 XXXXXX (9 dígitos) o 600 XXXXXXX (10 dígitos)
  if (national.startsWith('600') && (national.length === 9 || national.length === 10)) {
    return {
      nationalNumber: national,
      type: 'shared_cost',
      areaCode: '600',
    };
  }

  throw new Error('Número telefónico chileno inválido');
}

/**
 * Normaliza cualquier número telefónico chileno (móvil, red fija regional o líneas especiales)
 * al formato canónico internacional E.164 (+569XXXXXXXX, +56221234567, etc.).
 *
 * @param phone - Cadena con el número telefónico en cualquier formato.
 * @returns Cadena normalizada en estándar E.164.
 *
 * @example
 * ```ts
 * normalizePhone("9 1234 5678"); // "+56912345678"
 * normalizePhone("2 2123 4567"); // "+56221234567"
 * normalizePhone("09 1234 5678"); // "+56912345678"
 * ```
 */
export function normalizePhone(phone: string): string {
  const parsed = parsePhone(phone);
  return `+56${parsed.nationalNumber}`;
}

/**
 * Normaliza específicamente un número móvil celular chileno al estándar E.164 (+569XXXXXXXX).
 *
 * @param phone - Cadena con el número móvil.
 * @returns Cadena normalizada en estándar E.164 (+569XXXXXXXX).
 * @throws {Error} Si el número no corresponde a un teléfono móvil celular.
 *
 * @example
 * ```ts
 * normalizeMobile("912345678"); // "+56912345678"
 * normalizeMobile("1234 5678"); // "+56912345678"
 * ```
 */
export function normalizeMobile(phone: string): string {
  const parsed = parsePhone(phone);
  if (parsed.type !== 'mobile') {
    throw new Error('El número telefónico no corresponde a una línea móvil');
  }
  return `+56${parsed.nationalNumber}`;
}

/**
 * Verifica si un número telefónico es válido según el plan de numeración de SUBTEL.
 *
 * @param phone - Cadena a validar.
 * @returns `true` si el número es válido y normalizable a E.164, `false` en caso contrario.
 *
 * @example
 * ```ts
 * validatePhone("+56 9 1234 5678"); // true
 * validatePhone("12345"); // false
 * ```
 */
export function validatePhone(phone: string): boolean {
  try {
    parsePhone(phone);
    return true;
  } catch {
    return false;
  }
}

/**
 * Formatea un número telefónico chileno en el estilo visual solicitado.
 *
 * @param phone - Cadena con el número telefónico.
 * @param style - Estilo de formateo visual ('international' | 'national' | 'e164'). Por defecto 'international'.
 * @returns Cadena formateada visualmente.
 *
 * @example
 * ```ts
 * formatPhone("912345678"); // "+56 9 1234 5678"
 * formatPhone("912345678", "national"); // "9 1234 5678"
 * formatPhone("912345678", "e164"); // "+56912345678"
 * ```
 */
export function formatPhone(phone: string, style: PhoneFormatStyle = 'international'): string {
  const parsed = parsePhone(phone);
  const nat = parsed.nationalNumber;

  switch (style) {
    case 'e164':
      return `+56${nat}`;

    case 'national':
      switch (parsed.type) {
        case 'mobile':
          // 9 1234 5678
          return `${nat[0]} ${nat.slice(1, 5)} ${nat.slice(5)}`;
        case 'landline':
          if (parsed.areaCode === '2') {
            // 2 2123 4567
            return `${nat[0]} ${nat.slice(1, 5)} ${nat.slice(5)}`;
          }
          // 32 212 3456
          return `${nat.slice(0, 2)} ${nat.slice(2, 5)} ${nat.slice(5)}`;
        case 'voip':
          // 44 123 4567
          return `${nat.slice(0, 2)} ${nat.slice(2, 5)} ${nat.slice(5)}`;
        case 'toll_free':
          // 800 123 456
          return `${nat.slice(0, 3)} ${nat.slice(3, 6)} ${nat.slice(6)}`;
        case 'shared_cost':
          // 600 600 1234
          return `${nat.slice(0, 3)} ${nat.slice(3, 6)} ${nat.slice(6)}`;
        default:
          return nat;
      }

    case 'international':
    default:
      switch (parsed.type) {
        case 'mobile':
          // +56 9 1234 5678
          return `+56 ${nat[0]} ${nat.slice(1, 5)} ${nat.slice(5)}`;
        case 'landline':
          if (parsed.areaCode === '2') {
            // +56 2 2123 4567
            return `+56 ${nat[0]} ${nat.slice(1, 5)} ${nat.slice(5)}`;
          }
          // +56 32 212 3456
          return `+56 ${nat.slice(0, 2)} ${nat.slice(2, 5)} ${nat.slice(5)}`;
        case 'voip':
          // +56 44 123 4567
          return `+56 ${nat.slice(0, 2)} ${nat.slice(2, 5)} ${nat.slice(5)}`;
        case 'toll_free':
          // +56 800 123 456
          return `+56 ${nat.slice(0, 3)} ${nat.slice(3, 6)} ${nat.slice(6)}`;
        case 'shared_cost':
          // +56 600 600 1234
          return `+56 ${nat.slice(0, 3)} ${nat.slice(3, 6)} ${nat.slice(6)}`;
        default:
          return `+56 ${nat}`;
      }
  }
}

/**
 * Retorna la clasificación oficial del número telefónico chileno según la SUBTEL.
 * Detecta móviles, red fija, líneas 800, líneas 600 y números cortos de emergencia.
 *
 * @param phone - Cadena a analizar.
 * @returns Clasificación oficial ('mobile' | 'landline' | 'toll_free' | 'shared_cost' | 'emergency' | 'unknown').
 *
 * @example
 * ```ts
 * getPhoneType("+56 9 1234 5678"); // "mobile"
 * getPhoneType("221234567"); // "landline"
 * getPhoneType("133"); // "emergency"
 * ```
 */
export function getPhoneType(phone: string): PhoneType {
  if (typeof phone !== 'string') {
    return 'unknown';
  }

  const trimmed = phone.trim();

  // Verificar si corresponde a un número corto de emergencia (ej: 133, 131, etc.)
  let cleanDigits = '';
  let hasOnlyDigitsOrSeparators = true;
  for (let i = 0; i < trimmed.length; i++) {
    const ch = trimmed[i];
    if (ch >= '0' && ch <= '9') {
      cleanDigits += ch;
    } else if (ch !== ' ' && ch !== '-') {
      hasOnlyDigitsOrSeparators = false;
    }
  }

  if (hasOnlyDigitsOrSeparators && cleanDigits.length === 3) {
    if (EMERGENCY_NUMBERS.has(cleanDigits)) {
      return 'emergency';
    }
  }

  try {
    const parsed = parsePhone(phone);
    return parsed.type;
  } catch {
    return 'unknown';
  }
}

/**
 * Verifica si una cadena corresponde a un número corto de emergencia oficial en Chile.
 *
 * @param phone - Cadena a verificar.
 * @returns `true` si es un número de emergencia reconocido, `false` en caso contrario.
 *
 * @example
 * ```ts
 * isEmergencyNumber("133"); // true
 * isEmergencyNumber("123"); // false
 * ```
 */
export function isEmergencyNumber(phone: string): boolean {
  return getPhoneType(phone) === 'emergency';
}
