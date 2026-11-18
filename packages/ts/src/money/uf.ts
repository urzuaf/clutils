/**
 * @file uf.ts
 * @description Formateo de Unidades de Fomento (UF) segun la convencion de la CMF.
 */

import { formatThousands } from './clp.js';

export interface UFFormatOptions {
  /**
   * Cantidad de decimales a mostrar.
   * @default 2
   */
  decimals?: number;

  /**
   * Simbolo o prefijo de la unidad.
   * @default "UF"
   */
  symbol?: string;

  /**
   * Si es true, omite completamente el simbolo de la unidad.
   * @default false
   */
  noSymbol?: boolean;
}

/**
 * Formatea un valor en Unidades de Fomento (UF) segun la convencion oficial chilena,
 * utilizando puntos para separar los miles y coma para los decimales.
 *
 * @param amount - Valor numerico en UF.
 * @param options - Opciones de formateo.
 * @returns Cadena formateada (ej: "UF 38.450,25").
 *
 * @example
 * ```ts
 * formatUF(38450.25); // "UF 38.450,25"
 * formatUF(38450.2548, { decimals: 4 }); // "UF 38.450,2548"
 * formatUF(1.5, { noSymbol: true }); // "1,50"
 * ```
 */
export function formatUF(amount: number, options?: UFFormatOptions): string {
  const decimals = options?.decimals !== undefined && options.decimals > 0 ? options.decimals : 2;
  const symbol = options?.noSymbol ? '' : (options?.symbol ?? 'UF');

  const isNegative = amount < 0;
  const absAmount = Math.abs(amount);

  const formattedFloat = absAmount.toFixed(decimals);
  const dotIdx = formattedFloat.indexOf('.');
  const intPart = dotIdx !== -1 ? formattedFloat.slice(0, dotIdx) : formattedFloat;
  const decPart = dotIdx !== -1 ? formattedFloat.slice(dotIdx + 1) : '';

  const formattedInt = formatThousands(BigInt(intPart));
  const numStr = decimals > 0 && decPart.length > 0 ? `${formattedInt},${decPart}` : formattedInt;

  const prefix = symbol.length > 0 ? `${symbol} ` : '';
  const sign = isNegative ? '-' : '';

  return `${prefix}${sign}${numStr}`;
}
