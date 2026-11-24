using System;

namespace Clutils
{
    /// <summary>
    /// Códigos numéricos de error para el diagnóstico de validación de un RUT chileno.
    /// </summary>
    public enum RutErrorCode
    {
        /// <summary>Sin error (RUT válido).</summary>
        None = 0,
        /// <summary>El RUT está vacío o contiene solo espacios en blanco.</summary>
        Empty = 1,
        /// <summary>Puntos o guiones mal ubicados, repetidos o en exceso.</summary>
        MalformedFormat = 2,
        /// <summary>Longitud fuera del rango permitido (debe tener entre 1 y 9 dígitos de cuerpo + DV).</summary>
        InvalidLength = 3,
        /// <summary>El cuerpo contiene letras o caracteres no numéricos.</summary>
        InvalidCharacters = 4,
        /// <summary>El dígito verificador no coincide según el algoritmo de Módulo 11.</summary>
        InvalidDv = 5,
    }

    /// <summary>
    /// Representa el resultado detallado de un error de validación de RUT.
    /// </summary>
    public sealed class RutValidationError
    {
        /// <summary>Código numérico constante para serialización o i18n.</summary>
        public RutErrorCode Code { get; }

        /// <summary>Mensaje descriptivo en español del motivo del fallo.</summary>
        public string Message { get; }

        /// <summary>Inicializa un nuevo error de validación de RUT.</summary>
        public RutValidationError(RutErrorCode code, string message)
        {
            Code = code;
            Message = message;
        }
    }

    /// <summary>
    /// Opciones para personalizar la representación visual de un RUT.
    /// </summary>
    public sealed class RutFormatOptions
    {
        /// <summary>Si incluye puntos separadores de miles (por defecto true).</summary>
        public bool Dots { get; set; } = true;

        /// <summary>Si incluye guión antes del dígito verificador (por defecto true).</summary>
        public bool Hyphen { get; set; } = true;
    }

    /// <summary>
    /// Opciones para ofuscar un RUT por privacidad.
    /// </summary>
    public sealed class RutMaskOptions
    {
        /// <summary>Carácter de reemplazo (por defecto '*').</summary>
        public char MaskChar { get; set; } = '*';

        /// <summary>Cantidad de caracteres visibles al final (por defecto 4: últimos 3 del cuerpo + DV).</summary>
        public int VisibleCount { get; set; } = 4;

        /// <summary>Si conserva el formato formal con puntos y guión (por defecto true).</summary>
        public bool Formatted { get; set; } = true;
    }

    /// <summary>
    /// Representa la descomposición atómica de un RUT en cuerpo y dígito verificador.
    /// </summary>
    public sealed class RutParts
    {
        /// <summary>Cuerpo numérico en formato string para preservar ceros iniciales.</summary>
        public string Body { get; }

        /// <summary>Dígito verificador normalizado en mayúscula ("0"-"9" o "K").</summary>
        public string Dv { get; }

        /// <summary>Inicializa las partes de un RUT.</summary>
        public RutParts(string body, string dv)
        {
            Body = body;
            Dv = dv;
        }

        /// <summary>
        /// Retorna el cuerpo convertido a número entero.
        /// Lanza <see cref="FormatException"/> si contiene caracteres no numéricos.
        /// </summary>
        public int BodyInt()
        {
            if (string.IsNullOrEmpty(Body))
            {
                throw new FormatException("el cuerpo del RUT contiene caracteres no numéricos");
            }

            for (int i = 0; i < Body.Length; i++)
            {
                char ch = Body[i];
                if (ch < '0' || ch > '9')
                {
                    throw new FormatException("el cuerpo del RUT contiene caracteres no numéricos");
                }
            }

            return int.Parse(Body, System.Globalization.CultureInfo.InvariantCulture);
        }
    }

    /// <summary>
    /// Rangos numéricos canónicos sincronizados con la fuente única de verdad <c>data/rut/ranges.json</c>.
    /// </summary>
    public static class RutRanges
    {
        /// <summary>Límite inferior habitual para personas naturales vivas en Chile (5.000.000).</summary>
        public const int DefaultPersonMin = 5_000_000;

        /// <summary>Límite superior habitual para personas naturales en Chile (28.000.000).</summary>
        public const int DefaultPersonMax = 28_000_000;

        /// <summary>Límite inferior habitual para personas jurídicas y empresas en Chile (70.000.000).</summary>
        public const int DefaultBusinessMin = 70_000_000;

        /// <summary>Límite superior habitual para personas jurídicas y empresas en Chile (80.000.000).</summary>
        public const int DefaultBusinessMax = 80_000_000;

        /// <summary>Límite inferior seguro para generación general (1.000.000).</summary>
        public const int DefaultGeneralMin = 1_000_000;

        /// <summary>Límite superior seguro para generación general (28.000.000).</summary>
        public const int DefaultGeneralMax = 28_000_000;
    }
}
