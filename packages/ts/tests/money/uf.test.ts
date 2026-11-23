import { describe, it, expect } from 'vitest';
import * as fs from 'node:fs';
import * as path from 'node:path';
import { formatUF } from '../../src/money/uf.js';

interface UFFormattingCase {
  amount: number;
  decimals: number;
  expected_default: string;
  expected_no_symbol: string;
  description: string;
}

interface UFTestSuite {
  uf_formatting: UFFormattingCase[];
}

function loadSuite(): UFTestSuite {
  const filePath = path.resolve(__dirname, '../../../../specs/tests/money_formatting.json');
  const content = fs.readFileSync(filePath, 'utf-8');
  return JSON.parse(content);
}

describe('UF Formatting Suite (specs/tests/money_formatting.json)', () => {
  const suite = loadSuite();

  for (const tc of suite.uf_formatting) {
    it(`formatUF (default): ${tc.description}`, () => {
      expect(formatUF(tc.amount, { decimals: tc.decimals })).toBe(tc.expected_default);
    });

    it(`formatUF (no symbol): ${tc.description}`, () => {
      expect(formatUF(tc.amount, { decimals: tc.decimals, noSymbol: true })).toBe(tc.expected_no_symbol);
    });
  }

  it('maneja NaN, Infinity y valores muy grandes (>= 1e20) sin lanzar error', () => {
    expect(formatUF(NaN)).toBe('UF 0,00');
    expect(formatUF(Infinity)).toBe('UF 0,00');
    expect(formatUF(-Infinity)).toBe('UF 0,00');
    expect(formatUF(1e20)).toBe('UF 100.000.000.000.000.000.000,00');
    expect(formatUF(1e21)).toBe('UF 1.000.000.000.000.000.000.000,00');
  });
});
