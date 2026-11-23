import { describe, it, expect } from 'vitest';
import * as fs from 'node:fs';
import * as path from 'node:path';
import {
  normalizePhone,
  normalizeMobile,
  validatePhone,
  formatPhone,
  getPhoneType,
  isEmergencyNumber,
} from '../../src/telecom/index.js';

interface NormalizationCase {
  input: string;
  expected_e164: string;
  valid: boolean;
  description: string;
}

interface FormattingCase {
  input: string;
  expected_international: string;
  expected_national: string;
  expected_e164: string;
  description: string;
}

interface PhoneTypeCase {
  input: string;
  expected_type: string;
  description: string;
}

interface TestSuite {
  description: string;
  normalization: NormalizationCase[];
  formatting: FormattingCase[];
  phone_types: PhoneTypeCase[];
}

const specPath = path.resolve(__dirname, '../../../../specs/tests/telecom_phone.json');
const testSuite: TestSuite = JSON.parse(fs.readFileSync(specPath, 'utf-8'));

describe('Telecom Phone Suite (specs/tests/telecom_phone.json)', () => {
  describe('normalizePhone', () => {
    testSuite.normalization.forEach((tc) => {
      it(tc.description, () => {
        if (tc.valid) {
          expect(normalizePhone(tc.input)).toBe(tc.expected_e164);
        } else {
          expect(() => normalizePhone(tc.input)).toThrow();
        }
      });
    });
  });

  describe('validatePhone', () => {
    testSuite.normalization.forEach((tc) => {
      it(tc.description, () => {
        expect(validatePhone(tc.input)).toBe(tc.valid);
      });
    });
  });

  describe('formatPhone', () => {
    testSuite.formatting.forEach((tc) => {
      it(`${tc.description} (international)`, () => {
        expect(formatPhone(tc.input, 'international')).toBe(tc.expected_international);
      });

      it(`${tc.description} (default style)`, () => {
        expect(formatPhone(tc.input)).toBe(tc.expected_international);
      });

      it(`${tc.description} (national)`, () => {
        expect(formatPhone(tc.input, 'national')).toBe(tc.expected_national);
      });

      it(`${tc.description} (e164)`, () => {
        expect(formatPhone(tc.input, 'e164')).toBe(tc.expected_e164);
      });
    });

    it('debe arrojar error ante entrada inválida', () => {
      expect(() => formatPhone('invalido')).toThrow('Número telefónico chileno inválido');
    });
  });

  describe('getPhoneType', () => {
    testSuite.phone_types.forEach((tc) => {
      it(tc.description, () => {
        expect(getPhoneType(tc.input)).toBe(tc.expected_type);
      });
    });
  });

  describe('normalizeMobile', () => {
    it('normaliza móvil válido a E.164', () => {
      expect(normalizeMobile('912345678')).toBe('+56912345678');
      expect(normalizeMobile('1234 5678')).toBe('+56912345678');
    });

    it('rechaza red fija arrojando error descriptivo', () => {
      expect(() => normalizeMobile('221234567')).toThrow(
        'El número telefónico no corresponde a una línea móvil'
      );
    });

    it('rechaza número con longitud inválida', () => {
      expect(() => normalizeMobile('12345')).toThrow();
    });
  });

  describe('isEmergencyNumber', () => {
    it('detecta números oficiales de emergencia', () => {
      expect(isEmergencyNumber('131')).toBe(true);
      expect(isEmergencyNumber('132')).toBe(true);
      expect(isEmergencyNumber('133')).toBe(true);
      expect(isEmergencyNumber('134')).toBe(true);
      expect(isEmergencyNumber('149')).toBe(true);
      expect(isEmergencyNumber('911')).toBe(true);
      expect(isEmergencyNumber('143')).toBe(true);
      expect(isEmergencyNumber('1411')).toBe(true);
      expect(isEmergencyNumber('1412')).toBe(true);
    });

    it('retorna false para números cortos no asignados', () => {
      expect(isEmergencyNumber('123')).toBe(false);
      expect(isEmergencyNumber('999')).toBe(false);
      expect(isEmergencyNumber('1499')).toBe(false);
    });
  });
});
