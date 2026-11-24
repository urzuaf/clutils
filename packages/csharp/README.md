# Clutils (.NET / C#)

Suite de utilidades atómicas y sin dependencias externas para proyectos .NET que operan en o con **Chile**.

- **Cero dependencias externas:** basado únicamente en `netstandard2.0` (compatible con .NET Framework 4.6.1+, .NET Core 2+, .NET 5/6/7/8/9/10/11).
- **100% offline y determinista.**
- **Paridad garantizada:** validado contra los mismos vectores de prueba compartidos que Go y TypeScript.

## Instalación

```bash
dotnet add package Clutils
```

## Uso (`Clutils.Rut`)

```csharp
using Clutils;

// Validación simple y detallada
bool isValid = Rut.Validate("12.345.678-5"); // true

var err = Rut.ValidateDetailed("18.12A.458-K");
if (err != null)
{
    // err.Code == RutErrorCode.InvalidCharacters (4)
    Console.WriteLine(err.Message);
}

// Limpieza, cálculo de DV, formateo, split y masking
Rut.Clean(" 18.123.458-k ");                        // "18123458K"
Rut.CalculateDv(18123458);                          // "K"
Rut.Format("18123458k");                            // "18.123.458-K"
Rut.Format("18123458k", new RutFormatOptions { Dots = false }); // "18123458-K"

var parts = Rut.Split("09.123.456-4");
// parts.Body == "09123456", parts.Dv == "4", parts.BodyInt() == 9123456

Rut.Mask("18123458K");                              // "**.***.458-K"
Rut.Mask("18123458K", new RutMaskOptions { MaskChar = 'X' }); // "XX.XXX.458-K"

// Generación sintética de RUTs válidos
Rut.Generate();         // Rango general (1M - 28M)
Rut.GeneratePerson();   // Persona natural (5M - 28M)
Rut.GenerateBusiness(); // Empresa (70M - 80M)
```

## Licencia

Distribuido bajo licencia **MIT**.
