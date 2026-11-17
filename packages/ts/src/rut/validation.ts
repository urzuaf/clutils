/**
 * Códigos numéricos de error para el diagnóstico de validación de un RUT chileno.
 */
export enum RutErrorCode {
  /** Sin error (RUT válido). */
  None = 0,
  /** El RUT está vacío o contiene solo espacios en blanco. */
  Empty = 1,
  /** Puntos o guiones mal ubicados, repetidos o en exceso. */
  MalformedFormat = 2,
  /** Longitud fuera del rango permitido (debe tener entre 1 y 9 dígitos de cuerpo + DV). */
  InvalidLength = 3,
  /** El cuerpo contiene letras o caracteres no numéricos. */
  InvalidCharacters = 4,
  /** El dígito verificador no coincide según el algoritmo de Módulo 11. */
  InvalidDv = 5,
}

/**
 * Representa el resultado detallado de un error de validación de RUT.
 */
export interface ValidationError {
  /** Código numérico constante para serialización o i18n. */
  code: RutErrorCode;
  /** Mensaje descriptivo en español del motivo del fallo. */
  message: string;
}

/**
 * Sanitiza una cadena de RUT eliminando puntos, guiones y espacios en blanco,
 * y normalizando el dígito verificador 'k' a mayúscula 'K'.
 *
 * @param inputRut - Cadena de entrada con el RUT en cualquier formato.
 * @returns Cadena normalizada continua en mayúscula (ej: "18123458K").
 *
 * @example
 * ```ts
 * clean("  18.123.458-k  "); // "18123458K"
 * ```
 */
export function clean(inputRut: string): string {
  if (!inputRut) return '';

  let cleaned = '';
  for (let i = 0; i < inputRut.length; i++) {
    const char = inputRut[i];
    if (char === '.' || char === '-' || char === ' ' || char === '\t' || char === '\n' || char === '\r') {
      continue;
    }
    if (char === 'k') {
      cleaned += 'K';
      continue;
    }
    cleaned += char;
  }

  return cleaned;
}

/**
 * Calcula el Dígito Verificador oficial para un cuerpo numérico
 * aplicando el algoritmo de Módulo 11 con ponderadores cíclicos [2..7].
 *
 * @param body - Cuerpo numérico del RUT como entero o cadena numérica.
 * @returns El carácter del dígito verificador ("0"-"9" o "K"), o cadena vacía si es inválido.
 *
 * @example
 * ```ts
 * calculateDv(18123458); // "K"
 * calculateDv(18123466); // "0"
 * calculateDv("12345678"); // "5"
 * ```
 */
export function calculateDv(body: number | string): string {
  const num = typeof body === 'number' ? body : parseInt(body, 10);
  if (isNaN(num) || num <= 0) {
    return '';
  }

  let sum = 0;
  let multiplier = 2;

  for (let n = num; n > 0; n = Math.floor(n / 10)) {
    const digit = n % 10;
    sum += digit * multiplier;

    multiplier++;
    if (multiplier > 7) {
      multiplier = 2;
    }
  }

  const dv = 11 - (sum % 11);

  if (dv === 11) return '0';
  if (dv === 10) return 'K';
  return String(dv);
}

/**
 * Valida exhaustivamente un RUT y retorna null si es válido,
 * o un objeto ValidationError con el código numérico y motivo exacto del fallo.
 *
 * @param inputRut - Cadena con el RUT a evaluar.
 * @returns ValidationError con el detalle del error o null si es válido.
 *
 * @example
 * ```ts
 * const err = validateDetailed("18.12A.458-K");
 * if (err) {
 *   console.log(err.code); // RutErrorCode.InvalidCharacters (4)
 * }
 * ```
 */
export function validateDetailed(inputRut: string): ValidationError | null {
  if (!inputRut || inputRut.trim().length === 0) {
    return {
      code: RutErrorCode.Empty,
      message: 'el RUT no puede estar vacío',
    };
  }

  // 1. Control de separadores repetidos o malformados
  if (inputRut.includes('--') || inputRut.includes('..')) {
    return {
      code: RutErrorCode.MalformedFormat,
      message: 'el RUT contiene puntos o guiones repetidos',
    };
  }

  let hyphenCount = 0;
  let dotCount = 0;
  for (let i = 0; i < inputRut.length; i++) {
    if (inputRut[i] === '-') hyphenCount++;
    if (inputRut[i] === '.') dotCount++;
  }

  if (hyphenCount > 1 || dotCount > 2) {
    return {
      code: RutErrorCode.MalformedFormat,
      message: 'el RUT excede la cantidad máxima de puntos o guiones',
    };
  }

  const hyphenIdx = inputRut.indexOf('-');
  if (hyphenIdx !== -1) {
    const afterHyphen = clean(inputRut.slice(hyphenIdx + 1));
    if (afterHyphen.length !== 1) {
      return {
        code: RutErrorCode.MalformedFormat,
        message: 'el guión del RUT debe separar únicamente el dígito verificador',
      };
    }
  }

  const cleaned = clean(inputRut);

  // 2. Control de longitud
  if (cleaned.length < 2 || cleaned.length > 10) {
    return {
      code: RutErrorCode.InvalidLength,
      message: 'la longitud del RUT es inválida',
    };
  }

  const lastIndex = cleaned.length - 1;
  const bodyStr = cleaned.slice(0, lastIndex);
  const dv = cleaned[lastIndex];

  // 3. Control de cuerpo numérico
  for (let i = 0; i < bodyStr.length; i++) {
    const code = bodyStr.charCodeAt(i);
    if (code < 48 || code > 57) {
      return {
        code: RutErrorCode.InvalidCharacters,
        message: 'el cuerpo del RUT contiene caracteres no numéricos',
      };
    }
  }

  const numBody = parseInt(bodyStr, 10);

  // 4. Control de Dígito Verificador Módulo 11
  if (calculateDv(numBody) !== dv) {
    return {
      code: RutErrorCode.InvalidDv,
      message: 'el dígito verificador no coincide',
    };
  }

  return null;
}

/**
 * Verifica si una cadena representa un RUT formal y matemáticamente válido según la normativa chilena.
 * Retorna true si es válido, o false si está malformado, vacío o su DV es incorrecto.
 *
 * @param inputRut - Cadena con el RUT a validar.
 * @returns true si el RUT es válido, false en caso contrario.
 *
 * @example
 * ```ts
 * validate("12.345.678-5"); // true
 * validate("18.123.458-k"); // true
 * validate("12.345.678-0"); // false
 * ```
 */
export function validate(inputRut: string): boolean {
  return validateDetailed(inputRut) === null;
}
