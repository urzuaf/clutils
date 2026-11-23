# @urzuaf/clutils

Suite de utilidades canónicas, atómicas y sin dependencias externas para proyectos TypeScript / JavaScript que operan en o con **Chile**.

[![npm version](https://img.shields.io/npm/v/@urzuaf/clutils.svg)](https://www.npmjs.com/package/@urzuaf/clutils)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](./LICENSE)
[![Dependencies](https://img.shields.io/badge/dependencies-0-success.svg)](#)

---

## Características

- **Cero dependencias:** 0 dependencias en tiempo de ejecución.
- **Tipado estricto:** 100% TypeScript con archivos de definición `.d.ts` y mapas de fuentes.
- **Dual Export:** Soporte nativo para **ESM** (`import`) y **CommonJS** (`require`).
- **Tree-shaking granular:** Subpaths dedicados por dominio (`@urzuaf/clutils/rut`, `/money`, `/date`, `/telecom`, `/geo`).
- **Universal:** Compatible con Node.js, Bun, Deno, navegadores y Edge Runtimes.

---

## Instalación

```bash
# npm
npm install @urzuaf/clutils

# pnpm
pnpm add @urzuaf/clutils

# bun
bun add @urzuaf/clutils

# yarn
yarn add @urzuaf/clutils
```

---

## Formas de importación

### 1. Por subpath (recomendado para tree-shaking)
```ts
import { validate, format, mask } from '@urzuaf/clutils/rut';
import { formatCLP, parseCLP, calculateIVA, toWords } from '@urzuaf/clutils/money';
import { parseDate, formatDate, nowChile } from '@urzuaf/clutils/date';
import { normalizePhone, formatPhone, isEmergencyNumber } from '@urzuaf/clutils/telecom';
import { getRegiones, findComuna, searchComunas } from '@urzuaf/clutils/geo';
```

### 2. Por namespace desde la raíz
```ts
import { rut, money, date, telecom, geo } from '@urzuaf/clutils';

rut.validate('18.123.458-K');
money.formatCLP(1500000);
date.formatDate(date.parseDate('18/09/2026'), 'formal');
telecom.normalizePhone('9 1234 5678');
geo.findComuna('nunoa');
```

### 3. CommonJS
```js
const { rut, money } = require('@urzuaf/clutils');
const { findComuna } = require('@urzuaf/clutils/geo');
```

---

## 1. RUT / RUN (`@urzuaf/clutils/rut`)

```ts
import {
  validate,
  validateDetailed,
  RutErrorCode,
  clean,
  calculateDv,
  format,
  split,
  mask,
  generate,
  generatePerson,
  generateBusiness,
  generateBetween,
} from '@urzuaf/clutils/rut';

validate('12.345.678-5');   // true
validate('12.345.678-0');   // false

const err = validateDetailed('18.12A.458-K');
if (err) {
  console.log(err.code);    // RutErrorCode.InvalidCharacters (4)
  console.log(err.message); // "el cuerpo del RUT contiene caracteres no numéricos"
}

clean('  18.123.458-k  ');  // "18123458K"
calculateDv(18123458);      // "K"

format('18123458k');                             // "18.123.458-K"
format('18123458k', { dots: false });            // "18123458-K"
format('18123458k', { hyphen: false });          // "18.123.458K"

const parts = split('09.123.456-4');
parts.body;      // "09123456"
parts.dv;        // "4"
parts.bodyInt(); // 9123456

mask('18123458K');                               // "**.***.458-K"
mask('18123458K', { maskChar: 'X' });            // "XX.XXX.458-K"
mask('18123458K', { formatted: false });         // "*****458K"

generate();                   // RUT válido aleatorio (1M - 28M)
generatePerson();             // Persona natural (5M - 28M)
generateBusiness();           // Persona jurídica / empresa (70M - 80M)
generateBetween(10e6, 20e6);  // Rango personalizado
```

---

## 2. Dinero e Impuestos (`@urzuaf/clutils/money`)

```ts
import {
  formatCLP,
  parseCLP,
  formatUF,
  formatUTM,
  calculateIVA,
  toWords,
} from '@urzuaf/clutils/money';

formatCLP(1500000);                                     // "$1.500.000"
formatCLP(1500000, { symbol: 'CLP', space: true });     // "CLP 1.500.000"
formatCLP(-2500, { negativeFormat: 'parentheses' });    // "($2.500)"

parseCLP('$ 1.500.000');   // 1500000
parseCLP('-CLP 2.500');    // -2500
parseCLP('($2.500)');      // -2500

formatUF(38450.25);                         // "UF 38.450,25"
formatUF(38450.2548, { decimals: 4 });      // "UF 38.450,2548"
formatUTM(67294);                           // "UTM 67.294"

// Regla de cuadratura del SII (Neto + IVA === Bruto)
calculateIVA(1000, false); // { neto: 1000, iva: 190, bruto: 1190 }
calculateIVA(1000, true);  // { neto: 840, iva: 160, bruto: 1000 }

toWords(1500000);                   // "un millón quinientos mil pesos"
toWords(2000000);                   // "dos millones de pesos"
toWords(100, { noCurrency: true }); // "cien"
```

---

## 3. Fecha y Husos Horarios (`@urzuaf/clutils/date`)

```ts
import {
  parseDate,
  formatDate,
  isWeekend,
  nowChile,
  nowMagallanes,
  nowEasterIsland,
  toChileTime,
  toMagallanesTime,
  toEasterIslandTime,
  TZ_SANTIAGO,
  TZ_PUNTA_ARENAS,
  TZ_EASTER,
} from '@urzuaf/clutils/date';

const d = parseDate('18/09/2026');
formatDate(d, 'numeric');        // "18/09/2026"
formatDate(d, 'numeric_hyphen'); // "18-09-2026"
formatDate(d, 'formal');         // "18 de septiembre de 2026"
formatDate(d, 'full');           // "Viernes, 18 de septiembre de 2026"
isWeekend(d);                    // false

nowChile();                      // America/Santiago (UTC-4 invierno / UTC-3 verano)
nowMagallanes();                 // America/Punta_Arenas (UTC-3 permanente)
nowEasterIsland();               // Pacific/Easter (UTC-6 invierno / UTC-5 verano)

const utc = new Date('2026-07-15T15:00:00Z');
toChileTime(utc);
toMagallanesTime(utc);
toEasterIslandTime(utc);
```

---

## 4. Telecomunicaciones (`@urzuaf/clutils/telecom`)

```ts
import {
  normalizePhone,
  normalizeMobile,
  validatePhone,
  formatPhone,
  getPhoneType,
  isEmergencyNumber,
} from '@urzuaf/clutils/telecom';

normalizePhone('9 1234 5678');         // "+56912345678"
normalizePhone('02 2123 4567');        // "+56221234567"
normalizeMobile('5612 3456');          // "+56956123456"

validatePhone('+56 9 1234 5678');      // true
validatePhone('12345');                // false

formatPhone('912345678');                  // "+56 9 1234 5678"
formatPhone('912345678', 'national');      // "9 1234 5678"
formatPhone('322123456', 'international'); // "+56 32 212 3456"

getPhoneType('+56 9 1234 5678');       // "mobile"
getPhoneType('+56 44 123 4567');       // "voip"
getPhoneType('800 123 456');           // "toll_free"
getPhoneType('600 600 1234');          // "shared_cost"
getPhoneType('133');                   // "emergency"

isEmergencyNumber('133');              // true (Carabineros)
isEmergencyNumber('1411');             // true (Metro de Santiago)
```

---

## 5. Geografía y DPA Oficial (`@urzuaf/clutils/geo`)

Catálogo completo oficial de Chile (SUBDERE / INE): **16 regiones**, **56 provincias** y **346 comunas** con Códigos Únicos Territoriales (CUT).

```ts
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
} from '@urzuaf/clutils/geo';

getRegiones();                  // 16 regiones ordenadas de Norte a Sur
findRegion('CL-RM');            // { code: "13", isoCode: "CL-RM", name: "Metropolitana de Santiago", ... }
findRegion('VIII');             // { code: "08", isoCode: "CL-BI", name: "Biobío", ... }

getProvinciasByRegion('16');    // Diguillín, Itata, Punilla
getComunasByRegion('CL-NB');    // 21 comunas de Ñuble
getComunasByProvincia('131');   // 32 comunas de la Provincia de Santiago

findComuna('13120');            // { cut: "13120", name: "Ñuñoa", provinciaCode: "131", regionCode: "13" }
findComuna('nunoa');            // Tolerante a tildes, eñes, mayúsculas y Unicode NFD
searchComunas('San');           // Múltiples coincidencias ordenadas por relevancia (autocompletado)

getRegionByComuna('nunoa');     // Metropolitana de Santiago
getProvinciaByComuna('chillan'); // Diguillín
getRegionByProvincia('161');    // Ñuble
```

---

## Licencia

Distribuido bajo licencia **MIT**. Consulta el archivo [LICENSE](./LICENSE) para más detalles.
