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

  const trillones = val / 1_000_000_000_000_000_000n;
  const remTrillones = val % 1_000_000_000_000_000_000n;

  const billonesTotal = remTrillones / 1_000_000_000_000n;
  const remBillones = remTrillones % 1_000_000_000_000n;

  const millonesTotal = remBillones / 1_000_000n;
  const remMillones = remBillones % 1_000_000n;

  const miles = remMillones / 1_000n;
  const unidades = remMillones % 1_000n;

  const parts: string[] = [];

  // 1. Trillones (10^18)
  if (trillones > 0n) {
    if (trillones === 1n) {
      parts.push('un trillón');
    } else if (trillones < 1000n) {
      parts.push(`${convertGroup(Number(trillones))} trillones`);
    } else {
      parts.push(`${numberToWords(trillones)} trillones`);
    }
  }

  // 2. Billones (10^12)
  if (billonesTotal > 0n) {
    if (billonesTotal === 1n) {
      parts.push('un billón');
    } else {
      const milesDeBillones = billonesTotal / 1_000n;
      const restoBillones = billonesTotal % 1_000n;

      const billParts: string[] = [];
      if (milesDeBillones > 0n) {
        if (milesDeBillones === 1n) {
          billParts.push('mil');
        } else {
          billParts.push(convertGroup(Number(milesDeBillones)), 'mil');
        }
      }
      if (restoBillones > 0n) {
        billParts.push(convertGroup(Number(restoBillones)));
      }
      parts.push(`${billParts.join(' ')} billones`);
    }
  }

  // 3. Millones (10^6)
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

  // 4. Miles (10^3)
  if (miles > 0n) {
    if (miles === 1n) {
      parts.push('mil');
    } else {
      parts.push(`${convertGroup(Number(miles))} mil`);
    }
  }

  // 5. Unidades (10^0)
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

  let normalized: bigint;
  if (typeof amount === 'bigint') {
    normalized = amount;
  } else if (!Number.isFinite(amount)) {
    normalized = 0n;
  } else {
    const rounded = Math.sign(amount) * Math.round(Math.abs(amount));
    normalized = BigInt(rounded);
  }

  const isNegative = normalized < 0n;
  const absVal = isNegative ? -normalized : normalized;

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
