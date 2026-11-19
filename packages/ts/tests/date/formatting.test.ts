import { describe, it, expect } from 'vitest';
import * as fs from 'node:fs';
import * as path from 'node:path';
import { parseDate, formatDate, isWeekend } from '../../src/date/formatting.js';

interface ParseDateCase {
  input: string;
  expected_year: number;
  expected_month: number;
  expected_day: number;
  valid: boolean;
  description: string;
}

interface FormatDateCase {
  year: number;
  month: number;
  day: number;
  expected_numeric: string;
  expected_numeric_hyphen: string;
  expected_formal: string;
  expected_full: string;
  description: string;
}

interface IsWeekendCase {
  year: number;
  month: number;
  day: number;
  expected: boolean;
  description: string;
}

interface DateTestSuite {
  parse_date: ParseDateCase[];
  format_date: FormatDateCase[];
  is_weekend: IsWeekendCase[];
}

function loadSuite(): DateTestSuite {
  const filePath = path.resolve(__dirname, '../../../../specs/tests/date_formatting.json');
  const content = fs.readFileSync(filePath, 'utf-8');
  return JSON.parse(content);
}

describe('Parse Date Suite (specs/tests/date_formatting.json)', () => {
  const suite = loadSuite();

  for (const tc of suite.parse_date) {
    it(tc.description, () => {
      if (tc.valid) {
        const d = parseDate(tc.input);
        expect(d.getFullYear()).toBe(tc.expected_year);
        expect(d.getMonth() + 1).toBe(tc.expected_month);
        expect(d.getDate()).toBe(tc.expected_day);
      } else {
        expect(() => parseDate(tc.input)).toThrow();
      }
    });
  }
});

describe('Format Date Suite (specs/tests/date_formatting.json)', () => {
  const suite = loadSuite();

  for (const tc of suite.format_date) {
    const testDate = new Date(tc.year, tc.month - 1, tc.day, 12, 0, 0);

    it(`formatDate (default / numeric): ${tc.description}`, () => {
      expect(formatDate(testDate)).toBe(tc.expected_numeric);
      expect(formatDate(testDate, 'numeric')).toBe(tc.expected_numeric);
    });

    it(`formatDate (numeric_hyphen): ${tc.description}`, () => {
      expect(formatDate(testDate, 'numeric_hyphen')).toBe(tc.expected_numeric_hyphen);
    });

    it(`formatDate (formal): ${tc.description}`, () => {
      expect(formatDate(testDate, 'formal')).toBe(tc.expected_formal);
    });

    it(`formatDate (full): ${tc.description}`, () => {
      expect(formatDate(testDate, 'full')).toBe(tc.expected_full);
    });
  }
});

describe('Is Weekend Suite (specs/tests/date_formatting.json)', () => {
  const suite = loadSuite();

  for (const tc of suite.is_weekend) {
    it(tc.description, () => {
      const testDate = new Date(tc.year, tc.month - 1, tc.day, 12, 0, 0);
      expect(isWeekend(testDate)).toBe(tc.expected);
    });
  }
});
