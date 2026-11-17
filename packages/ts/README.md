# clutils (TypeScript / JavaScript)

Suite de utilidades canónicas, de alto rendimiento y cero dependencias externas para proyectos que operan en o con Chile.

[![npm version](https://img.shields.io/npm/v/clutils.svg)](https://www.npmjs.com/package/clutils)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Dependencies](https://img.shields.io/badge/dependencies-0-success.svg)](#)

---

## Características

- **Cero dependencias:** 0 dependencias en tiempo de ejecución.
- **Tipado estricto:** 100% TypeScript con archivos de definición `.d.ts` y mapas de tipos completos.
- **Dual Export:** Soporte nativo y optimizado para **ESM** (`import`) y **CommonJS** (`require`).

---

## Instalación

```bash
# npm
npm install clutils

# pnpm
pnpm add clutils

# bun
bun add clutils

# yarn
yarn add clutils
```

---

## Módulo de RUT (`clutils/rut`)

Manejo integral de Rol Único Tributario (RUT / RUN) chileno: validación detallada, sanitización, formateo, enmascaramiento y generación sintética.

### Importación

```ts
// ESM / TypeScript
import { 
  validate, 
  validateDetailed, 
  RutErrorCode, 
  clean, 
  format, 
  split, 
  mask, 
  generate,
  generatePerson,
  generateBusiness
} from 'clutils/rut';

// O desde la raíz:
// import { rut } from 'clutils';
```

```javascript
// CommonJS
const { validate, clean, format, mask } = require('clutils/rut');
```

---

### 1. Validación

#### Validación simple (`validate`)
Devuelve un valor booleano (`true` o `false`):

```ts
validate('12.345.678-5'); // true
validate('123456785');     // true
validate('12345678-0');     // false (DV inválido)
validate('12.34.5678-5');   // false (formato con puntos mal ubicados)
```

#### Validación detallada con códigos de error (`validateDetailed`)
Identifica con exactitud por qué falló un RUT:

```ts
const result = validateDetailed('18.12A.458-K');

if (!result.valid) {
  switch (result.error?.code) {
    case RutErrorCode.Empty:             // 1: El RUT está vacío
      console.log('Ingresa un RUT');
      break;
    case RutErrorCode.MalformedFormat:   // 2: Puntos o guiones mal ubicados
      console.log('El formato de puntos o guión es incorrecto');
      break;
    case RutErrorCode.InvalidLength:     // 3: Longitud fuera de rango (1M a 99M)
      console.log('Longitud de RUT inválida');
      break;
    case RutErrorCode.InvalidCharacters: // 4: Caracteres no numéricos en el cuerpo
      console.log('El cuerpo contiene caracteres no válidos');
      break;
    case RutErrorCode.InvalidDv:         // 5: Dígito verificador no coincide con módulo 11
      console.log(`Dígito verificador incorrecto (esperado: ${result.error.expectedDv})`);
      break;
  }
}
```

---

### 2. Limpieza y Normalización (`clean`)

Elimina puntos, espacios y caracteres separadores, conservando el cuerpo numérico y la letra verificadora mayúscula:

```ts
clean(' 12.345.678-k '); // "12345678K"
clean('7.654.321-0');    // "76543210"
```

---

### 3. Formateo Flexible (`format`)

Aplica diferentes estilos visuales de forma declarativa:

```ts
// Con puntos y guión (por defecto)
format('123456785'); 
// => "12.345.678-5"

// Sin puntos (solo guión)
format('123456785', { dots: false }); 
// => "12345678-5"

// Solo números y DV en mayúscula (sin puntos ni guión)
format('12.345.678-k', { dots: false, hyphen: false }); 
// => "12345678K"
```

---

### 4. División de Partes (`split`)

Separa el cuerpo numérico y el dígito verificador, proporcionando métodos para obtener el cuerpo como entero y el RUT normalizado:

```ts
const parts = split('12.345.678-5');

console.log(parts.body);    // "12345678"
console.log(parts.dv);      // "5"
console.log(parts.bodyInt); // 12345678 (número)
```

---

### 5. Enmascaramiento Seguro (`mask`)

Oculta los primeros dígitos del cuerpo para interfaces visuales sensibles (tarjetas, comprobantes, balances):

```ts
// Por defecto: muestra los últimos 4 dígitos del cuerpo y el DV, formateado
mask('123456785');
// => "*.345.678-5"

// Sin puntos
mask('123456785', { formatted: false });
// => "***5678-5"

// Personalizando caracteres y cantidad de dígitos visibles
mask('12.345.678-5', { visibleCount: 2, maskChar: 'X' });
// => "XX.XXX.X78-5"
```

---

### 6. Generación Sintética de RUTs (`generate`)

Generación determinista y válida de RUTs para suites de pruebas unitarias o seeds de bases de datos:

```ts
// RUT aleatorio dentro del rango general válido (1.000.000 - 28.000.000)
const rutPrueba = generate(); 

// RUT en rango de personas naturales (5.000.000 - 28.000.000)
const rutPersona = generatePerson(); 

// RUT en rango de empresas / personas jurídicas (60.000.000 - 90.000.000)
const rutEmpresa = generateBusiness(); 

// RUT dentro de un rango numérico específico
const rutRango = generateBetween(10000000, 15000000, { dots: false }); 
// => "12345678-5"
```

---

## 📄 Licencia

MIT © urzuaf
