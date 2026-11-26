using System;
using System.Globalization;
using System.Text;

namespace Clutils
{
    public static partial class Money
    {
        private static bool IsAllZeros(string s)
        {
            for (int i = 0; i < s.Length; i++)
            {
                if (s[i] != '0')
                {
                    return false;
                }
            }

            return true;
        }

        /// <summary>
        /// Formatea un valor en Unidades de Fomento (UF) según la convención oficial chilena (CMF),
        /// utilizando puntos para separador de miles, coma para decimales y redondeo aritmético
        /// estándar (<see cref="MidpointRounding.AwayFromZero"/>) consistente con Go y TypeScript.
        /// </summary>
        /// <example>
        /// <code>
        /// Money.FormatUF(38450.25); // "UF 38.450,25"
        /// Money.FormatUF(38450.2548, new UfFormatOptions { Decimals = 4 }); // "UF 38.450,2548"
        /// </code>
        /// </example>
        public static string FormatUF(double amount, UfFormatOptions? options = null)
        {
            int decimals = options != null && options.Decimals > 0 ? options.Decimals : 2;
            string symbol = options != null && options.NoSymbol
                ? string.Empty
                : (string.IsNullOrEmpty(options?.Symbol) ? "UF" : options!.Symbol);

            if (double.IsNaN(amount) || double.IsInfinity(amount))
            {
                amount = 0.0;
            }

            bool isNegative = amount < 0.0;
            double absAmount = Math.Abs(amount);

            string integerPartStr;
            string decimalPartStr;

            // Redondeo aritmético explícito AwayFromZero (C# Math.Round por defecto usa ToEven)
            double factor = Math.Pow(10.0, decimals);
            double scaled = absAmount * factor;

            if (decimals <= 12 && !double.IsInfinity(scaled) && scaled <= 9e15)
            {
                ulong roundedInt = (ulong)Math.Round(scaled, MidpointRounding.AwayFromZero);
                string rawDigits = roundedInt.ToString(CultureInfo.InvariantCulture);
                if (rawDigits.Length <= decimals)
                {
                    integerPartStr = "0";
                    decimalPartStr = new string('0', decimals - rawDigits.Length) + rawDigits;
                }
                else
                {
                    int splitPos = rawDigits.Length - decimals;
                    integerPartStr = rawDigits.Substring(0, splitPos);
                    decimalPartStr = rawDigits.Substring(splitPos);
                }
            }
            else
            {
                int fmtDecimals = decimals > 99 ? 99 : decimals;
                string formattedFloat = absAmount.ToString("F" + fmtDecimals.ToString(CultureInfo.InvariantCulture), CultureInfo.InvariantCulture);
                int dotIdx = formattedFloat.IndexOf('.');
                if (dotIdx >= 0)
                {
                    integerPartStr = formattedFloat.Substring(0, dotIdx);
                    decimalPartStr = formattedFloat.Substring(dotIdx + 1);
                }
                else
                {
                    integerPartStr = formattedFloat;
                    decimalPartStr = string.Empty;
                }

                if (decimals > fmtDecimals)
                {
                    decimalPartStr += new string('0', decimals - fmtDecimals);
                }
            }

            // Evitar "-0,00" cuando un valor negativo pequeño redondea a cero
            if (isNegative && IsAllZeros(integerPartStr) && IsAllZeros(decimalPartStr))
            {
                isNegative = false;
            }

            string formattedInt = FormatThousandsStr(integerPartStr);

            var sb = new StringBuilder(symbol.Length + formattedInt.Length + decimalPartStr.Length + 4);
            if (symbol.Length > 0)
            {
                sb.Append(symbol);
                sb.Append(' ');
            }

            if (isNegative)
            {
                sb.Append('-');
            }

            sb.Append(formattedInt);
            if (decimals > 0 && decimalPartStr.Length > 0)
            {
                sb.Append(',');
                sb.Append(decimalPartStr);
            }

            return sb.ToString();
        }
    }
}
