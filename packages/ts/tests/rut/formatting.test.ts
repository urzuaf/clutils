import { describe, it, expect } from 'vitest';
import * as fs from 'node:fs';
import * as path from 'node:path';
import { format } from '../../src/rut/formatting.js';

interface FormattingTestCase {
  input: string;
  expected_default: string;
  expected_no_dots: string;
  expected_no_hyphen: string;
  expected_plain: string;
  description: string;
}

interface FormattingTestSuite {
  description: string;
  cases: FormattingTestCase[];
}

function loadFormattingTestCases(): FormattingTestCase[] {
  const filePath = path.resolve(__dirname, '../../../../specs/tests/rut_formatting.json');
  const content = fs.readFileSync(filePath, 'utf-8');
  const suite: FormattingTestSuite = JSON.parse(content);
  return suite.cases;
}

describe('RUT Formatting Suite (specs/tests/rut_formatting.json)', () => {
  const cases = loadFormattingTestCases();

  for (const tc of cases) {
    it(`Format (default): ${tc.description}`, () => {
      expect(format(tc.input)).toBe(tc.expected_default);
    });

    it(`Format (no dots): ${tc.description}`, () => {
      expect(format(tc.input, { dots: false, hyphen: true })).toBe(tc.expected_no_dots);
    });

    it(`Format (no hyphen): ${tc.description}`, () => {
      expect(format(tc.input, { dots: true, hyphen: false })).toBe(tc.expected_no_hyphen);
    });

    it(`Format (plain): ${tc.description}`, () => {
      expect(format(tc.input, { dots: false, hyphen: false })).toBe(tc.expected_plain);
    });
  }
});
