import { describe, it, expect } from 'vitest';
import * as fs from 'node:fs';
import * as path from 'node:path';
import { calculateIVA } from '../../src/money/iva.js';

interface IVACalculationCase {
  input_amount: number;
  is_bruto: boolean;
  expected_neto: number;
  expected_iva: number;
  expected_bruto: number;
  description: string;
}

interface IVATestSuite {
  iva_calculation: IVACalculationCase[];
}

function loadSuite(): IVATestSuite {
  const filePath = path.resolve(__dirname, '../../../../specs/tests/money_formatting.json');
  const content = fs.readFileSync(filePath, 'utf-8');
  return JSON.parse(content);
}

describe('IVA Calculation Suite (specs/tests/money_formatting.json)', () => {
  const suite = loadSuite();

  for (const tc of suite.iva_calculation) {
    it(`calculateIVA: ${tc.description}`, () => {
      const result = calculateIVA(tc.input_amount, tc.is_bruto);

      expect(result.neto).toBe(tc.expected_neto);
      expect(result.iva).toBe(tc.expected_iva);
      expect(result.bruto).toBe(tc.expected_bruto);

      // Invariante sagrado del SII
      expect(result.neto + result.iva).toBe(result.bruto);
    });
  }

  it('calculateIVA con tasa personalizada', () => {
    const resNeto = calculateIVA(1000, false, { rate: 0.10 });
    expect(resNeto).toEqual({ neto: 1000, iva: 100, bruto: 1100 });

    const resBruto = calculateIVA(1100, true, { rate: 0.10 });
    expect(resBruto).toEqual({ neto: 1000, iva: 100, bruto: 1100 });
  });

  it('evita -0 en IVA de montos negativos pequeños', () => {
    const res = calculateIVA(-1, false);
    expect(Object.is(res.iva, 0)).toBe(true);
    expect(res).toEqual({ neto: -1, iva: 0, bruto: -1 });
  });
});
