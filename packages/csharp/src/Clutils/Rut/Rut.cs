using System;
using System.Globalization;
using System.Text;
using System.Threading;

namespace Clutils
{
    /// <summary>
    /// Utilidades canónicas para sanitizar, validar, formatear, enmascarar y generar RUT / RUN chilenos.
    /// </summary>
    public static class Rut
    {
        private static readonly ThreadLocal<Random> ThreadRandom =
            new ThreadLocal<Random>(() => new Random(Guid.NewGuid().GetHashCode()));

        /// <summary>
        /// Sanitiza una cadena de RUT eliminando puntos, guiones y espacios en blanco,
        /// y normalizando el dígito verificador 'k' a mayúscula 'K'.
        /// </summary>
        /// <example>
        /// <code>
        /// Rut.Clean(" 18.123.458-k "); // "18123458K"
        /// </code>
        /// </example>
        public static string Clean(string? inputRut)
        {
            if (string.IsNullOrEmpty(inputRut))
            {
                return string.Empty;
            }

            var sb = new StringBuilder(inputRut!.Length);
            for (int i = 0; i < inputRut.Length; i++)
            {
                char ch = inputRut[i];
                if (ch == '.' || ch == '-' || ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r')
                {
                    continue;
                }

                if (ch == 'k')
                {
                    sb.Append('K');
                    continue;
                }

                sb.Append(ch);
            }

            return sb.ToString();
        }

        /// <summary>
        /// Calcula el Dígito Verificador oficial para un cuerpo numérico
        /// aplicando el algoritmo de Módulo 11 con ponderadores cíclicos [2..7].
        /// Retorna "0"-"9" o "K", o cadena vacía si el cuerpo es menor o igual a cero.
        /// </summary>
        /// <example>
        /// <code>
        /// Rut.CalculateDv(18123458); // "K"
        /// Rut.CalculateDv(12345678); // "5"
        /// </code>
        /// </example>
        public static string CalculateDv(int body)
        {
            if (body <= 0)
            {
                return string.Empty;
            }

            int sum = 0;
            int multiplier = 2;

            for (int n = body; n > 0; n /= 10)
            {
                int digit = n % 10;
                sum += digit * multiplier;

                multiplier++;
                if (multiplier > 7)
                {
                    multiplier = 2;
                }
            }

            int dv = 11 - (sum % 11);
            if (dv == 11) return "0";
            if (dv == 10) return "K";
            return dv.ToString(CultureInfo.InvariantCulture);
        }

        /// <summary>
        /// Calcula el Dígito Verificador oficial para un cuerpo numérico en formato texto
        /// aplicando el algoritmo de Módulo 11. Retorna cadena vacía si el texto no es estrictamente numérico.
        /// </summary>
        public static string CalculateDv(string? body)
        {
            if (string.IsNullOrEmpty(body) || !IsStrictDigits(body!))
            {
                return string.Empty;
            }

            if (!int.TryParse(body, NumberStyles.None, CultureInfo.InvariantCulture, out int num) || num <= 0)
            {
                return string.Empty;
            }

            return CalculateDv(num);
        }

        /// <summary>
        /// Valida exhaustivamente un RUT y retorna <c>null</c> si es válido,
        /// o un <see cref="RutValidationError"/> con el código numérico y motivo exacto del fallo.
        /// </summary>
        /// <example>
        /// <code>
        /// var err = Rut.ValidateDetailed("18.12A.458-K");
        /// if (err != null &amp;&amp; err.Code == RutErrorCode.InvalidCharacters) { ... }
        /// </code>
        /// </example>
        public static RutValidationError? ValidateDetailed(string? inputRut)
        {
            if (string.IsNullOrWhiteSpace(inputRut))
            {
                return new RutValidationError(RutErrorCode.Empty, "el RUT no puede estar vacío");
            }

            string raw = inputRut!;
            if (raw.IndexOf("--", StringComparison.Ordinal) >= 0 || raw.IndexOf("..", StringComparison.Ordinal) >= 0)
            {
                return new RutValidationError(RutErrorCode.MalformedFormat, "el RUT contiene puntos o guiones repetidos");
            }

            int hyphenCount = 0;
            int dotCount = 0;
            for (int i = 0; i < raw.Length; i++)
            {
                if (raw[i] == '-') hyphenCount++;
                if (raw[i] == '.') dotCount++;
            }

            if (hyphenCount > 1 || dotCount > 2)
            {
                return new RutValidationError(RutErrorCode.MalformedFormat, "el RUT excede la cantidad máxima de puntos o guiones");
            }

            string trimmed = raw.Trim();
            int hyphenIdx = trimmed.IndexOf('-');
            string bodyWithDots = string.Empty;

            if (hyphenIdx >= 0)
            {
                string before = trimmed.Substring(0, hyphenIdx);
                string after = trimmed.Substring(hyphenIdx + 1);
                if (after.IndexOf('.') >= 0)
                {
                    return new RutValidationError(RutErrorCode.MalformedFormat, "no se permiten puntos después del guión del RUT");
                }

                string afterHyphen = Clean(after);
                if (afterHyphen.Length != 1)
                {
                    return new RutValidationError(RutErrorCode.MalformedFormat, "el guión del RUT debe separar únicamente el dígito verificador");
                }

                bodyWithDots = StripWhitespace(before);
            }
            else
            {
                string noSpace = StripWhitespace(trimmed);
                if (noSpace.IndexOf('.') >= 0)
                {
                    if (noSpace.Length < 2 || noSpace[noSpace.Length - 1] == '.')
                    {
                        return new RutValidationError(RutErrorCode.MalformedFormat, "ubicación inválida de puntos separadores en el RUT");
                    }

                    bodyWithDots = noSpace.Substring(0, noSpace.Length - 1);
                }
            }

            if (bodyWithDots.IndexOf('.') >= 0)
            {
                string[] groups = bodyWithDots.Split('.');
                if (groups[0].Length < 1 || groups[0].Length > 3)
                {
                    return new RutValidationError(RutErrorCode.MalformedFormat, "ubicación inválida de puntos separadores en el RUT");
                }

                for (int i = 1; i < groups.Length; i++)
                {
                    if (groups[i].Length != 3)
                    {
                        return new RutValidationError(RutErrorCode.MalformedFormat, "ubicación inválida de puntos separadores en el RUT");
                    }
                }
            }

            string cleaned = Clean(raw);
            if (cleaned.Length < 2 || cleaned.Length > 10)
            {
                return new RutValidationError(RutErrorCode.InvalidLength, "la longitud del RUT es inválida");
            }

            int lastIndex = cleaned.Length - 1;
            string bodyStr = cleaned.Substring(0, lastIndex);
            string dv = cleaned.Substring(lastIndex, 1);

            if (!IsStrictDigits(bodyStr))
            {
                return new RutValidationError(RutErrorCode.InvalidCharacters, "el cuerpo del RUT contiene caracteres no numéricos");
            }

            if (!int.TryParse(bodyStr, NumberStyles.None, CultureInfo.InvariantCulture, out int numBody))
            {
                return new RutValidationError(RutErrorCode.InvalidCharacters, "el cuerpo del RUT contiene caracteres no numéricos");
            }

            if (CalculateDv(numBody) != dv)
            {
                return new RutValidationError(RutErrorCode.InvalidDv, "el dígito verificador no coincide");
            }

            return null;
        }

        /// <summary>
        /// Verifica si una cadena representa un RUT formal y matemáticamente válido según la normativa chilena.
        /// </summary>
        /// <example>
        /// <code>
        /// Rut.Validate("12.345.678-5"); // true
        /// Rut.Validate("18.123.458-k"); // true
        /// Rut.Validate("12.345.678-0"); // false
        /// </code>
        /// </example>
        public static bool Validate(string? inputRut) => ValidateDetailed(inputRut) == null;

        /// <summary>
        /// Convierte una cadena de RUT en una representación visual estándar (ej: "18.123.458-K").
        /// </summary>
        /// <example>
        /// <code>
        /// Rut.Format("18123458k"); // "18.123.458-K"
        /// Rut.Format("18123458k", new RutFormatOptions { Dots = false }); // "18123458-K"
        /// </code>
        /// </example>
        public static string Format(string? inputRut, RutFormatOptions? options = null)
        {
            string cleaned = Clean(inputRut);
            if (cleaned.Length < 2)
            {
                return cleaned;
            }

            bool dots = options?.Dots ?? true;
            bool hyphen = options?.Hyphen ?? true;

            int lastIndex = cleaned.Length - 1;
            int bodyLen = lastIndex;

            var sb = new StringBuilder(cleaned.Length + 4);
            for (int i = 0; i < bodyLen; i++)
            {
                if (dots && i > 0 && (bodyLen - i) % 3 == 0)
                {
                    sb.Append('.');
                }

                sb.Append(cleaned[i]);
            }

            if (hyphen)
            {
                sb.Append('-');
            }

            sb.Append(cleaned[lastIndex]);
            return sb.ToString();
        }

        /// <summary>
        /// Divide un RUT en su cuerpo numérico y dígito verificador.
        /// Lanza <see cref="ArgumentException"/> si la cadena contiene menos de 2 caracteres limpios.
        /// </summary>
        /// <example>
        /// <code>
        /// var parts = Rut.Split("18.123.458-K");
        /// // parts.Body == "18123458", parts.Dv == "K", parts.BodyInt() == 18123458
        /// </code>
        /// </example>
        public static RutParts Split(string? inputRut)
        {
            string cleaned = Clean(inputRut);
            if (cleaned.Length < 2)
            {
                throw new ArgumentException("el RUT no contiene cuerpo y dígito verificador suficientes", nameof(inputRut));
            }

            int lastIndex = cleaned.Length - 1;
            return new RutParts(cleaned.Substring(0, lastIndex), cleaned.Substring(lastIndex, 1));
        }

        /// <summary>
        /// Ofusca los caracteres iniciales de un RUT para proteger la privacidad del titular
        /// según la Ley 19.628, manteniendo visibles los últimos N caracteres.
        /// </summary>
        /// <example>
        /// <code>
        /// Rut.Mask("18123458K"); // "**.***.458-K"
        /// Rut.Mask("18.123.458-k", new RutMaskOptions { Formatted = false }); // "*****458K"
        /// </code>
        /// </example>
        public static string Mask(string? inputRut, RutMaskOptions? options = null)
        {
            string cleaned = Clean(inputRut);
            if (cleaned.Length < 2)
            {
                return cleaned;
            }

            char maskChar = options == null || options.MaskChar == '\0' ? '*' : options.MaskChar;
            int visibleCount = Math.Max(options?.VisibleCount ?? 4, 0);
            bool formatted = options?.Formatted ?? true;

            if (!formatted)
            {
                if (visibleCount >= cleaned.Length)
                {
                    return cleaned;
                }

                int maskLen = Math.Max(cleaned.Length - visibleCount, 0);
                return new string(maskChar, maskLen) + cleaned.Substring(maskLen);
            }

            string formattedStr = Format(cleaned);
            if (visibleCount >= cleaned.Length)
            {
                return formattedStr;
            }

            int threshold = cleaned.Length - visibleCount;
            var sb = new StringBuilder(formattedStr.Length);
            int seenAlphanumeric = 0;

            for (int i = 0; i < formattedStr.Length; i++)
            {
                char ch = formattedStr[i];
                if (ch == '.' || ch == '-')
                {
                    sb.Append(ch);
                    continue;
                }

                sb.Append(seenAlphanumeric < threshold ? maskChar : ch);
                seenAlphanumeric++;
            }

            return sb.ToString();
        }

        /// <summary>
        /// Genera un RUT aleatorio y matemáticamente válido en el rango general (1M a 28M).
        /// </summary>
        public static string Generate(RutFormatOptions? options = null) =>
            GenerateBetween(RutRanges.DefaultGeneralMin, RutRanges.DefaultGeneralMax, options);

        /// <summary>
        /// Genera un RUT matemáticamente válido dentro de un rango numérico específico.
        /// Corrige automáticamente límites negativos o invertidos.
        /// </summary>
        public static string GenerateBetween(int min, int max, RutFormatOptions? options = null)
        {
            int lower = min < RutRanges.DefaultGeneralMin ? RutRanges.DefaultGeneralMin : min;
            int upper = max < RutRanges.DefaultGeneralMin ? RutRanges.DefaultGeneralMax : max;

            if (lower > upper)
            {
                int tmp = lower;
                lower = upper;
                upper = tmp;
            }

            int body;
            if (upper < int.MaxValue)
            {
                body = ThreadRandom.Value!.Next(lower, upper + 1);
            }
            else
            {
                long offset = (long)(ThreadRandom.Value!.NextDouble() * ((double)upper - lower + 1.0));
                long candidate = (long)lower + offset;
                body = candidate > int.MaxValue ? int.MaxValue : (int)candidate;
            }
            string dv = CalculateDv(body);
            return Format(body.ToString(CultureInfo.InvariantCulture) + dv, options);
        }

        /// <summary>
        /// Genera un RUT aleatorio representativo de personas naturales vivas en Chile (rango 5M a 28M).
        /// </summary>
        public static string GeneratePerson(RutFormatOptions? options = null) =>
            GenerateBetween(RutRanges.DefaultPersonMin, RutRanges.DefaultPersonMax, options);

        /// <summary>
        /// Genera un RUT aleatorio representativo de personas jurídicas y empresas en Chile (rango 70M a 80M).
        /// </summary>
        public static string GenerateBusiness(RutFormatOptions? options = null) =>
            GenerateBetween(RutRanges.DefaultBusinessMin, RutRanges.DefaultBusinessMax, options);

        private static string StripWhitespace(string s)
        {
            var sb = new StringBuilder(s.Length);
            for (int i = 0; i < s.Length; i++)
            {
                char ch = s[i];
                if (ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r')
                {
                    continue;
                }

                sb.Append(ch);
            }

            return sb.ToString();
        }

        private static bool IsStrictDigits(string s)
        {
            if (s.Length == 0)
            {
                return false;
            }

            for (int i = 0; i < s.Length; i++)
            {
                char ch = s[i];
                if (ch < '0' || ch > '9')
                {
                    return false;
                }
            }

            return true;
        }
    }
}
