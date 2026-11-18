import { describe, it, expect } from 'vitest';
import * as fs from 'node:fs';
import * as path from 'node:path';
import { formatUTM } from '../../src/money/utm.js';

interface UTMFormattingCase {
  amount: number;
  expected_default: string;
  expected_no_symbol: string;
  description: string;
}

interface UTMTestSuite {
  utm_formatting: UTMFormattingCase[];
}

function loadSuite(): UTMTestSuite {
  const filePath = path.resolve(__dirname, '../../../../specs/tests/money_formatting.json');
  const content = fs.readFileSync(filePath, 'utf-8');
  return JSON.parse(content);
}

describe('UTM Formatting Suite (specs/tests/money_formatting.json)', () => {
  const suite = loadSuite();

  for (const tc of suite.utm_formatting) {
    it(`formatUTM (default): ${tc.description}`, () => {
      expect(formatUTM(tc.amount)).toBe(tc.expected_default);
    });

    it(`formatUTM (no symbol): ${tc.description}`, () => {
      expect(formatUTM(tc.amount, { noSymbol: true })).toBe(tc.expected_no_symbol);
    });
  }
});
