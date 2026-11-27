import { calculateDv } from './validation.js';
import { format, type FormatOptions } from './formatting.js';
import {
  DEFAULT_GENERAL_MIN,
  DEFAULT_GENERAL_MAX,
  DEFAULT_PERSON_MIN,
  DEFAULT_PERSON_MAX,
  DEFAULT_BUSINESS_MIN,
  DEFAULT_BUSINESS_MAX,
} from './ranges.js';

export * from './ranges.js';

/**
 * Genera un número entero aleatorio en el rango [min, max] inclusive.
 */
function randomInt(min: number, max: number): number {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

/**
 * Genera un RUT aleatorio y matemáticamente válido con formato estándar.
 * Por defecto utiliza el rango seguro habitual de personas naturales.
 *
 * @param options - Opciones de formato visual.
 * @returns Cadena con un RUT válido garantizado.
 *
 * @example
 * ```ts
 * generate(); // ej: "18.123.458-K"
 * generate({ dots: false, hyphen: true }); // ej: "18123458-K"
 * ```
 */
export function generate(options?: FormatOptions): string {
  return generateBetween(DEFAULT_GENERAL_MIN, DEFAULT_GENERAL_MAX, options);
}

/**
 * Genera un RUT matemáticamente válido dentro de un rango numérico específico.
 * Corrige de forma automática valores negativos o límites invertidos para garantizar seguridad total.
 *
 * @param min - Límite numérico inferior.
 * @param max - Límite numérico superior.
 * @param options - Opciones de formato visual.
 * @returns Cadena con el RUT generado.
 *
 * @example
 * ```ts
 * generateBetween(1_000_000, 10_000_000); // RUT bajo/antiguo
 * generateBetween(50_000_000, 99_000_000); // RUT en rango de personas jurídicas
 * ```
 */
export function generateBetween(min: number, max: number, options?: FormatOptions): string {
  const safeMin = Number.isFinite(min) ? Math.trunc(min) : DEFAULT_GENERAL_MIN;
  const safeMax = Number.isFinite(max) ? Math.trunc(max) : DEFAULT_GENERAL_MAX;
  let lower = safeMin < DEFAULT_GENERAL_MIN ? DEFAULT_GENERAL_MIN : safeMin;
  let upper = safeMax < DEFAULT_GENERAL_MIN ? DEFAULT_GENERAL_MAX : safeMax;

  if (lower > upper) {
    const temp = lower;
    lower = upper;
    upper = temp;
  }

  const body = randomInt(lower, upper);
  const dv = calculateDv(body);
  const rut = `${body}${dv}`;

  return format(rut, options);
}

/**
 * Genera un RUT aleatorio representativo de personas naturales vivas en Chile (rango 5M a 28M).
 *
 * @param options - Opciones de formato visual.
 * @returns Cadena con el RUT generado.
 *
 * @example
 * ```ts
 * generatePerson(); // ej: "21.458.123-5"
 * ```
 */
export function generatePerson(options?: FormatOptions): string {
  return generateBetween(DEFAULT_PERSON_MIN, DEFAULT_PERSON_MAX, options);
}

/**
 * Genera un RUT aleatorio representativo de personas jurídicas y empresas en Chile (rango 70M a 80M).
 *
 * @param options - Opciones de formato visual.
 * @returns Cadena con el RUT de empresa generado.
 *
 * @example
 * ```ts
 * generateBusiness(); // ej: "76.458.123-0"
 * ```
 */
export function generateBusiness(options?: FormatOptions): string {
  return generateBetween(DEFAULT_BUSINESS_MIN, DEFAULT_BUSINESS_MAX, options);
}
