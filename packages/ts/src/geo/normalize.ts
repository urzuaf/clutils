/**
 * @file normalize.ts
 * @description Utilidades de normalización y motores de búsqueda geográfica tolerantes a tildes, eñes y mayúsculas.
 */

/**
 * Limpia y normaliza una cadena para búsquedas insensibles a mayúsculas,
 * tildes, eñes, signos de puntuación y espacios redundantes.
 */
export function normalizeSearch(s: string): string {
  const trimmed = s.trim();
  if (!trimmed) {
    return '';
  }

  let res = '';
  const lower = trimmed.toLowerCase();

  for (let i = 0; i < lower.length; i++) {
    const ch = lower[i];
    switch (ch) {
      case 'á':
      case 'à':
      case 'ä':
      case 'â':
        res += 'a';
        break;
      case 'é':
      case 'è':
      case 'ë':
      case 'ê':
        res += 'e';
        break;
      case 'í':
      case 'ì':
      case 'ï':
      case 'î':
        res += 'i';
        break;
      case 'ó':
      case 'ò':
      case 'ö':
      case 'ô':
        res += 'o';
        break;
      case 'ú':
      case 'ù':
      case 'ü':
      case 'û':
        res += 'u';
        break;
      case 'ñ':
        res += 'n';
        break;
      case "'":
      case '.':
      case '-':
      case '_':
      case '/':
        // Omitir puntuación
        break;
      default:
        res += ch;
        break;
    }
  }

  return res.trim().split(/\s+/).filter(Boolean).join(' ');
}

/**
 * Motor genérico de búsqueda territorial para retornar un único resultado:
 * 1. Coincidencia por identificadores exactos (códigos INE, ISOs, abreviaturas, CUTs).
 * 2. Coincidencia exacta por nombre normalizado (sin tildes, eñes ni mayúsculas).
 * 3. Coincidencia parcial contenida (subcadena) en el nombre normalizado.
 */
export function searchEntity<T>(
  query: string,
  items: readonly T[],
  getIdents: (item: T) => string[],
  getName: (item: T) => string
): T | null {
  const trimmed = query.trim();
  if (!trimmed) {
    return null;
  }

  const normQuery = normalizeSearch(query);

  let padded = trimmed;
  if (trimmed.length === 1 && trimmed >= '0' && trimmed <= '9') {
    padded = '0' + trimmed;
  }

  // Paso 1: Coincidencia por identificadores exactos
  const trimmedLower = trimmed.toLowerCase();
  const paddedLower = padded.toLowerCase();

  for (const item of items) {
    const idents = getIdents(item);
    for (const id of idents) {
      if (!id) continue;
      const idLower = id.toLowerCase();
      if (idLower === trimmedLower || idLower === paddedLower || normalizeSearch(id) === normQuery) {
        return item;
      }
    }
  }

  if (!normQuery) {
    return null;
  }

  // Paso 2: Coincidencia exacta por nombre normalizado
  for (const item of items) {
    if (normalizeSearch(getName(item)) === normQuery) {
      return item;
    }
  }

  // Paso 3: Coincidencia parcial contenida
  for (const item of items) {
    if (normalizeSearch(getName(item)).includes(normQuery)) {
      return item;
    }
  }

  return null;
}

/**
 * Busca y retorna todas las entidades coincidentes como un subconjunto,
 * ordenando primero coincidencias exactas, luego coincidencias por prefijo y finalmente subcadenas.
 */
export function searchAllEntities<T>(
  query: string,
  items: readonly T[],
  getIdents: (item: T) => string[],
  getName: (item: T) => string
): T[] {
  const trimmed = query.trim();
  if (!trimmed) {
    return [];
  }

  const normQuery = normalizeSearch(query);

  let padded = trimmed;
  if (trimmed.length === 1 && trimmed >= '0' && trimmed <= '9') {
    padded = '0' + trimmed;
  }

  const exactMatches: T[] = [];
  const prefixMatches: T[] = [];
  const substringMatches: T[] = [];
  const seen = new Set<number>();

  const trimmedLower = trimmed.toLowerCase();
  const paddedLower = padded.toLowerCase();

  // 1. Identificadores exactos
  for (let i = 0; i < items.length; i++) {
    const item = items[i];
    let matched = false;
    for (const id of getIdents(item)) {
      if (!id) continue;
      const idLower = id.toLowerCase();
      if (idLower === trimmedLower || idLower === paddedLower || normalizeSearch(id) === normQuery) {
        matched = true;
        break;
      }
    }
    if (matched && !seen.has(i)) {
      seen.add(i);
      exactMatches.push(item);
    }
  }

  if (!normQuery) {
    return exactMatches;
  }

  // 2. Coincidencia exacta por nombre
  for (let i = 0; i < items.length; i++) {
    if (seen.has(i)) continue;
    const item = items[i];
    if (normalizeSearch(getName(item)) === normQuery) {
      seen.add(i);
      exactMatches.push(item);
    }
  }

  // 3. Coincidencia por prefijo del nombre
  for (let i = 0; i < items.length; i++) {
    if (seen.has(i)) continue;
    const item = items[i];
    if (normalizeSearch(getName(item)).startsWith(normQuery)) {
      seen.add(i);
      prefixMatches.push(item);
    }
  }

  // 4. Coincidencia por subcadena dentro del nombre
  for (let i = 0; i < items.length; i++) {
    if (seen.has(i)) continue;
    const item = items[i];
    if (normalizeSearch(getName(item)).includes(normQuery)) {
      seen.add(i);
      substringMatches.push(item);
    }
  }

  return [...exactMatches, ...prefixMatches, ...substringMatches];
}
