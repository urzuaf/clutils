namespace Clutils
{
    /// <summary>
    /// Estilo de representación para montos negativos en CLP.
    /// </summary>
    public enum NegativeFormat
    {
        /// <summary>Representa montos negativos con signo menos (ej: "-$1.500").</summary>
        MinusSign = 0,

        /// <summary>Representa montos negativos en formato contable entre paréntesis (ej: "($1.500)").</summary>
        Parentheses = 1,
    }

    /// <summary>
    /// Opciones de formateo para pesos chilenos (CLP).
    /// </summary>
    public sealed class ClpFormatOptions
    {
        /// <summary>Símbolo o prefijo de la moneda (por defecto "$").</summary>
        public string Symbol { get; set; } = "$";

        /// <summary>Si es true, omite completamente el símbolo de la moneda (por defecto false).</summary>
        public bool NoSymbol { get; set; }

        /// <summary>Si incluye un espacio entre el símbolo y el monto (ej: "$ 1.500", por defecto false).</summary>
        public bool Space { get; set; }

        /// <summary>Estilo para números negativos (por defecto <see cref="NegativeFormat.MinusSign"/>).</summary>
        public NegativeFormat NegativeFormat { get; set; } = NegativeFormat.MinusSign;
    }

    /// <summary>
    /// Opciones de formateo para Unidades de Fomento (UF).
    /// </summary>
    public sealed class UfFormatOptions
    {
        /// <summary>Cantidad de decimales a formatear (por defecto 2).</summary>
        public int Decimals { get; set; } = 2;

        /// <summary>Símbolo o prefijo de la unidad (por defecto "UF").</summary>
        public string Symbol { get; set; } = "UF";

        /// <summary>Si es true, omite completamente el símbolo de la unidad (por defecto false).</summary>
        public bool NoSymbol { get; set; }
    }

    /// <summary>
    /// Opciones de formateo para Unidades Tributarias Mensuales (UTM).
    /// </summary>
    public sealed class UtmFormatOptions
    {
        /// <summary>Símbolo o prefijo de la unidad (por defecto "UTM").</summary>
        public string Symbol { get; set; } = "UTM";

        /// <summary>Si es true, omite completamente el símbolo de la unidad (por defecto false).</summary>
        public bool NoSymbol { get; set; }
    }

    /// <summary>
    /// Opciones para personalizar la tasa de cálculo del IVA.
    /// </summary>
    public sealed class IvaOptions
    {
        /// <summary>Tasa decimal del impuesto (por defecto 0.19 para 19%).</summary>
        public double Rate { get; set; } = Money.DefaultIvaRate;
    }

    /// <summary>
    /// Desglose contable exacto de una transacción garantizando la cuadratura del SII (Neto + IVA == Bruto).
    /// </summary>
    public sealed class IvaBreakdown
    {
        /// <summary>Valor base antes de impuestos.</summary>
        public long Neto { get; }

        /// <summary>Monto correspondiente al impuesto.</summary>
        public long Iva { get; }

        /// <summary>Total final que incluye el impuesto.</summary>
        public long Bruto { get; }

        /// <summary>Inicializa un desglose contable de IVA.</summary>
        public IvaBreakdown(long neto, long iva, long bruto)
        {
            Neto = neto;
            Iva = iva;
            Bruto = bruto;
        }
    }

    /// <summary>
    /// Opciones para la conversión de montos a palabras en español chileno.
    /// </summary>
    public sealed class ToWordsOptions
    {
        /// <summary>Si es true, omite la denominación monetaria ("peso", "pesos", "de pesos").</summary>
        public bool NoCurrency { get; set; }
    }
}
