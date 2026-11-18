/**
 * @file utm.ts
 * @description Formateo de Unidades Tributarias Mensuales (UTM).
 */

import { formatThousands } from './clp.js';

export interface UTMFormatOptions {
  /**
   * Simbolo o prefijo de la unidad.
   * @default "UTM"
   */
  symbol?: string;

  /**
   * Si es true, omite completamente el simbolo de la unidad.
   * @default false
   */
  noSymbol?: boolean;
}

/**
 * Formatea un monto entero en Unidades Tributarias Mensuales (UTM),
 * utilizando puntos para separar los miles y sin decimales.
 *
 * @param amount - Monto entero en UTM.
 * @param options - Opciones de formateo.
 * @returns Cadena formateada (ej: "UTM 67.294").
 *
 * @example
 * ```ts
 * formatUTM(67294); // "UTM 67.294"
 * formatUTM(67294, { noSymbol: true }); // "67.294"
 * ```
 */
export function formatUTM(amount: number | bigint, options?: UTMFormatOptions): string {
  const symbol = options?.noSymbol ? '' : (options?.symbol ?? 'UTM');

  const isNegative = amount < 0;
  const absAmount = isNegative ? (typeof amount === 'bigint' ? -amount : Math.abs(amount)) : amount;
  const numStr = formatThousands(absAmount);

  const prefix = symbol.length > 0 ? `${symbol} ` : '';
  const sign = isNegative ? '-' : '';

  return `${prefix}${sign}${numStr}`;
}
