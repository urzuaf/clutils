/**
 * @file utm.ts
 * @description Formateo de Unidades Tributarias Mensuales (UTM).
 */

import { formatThousandsStr } from './clp.js';

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

  let isNegative: boolean;
  let absDigits: string;
  if (typeof amount === 'bigint') {
    isNegative = amount < 0n;
    absDigits = (isNegative ? -amount : amount).toString();
  } else if (!Number.isFinite(amount)) {
    isNegative = false;
    absDigits = '0';
  } else {
    const rounded = Math.sign(amount) * Math.round(Math.abs(amount));
    isNegative = rounded < 0;
    const absVal = Math.abs(rounded);
    absDigits = absVal >= 1e21 ? BigInt(absVal).toString() : absVal.toString();
  }

  const numStr = formatThousandsStr(absDigits);

  const prefix = symbol.length > 0 ? `${symbol} ` : '';
  const sign = isNegative ? '-' : '';

  return `${prefix}${sign}${numStr}`;
}
