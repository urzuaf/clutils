# clutils

El objetivo de `clutils` es evitar reimplementar desde cero en cada proyecto validaciones, formatos y reglas normativas locales, ofreciendo funciones atómicas, sin dependencias externas y fuertemente testeadas.

---

## 📦 Ecosistemas Soportados

- **Go:** `github.com/urzuaf/clutils/packages/go`
- **TypeScript / JavaScript:** `@urzuaf/clutils` (npm)

---

## 🚀 Instalación

### Go
```bash
go get github.com/urzuaf/clutils/packages/go
```

### TypeScript / JavaScript (Node, Bun, Deno, Navegador)
```bash
npm install @urzuaf/clutils
```

---

## 🇨🇱 Módulo de RUT (`rut`)

El módulo `rut` proporciona herramientas completas para sanitizar, validar, formatear y generar Rol Único Tributario (RUT / RUN).

### Importación

**Go:**
```go
import "github.com/urzuaf/clutils/packages/go/rut"
```

**TypeScript / JavaScript:**
```ts
import { validate, clean, format, split, mask, generate } from '@urzuaf/clutils/rut';
```

---

### 1. Validación

**Go:**
```go
isValid := rut.Validate("12.345.678-5") // true

if err := rut.ValidateDetailed("18.12A.458-K"); err != nil {
    switch err.Code {
    case rut.ErrCodeEmpty:              // 1: RUT vacío
    case rut.ErrCodeMalformedFormat:    // 2: Puntos o guiones mal ubicados
    case rut.ErrCodeInvalidLength:      // 3: Longitud fuera de rango
    case rut.ErrCodeInvalidCharacters:  // 4: Letras en el cuerpo
    case rut.ErrCodeInvalidDv:          // 5: Dígito verificador no coincide
    }
}
```

**TypeScript / JavaScript:**
```ts
const isValid = validate("12.345.678-5"); // true

const error = validateDetailed("18.12A.458-K");
if (error) {
    console.log(error.code); // 4 (RutErrorCode.InvalidCharacters)
    console.log(error.message);
}
```

---

### 2. Limpieza y Normalización

**Go:**
```go
clean := rut.Clean("  18.123.458-k  ") // "18123458K"
```

**TypeScript / JavaScript:**
```ts
const cleaned = clean("  18.123.458-k  "); // "18123458K"
```

---

### 3. Formateo Visual

**Go:**
```go
formatted := rut.Format("18123458K") // "18.123.458-K"
noDots := rut.Format("18123458K", rut.FormatOptions{Dots: false, Hyphen: true}) // "18123458-K"
```

**TypeScript / JavaScript:**
```ts
const formatted = format("18123458K"); // "18.123.458-K"
const noDots = format("18123458K", { dots: false, hyphen: true }); // "18123458-K"
```

---

### 4. Separador de Cuerpo y Dígito Verificador (`Split`)

**Go:**
```go
parts, err := rut.Split("18.123.458-K")
// parts.Body ➔ "18123458", parts.Dv ➔ "K"
bodyInt, _ := parts.BodyInt() // 18123458 (int)
```

**TypeScript / JavaScript:**
```ts
const parts = split("18.123.458-K");
// parts.body ➔ "18123458", parts.dv ➔ "K"
const bodyNum = parts.bodyInt(); // 18123458 (number)
```

---

### 5. Ofuscación de Privacidad (`Mask`)

**Go:**
```go
masked := rut.Mask("18123458K") // "**.***.458-K"
plain := rut.Mask("18123458K", rut.MaskOptions{Formatted: false}) // "*****458K"
```

**TypeScript / JavaScript:**
```ts
const masked = mask("18123458K"); // "**.***.458-K"
const plain = mask("18123458K", { formatted: false }); // "*****458K"
```

---

### 6. Generador de RUTs Válidos para Testing

**Go:**
```go
r1 := rut.Generate()         // ej: "19.458.123-5"
p1 := rut.GeneratePerson()   // rango 5M a 28M
b1 := rut.GenerateBusiness() // rango 70M a 80M
```

**TypeScript / JavaScript:**
```ts
const r1 = generate();         // ej: "19.458.123-5"
const p1 = generatePerson();   // rango 5M a 28M
const b1 = generateBusiness(); // rango 70M a 80M
```

---

## 📄 Licencia

MIT License © 2026
