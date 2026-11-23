import { describe, it, expect } from 'vitest';
import * as fs from 'node:fs';
import * as path from 'node:path';
import {
  getRegiones,
  findRegion,
  getProvincias,
  getProvinciasByRegion,
  findProvincia,
  searchProvincias,
  getRegionByProvincia,
  getComunas,
  getComunasByRegion,
  getComunasByProvincia,
  findComuna,
  searchComunas,
  getRegionByComuna,
  getProvinciaByComuna,
  normalizeSearch,
} from '../../src/geo/index.js';

interface RegionCommuneCount {
  region_code: string;
  iso_code: string;
  name: string;
  expected_communes: number;
  expected_provinces: number;
}

interface FilteringCase {
  description: string;
  region_query?: string;
  expected_count?: number;
  sample_commune_cuts?: string[];
  provinces_region_query?: string;
  expected_provinces_count?: number;
  expected_province_names?: string[];
}

interface FindComunaCase {
  query: string;
  expected_cut: string | null;
  expected_name: string | null;
  expected_region_code: string | null;
  description: string;
}

interface FindRegionCase {
  query: string;
  expected_code: string | null;
  expected_iso: string | null;
  expected_name: string | null;
  description: string;
}

interface ReverseLookupCase {
  comuna_query: string;
  expected_region_code: string | null;
  expected_region_name: string | null;
  description: string;
}

interface GeoTestSuite {
  description: string;
  source: string;
  counts: {
    expected_regions: number;
    expected_provinces: number;
    expected_communes: number;
    region_commune_counts: RegionCommuneCount[];
  };
  filtering: FilteringCase[];
  find_comuna: FindComunaCase[];
  find_region: FindRegionCase[];
  get_region_by_comuna: ReverseLookupCase[];
}

const specPath = path.resolve(__dirname, '../../../../specs/tests/geo_dpa.json');
const testSuite: GeoTestSuite = JSON.parse(fs.readFileSync(specPath, 'utf-8'));

describe('Geo DPA Suite (specs/tests/geo_dpa.json)', () => {
  describe('Estructura canónica y conteos oficiales', () => {
    it('debe tener exactamente 16 regiones', () => {
      const regiones = getRegiones();
      expect(regiones).toHaveLength(testSuite.counts.expected_regions);
      expect(regiones[0].order).toBe(1);
      expect(regiones[15].order).toBe(16);
    });

    it('debe tener exactamente 56 provincias', () => {
      const provincias = getProvincias();
      expect(provincias).toHaveLength(testSuite.counts.expected_provinces);
    });

    it('debe tener exactamente 346 comunas', () => {
      const comunas = getComunas();
      expect(comunas).toHaveLength(testSuite.counts.expected_communes);
    });

    it('debe cumplir con la distribución de comunas y provincias por región', () => {
      for (const r of testSuite.counts.region_commune_counts) {
        const comunas = getComunasByRegion(r.region_code);
        expect(comunas).toHaveLength(r.expected_communes);

        const provincias = getProvinciasByRegion(r.region_code);
        expect(provincias).toHaveLength(r.expected_provinces);
      }
    });
  });

  describe('Filtrado jerárquico', () => {
    testSuite.filtering.forEach((tc) => {
      it(tc.description, () => {
        if (tc.region_query !== undefined) {
          const comunas = getComunasByRegion(tc.region_query);
          expect(comunas).toHaveLength(tc.expected_count!);
          if (tc.sample_commune_cuts && tc.sample_commune_cuts.length > 0) {
            const cuts = comunas.map((c) => c.cut);
            for (const expectedCut of tc.sample_commune_cuts) {
              expect(cuts).toContain(expectedCut);
            }
          }
        }

        if (tc.provinces_region_query !== undefined) {
          const provs = getProvinciasByRegion(tc.provinces_region_query);
          expect(provs).toHaveLength(tc.expected_provinces_count!);
          if (tc.expected_province_names && tc.expected_province_names.length > 0) {
            const names = provs.map((p) => p.name).sort();
            const expectedSorted = [...tc.expected_province_names].sort();
            expect(names).toEqual(expectedSorted);
          }
        }
      });
    });

    it('getComunasByProvincia debe filtrar comunas por provincia correctamente', () => {
      const comunasStgo = getComunasByProvincia('131');
      expect(comunasStgo.length).toBe(32); // 32 comunas de la provincia de Santiago

      const comunasDiguillin = getComunasByProvincia('diguillin');
      expect(comunasDiguillin.length).toBe(9); // 9 comunas de la provincia de Diguillín (Ñuble)

      const inexistente = getComunasByProvincia('999');
      expect(inexistente).toHaveLength(0);
    });
  });

  describe('findComuna', () => {
    testSuite.find_comuna.forEach((tc) => {
      it(tc.description, () => {
        const result = findComuna(tc.query);
        if (tc.expected_cut === null) {
          expect(result).toBeNull();
        } else {
          expect(result).not.toBeNull();
          expect(result?.cut).toBe(tc.expected_cut);
          expect(result?.name).toBe(tc.expected_name);
          expect(result?.regionCode).toBe(tc.expected_region_code);
        }
      });
    });
  });

  describe('findRegion', () => {
    testSuite.find_region.forEach((tc) => {
      it(tc.description, () => {
        const result = findRegion(tc.query);
        if (tc.expected_code === null) {
          expect(result).toBeNull();
        } else {
          expect(result).not.toBeNull();
          expect(result?.code).toBe(tc.expected_code);
          expect(result?.isoCode).toBe(tc.expected_iso);
          expect(result?.name).toBe(tc.expected_name);
        }
      });
    });
  });

  describe('findProvincia y searchProvincias', () => {
    it('debe encontrar provincia por código INE exacto', () => {
      const prov = findProvincia('131');
      expect(prov).not.toBeNull();
      expect(prov?.name).toBe('Santiago');
      expect(prov?.regionCode).toBe('13');
    });

    it('debe encontrar provincia insensible a mayúsculas y tildes', () => {
      const prov1 = findProvincia('diguillin');
      expect(prov1).not.toBeNull();
      expect(prov1?.name).toBe('Diguillín');
      expect(prov1?.regionCode).toBe('16');

      const prov2 = findProvincia('CONCEPCION');
      expect(prov2).not.toBeNull();
      expect(prov2?.name).toBe('Concepción');
    });

    it('debe retornar null para provincias inexistentes o vacías', () => {
      expect(findProvincia('narnia')).toBeNull();
      expect(findProvincia('')).toBeNull();
      expect(findProvincia('   ')).toBeNull();
    });

    it('searchProvincias debe retornar subconjunto ordenado de coincidencias', () => {
      const results = searchProvincias('San');
      expect(results.length).toBeGreaterThan(0);
      expect(results.some((p) => p.name === 'San Antonio')).toBe(true);
      expect(results.some((p) => p.name === 'Santiago')).toBe(true);

      const empty = searchProvincias('xyz999');
      expect(empty).toHaveLength(0);
    });
  });

  describe('searchComunas', () => {
    it('debe retornar múltiples coincidencias ordenadas por relevancia para autocompletado', () => {
      const matches = searchComunas('San');
      expect(matches.length).toBeGreaterThan(10);
      // Coincidencias por prefijo ("San Bernardo", "San Miguel", etc.)
      const names = matches.map((c) => c.name);
      expect(names).toContain('San Bernardo');
      expect(names).toContain('San Miguel');
      expect(names).toContain('San Joaquín');
    });

    it('debe retornar comuna única si es coincidencia exacta', () => {
      const matches = searchComunas('13120');
      expect(matches).toHaveLength(1);
      expect(matches[0].name).toBe('Ñuñoa');
    });

    it('debe retornar arreglo vacío para consultas vacías o sin coincidencias', () => {
      expect(searchComunas('')).toEqual([]);
      expect(searchComunas('   ')).toEqual([]);
      expect(searchComunas('comuna_inexistente_999')).toEqual([]);
    });
  });

  describe('Resolución inversa: getRegionByComuna, getProvinciaByComuna y getRegionByProvincia', () => {
    testSuite.get_region_by_comuna.forEach((tc) => {
      it(tc.description, () => {
        const reg = getRegionByComuna(tc.comuna_query);
        if (tc.expected_region_code === null) {
          expect(reg).toBeNull();
        } else {
          expect(reg).not.toBeNull();
          expect(reg?.code).toBe(tc.expected_region_code);
          expect(reg?.name).toBe(tc.expected_region_name);
        }
      });
    });

    it('getProvinciaByComuna debe resolver la provincia oficial de la comuna', () => {
      const prov1 = getProvinciaByComuna('Ñuñoa');
      expect(prov1).not.toBeNull();
      expect(prov1?.name).toBe('Santiago');
      expect(prov1?.code).toBe('131');

      const prov2 = getProvinciaByComuna('chillan');
      expect(prov2).not.toBeNull();
      expect(prov2?.name).toBe('Diguillín');
      expect(prov2?.code).toBe('161');

      expect(getProvinciaByComuna('inexistente')).toBeNull();
    });

    it('getRegionByProvincia debe resolver la región oficial de la provincia', () => {
      const reg1 = getRegionByProvincia('161');
      expect(reg1).not.toBeNull();
      expect(reg1?.name).toBe('Ñuble');
      expect(reg1?.code).toBe('16');

      const reg2 = getRegionByProvincia('santiago');
      expect(reg2).not.toBeNull();
      expect(reg2?.name).toBe('Metropolitana de Santiago');

      expect(getRegionByProvincia('inexistente')).toBeNull();
    });
  });

  describe('normalizeSearch', () => {
    it('debe remover tildes, eñes y caracteres especiales correctamente', () => {
      expect(normalizeSearch('Ñuñoa')).toBe('nunoa');
      expect(normalizeSearch('Concepción')).toBe('concepcion');
      expect(normalizeSearch("O'Higgins")).toBe('ohiggins');
      expect(normalizeSearch('  Valparaíso - Viña  ')).toBe('valparaiso vina');
      expect(normalizeSearch('')).toBe('');
      expect(normalizeSearch('   ')).toBe('');
    });
  });
});
