using System.Text;

namespace Clutils
{
    public static partial class Money
    {
        /// <summary>
        /// Formatea un monto entero en Unidades Tributarias Mensuales (UTM),
        /// utilizando puntos como separador de miles y sin decimales según la normativa tributaria chilena.
        /// </summary>
        /// <example>
        /// <code>
        /// Money.FormatUTM(67294); // "UTM 67.294"
        /// Money.FormatUTM(67294, new UtmFormatOptions { NoSymbol = true }); // "67.294"
        /// </code>
        /// </example>
        public static string FormatUTM(long amount, UtmFormatOptions? options = null)
        {
            string symbol = options != null && options.NoSymbol
                ? string.Empty
                : (string.IsNullOrEmpty(options?.Symbol) ? "UTM" : options!.Symbol);

            bool isNegative = amount < 0;
            string numStr = FormatThousands(AbsUnsigned(amount));

            var sb = new StringBuilder(symbol.Length + numStr.Length + 3);
            if (symbol.Length > 0)
            {
                sb.Append(symbol);
                sb.Append(' ');
            }

            if (isNegative)
            {
                sb.Append('-');
            }

            sb.Append(numStr);
            return sb.ToString();
        }
    }
}
