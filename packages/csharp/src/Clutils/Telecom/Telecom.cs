using System;
using System.Collections.Generic;
using System.Text;

namespace Clutils
{
    /// <summary>
    /// Utilidades canónicas para normalización E.164, validación, formateo visual y clasificación
    /// de números telefónicos chilenos según el plan de numeración de la SUBTEL.
    /// </summary>
    public static class Telecom
    {
        private static readonly HashSet<string> RegionalAreaCodes = new HashSet<string>(StringComparer.Ordinal)
        {
            "32", // Valparaíso / Isla de Pascua
            "33", // Quillota
            "34", // Los Andes / San Felipe
            "35", // San Antonio
            "41", // Concepción
            "42", // Chillán
            "43", // Los Ángeles
            "45", // Temuco
            "51", // La Serena / Coquimbo
            "52", // Copiapó
            "53", // Ovalle
            "55", // Antofagasta / Calama
            "57", // Iquique
            "58", // Arica
            "61", // Punta Arenas
            "63", // Valdivia
            "64", // Osorno
            "65", // Puerto Montt
            "67", // Coyhaique
            "71", // Talca
            "72", // Rancagua / O'Higgins
            "73", // Linares
            "75", // Curicó
        };

        private static readonly HashSet<string> EmergencyNumbers = new HashSet<string>(StringComparer.Ordinal)
        {
            "130",  // CONAF Incendios Forestales
            "131",  // SAMU Ambulancia
            "132",  // Bomberos
            "133",  // Carabineros
            "134",  // PDI Policía de Investigaciones
            "135",  // Fono Drogas
            "136",  // Cuerpo de Socorro Andino
            "137",  // Rescate Marítimo Armada
            "138",  // Rescate Aéreo FACh
            "139",  // Fono Información Policial
            "143",  // Fono Acción
            "147",  // Fono Niños
            "149",  // Fono Familia Carabineros
            "911",  // Central de Carabineros más cercana (móviles)
            "1404", // Servicio de Rescate
            "1407", // Fono Emergencia (Mutual de Seguridad)
            "1411", // Fono Emergencia (Metro de Santiago)
            "1412", // Fono Drogas (SENDA)
            "1488", // Fono Denuncia (Metro de Santiago)
        };

        private sealed class ParsedPhone
        {
            public string NationalNumber { get; }
            public PhoneType Type { get; }
            public string AreaCode { get; }

            public ParsedPhone(string nationalNumber, PhoneType type, string areaCode)
            {
                NationalNumber = nationalNumber;
                Type = type;
                AreaCode = areaCode;
            }
        }

        private static ParsedPhone ParsePhone(string? phone)
        {
            if (string.IsNullOrWhiteSpace(phone))
            {
                throw new FormatException("el número telefónico no puede estar vacío");
            }

            string trimmed = phone!.Trim();

            var digitsBuilder = new StringBuilder(trimmed.Length);
            for (int i = 0; i < trimmed.Length; i++)
            {
                char ch = trimmed[i];
                if (ch >= '0' && ch <= '9')
                {
                    digitsBuilder.Append(ch);
                }
                else if (ch != '+' && ch != '-' && ch != ' ' && ch != '(' && ch != ')' && ch != '.' && ch != '/')
                {
                    throw new FormatException("número telefónico chileno inválido");
                }
            }

            string digits = digitsBuilder.ToString();
            if (digits.Length == 0)
            {
                throw new FormatException("el número telefónico no puede estar vacío");
            }

            string national;

            // Prefijo país 56 solo si tiene al menos 10 dígitos (evita colisión con móviles de 8 dígitos tipo "5612 3456")
            if (digits.StartsWith("56", StringComparison.Ordinal) && digits.Length >= 10)
            {
                string rest = digits.Substring(2);
                if (rest.Length == 8)
                {
                    national = "9" + rest;
                }
                else if (rest.Length == 9)
                {
                    national = rest;
                }
                else if (rest.Length == 10 && rest.StartsWith("600", StringComparison.Ordinal))
                {
                    national = rest;
                }
                else
                {
                    throw new FormatException("número telefónico chileno inválido");
                }
            }
            else if (digits.StartsWith("0", StringComparison.Ordinal))
            {
                string rest = digits.Substring(1);
                if (rest.Length == 9)
                {
                    national = rest;
                }
                else
                {
                    throw new FormatException("número telefónico chileno inválido");
                }
            }
            else if (digits.Length == 8)
            {
                national = "9" + digits;
            }
            else if (digits.Length == 9)
            {
                national = digits;
            }
            else if (digits.Length == 10 && digits.StartsWith("600", StringComparison.Ordinal))
            {
                national = digits;
            }
            else
            {
                throw new FormatException("número telefónico chileno inválido");
            }

            // 1. Móviles: inician con 9 y tienen 9 dígitos
            if (national.Length == 9 && national[0] == '9')
            {
                return new ParsedPhone(national, PhoneType.Mobile, "9");
            }

            // 2. Red Fija Santiago (RM): inicia con 2 y tiene 9 dígitos
            if (national.Length == 9 && national[0] == '2')
            {
                return new ParsedPhone(national, PhoneType.Landline, "2");
            }

            // 3. Red Fija Regional: código de área de 2 dígitos y 9 dígitos en total
            if (national.Length == 9)
            {
                string prefix2 = national.Substring(0, 2);
                if (RegionalAreaCodes.Contains(prefix2))
                {
                    return new ParsedPhone(national, PhoneType.Landline, prefix2);
                }
            }

            // 4. Telefonía IP (VoIP): código de área virtual 44 y 9 dígitos en total
            if (national.Length == 9 && national.StartsWith("44", StringComparison.Ordinal))
            {
                return new ParsedPhone(national, PhoneType.VoIP, "44");
            }

            // 5. Toll-Free: 800 XXXXXX (9 dígitos)
            if (national.Length == 9 && national.StartsWith("800", StringComparison.Ordinal))
            {
                return new ParsedPhone(national, PhoneType.TollFree, "800");
            }

            // 6. Shared-Cost: 600 XXXXXX (9 dígitos) o 600 XXXXXXX (10 dígitos)
            if ((national.Length == 9 || national.Length == 10) && national.StartsWith("600", StringComparison.Ordinal))
            {
                return new ParsedPhone(national, PhoneType.SharedCost, "600");
            }

            throw new FormatException("número telefónico chileno inválido");
        }

        /// <summary>
        /// Normaliza cualquier número telefónico chileno (móvil, red fija o líneas especiales)
        /// al formato canónico internacional E.164 (<c>+569XXXXXXXX</c>, <c>+56221234567</c>, etc.).
        /// </summary>
        /// <exception cref="FormatException">Si el número es vacío o inválido según SUBTEL.</exception>
        /// <example>
        /// <code>
        /// Telecom.NormalizePhone("9 1234 5678"); // "+56912345678"
        /// Telecom.NormalizePhone("02 2123 4567"); // "+56221234567"
        /// </code>
        /// </example>
        public static string NormalizePhone(string? phone) => "+56" + ParsePhone(phone).NationalNumber;

        /// <summary>
        /// Normaliza específicamente un número móvil celular chileno al estándar E.164 (<c>+569XXXXXXXX</c>).
        /// </summary>
        /// <exception cref="FormatException">Si el número es inválido o no corresponde a telefonía móvil.</exception>
        public static string NormalizeMobile(string? phone)
        {
            ParsedPhone parsed = ParsePhone(phone);
            if (parsed.Type != PhoneType.Mobile)
            {
                throw new FormatException("el número telefónico no corresponde a una línea móvil");
            }

            return "+56" + parsed.NationalNumber;
        }

        /// <summary>
        /// Verifica si un número telefónico es válido y normalizable a E.164 según el plan de numeración de SUBTEL.
        /// </summary>
        public static bool ValidatePhone(string? phone)
        {
            try
            {
                ParsePhone(phone);
                return true;
            }
            catch (FormatException)
            {
                return false;
            }
        }

        /// <summary>
        /// Formatea un número telefónico chileno en el estilo visual solicitado.
        /// </summary>
        /// <example>
        /// <code>
        /// Telecom.FormatPhone("912345678"); // "+56 9 1234 5678"
        /// Telecom.FormatPhone("322123456", PhoneFormatStyle.National); // "32 212 3456"
        /// </code>
        /// </example>
        public static string FormatPhone(string? phone, PhoneFormatStyle style = PhoneFormatStyle.International)
        {
            ParsedPhone parsed = ParsePhone(phone);
            string nat = parsed.NationalNumber;

            if (style == PhoneFormatStyle.E164)
            {
                return "+56" + nat;
            }

            string formattedNational;
            switch (parsed.Type)
            {
                case PhoneType.Mobile:
                    formattedNational = nat.Substring(0, 1) + " " + nat.Substring(1, 4) + " " + nat.Substring(5);
                    break;
                case PhoneType.Landline:
                    formattedNational = parsed.AreaCode == "2"
                        ? nat.Substring(0, 1) + " " + nat.Substring(1, 4) + " " + nat.Substring(5)
                        : nat.Substring(0, 2) + " " + nat.Substring(2, 3) + " " + nat.Substring(5);
                    break;
                case PhoneType.VoIP:
                    formattedNational = nat.Substring(0, 2) + " " + nat.Substring(2, 3) + " " + nat.Substring(5);
                    break;
                case PhoneType.TollFree:
                case PhoneType.SharedCost:
                    formattedNational = nat.Substring(0, 3) + " " + nat.Substring(3, 3) + " " + nat.Substring(6);
                    break;
                default:
                    formattedNational = nat;
                    break;
            }

            return style == PhoneFormatStyle.National ? formattedNational : "+56 " + formattedNational;
        }

        /// <summary>
        /// Retorna la clasificación oficial del número telefónico chileno según SUBTEL
        /// (móvil, red fija, VoIP, 800, 600, emergencia o desconocido).
        /// </summary>
        public static PhoneType GetPhoneType(string? phone)
        {
            if (string.IsNullOrWhiteSpace(phone))
            {
                return PhoneType.Unknown;
            }

            string trimmed = phone!.Trim();
            var cleanDigits = new StringBuilder(trimmed.Length);
            bool hasOnlyDigitsOrSeparators = true;

            for (int i = 0; i < trimmed.Length; i++)
            {
                char ch = trimmed[i];
                if (ch >= '0' && ch <= '9')
                {
                    cleanDigits.Append(ch);
                }
                else if (ch != ' ' && ch != '-')
                {
                    hasOnlyDigitsOrSeparators = false;
                }
            }

            string cleanStr = cleanDigits.ToString();
            if (hasOnlyDigitsOrSeparators && (cleanStr.Length == 3 || cleanStr.Length == 4))
            {
                if (EmergencyNumbers.Contains(cleanStr))
                {
                    return PhoneType.Emergency;
                }
            }

            try
            {
                return ParsePhone(phone).Type;
            }
            catch (FormatException)
            {
                return PhoneType.Unknown;
            }
        }

        /// <summary>
        /// Verifica si una cadena corresponde a un número corto de emergencia oficial en Chile.
        /// </summary>
        public static bool IsEmergencyNumber(string? phone) => GetPhoneType(phone) == PhoneType.Emergency;
    }
}
