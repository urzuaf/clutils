# clutils

Suite de utilidades atómicas y sin dependencias externas para proyectos que operan en o con **Chile**.

El objetivo de `clutils` es evitar reimplementar en cada proyecto las mismas validaciones, formatos, cálculos normativos y catálogos territoriales chilenos, garantizando paridad exacta entre ecosistemas mediante una fuente única de verdad (`data/`) y vectores de prueba compartidos (`specs/tests/`).

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](./LICENSE)
[![Zero Dependencies](https://img.shields.io/badge/dependencies-0-success.svg)](#)

---

## Ecosistemas Soportados (`v1.0.0`)

- **Go:** `github.com/urzuaf/clutils/packages/go`
- **TypeScript / JavaScript:** `@urzuaf/clutils` (ESM + CJS + `.d.ts`, compatible con Node.js, Bun, Deno, navegadores y Edge Runtimes)

## Instalación

**Go:**
```bash
go get github.com/urzuaf/clutils/packages/go@v1.0.0
```

**TypeScript / JavaScript:**
```bash
npm install @urzuaf/clutils
```

---

## Dominios incluidos

| Dominio | Descripción | Go | TypeScript (Subpath) |
| :--- | :--- | :--- | :--- |
| **`rut`** | Validación, errores detallados, limpieza, formato, split, masking y generación sintética | `.../packages/go/rut` | `@urzuaf/clutils/rut` |
| **`money`** | CLP, UF (CMF), UTM, IVA 19% (cuadratura SII) y montos a palabras | `.../packages/go/money` | `@urzuaf/clutils/money` |
| **`date`** | 3 husos horarios oficiales, formatos locales, parseo estricto `DD/MM/YYYY` y fin de semana | `.../packages/go/date` | `@urzuaf/clutils/date` |
| **`telecom`** | Normalización E.164 (`+56...`), móviles, red fija por área, VoIP `44`, `800`/`600` y emergencias | `.../packages/go/telecom` | `@urzuaf/clutils/telecom` |
| **`geo`** | DPA oficial SUBDERE/INE (16 regiones, 56 provincias, 346 comunas con CUT), jerarquía y búsqueda | `.../packages/go/geo` | `@urzuaf/clutils/geo` |

> En TypeScript también puedes importar los 5 namespaces desde la raíz:
> ```ts
> import { rut, money, date, telecom, geo } from '@urzuaf/clutils';
> ```
> Para el mejor *tree-shaking* en frontend, prefiere los subpaths (`@urzuaf/clutils/<dominio>`).

---

## 1. RUT / RUN (`rut`)

### Go
```go
import "github.com/urzuaf/clutils/packages/go/rut"

rut.Validate("12.345.678-5")     // true
rut.Clean(" 18.123.458-k ")      // "18123458K"
rut.Format("18123458k")          // "18.123.458-K"
rut.Mask("18123458K")            // "**.***.458-K"
parts, _ := rut.Split("18.123.458-K") // Body: "18123458", Dv: "K"

if err := rut.ValidateDetailed("18.12A.458-K"); err != nil {
    // err.Code == rut.ErrCodeInvalidCharacters (4)
}

rut.Generate()         // RUT aleatorio válido (1M - 28M)
rut.GeneratePerson()   // Persona natural (5M - 28M)
rut.GenerateBusiness() // Empresa (70M - 80M)
```

### TypeScript / JavaScript
```ts
import {
  validate,
  validateDetailed,
  RutErrorCode,
  clean,
  format,
  split,
  mask,
  generatePerson,
  generateBusiness,
} from '@urzuaf/clutils/rut';

validate('12.345.678-5'); // true
clean(' 18.123.458-k ');  // "18123458K"
format('18123458k');      // "18.123.458-K"
mask('18123458K');        // "**.***.458-K"
split('18.123.458-K');    // { body: "18123458", dv: "K", bodyInt: [Function] }

const err = validateDetailed('18.12A.458-K');
if (err?.code === RutErrorCode.InvalidCharacters) {
  console.log(err.message);
}

generatePerson();   // RUT persona natural válido (5M - 28M)
generateBusiness(); // RUT empresa válido (70M - 80M)
```

---

## 2. Dinero e Impuestos (`money`)

Formateo y parseo de pesos chilenos (CLP), UF (convención CMF con 2 o 4 decimales), UTM, desglose de IVA 19% bajo cuadratura del SII (`Neto + IVA = Bruto`) y conversión de montos a palabras formales en español chileno.

### Go
```go
import "github.com/urzuaf/clutils/packages/go/money"

money.FormatCLP(1500000) // "$1.500.000"
money.FormatCLP(-2500, money.CLPFormatOptions{NegativeFormat: money.NegativeFormatParentheses}) // "($2.500)"
val, _ := money.ParseCLP("-CLP 2.500") // -2500

money.FormatUF(38450.25) // "UF 38.450,25"
money.FormatUF(38450.2548, money.UFFormatOptions{Decimals: 4}) // "UF 38.450,2548"
money.FormatUTM(67294) // "UTM 67.294"

iva := money.CalculateIVA(1000, true) // {Neto: 840, IVA: 160, Bruto: 1000}
money.ToWords(2000000) // "dos millones de pesos"
```

### TypeScript / JavaScript
```ts
import { formatCLP, parseCLP, formatUF, formatUTM, calculateIVA, toWords } from '@urzuaf/clutils/money';

formatCLP(1500000); // "$1.500.000"
formatCLP(-2500, { negativeFormat: 'parentheses' }); // "($2.500)"
parseCLP('-CLP 2.500'); // -2500

formatUF(38450.25); // "UF 38.450,25"
formatUF(38450.2548, { decimals: 4 }); // "UF 38.450,2548"
formatUTM(67294); // "UTM 67.294"

calculateIVA(1000, true); // { neto: 840, iva: 160, bruto: 1000 }
toWords(2000000); // "dos millones de pesos"
```

---

## 3. Fecha y Husos Horarios (`date`)

Soporte para las 3 zonas horarias legales de Chile (`America/Santiago`, `America/Punta_Arenas` en UTC-3 permanente, y `Pacific/Easter`), parseo estricto `DD/MM/YYYY` sin inversión anglosajona y formateo local.

### Go
```go
import "github.com/urzuaf/clutils/packages/go/date"

d, _ := date.ParseDate("18/09/2026")
date.FormatDate(d, date.DateFormatNumeric) // "18/09/2026"
date.FormatDate(d, date.DateFormatFormal)  // "18 de septiembre de 2026"
date.FormatDate(d, date.DateFormatFull)    // "Viernes, 18 de septiembre de 2026"
date.IsWeekend(d)                          // false

date.NowChile()       // Hora actual en America/Santiago
date.NowMagallanes()  // Hora actual en America/Punta_Arenas
date.NowEasterIsland() // Hora actual en Pacific/Easter
```

### TypeScript / JavaScript
```ts
import {
  parseDate,
  formatDate,
  isWeekend,
  nowChile,
  nowMagallanes,
  nowEasterIsland,
} from '@urzuaf/clutils/date';

const d = parseDate('18/09/2026');
formatDate(d, 'numeric'); // "18/09/2026"
formatDate(d, 'formal');  // "18 de septiembre de 2026"
formatDate(d, 'full');    // "Viernes, 18 de septiembre de 2026"
isWeekend(d);             // false

nowChile();
nowMagallanes();
nowEasterIsland();
```

---

## 4. Telecomunicaciones (`telecom`)

Normalización al estándar internacional E.164 (`+56...`), formateo visual y clasificación según el plan de numeración de la SUBTEL (móviles, red fija Santiago y regional, VoIP `44`, líneas `800`/`600` y números cortos de emergencia).

### Go
```go
import "github.com/urzuaf/clutils/packages/go/telecom"

telecom.NormalizePhone("9 1234 5678")     // "+56912345678", nil
telecom.NormalizeMobile("5612 3456")      // "+56956123456", nil
telecom.FormatPhone("912345678")          // "+56 9 1234 5678", nil
telecom.FormatPhone("322123456", telecom.FormatNational) // "32 212 3456", nil
telecom.GetPhoneType("+56 44 123 4567")   // telecom.PhoneTypeVoIP ("voip")
telecom.IsEmergencyNumber("133")          // true
telecom.IsEmergencyNumber("1411")         // true
```

### TypeScript / JavaScript
```ts
import {
  normalizePhone,
  normalizeMobile,
  formatPhone,
  getPhoneType,
  isEmergencyNumber,
} from '@urzuaf/clutils/telecom';

normalizePhone('9 1234 5678');         // "+56912345678"
normalizeMobile('5612 3456');          // "+56956123456"
formatPhone('912345678');              // "+56 9 1234 5678"
formatPhone('322123456', 'national');  // "32 212 3456"
getPhoneType('+56 44 123 4567');       // "voip"
isEmergencyNumber('133');              // true
isEmergencyNumber('1411');             // true
```

---

## 5. Geografía y DPA Oficial (`geo`)

División Político-Administrativa oficial de Chile (SUBDERE / INE): **16 regiones**, **56 provincias** y **346 comunas** con Códigos Únicos Territoriales (CUT), filtrado jerárquico, resolución inversa y búsqueda tolerante a tildes, eñes, mayúsculas, números romanos y Unicode NFD.

### Go
```go
import "github.com/urzuaf/clutils/packages/go/geo"

regiones := geo.GetRegiones()            // 16 regiones ordenadas de Norte a Sur
rm, _ := geo.FindRegion("CL-RM")         // Metropolitana de Santiago
biobio, _ := geo.FindRegion("VIII")      // Biobío

comunasNuble := geo.GetComunasByRegion("16")     // 21 comunas de Ñuble
provStgo := geo.GetProvinciasByRegion("CL-RM")   // 6 provincias de RM

nunoa, _ := geo.FindComuna("nunoa")              // {Cut: "13120", Name: "Ñuñoa", ...}
matches := geo.SearchComunas("San")              // Coincidencias ordenadas para autocompletado
reg, _ := geo.GetRegionByComuna("13120")         // Metropolitana de Santiago
prov, _ := geo.GetProvinciaByComuna("chillan")   // Diguillín
```

### TypeScript / JavaScript
```ts
import {
  getRegiones,
  findRegion,
  getProvinciasByRegion,
  getComunasByRegion,
  findComuna,
  searchComunas,
  getRegionByComuna,
  getProvinciaByComuna,
} from '@urzuaf/clutils/geo';

const regiones = getRegiones();          // 16 regiones ordenadas de Norte a Sur
const rm = findRegion('CL-RM');          // Metropolitana de Santiago
const biobio = findRegion('VIII');       // Biobío

const comunasNuble = getComunasByRegion('16');     // 21 comunas de Ñuble
const provStgo = getProvinciasByRegion('CL-RM');   // 6 provincias de RM

const nunoa = findComuna('nunoa');                 // { cut: "13120", name: "Ñuñoa", ... }
const matches = searchComunas('San');              // Coincidencias ordenadas para autocompletado
const reg = getRegionByComuna('13120');            // Metropolitana de Santiago
const prov = getProvinciaByComuna('chillan');      // Diguillín
```

---

## Arquitectura del Monorepo

```text
clutils/
├── data/            # Fuente única de verdad (DPA oficial INE/SUBDERE, rangos RUT)
├── specs/tests/     # Vectores de prueba compartidos entre todos los lenguajes
└── packages/
    ├── go/          # Módulo Go estándar (github.com/urzuaf/clutils/packages/go)
    └── ts/          # Paquete npm (@urzuaf/clutils) con ESM, CJS y .d.ts
```

## Licencia

Distribuido bajo licencia **MIT**. Consulta el archivo [LICENSE](./LICENSE) para más detalles.
