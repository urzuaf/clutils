import { describe, it, expect } from 'vitest';
import * as fs from 'node:fs';
import * as path from 'node:path';
import { toWords } from '../../src/money/words.js';

interface ToWordsCase {
  amount: number;
  expected: string;
  expected_no_currency: string;
  description: string;
}

interface ToWordsTestSuite {
  to_words: ToWordsCase[];
}

function loadSuite(): ToWordsTestSuite {
  const filePath = path.resolve(__dirname, '../../../../specs/tests/money_formatting.json');
  const content = fs.readFileSync(filePath, 'utf-8');
  return JSON.parse(content);
}

describe('toWords Suite (specs/tests/money_formatting.json)', () => {
  const suite = loadSuite();

  for (const tc of suite.to_words) {
    it(`toWords (con moneda): ${tc.description}`, () => {
      expect(toWords(tc.amount)).toBe(tc.expected);
    });

    it(`toWords (sin moneda): ${tc.description}`, () => {
      expect(toWords(tc.amount, { noCurrency: true })).toBe(tc.expected_no_currency);
    });
  }
});
