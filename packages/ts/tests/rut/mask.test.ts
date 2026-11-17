import { describe, it, expect } from 'vitest';
import * as fs from 'node:fs';
import * as path from 'node:path';
import { mask } from '../../src/rut/formatting.js';

interface MaskTestCase {
  input: string;
  expected_default: string;
  expected_plain: string;
  description: string;
}

interface MaskTestSuite {
  description: string;
  cases: MaskTestCase[];
}

function loadMaskTestCases(): MaskTestCase[] {
  const filePath = path.resolve(__dirname, '../../../../specs/tests/rut_mask.json');
  const content = fs.readFileSync(filePath, 'utf-8');
  const suite: MaskTestSuite = JSON.parse(content);
  return suite.cases;
}

describe('RUT Mask Suite (specs/tests/rut_mask.json)', () => {
  const cases = loadMaskTestCases();

  for (const tc of cases) {
    it(`Mask (default): ${tc.description}`, () => {
      expect(mask(tc.input)).toBe(tc.expected_default);
    });

    it(`Mask (plain): ${tc.description}`, () => {
      expect(mask(tc.input, { formatted: false })).toBe(tc.expected_plain);
    });
  }

  it('permite carácter de ofuscación personalizado', () => {
    expect(mask('18123458K', { maskChar: 'X', visibleCount: 4, formatted: true })).toBe('XX.XXX.458-K');
  });

  it('permite cantidad personalizada de caracteres visibles', () => {
    expect(mask('18123458K', { visibleCount: 2, formatted: true })).toBe('**.***.**8-K');
  });
});
