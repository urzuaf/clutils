import { describe, it, expect } from 'vitest';
import * as fs from 'node:fs';
import * as path from 'node:path';
import {
  generate,
  generatePerson,
  generateBusiness,
  generateBetween,
  DEFAULT_PERSON_MIN,
  DEFAULT_PERSON_MAX,
  DEFAULT_BUSINESS_MIN,
  DEFAULT_BUSINESS_MAX,
  DEFAULT_GENERAL_MIN,
  DEFAULT_GENERAL_MAX,
} from '../../src/rut/generation.js';
import { validate } from '../../src/rut/validation.js';
import { split } from '../../src/rut/formatting.js';

describe('RUT Generation Suite', () => {
  it('genera 100 RUTs válidos consecutivos por defecto', () => {
    for (let i = 0; i < 100; i++) {
      const generated = generate();
      expect(validate(generated)).toBe(true);
    }
  });

  it('respeta opciones de formato plano', () => {
    const generated = generate({ dots: false, hyphen: false });
    expect(validate(generated)).toBe(true);
    expect(generated).not.toContain('.');
    expect(generated).not.toContain('-');
  });

  it('generatePerson produce RUTs en rango [5M-28M] y válidos', () => {
    for (let i = 0; i < 50; i++) {
      const generated = generatePerson();
      expect(validate(generated)).toBe(true);
      const parts = split(generated);
      const body = parts.bodyInt();
      expect(body).toBeGreaterThanOrEqual(5_000_000);
      expect(body).toBeLessThanOrEqual(28_000_000);
    }
  });

  it('generateBusiness produce RUTs en rango [60M-90M] y válidos', () => {
    for (let i = 0; i < 50; i++) {
      const generated = generateBusiness();
      expect(validate(generated)).toBe(true);
      const parts = split(generated);
      const body = parts.bodyInt();
      expect(body).toBeGreaterThanOrEqual(60_000_000);
      expect(body).toBeLessThanOrEqual(90_000_000);
    }
  });

  it('generateBetween es seguro con límites invertidos o negativos', () => {
    const inverted = generateBetween(80_000_000, 20_000_000);
    expect(validate(inverted)).toBe(true);

    const negative = generateBetween(-100, -50);
    expect(validate(negative)).toBe(true);
  });

  it('constantes en TypeScript coinciden 1:1 con data/rut/ranges.json', () => {
    const filePath = path.resolve(__dirname, '../../../../data/rut/ranges.json');
    const content = fs.readFileSync(filePath, 'utf-8');
    const schema = JSON.parse(content);

    expect(DEFAULT_PERSON_MIN).toBe(schema.ranges.person.min);
    expect(DEFAULT_PERSON_MAX).toBe(schema.ranges.person.max);
    expect(DEFAULT_BUSINESS_MIN).toBe(schema.ranges.business.min);
    expect(DEFAULT_BUSINESS_MAX).toBe(schema.ranges.business.max);
    expect(DEFAULT_GENERAL_MIN).toBe(schema.ranges.general.min);
    expect(DEFAULT_GENERAL_MAX).toBe(schema.ranges.general.max);
  });
});
