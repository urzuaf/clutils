/**
 * @file words.ts
 * @description Conversion de montos numericos a palabras formales en espanol chileno.
 */

export interface ToWordsOptions {
  /**
   * Si es true, omite la denominacion monetaria ("peso", "pesos", "de pesos").
   * @default false
   */
  noCurrency?: boolean;
}

const unitsWords = [
  '', 'un', 'dos', 'tres', 'cuatro', 'cinco', 'seis', 'siete', 'ocho', 'nueve',
  'diez', 'once', 'doce', 'trece', 'catorce', 'quince', 'dieciséis', 'diecisiete', 'dieciocho', 'diecinueve',
  'veinte', 'veintiún', 'veintidós', 'veintitrés', 'veinticuatro', 'veinticinco', 'veintiséis', 'veintisiete', 'veintiocho', 'veintinueve',
];

const tensWords = [
  '', '', '', 'treinta', 'cuarenta', 'cincuenta', 'sesenta', 'setenta', 'ochenta', 'noventa',
];

const hundredsWords = [
  '', 'ciento', 'doscientos', 'trescientos', 'cuatrocientos', 'quinientos', 'seiscientos', 'setecientos', 'ochocientos', 'novecientos',
];

/**
 * Convierte un entero entre 1 y 999 a palabras.
 */
function convertGroup(n: number): string {
  if (n === 0) return '';
  if (n === 100) return 'cien';

  const parts: string[] = [];
  const c = Math.floor(n / 100);
  const d = Math.floor((n % 100) / 10);
  const u = n % 10;

  if (c > 0) {
    parts.push(hundredsWords[c]);
  }

  const du = n % 100;
  if (du > 0) {
    if (du < 30) {
      parts.push(unitsWords[du]);
    } else {
      if (u === 0) {
        parts.push(tensWords[d]);
      } else {
        parts.push(`${tensWords[d]} y ${unitsWords[u]}`);
      }
    }
  }

  return parts.join(' ');
}

/**
 * Convierte un bigint a palabras en espanol segun la escala larga.
 */
function numberToWords(val: bigint): string {
  if (val === 0n) return 'cero';
  if (val === 1n) return 'un';

  const billones = val / 1_000_000_000_000n;
  const remBillones = val % 1_000_000_000_000n;

  const millonesTotal = remBillones / 1_000_000n;
  const remMillones = remBillones % 1_000_000n;

  const miles = remMillones / 1_000n;
  const unidades = remMillones % 1_000n;

  const parts: string[] = [];

  // 1. Billones (10^12)
  if (billones > 0n) {
    if (billones === 1n) {
      parts.push('un billón');
    } else {
      parts.push(`${convertGroup(Number(billones))} billones`);
    }
  }

  // 2. Millones (hasta 999.999 millones)
  if (millonesTotal > 0n) {
    if (millonesTotal === 1n) {
      parts.push('un millón');
    } else {
      const milesDeMillones = millonesTotal / 1_000n;
      const restoMillones = millonesTotal % 1_000n;

      const millParts: string[] = [];
      if (milesDeMillones > 0n) {
        if (milesDeMillones === 1n) {
          millParts.push('mil');
        } else {
          millParts.push(convertGroup(Number(milesDeMillones)), 'mil');
        }
      }
      if (restoMillones > 0n) {
        millParts.push(convertGroup(Number(restoMillones)));
      }
      parts.push(`${millParts.join(' ')} millones`);
    }
  }

  // 3. Miles
  if (miles > 0n) {
    if (miles === 1n) {
      parts.push('mil');
    } else {
      parts.push(`${convertGroup(Number(miles))} mil`);
    }
  }

  // 4. Unidades
  if (unidades > 0n) {
    parts.push(convertGroup(Number(unidades)));
  }

  return parts.join(' ');
}

/**
 * Convierte un monto en pesos chilenos a su representacion textual formal en espanol,
 * respetando las reglas ortograficas, gramaticales y bancarias chilenas (cheques, pagares y contratos).
 *
 * @param amount - Monto entero en pesos chilenos.
 * @param options - Opciones de conversion.
 * @returns Cadena en palabras (ej: "un millón quinientos mil pesos").
 *
 * @example
 * ```ts
 * toWords(1500000); // "un millón quinientos mil pesos"
 * toWords(2000000); // "dos millones de pesos"
 * toWords(1);       // "un peso"
 * toWords(100, { noCurrency: true }); // "cien"
 * ```
 */
export function toWords(amount: number | bigint, options?: ToWordsOptions): string {
  const noCurrency = options?.noCurrency ?? false;
  const isNegative = amount < 0;
  const absVal = isNegative ? (typeof amount === 'bigint' ? -amount : BigInt(-amount)) : BigInt(amount);

  const text = numberToWords(absVal);

  if (noCurrency) {
    return isNegative ? `menos ${text}` : text;
  }

  let currencySuffix: string;
  if (absVal === 1n) {
    currencySuffix = 'peso';
  } else if (absVal >= 1_000_000n && absVal % 1_000_000n === 0n) {
    currencySuffix = 'de pesos';
  } else {
    currencySuffix = 'pesos';
  }

  return isNegative ? `menos ${text} ${currencySuffix}` : `${text} ${currencySuffix}`;
}
