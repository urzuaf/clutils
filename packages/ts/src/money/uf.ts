/**
 * @file uf.ts
 * @description Formateo de Unidades de Fomento (UF) segun la convencion de la CMF.
 */

import { formatThousandsStr } from './clp.js';

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

function isAllZeros(s: string): boolean {
  for (let i = 0; i < s.length; i++) {
    if (s[i] !== '0') return false;
  }
  return true;
}

/**
 * Formatea un valor en Unidades de Fomento (UF) segun la convencion oficial chilena,
 * utilizando puntos para separar los miles, coma para los decimales y redondeo estandar
 * aritmetico (round half away from zero) consistente con Go.
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
  const decimals = options?.decimals !== undefined && options.decimals > 0 ? Math.trunc(options.decimals) : 2;
  const symbol = options?.noSymbol ? '' : (options?.symbol ?? 'UF');

  const safeAmount = Number.isFinite(amount) ? amount : 0;
  let isNegative = safeAmount < 0;
  const absAmount = Math.abs(safeAmount);

  let intPart: string;
  let decPart: string;

  const factor = Math.pow(10, decimals);
  const scaled = absAmount * factor;

  if (decimals <= 12 && Number.isFinite(scaled) && scaled <= 9e15) {
    const roundedInt = Math.round(scaled);
    const rawDigits = roundedInt.toString();
    if (rawDigits.length <= decimals) {
      intPart = '0';
      decPart = '0'.repeat(decimals - rawDigits.length) + rawDigits;
    } else {
      const splitPos = rawDigits.length - decimals;
      intPart = rawDigits.slice(0, splitPos);
      decPart = rawDigits.slice(splitPos);
    }
  } else if (absAmount >= 1e21) {
    // Números >= 1e21 en float64 ya son enteros exactos (sin fracción)
    intPart = BigInt(absAmount).toString();
    decPart = '0'.repeat(decimals);
  } else {
    const formattedFloat = absAmount.toFixed(decimals);
    const dotIdx = formattedFloat.indexOf('.');
    intPart = dotIdx !== -1 ? formattedFloat.slice(0, dotIdx) : formattedFloat;
    decPart = dotIdx !== -1 ? formattedFloat.slice(dotIdx + 1) : '';
  }

  // Evitar "-0,00" cuando un valor negativo pequeño redondea a cero
  if (isNegative && isAllZeros(intPart) && isAllZeros(decPart)) {
    isNegative = false;
  }

  const formattedInt = formatThousandsStr(intPart);
  const numStr = decimals > 0 && decPart.length > 0 ? `${formattedInt},${decPart}` : formattedInt;

  const prefix = symbol.length > 0 ? `${symbol} ` : '';
  const sign = isNegative ? '-' : '';

  return `${prefix}${sign}${numStr}`;
}
