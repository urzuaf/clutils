/**
 * @file iva.ts
 * @description Calculo y desglose de IVA (19%) bajo la regla de cuadratura del SII.
 */

/** Tasa legal estandar de IVA en Chile (19%). */
export const DEFAULT_IVA_RATE = 0.19;

/** Desglose contable de una transaccion comercial. */
export interface IVABreakdown {
  /** Monto base antes de impuestos. */
  neto: number;
  /** Monto del impuesto correspondiente. */
  iva: number;
  /** Monto total final que incluye impuesto. */
  bruto: number;
}

export interface IVAOptions {
  /**
   * Tasa decimal del impuesto.
   * @default 0.19
   */
  rate?: number;
}

/**
 * Desglosa un monto en Neto, IVA y Bruto aplicando la regla de cuadratura oficial del SII
 * garantizando que Neto + IVA === Bruto sin descuadre por redondeo de punto flotante.
 *
 * @param amount - Monto monetario entero.
 * @param isBruto - Si true, el monto ingresado ya contiene impuesto (Bruto). Si false, es el valor base (Neto).
 * @param options - Opciones adicionales (ej. tasa personalizada).
 * @returns Desglose contable { neto, iva, bruto }.
 *
 * @example
 * ```ts
 * calculateIVA(1000, false); // { neto: 1000, iva: 190, bruto: 1190 }
 * calculateIVA(1000, true);  // { neto: 840, iva: 160, bruto: 1000 }
 * ```
 */
/**
 * Redondea un numero alejandose del cero (estandar contable simetrico / round half away from zero).
 */
function roundHalfAway(val: number): number {
  const rounded = Math.sign(val) * Math.round(Math.abs(val));
  return rounded === 0 ? 0 : rounded;
}

export function calculateIVA(amount: number, isBruto: boolean, options?: IVAOptions): IVABreakdown {
  const rate =
    options?.rate !== undefined && Number.isFinite(options.rate) && options.rate > 0
      ? options.rate
      : DEFAULT_IVA_RATE;

  if (!Number.isFinite(amount) || amount === 0) {
    return { neto: 0, iva: 0, bruto: 0 };
  }

  if (isBruto) {
    const bruto = amount;
    const neto = roundHalfAway(bruto / (1 + rate));
    const iva = bruto - neto;
    return { neto, iva: iva === 0 ? 0 : iva, bruto };
  }

  const neto = amount;
  const iva = roundHalfAway(neto * rate);
  const bruto = neto + iva;
  return { neto, iva, bruto: bruto === 0 ? 0 : bruto };
}
