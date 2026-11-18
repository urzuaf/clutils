import { describe, it, expect } from 'vitest';
import * as fs from 'node:fs';
import * as path from 'node:path';
import { formatCLP, parseCLP } from '../../src/money/clp.js';

interface CLPFormattingCase {
  amount: number;
  expected_default: string;
  expected_space: string;
  expected_symbol_clp: string;
  expected_no_symbol: string;
  expected_parentheses?: string;
  description: string;
}

interface CLPParsingCase {
  input: string;
  expected: number;
  valid: boolean;
  description: string;
}

interface MoneyTestSuite {
  description: string;
  clp_formatting: CLPFormattingCase[];
  clp_parsing: CLPParsingCase[];
}

function loadSuite(): MoneyTestSuite {
  const filePath = path.resolve(__dirname, '../../../../specs/tests/money_formatting.json');
  const content = fs.readFileSync(filePath, 'utf-8');
  return JSON.parse(content);
}

describe('CLP Formatting Suite (specs/tests/money_formatting.json)', () => {
  const suite = loadSuite();

  for (const tc of suite.clp_formatting) {
    it(`formatCLP (default): ${tc.description}`, () => {
      expect(formatCLP(tc.amount)).toBe(tc.expected_default);
    });

    it(`formatCLP (space): ${tc.description}`, () => {
      expect(formatCLP(tc.amount, { space: true })).toBe(tc.expected_space);
    });

    it(`formatCLP (symbol CLP): ${tc.description}`, () => {
      expect(formatCLP(tc.amount, { symbol: 'CLP', space: true })).toBe(tc.expected_symbol_clp);
    });

    it(`formatCLP (no symbol): ${tc.description}`, () => {
      expect(formatCLP(tc.amount, { noSymbol: true })).toBe(tc.expected_no_symbol);
    });

    if (tc.expected_parentheses) {
      it(`formatCLP (parentheses): ${tc.description}`, () => {
        expect(formatCLP(tc.amount, { negativeFormat: 'parentheses' })).toBe(tc.expected_parentheses);
      });
    }
  }
});

describe('CLP Parsing Suite (specs/tests/money_formatting.json)', () => {
  const suite = loadSuite();

  for (const tc of suite.clp_parsing) {
    it(`parseCLP: ${tc.description}`, () => {
      if (tc.valid) {
        expect(parseCLP(tc.input)).toBe(tc.expected);
      } else {
        expect(() => parseCLP(tc.input)).toThrow();
      }
    });
  }
});
