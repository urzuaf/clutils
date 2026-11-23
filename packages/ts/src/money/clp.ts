/**
 * @file clp.ts
 * @description Formateo y parseo de pesos chilenos (CLP) segun las convenciones locales.
 */

export type NegativeFormat = '-sign' | 'parentheses';

export interface CLPFormatOptions {
  /**
   * Simbolo o prefijo de la moneda.
   * @default "$"
   */
  symbol?: string;

  /**
   * Si es true, omite completamente el simbolo de la moneda.
   * @default false
   */
  noSymbol?: boolean;

  /**
   * Si incluye un espacio entre el simbolo y el monto (ej: "$ 1.500.000").
   * @default false
   */
  space?: boolean;

  /**
   * Estilo de representacion para montos negativos.
   * @default "-sign"
   */
  negativeFormat?: NegativeFormat;
}

/**
 * Formatea una cadena de dígitos enteros positivos con separador de miles '.' de forma lineal y sin regex.
 */
export function formatThousandsStr(s: string): string {
  const n = s.length;
  if (n <= 3) return s;

  const dots = Math.floor((n - 1) / 3);
  const res = new Array<string>(n + dots);
  let j = res.length - 1;
  let count = 0;

  for (let i = n - 1; i >= 0; i--) {
    res[j--] = s[i];
    count++;
    if (count === 3 && i > 0) {
      res[j--] = '.';
      count = 0;
    }
  }

  return res.join('');
}

/**
 * Formatea un entero positivo con separador de miles '.' de forma lineal y sin expresiones regulares.
 */
export function formatThousands(val: number | bigint): string {
  return formatThousandsStr(val.toString());
}

/**
 * Normaliza un monto numérico (entero o decimal) a su representación entera positiva y signo.
 */
function normalizeIntegerAmount(amount: number | bigint): { isNegative: boolean; absDigits: string } {
  if (typeof amount === 'bigint') {
    const isNegative = amount < 0n;
    const absVal = isNegative ? -amount : amount;
    return { isNegative, absDigits: absVal.toString() };
  }

  if (!Number.isFinite(amount)) {
    return { isNegative: false, absDigits: '0' };
  }

  const rounded = Math.sign(amount) * Math.round(Math.abs(amount));
  const isNegative = rounded < 0;
  const absVal = Math.abs(rounded);
  // Si el número es muy grande (>= 1e21), expandir notación científica con BigInt seguro
  const absDigits = absVal >= 1e21 ? BigInt(absVal).toString() : absVal.toString();
  return { isNegative, absDigits };
}

/**
 * Formatea un valor entero a pesos chilenos (CLP) segun la convencion oficial,
 * utilizando puntos para separar los miles y sin centavos.
 *
 * @param amount - Monto en pesos chilenos (entero).
 * @param options - Opciones de formateo.
 * @returns Cadena formateada (ej: "$1.500.000").
 *
 * @example
 * ```ts
 * formatCLP(1500000); // "$1.500.000"
 * formatCLP(1500000, { symbol: "CLP", space: true }); // "CLP 1.500.000"
 * formatCLP(-2500, { negativeFormat: "parentheses" }); // "($2.500)"
 * ```
 */
export function formatCLP(amount: number | bigint, options?: CLPFormatOptions): string {
  const symbol = options?.noSymbol ? '' : (options?.symbol ?? '$');
  const space = options?.space ?? false;
  const negativeFormat = options?.negativeFormat ?? '-sign';

  const { isNegative, absDigits } = normalizeIntegerAmount(amount);
  const numStr = formatThousandsStr(absDigits);

  const spaceStr = space && symbol.length > 0 ? ' ' : '';

  if (!isNegative) {
    return `${symbol}${spaceStr}${numStr}`;
  }

  if (negativeFormat === 'parentheses') {
    return `(${symbol}${spaceStr}${numStr})`;
  }

  return `-${symbol}${spaceStr}${numStr}`;
}

/**
 * Extrae y convierte cualquier representacion en texto de pesos chilenos a un entero limpio,
 * reconociendo signos negativos estandar y contables, y rechazando entradas malformadas.
 *
 * @param text - Cadena con el monto formateado o sucio (ej: "$ 1.500.000", "($2.500)", "-CLP 2.500").
 * @returns Monto entero en pesos.
 * @throws {Error} Si el texto esta vacio, no contiene digitos o presenta formato invalido.
 */
export function parseCLP(text: string): number {
  const trimmed = text.trim();
  if (trimmed.length === 0) {
    throw new Error('El texto está vacío');
  }

  let isNegative = false;
  let hasParen = false;
  let digitsStarted = false;
  let digitsEnded = false;
  let minusCount = 0;
  let dollarCount = 0;
  let clpCount = 0;
  let openParens = 0;
  let closeParens = 0;
  let rawNum = '';

  for (let i = 0; i < trimmed.length; i++) {
    const ch = trimmed[i];

    // Reconocer token "CLP" (prefijo, sufijo o tras signo/paréntesis)
    if (
      (ch === 'C' || ch === 'c') &&
      i + 3 <= trimmed.length &&
      trimmed.slice(i, i + 3).toUpperCase() === 'CLP'
    ) {
      clpCount++;
      if (clpCount > 1) {
        throw new Error('Múltiples prefijos CLP inválidos');
      }
      if (digitsStarted) {
        digitsEnded = true;
      }
      i += 2;
      continue;
    }

    if (ch >= '0' && ch <= '9') {
      if (closeParens > 0) {
        throw new Error('Dígito posterior al paréntesis de cierre');
      }
      if (digitsEnded) {
        throw new Error('Formato numérico discontinuo inválido');
      }
      digitsStarted = true;
      rawNum += ch;
    } else if (ch === '.') {
      if (closeParens > 0 || digitsEnded) {
        throw new Error('Separador de miles en posición inválida');
      }
      digitsStarted = true;
      rawNum += ch;
    } else if (ch === '$') {
      if (digitsStarted) {
        throw new Error("Símbolo '$' no puede ir después de dígitos");
      }
      dollarCount++;
      if (dollarCount > 1) {
        throw new Error("Múltiples símbolos '$' inválidos");
      }
    } else if (ch === ' ' || ch === '\t') {
      if (digitsStarted) {
        digitsEnded = true;
      }
    } else if (ch === '-') {
      if (digitsStarted) {
        throw new Error('Signo menos no puede ir después de dígitos');
      }
      minusCount++;
      isNegative = true;
    } else if (ch === '(') {
      if (digitsStarted) {
        throw new Error('Paréntesis de apertura no puede ir después de dígitos');
      }
      openParens++;
    } else if (ch === ')') {
      if (openParens === 0) {
        throw new Error('Paréntesis de cierre sin apertura previa');
      }
      closeParens++;
      if (digitsStarted) {
        digitsEnded = true;
      }
    } else {
      throw new Error(`Carácter no válido en monto CLP: "${ch}"`);
    }
  }

  if (openParens > 0 || closeParens > 0) {
    if (openParens !== 1 || closeParens !== 1) {
      throw new Error('Paréntesis contables desbalanceados');
    }
    if (minusCount > 0) {
      throw new Error("Formato negativo ambiguo (combina '-' y paréntesis)");
    }
    hasParen = true;
  }

  if (minusCount > 1) {
    throw new Error('Múltiples signos menos inválidos');
  }

  if (rawNum.length === 0) {
    throw new Error('El texto no contiene dígitos numéricos válidos');
  }

  let digits: string;
  if (rawNum.includes('.')) {
    const groups = rawNum.split('.');
    if (groups[0].length < 1 || groups[0].length > 3) {
      throw new Error(`Separadores de miles mal ubicados: "${rawNum}"`);
    }
    for (let i = 1; i < groups.length; i++) {
      if (groups[i].length !== 3) {
        throw new Error(`Separadores de miles mal ubicados: "${rawNum}"`);
      }
    }
    digits = groups.join('');
  } else {
    digits = rawNum;
  }

  let val = parseInt(digits, 10);
  if (Number.isNaN(val) || !Number.isSafeInteger(val)) {
    throw new Error('Desborde o error al parsear el número entero');
  }

  if (isNegative || hasParen) {
    val = -val;
  }

  return val === 0 ? 0 : val;
}
