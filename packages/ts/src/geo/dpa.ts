/**
 * @file dpa.ts
 * @description Funciones canónicas para consulta, filtrado jerárquico y búsqueda territorial
 * de la División Político-Administrativa de Chile (Regiones, Provincias y Comunas).
 */

import { REGIONES, PROVINCIAS, COMUNAS } from './data.js';
import { searchEntity, searchAllEntities } from './normalize.js';
import type { Region, Provincia, Comuna } from './types.js';

/** Mapeo de código oficial de región a sus números romanos históricos. */
const REGION_CODE_TO_ROMAN: Readonly<Record<string, string>> = Object.freeze({
  '01': 'I',
  '02': 'II',
  '03': 'III',
  '04': 'IV',
  '05': 'V',
  '06': 'VI',
  '07': 'VII',
  '08': 'VIII',
  '09': 'IX',
  '10': 'X',
  '11': 'XI',
  '12': 'XII',
  '13': 'RM',
  '14': 'XIV',
  '15': 'XV',
  '16': 'XVI',
});

// Garantizar inmutabilidad en runtime de los catálogos maestros
for (const r of REGIONES) Object.freeze(r);
Object.freeze(REGIONES);
for (const p of PROVINCIAS) Object.freeze(p);
Object.freeze(PROVINCIAS);
for (const c of COMUNAS) Object.freeze(c);
Object.freeze(COMUNAS);

// ============================================================================
// REGIONES
// ============================================================================

/**
 * Retorna una copia de las 16 regiones oficiales de Chile ordenadas geográficamente de Norte a Sur.
 */
export function getRegiones(): Region[] {
  return REGIONES.map((r) => ({ ...r }));
}

/**
 * Busca una región por su código oficial INE de 2 dígitos (ej: "13", "16", "05", "1"),
 * código internacional ISO (ej: "CL-RM", "CL-NB"), abreviatura oficial (ej: "RM", "NUBLE"),
 * número romano histórico (ej: "VIII", "XVI") o nombre oficial (tolerante a mayúsculas, tildes y eñes).
 * Retorna una copia de la Region encontrada o null si no existe coincidencia.
 */
export function findRegion(query: string): Region | null {
  const found = searchEntity(
    query,
    REGIONES,
    (r) => {
      const roman = REGION_CODE_TO_ROMAN[r.code];
      return roman ? [r.code, r.isoCode, r.abbreviation, roman] : [r.code, r.isoCode, r.abbreviation];
    },
    (r) => r.name
  );
  return found ? { ...found } : null;
}

// ============================================================================
// PROVINCIAS
// ============================================================================

/**
 * Retorna una copia de las 56 provincias oficiales de Chile.
 */
export function getProvincias(): Provincia[] {
  return PROVINCIAS.map((p) => ({ ...p }));
}

/**
 * Retorna las provincias pertenecientes a una región específica.
 * Acepta código de región (ej: "13", "16"), código ISO (ej: "CL-RM", "CL-NB") o nombre de región.
 * Si la región no existe o no tiene provincias, retorna un arreglo vacío.
 */
export function getProvinciasByRegion(regionQuery: string): Provincia[] {
  const reg = findRegion(regionQuery);
  if (!reg) {
    return [];
  }
  return PROVINCIAS.filter((p) => p.regionCode === reg.code).map((p) => ({ ...p }));
}

/**
 * Busca una provincia por su código INE de 3 dígitos (ej: "131", "161")
 * o por su nombre oficial (tolerante a mayúsculas, tildes y eñes).
 * Retorna una copia de la Provincia encontrada o null si no existe coincidencia.
 */
export function findProvincia(query: string): Provincia | null {
  const found = searchEntity(
    query,
    PROVINCIAS,
    (p) => [p.code],
    (p) => p.name
  );
  return found ? { ...found } : null;
}

/**
 * Busca y retorna todas las provincias que coincidan con el término de búsqueda,
 * ordenadas por relevancia (coincidencias exactas primero, luego prefijos y finalmente subcadenas).
 * Ideal para cajas de autocompletado y filtros de búsqueda interactivos en UI.
 */
export function searchProvincias(query: string): Provincia[] {
  return searchAllEntities(
    query,
    PROVINCIAS,
    (p) => [p.code],
    (p) => p.name
  ).map((p) => ({ ...p }));
}

/**
 * Resuelve la región oficial a la que pertenece una provincia dada.
 * Acepta código de provincia (ej: "131", "161") o nombre de provincia (ej: "diguillin", "santiago").
 * Retorna una copia de la Region encontrada o null si la provincia no existe.
 */
export function getRegionByProvincia(provinciaQuery: string): Region | null {
  const prov = findProvincia(provinciaQuery);
  if (!prov) {
    return null;
  }
  return findRegion(prov.regionCode);
}

// ============================================================================
// COMUNAS
// ============================================================================

/**
 * Retorna una copia de las 346 comunas oficiales de Chile.
 */
export function getComunas(): Comuna[] {
  return COMUNAS.map((c) => ({ ...c }));
}

/**
 * Retorna las comunas pertenecientes a una región específica.
 * Acepta código de región (ej: "13", "16"), código ISO (ej: "CL-RM", "CL-NB") o nombre de región.
 * Si la región no existe o no tiene comunas, retorna un arreglo vacío.
 */
export function getComunasByRegion(regionQuery: string): Comuna[] {
  const reg = findRegion(regionQuery);
  if (!reg) {
    return [];
  }
  return COMUNAS.filter((c) => c.regionCode === reg.code).map((c) => ({ ...c }));
}

/**
 * Retorna las comunas pertenecientes a una provincia específica.
 * Acepta código de provincia (ej: "131", "161") o nombre de provincia (ej: "santiago", "diguillin").
 * Si la provincia no existe o no tiene comunas, retorna un arreglo vacío.
 */
export function getComunasByProvincia(provinciaQuery: string): Comuna[] {
  const prov = findProvincia(provinciaQuery);
  if (!prov) {
    return [];
  }
  return COMUNAS.filter((c) => c.provinciaCode === prov.code).map((c) => ({ ...c }));
}

/**
 * Busca una comuna por su Código Único Territorial (CUT de 5 dígitos, ej: "13120", "16101")
 * o por su nombre oficial (tolerante a mayúsculas, tildes y eñes, ej: "nunoa", "vina del mar", "aysen").
 * Retorna una copia de la Comuna encontrada o null si no existe coincidencia.
 */
export function findComuna(query: string): Comuna | null {
  const found = searchEntity(
    query,
    COMUNAS,
    (c) => [c.cut],
    (c) => c.name
  );
  return found ? { ...found } : null;
}

/**
 * Busca y retorna todas las comunas que coincidan con el término de búsqueda,
 * ordenadas por relevancia (coincidencias exactas primero, luego prefijos y finalmente subcadenas).
 * Ideal para cajas de autocompletado y filtros de búsqueda interactivos en UI (ej: "San" -> San Bernardo, San Miguel...).
 */
export function searchComunas(query: string): Comuna[] {
  return searchAllEntities(
    query,
    COMUNAS,
    (c) => [c.cut],
    (c) => c.name
  ).map((c) => ({ ...c }));
}

/**
 * Resuelve la región oficial a la que pertenece una comuna dada.
 * Acepta código CUT (ej: "13120", "16101") o nombre de comuna (ej: "nunoa", "chillan").
 * Retorna una copia de la Region encontrada o null si la comuna no existe.
 */
export function getRegionByComuna(comunaQuery: string): Region | null {
  const comuna = findComuna(comunaQuery);
  if (!comuna) {
    return null;
  }
  return findRegion(comuna.regionCode);
}

/**
 * Resuelve la provincia oficial a la que pertenece una comuna dada.
 * Acepta código CUT (ej: "13120", "16101") o nombre de comuna (ej: "nunoa", "chillan").
 * Retorna una copia de la Provincia encontrada o null si la comuna no existe.
 */
export function getProvinciaByComuna(comunaQuery: string): Provincia | null {
  const comuna = findComuna(comunaQuery);
  if (!comuna) {
    return null;
  }
  return findProvincia(comuna.provinciaCode);
}
