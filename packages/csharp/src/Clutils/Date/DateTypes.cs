namespace Clutils
{
    /// <summary>
    /// Estilos visuales disponibles para el formateo de fechas chilenas.
    /// </summary>
    public enum DateFormatStyle
    {
        /// <summary>Formato numérico estándar chileno con barras (ej: "18/09/2026").</summary>
        Numeric = 0,

        /// <summary>Formato numérico chileno con guiones (ej: "18-09-2026").</summary>
        NumericHyphen = 1,

        /// <summary>Formato formal en español (ej: "18 de septiembre de 2026").</summary>
        Formal = 2,

        /// <summary>Formato formal completo con día de la semana (ej: "Viernes, 18 de septiembre de 2026").</summary>
        Full = 3,
    }
}
