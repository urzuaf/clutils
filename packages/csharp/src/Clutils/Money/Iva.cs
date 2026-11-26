using System;

namespace Clutils
{
    public static partial class Money
    {
        /// <summary>
        /// Desglosa un monto en Neto, IVA y Bruto aplicando la regla de cuadratura del SII
        /// (<c>Neto + IVA == Bruto</c>) con redondeo aritmético simétrico (<see cref="MidpointRounding.AwayFromZero"/>).
        /// </summary>
        /// <example>
        /// <code>
        /// Money.CalculateIVA(1000, false); // { Neto: 1000, Iva: 190, Bruto: 1190 }
        /// Money.CalculateIVA(1000, true);  // { Neto: 840,  Iva: 160, Bruto: 1000 }
        /// </code>
        /// </example>
        public static IvaBreakdown CalculateIVA(long amount, bool isBruto, IvaOptions? options = null)
        {
            double rate = options != null && !double.IsNaN(options.Rate) && !double.IsInfinity(options.Rate) && options.Rate > 0.0
                ? options.Rate
                : DefaultIvaRate;

            if (amount == 0)
            {
                return new IvaBreakdown(0, 0, 0);
            }

            if (isBruto)
            {
                long bruto = amount;
                long neto = (long)Math.Round(bruto / (1.0 + rate), MidpointRounding.AwayFromZero);
                long iva = bruto - neto;
                return new IvaBreakdown(neto, iva, bruto);
            }
            else
            {
                long neto = amount;
                long iva = (long)Math.Round(neto * rate, MidpointRounding.AwayFromZero);
                long bruto = neto + iva;
                return new IvaBreakdown(neto, iva, bruto);
            }
        }
    }
}
