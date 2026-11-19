import { describe, it, expect } from 'vitest';
import * as fs from 'node:fs';
import * as path from 'node:path';
import {
  TZ_SANTIAGO,
  TZ_PUNTA_ARENAS,
  TZ_EASTER,
  nowChile,
  nowMagallanes,
  nowEasterIsland,
  toChileTime,
  toMagallanesTime,
  toEasterIslandTime,
} from '../../src/date/timezones.js';

interface TimezoneCase {
  identifier: string;
  expected_iana: string;
  utc_timestamp: string;
  expected_hour: number;
  description: string;
}

interface DateTestSuite {
  timezones: TimezoneCase[];
}

function loadSuite(): DateTestSuite {
  const filePath = path.resolve(__dirname, '../../../../specs/tests/date_formatting.json');
  const content = fs.readFileSync(filePath, 'utf-8');
  return JSON.parse(content);
}

describe('Timezone Constants', () => {
  it('identificadores IANA oficiales', () => {
    expect(TZ_SANTIAGO).toBe('America/Santiago');
    expect(TZ_PUNTA_ARENAS).toBe('America/Punta_Arenas');
    expect(TZ_EASTER).toBe('Pacific/Easter');
  });
});

describe('Timezone Conversions (specs/tests/date_formatting.json)', () => {
  const suite = loadSuite();

  for (const tc of suite.timezones) {
    it(tc.description, () => {
      const utcDate = new Date(tc.utc_timestamp);
      let converted: Date;

      switch (tc.identifier) {
        case 'santiago':
          converted = toChileTime(utcDate);
          break;
        case 'punta_arenas':
          converted = toMagallanesTime(utcDate);
          break;
        case 'easter_island':
          converted = toEasterIslandTime(utcDate);
          break;
        default:
          throw new Error(`Zona desconocida: ${tc.identifier}`);
      }

      expect(converted.getHours()).toBe(tc.expected_hour);
    });
  }
});

describe('Now Helpers', () => {
  it('nowChile, nowMagallanes y nowEasterIsland retornan fechas validas', () => {
    const santiago = nowChile();
    const magallanes = nowMagallanes();
    const easter = nowEasterIsland();

    expect(santiago).toBeInstanceOf(Date);
    expect(magallanes).toBeInstanceOf(Date);
    expect(easter).toBeInstanceOf(Date);

    expect(Number.isNaN(santiago.getTime())).toBe(false);
    expect(Number.isNaN(magallanes.getTime())).toBe(false);
    expect(Number.isNaN(easter.getTime())).toBe(false);
  });
});
