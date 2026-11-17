# clutils

El objetivo de `clutils` es evitar reimplementar desde cero en cada proyecto validaciones, formatos y reglas normativas locales, ofreciendo funciones atómicas, sin dependencias externas y fuertemente testeadas.

---

## 📦 Ecosistemas Soportados

- **Go:** `github.com/urzuaf/clutils/packages/go`
- **TypeScript / JavaScript:** *(Próximamente vía npm)*

---

## 🚀 Instalación (Go)

```bash
go get github.com/urzuaf/clutils/packages/go
```

---

## 🇨🇱 Módulo de RUT (`rut`)

El módulo `rut` proporciona herramientas completas para sanitizar, validar, formatear y generar Rol Único Tributario (RUT / RUN).

```go
import "github.com/urzuaf/clutils/packages/go/rut"
```

### 1. Validación

```go
// Validación booleana simple
isValid := rut.Validate("12.345.678-5") // true
isInvalid := rut.Validate("12.345.678-0") // false

// Validación detallada con códigos numéricos de error (ideal para APIs y formularios)
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

### 2. Limpieza y Normalización

```go
// Quita puntos, guiones y espacios, normalizando 'k' a mayúscula 'K'
clean := rut.Clean("  18.123.458-k  ") // "18123458K"
```

### 3. Formateo Visual

```go
// Formato formal por defecto (con puntos y guión)
formatted := rut.Format("18123458K") // "18.123.458-K"

// Formato personalizable
noDots := rut.Format("18123458K", rut.FormatOptions{Dots: false, Hyphen: true}) // "18123458-K"
plain := rut.Format("18.123.458-K", rut.FormatOptions{Dots: false, Hyphen: false}) // "18123458K"
```

### 4. Separador de Cuerpo y Dígito Verificador (`Split`)

Útil para bases de datos bancarias y sistemas legados con columnas separadas para el cuerpo y el DV:

```go
parts, err := rut.Split("18.123.458-K")
// parts.Body ➔ "18123458" (preserva ceros iniciales como string)
// parts.Dv   ➔ "K"

// Obtener el cuerpo directamente como entero
bodyInt, _ := parts.BodyInt() // 18123458 (int)
```

### 5. Ofuscación de Privacidad (`Mask`)

Permite ocultar caracteres del RUT para cumplimiento de la Ley 19.628 de Protección de Datos Personales en pantallas públicas y comprobantes:

```go
// Por defecto: oculta los caracteres iniciales y preserva los últimos 4 caracteres
masked := rut.Mask("18123458K") // "**.***.458-K"

// Sin formato (plano)
plainMasked := rut.Mask("18123458K", rut.MaskOptions{Formatted: false}) // "*****458K"
```

### 6. Generador de RUTs Válidos para Testing

Genera RUTs sintéticos con Dígito Verificador matemáticamente válido para suites de pruebas unitarias, factories y seeders:

```go
// Generación aleatoria general
r1 := rut.Generate() // ej: "19.458.123-5"

// Personas naturales (rango 5M a 28M)
person := rut.GeneratePerson() // ej: "21.123.456-7"

// Personas jurídicas / Empresas (rango 60M a 90M)
business := rut.GenerateBusiness() // ej: "76.458.123-0"

// Rango personalizado
custom := rut.GenerateBetween(1_000_000, 10_000_000)
```

---

## 📄 Licencia

MIT License © 2026
