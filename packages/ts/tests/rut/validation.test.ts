import { describe, it, expect } from 'vitest';
import * as fs from 'node:fs';
import * as path from 'node:path';
import { clean, calculateDv, validate, validateDetailed, RutErrorCode } from '../../src/rut/validation.js';

interface TestCase {
  input: string;
  expected_clean: string;
  expected_body: string | null;
  expected_dv: string | null;
  is_valid: boolean;
  description: string;
}

interface TestSuite {
  description: string;
  cases: TestCase[];
}

function loadValidationTestCases(): TestCase[] {
  const filePath = path.resolve(__dirname, '../../../../specs/tests/rut_validation.json');
  const content = fs.readFileSync(filePath, 'utf-8');
  const suite: TestSuite = JSON.parse(content);
  return suite.cases;
}

describe('RUT Validation Suite (specs/tests/rut_validation.json)', () => {
  const cases = loadValidationTestCases();

  describe('clean()', () => {
    for (const tc of cases) {
      it(`Clean: ${tc.description}`, () => {
        expect(clean(tc.input)).toBe(tc.expected_clean);
      });
    }
  });

  describe('calculateDv()', () => {
    for (const tc of cases) {
      if (tc.expected_body !== null && tc.expected_dv !== null) {
        it(`CalculateDv: ${tc.description}`, () => {
          expect(calculateDv(tc.expected_body!)).toBe(tc.expected_dv);
        });
      }
    }

    it('retorna vacío para cuerpo <= 0', () => {
      expect(calculateDv(0)).toBe('');
      expect(calculateDv(-123)).toBe('');
    });
  });

  describe('validate()', () => {
    for (const tc of cases) {
      it(`Validate: ${tc.description}`, () => {
        expect(validate(tc.input)).toBe(tc.is_valid);
      });
    }
  });

  describe('validateDetailed()', () => {
    it('retorna null para RUTs válidos', () => {
      expect(validateDetailed('12.345.678-5')).toBeNull();
      expect(validateDetailed('18.123.458-k')).toBeNull();
    });

    it('diagnostica código Empty', () => {
      expect(validateDetailed('')?.code).toBe(RutErrorCode.Empty);
      expect(validateDetailed('   ')?.code).toBe(RutErrorCode.Empty);
    });

    it('diagnostica código MalformedFormat', () => {
      expect(validateDetailed('12--345-678-5')?.code).toBe(RutErrorCode.MalformedFormat);
      expect(validateDetailed('12..345..678-5')?.code).toBe(RutErrorCode.MalformedFormat);
      expect(validateDetailed('1.2.3.4.5-K')?.code).toBe(RutErrorCode.MalformedFormat);
      expect(validateDetailed('12-345-678-5')?.code).toBe(RutErrorCode.MalformedFormat);
    });

    it('diagnostica código InvalidLength', () => {
      expect(validateDetailed('1')?.code).toBe(RutErrorCode.InvalidLength);
      expect(validateDetailed('K')?.code).toBe(RutErrorCode.InvalidLength);
      expect(validateDetailed('12345678901-5')?.code).toBe(RutErrorCode.InvalidLength);
    });

    it('diagnostica código InvalidCharacters', () => {
      expect(validateDetailed('18.12A.458-K')?.code).toBe(RutErrorCode.InvalidCharacters);
      expect(validateDetailed('abcdefgh-K')?.code).toBe(RutErrorCode.InvalidCharacters);
    });

    it('diagnostica código InvalidDv', () => {
      expect(validateDetailed('12.345.678-0')?.code).toBe(RutErrorCode.InvalidDv);
      expect(validateDetailed('18.123.458-1')?.code).toBe(RutErrorCode.InvalidDv);
    });
  });

  describe('QA Regresiones: calculateDv estricto', () => {
    it('rechaza Infinity, NaN, decimales y strings con sufijos no numéricos', () => {
      expect(calculateDv(Infinity)).toBe('');
      expect(calculateDv(-Infinity)).toBe('');
      expect(calculateDv(NaN)).toBe('');
      expect(calculateDv(12345678.9)).toBe('');
      expect(calculateDv('12345678abc')).toBe('');
      expect(calculateDv('+12345678')).toBe('');
    });
  });
});
