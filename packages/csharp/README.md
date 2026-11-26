# Clutils (.NET / C#)

Suite de utilidades canónicas, atómicas y sin dependencias externas para proyectos .NET que operan en o con **Chile**.

- **Cero dependencias externas:** basado únicamente en `netstandard2.0` (compatible con .NET Framework 4.6.1+, .NET Core 2+, .NET 5/6/7/8/9/10/11).
- **100% offline y determinista.**
- **Paridad garantizada:** validado contra los mismos vectores de prueba compartidos (`specs/tests/`) que Go y TypeScript.

## Instalación

```bash
dotnet add package Clutils
```

---

## 1. RUT / RUN (`Clutils.Rut`)

```csharp
using Clutils;

bool isValid = Rut.Validate("12.345.678-5"); // true

var err = Rut.ValidateDetailed("18.12A.458-K");
if (err != null)
{
    // err.Code == RutErrorCode.InvalidCharacters (4)
    Console.WriteLine(err.Message);
}

Rut.Clean(" 18.123.458-k ");                                    // "18123458K"
Rut.CalculateDv(18123458);                                      // "K"
Rut.Format("18123458k");                                        // "18.123.458-K"
Rut.Format("18123458k", new RutFormatOptions { Dots = false }); // "18123458-K"

var parts = Rut.Split("09.123.456-4");
// parts.Body == "09123456", parts.Dv == "4", parts.BodyInt() == 9123456

Rut.Mask("18123458K");                                          // "**.***.458-K"
Rut.Mask("18123458K", new RutMaskOptions { MaskChar = 'X' });   // "XX.XXX.458-K"

Rut.Generate();         // Rango general (1M - 28M)
Rut.GeneratePerson();   // Persona natural (5M - 28M)
Rut.GenerateBusiness(); // Empresa (70M - 80M)
```

---

## 2. Dinero e Impuestos (`Clutils.Money`)

```csharp
using Clutils;

Money.FormatCLP(1500000); // "$1.500.000"
Money.FormatCLP(-2500, new ClpFormatOptions { NegativeFormat = NegativeFormat.Parentheses }); // "($2.500)"
Money.ParseCLP("-CLP 2.500"); // -2500

Money.FormatUF(38450.25);                                    // "UF 38.450,25"
Money.FormatUF(38450.2548, new UfFormatOptions { Decimals = 4 }); // "UF 38.450,2548"
Money.FormatUTM(67294);                                      // "UTM 67.294"

// Cuadratura oficial SII (Neto + Iva == Bruto) con redondeo AwayFromZero
var iva = Money.CalculateIVA(1000, isBruto: true); // Neto: 840, Iva: 160, Bruto: 1000

Money.ToWords(1500000); // "un millón quinientos mil pesos"
Money.ToWords(2000000); // "dos millones de pesos"
```

---

## 3. Fecha y Husos Horarios (`Clutils.Date`)

```csharp
using Clutils;

var dt = Date.ParseDate("18/09/2026");
Date.FormatDate(dt, DateFormatStyle.Numeric);       // "18/09/2026"
Date.FormatDate(dt, DateFormatStyle.NumericHyphen); // "18-09-2026"
Date.FormatDate(dt, DateFormatStyle.Formal);        // "18 de septiembre de 2026"
Date.FormatDate(dt, DateFormatStyle.Full);          // "Viernes, 18 de septiembre de 2026"
Date.IsWeekend(dt);                                 // false

Date.NowChile();       // America/Santiago (UTC-4 invierno / UTC-3 verano)
Date.NowMagallanes();  // America/Punta_Arenas (UTC-3 permanente)
Date.NowEasterIsland(); // Pacific/Easter (UTC-6 invierno / UTC-5 verano)
```

---

## 4. Telecomunicaciones (`Clutils.Telecom`)

```csharp
using Clutils;

Telecom.NormalizePhone("9 1234 5678");                    // "+56912345678"
Telecom.NormalizeMobile("5612 3456");                     // "+56956123456"
Telecom.ValidatePhone("+56 9 1234 5678");                 // true
Telecom.FormatPhone("322123456", PhoneFormatStyle.National); // "32 212 3456"
Telecom.GetPhoneType("+56 44 123 4567");                  // PhoneType.VoIP
Telecom.IsEmergencyNumber("133");                         // true (Carabineros)
Telecom.IsEmergencyNumber("1411");                        // true (Metro de Santiago)
```

---

## 5. Geografía y DPA Oficial (`Clutils.Geo`)

Catálogo oficial de Chile (SUBDERE / INE): **16 regiones**, **56 provincias** y **346 comunas** con Códigos Únicos Territoriales (CUT).

```csharp
using Clutils;

var regiones = Geo.GetRegiones();            // 16 regiones ordenadas de Norte a Sur
var rm = Geo.FindRegion("CL-RM");            // Metropolitana de Santiago
var biobio = Geo.FindRegion("VIII");         // Biobío

var comunasNuble = Geo.GetComunasByRegion("16");   // 21 comunas de Ñuble
var provStgo = Geo.GetProvinciasByRegion("CL-RM"); // 6 provincias de RM

var nunoa = Geo.FindComuna("nunoa");         // Cut: "13120", Name: "Ñuñoa"
var matches = Geo.SearchComunas("San");      // Coincidencias ordenadas para autocompletado
var reg = Geo.GetRegionByComuna("13120");    // Metropolitana de Santiago
var prov = Geo.GetProvinciaByComuna("chillan"); // Diguillín
```

---

## Licencia

Distribuido bajo licencia **MIT**.
