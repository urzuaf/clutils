import { clean } from './validation.js';

/**
 * Opciones para personalizar la representación visual de un RUT.
 */
export interface FormatOptions {
  /** Si incluye puntos separadores de miles (por defecto true). */
  dots?: boolean;
  /** Si incluye guión antes del dígito verificador (por defecto true). */
  hyphen?: boolean;
}

/**
 * Representa la descomposición atómica de un RUT en cuerpo y dígito verificador.
 */
export interface RutParts {
  /** Cuerpo numérico en formato string para preservar ceros iniciales. */
  body: string;
  /** Dígito verificador normalizado en mayúscula ("0"-"9" o "K"). */
  dv: string;
  /**
   * Retorna el cuerpo convertido a número entero.
   * Lanza un error si contiene caracteres no numéricos.
   */
  bodyInt(): number;
}

/**
 * Opciones para ofuscar un RUT por privacidad.
 */
export interface MaskOptions {
  /** Carácter de reemplazo (por defecto '*'). */
  maskChar?: string;
  /** Cantidad de caracteres visibles al final (por defecto 4: últimos 3 del cuerpo + DV). */
  visibleCount?: number;
  /** Si conserva el formato formal con puntos y guión (por defecto true). */
  formatted?: boolean;
}

/**
 * Convierte una cadena de RUT en una representación visual estándar.
 * Si no se especifican opciones, aplica el formato formal estándar chileno (ej: "18.123.458-K").
 *
 * @param inputRut - Cadena de RUT en cualquier formato.
 * @param options - Opciones de formato (dots, hyphen).
 * @returns Cadena con el formato visual solicitado.
 *
 * @example
 * ```ts
 * format("18123458K"); // "18.123.458-K"
 * format("18123458K", { dots: false, hyphen: true }); // "18123458-K"
 * format("18123458K", { dots: false, hyphen: false }); // "18123458K"
 * ```
 */
export function format(inputRut: string, options?: FormatOptions): string {
  const cleaned = clean(inputRut);
  if (cleaned.length < 2) {
    return cleaned;
  }

  const dots = options?.dots ?? true;
  const hyphen = options?.hyphen ?? true;

  const lastIndex = cleaned.length - 1;
  const body = cleaned.slice(0, lastIndex);
  const dv = cleaned[lastIndex];

  let formattedBody = '';
  const bodyLen = body.length;

  for (let i = 0; i < bodyLen; i++) {
    if (dots && i > 0 && (bodyLen - i) % 3 === 0) {
      formattedBody += '.';
    }
    formattedBody += body[i];
  }

  const separator = hyphen ? '-' : '';
  return `${formattedBody}${separator}${dv}`;
}

/**
 * Divide un RUT en su cuerpo numérico y dígito verificador.
 * Acepta entradas en cualquier formato habitual (con o sin puntos, guiones o espacios).
 * Lanza un error si la cadena contiene menos de 2 caracteres.
 *
 * @param inputRut - Cadena con el RUT a dividir.
 * @returns Objeto RutParts con body, dv y el método helper bodyInt().
 *
 * @example
 * ```ts
 * const parts = split("18.123.458-K");
 * console.log(parts.body); // "18123458"
 * console.log(parts.dv);   // "K"
 * console.log(parts.bodyInt()); // 18123458
 * ```
 */
export function split(inputRut: string): RutParts {
  const cleaned = clean(inputRut);
  if (cleaned.length < 2) {
    throw new Error('el RUT no contiene cuerpo y dígito verificador suficientes');
  }

  const lastIndex = cleaned.length - 1;
  const body = cleaned.slice(0, lastIndex);
  const dv = cleaned[lastIndex];

  return {
    body,
    dv,
    bodyInt(): number {
      for (let i = 0; i < body.length; i++) {
        const code = body.charCodeAt(i);
        if (code < 48 || code > 57) {
          throw new Error('el cuerpo del RUT contiene caracteres no numéricos');
        }
      }
      const parsed = parseInt(body, 10);
      if (isNaN(parsed)) {
        throw new Error('el cuerpo del RUT contiene caracteres no numéricos');
      }
      return parsed;
    },
  };
}

/**
 * Ofusca los caracteres iniciales de un RUT para proteger la privacidad del titular
 * según la Ley 19.628 de Protección de la Vida Privada / Datos Personales,
 * manteniendo visibles los últimos N caracteres.
 *
 * @param inputRut - Cadena con el RUT a enmascarar.
 * @param options - Opciones de máscara (maskChar, visibleCount, formatted).
 * @returns Cadena con el RUT ofuscado.
 *
 * @example
 * ```ts
 * mask("18123458K"); // "**.***.458-K"
 * mask("18.123.458-k", { formatted: false }); // "*****458K"
 * mask("18123458K", { maskChar: 'X' }); // "XX.XXX.458-K"
 * ```
 */
export function mask(inputRut: string, options?: MaskOptions): string {
  const cleaned = clean(inputRut);
  if (cleaned.length < 2) {
    return cleaned;
  }

  const maskChar = options?.maskChar ?? '*';
  const visibleCount = Math.max(options?.visibleCount ?? 4, 0);
  const formatted = options?.formatted ?? true;

  if (!formatted) {
    if (visibleCount >= cleaned.length) {
      return cleaned;
    }
    const maskLen = Math.max(cleaned.length - visibleCount, 0);
    return maskChar.repeat(maskLen) + cleaned.slice(maskLen);
  }

  const formattedStr = format(cleaned);
  if (visibleCount >= cleaned.length) {
    return formattedStr;
  }

  const threshold = cleaned.length - visibleCount;
  let result = '';
  let seenAlphanumeric = 0;

  for (let i = 0; i < formattedStr.length; i++) {
    const char = formattedStr[i];
    if (char === '.' || char === '-') {
      result += char;
      continue;
    }

    if (seenAlphanumeric < threshold) {
      result += maskChar;
    } else {
      result += char;
    }
    seenAlphanumeric++;
  }

  return result;
}
