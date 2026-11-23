import { describe, it, expect } from 'vitest';
import * as fs from 'node:fs';
import * as path from 'node:path';
import { split } from '../../src/rut/formatting.js';

interface SplitTestCase {
  input: string;
  expected_body: string | null;
  expected_dv: string | null;
  expected_int: number | null;
  is_error: boolean;
  description: string;
}

interface SplitTestSuite {
  description: string;
  cases: SplitTestCase[];
}

function loadSplitTestCases(): SplitTestCase[] {
  const filePath = path.resolve(__dirname, '../../../../specs/tests/rut_split.json');
  const content = fs.readFileSync(filePath, 'utf-8');
  const suite: SplitTestSuite = JSON.parse(content);
  return suite.cases;
}

describe('RUT Split Suite (specs/tests/rut_split.json)', () => {
  const cases = loadSplitTestCases();

  for (const tc of cases) {
    it(`Split: ${tc.description}`, () => {
      if (tc.is_error) {
        expect(() => split(tc.input)).toThrow();
      } else {
        const parts = split(tc.input);
        expect(parts.body).toBe(tc.expected_body);
        expect(parts.dv).toBe(tc.expected_dv);
        if (tc.expected_int !== null) {
          expect(parts.bodyInt()).toBe(tc.expected_int);
        }
      }
    });
  }

  it('bodyInt lanza error si el cuerpo contiene letras intercaladas (ej. 12A34-K)', () => {
    const parts = split('12A34-K');
    expect(() => parts.bodyInt()).toThrow();
  });
});
