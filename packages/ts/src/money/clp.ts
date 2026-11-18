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
 * Formatea un entero con separador de miles '.' de forma lineal y sin expresiones regulares.
 */
export function formatThousands(val: number | bigint): string {
  const s = val.toString();
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

  const isNegative = amount < 0;
  const absAmount = isNegative ? (typeof amount === 'bigint' ? -amount : Math.abs(amount)) : amount;
  const numStr = formatThousands(absAmount);

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
 * @param text - Cadena con el monto formateado o sucio (ej: "$ 1.500.000", "($2.500)").
 * @returns Monto entero en pesos.
 * @throws {Error} Si el texto esta vacio, no contiene digitos o presenta formato invalido.
 */
export function parseCLP(text: string): number {
  const trimmed = text.trim();
  if (trimmed.length === 0) {
    throw new Error('El texto está vacío');
  }

  let clean = trimmed;
  const cleanUpper = clean.toUpperCase();
  if (cleanUpper.startsWith('CLP')) {
    clean = clean.slice(3).trim();
  } else if (cleanUpper.endsWith('CLP')) {
    clean = clean.slice(0, -3).trim();
  }

  let isNegative = false;
  let hasParen = false;
  let minusCount = 0;
  let openParens = 0;
  let closeParens = 0;
  let digits = '';

  for (let i = 0; i < clean.length; i++) {
    const ch = clean[i];
    if (ch >= '0' && ch <= '9') {
      if (closeParens > 0) {
        throw new Error('Dígito posterior al paréntesis de cierre');
      }
      digits += ch;
    } else if (ch === '.') {
      continue;
    } else if (ch === '$' || ch === ' ' || ch === '\t') {
      continue;
    } else if (ch === '-') {
      if (digits.length > 0) {
        throw new Error('Signo menos no puede ir después de dígitos');
      }
      minusCount++;
      isNegative = true;
    } else if (ch === '(') {
      if (digits.length > 0) {
        throw new Error('Paréntesis de apertura no puede ir después de dígitos');
      }
      openParens++;
    } else if (ch === ')') {
      if (openParens === 0) {
        throw new Error('Paréntesis de cierre sin apertura previa');
      }
      closeParens++;
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

  if (digits.length === 0) {
    throw new Error('El texto no contiene dígitos numéricos válidos');
  }

  let val = parseInt(digits, 10);
  if (Number.isNaN(val)) {
    throw new Error('Error al parsear el número entero');
  }

  if (isNegative || hasParen) {
    val = -val;
  }

  return val;
}
