using System;
using System.Globalization;
using System.Text;

namespace Clutils
{
    public static partial class Money
    {
        /// <summary>
        /// Formatea una cadena de dígitos enteros positivos con separadores de miles '.' de forma lineal.
        /// </summary>
        internal static string FormatThousandsStr(string s)
        {
            int n = s.Length;
            if (n <= 3)
            {
                return s;
            }

            int dots = (n - 1) / 3;
            char[] res = new char[n + dots];
            int j = res.Length - 1;
            int count = 0;

            for (int i = n - 1; i >= 0; i--)
            {
                res[j--] = s[i];
                count++;
                if (count == 3 && i > 0)
                {
                    res[j--] = '.';
                    count = 0;
                }
            }

            return new string(res);
        }

        internal static string FormatThousands(ulong val) =>
            FormatThousandsStr(val.ToString(CultureInfo.InvariantCulture));

        internal static ulong AbsUnsigned(long amount) =>
            amount < 0 ? (ulong)(-(amount + 1)) + 1UL : (ulong)amount;

        /// <summary>
        /// Formatea un monto entero a pesos chilenos (CLP) según la convención local,
        /// utilizando puntos como separador de miles y sin centavos.
        /// </summary>
        /// <example>
        /// <code>
        /// Money.FormatCLP(1500000); // "$1.500.000"
        /// Money.FormatCLP(-2500, new ClpFormatOptions { NegativeFormat = NegativeFormat.Parentheses }); // "($2.500)"
        /// </code>
        /// </example>
        public static string FormatCLP(long amount, ClpFormatOptions? options = null)
        {
            string symbol = options != null && options.NoSymbol
                ? string.Empty
                : (string.IsNullOrEmpty(options?.Symbol) ? "$" : options!.Symbol);

            bool space = options?.Space ?? false;
            NegativeFormat negFormat = options?.NegativeFormat ?? NegativeFormat.MinusSign;

            bool isNegative = amount < 0;
            string numStr = FormatThousands(AbsUnsigned(amount));
            string spaceStr = space && symbol.Length > 0 ? " " : string.Empty;

            if (!isNegative)
            {
                return symbol + spaceStr + numStr;
            }

            if (negFormat == NegativeFormat.Parentheses)
            {
                return "(" + symbol + spaceStr + numStr + ")";
            }

            return "-" + symbol + spaceStr + numStr;
        }

        /// <summary>
        /// Extrae y convierte cualquier representación en texto de pesos chilenos a un entero limpio (<see cref="long"/>),
        /// reconociendo signos negativos estándar y contables, y rechazando entradas malformadas.
        /// </summary>
        /// <exception cref="FormatException">Si el texto está vacío, no contiene dígitos válidos o presenta formato inválido.</exception>
        /// <example>
        /// <code>
        /// Money.ParseCLP("$ 1.500.000"); // 1500000
        /// Money.ParseCLP("-CLP 2.500");  // -2500
        /// Money.ParseCLP("($2.500)");    // -2500
        /// </code>
        /// </example>
        public static long ParseCLP(string? amount)
        {
            if (string.IsNullOrWhiteSpace(amount))
            {
                throw new FormatException("el texto está vacío");
            }

            string trimmed = amount!.Trim();

            bool isNegative = false;
            bool hasParen = false;
            bool digitsStarted = false;
            bool digitsEnded = false;
            int minusCount = 0;
            int dollarCount = 0;
            int clpCount = 0;
            int openParens = 0;
            int closeParens = 0;
            var rawNum = new StringBuilder(trimmed.Length);

            for (int i = 0; i < trimmed.Length; i++)
            {
                char ch = trimmed[i];

                // Reconocer token "CLP" (prefijo, sufijo o tras signo/paréntesis)
                if ((ch == 'C' || ch == 'c') &&
                    i + 3 <= trimmed.Length &&
                    string.Compare(trimmed, i, "CLP", 0, 3, StringComparison.OrdinalIgnoreCase) == 0)
                {
                    clpCount++;
                    if (clpCount > 1)
                    {
                        throw new FormatException("múltiples prefijos CLP inválidos");
                    }

                    if (digitsStarted)
                    {
                        digitsEnded = true;
                    }

                    i += 2;
                    continue;
                }

                if (ch >= '0' && ch <= '9')
                {
                    if (closeParens > 0)
                    {
                        throw new FormatException("dígito posterior al paréntesis de cierre");
                    }

                    if (digitsEnded)
                    {
                        throw new FormatException("formato numérico discontinuo inválido");
                    }

                    digitsStarted = true;
                    rawNum.Append(ch);
                }
                else if (ch == '.')
                {
                    if (closeParens > 0 || digitsEnded)
                    {
                        throw new FormatException("separador de miles en posición inválida");
                    }

                    digitsStarted = true;
                    rawNum.Append(ch);
                }
                else if (ch == '$')
                {
                    if (digitsStarted)
                    {
                        throw new FormatException("símbolo '$' no puede ir después de dígitos");
                    }

                    dollarCount++;
                    if (dollarCount > 1)
                    {
                        throw new FormatException("múltiples símbolos '$' inválidos");
                    }
                }
                else if (ch == ' ' || ch == '\t')
                {
                    if (digitsStarted)
                    {
                        digitsEnded = true;
                    }
                }
                else if (ch == '-')
                {
                    if (digitsStarted)
                    {
                        throw new FormatException("signo menos no puede ir después de dígitos");
                    }

                    minusCount++;
                    isNegative = true;
                }
                else if (ch == '(')
                {
                    if (digitsStarted)
                    {
                        throw new FormatException("paréntesis de apertura no puede ir después de dígitos");
                    }

                    openParens++;
                }
                else if (ch == ')')
                {
                    if (openParens == 0)
                    {
                        throw new FormatException("paréntesis de cierre sin apertura previa");
                    }

                    closeParens++;
                    if (digitsStarted)
                    {
                        digitsEnded = true;
                    }
                }
                else
                {
                    throw new FormatException($"carácter no válido en monto CLP: '{ch}'");
                }
            }

            if (openParens > 0 || closeParens > 0)
            {
                if (openParens != 1 || closeParens != 1)
                {
                    throw new FormatException("paréntesis contables desbalanceados");
                }

                if (minusCount > 0)
                {
                    throw new FormatException("formato negativo ambiguo (combina '-' y paréntesis)");
                }

                hasParen = true;
            }

            if (minusCount > 1)
            {
                throw new FormatException("múltiples signos menos inválidos");
            }

            string numStr = rawNum.ToString();
            if (numStr.Length == 0)
            {
                throw new FormatException("el texto no contiene dígitos numéricos válidos");
            }

            string digits;
            if (numStr.IndexOf('.') >= 0)
            {
                string[] groups = numStr.Split('.');
                if (groups[0].Length < 1 || groups[0].Length > 3)
                {
                    throw new FormatException($"separadores de miles mal ubicados: \"{numStr}\"");
                }

                for (int i = 1; i < groups.Length; i++)
                {
                    if (groups[i].Length != 3)
                    {
                        throw new FormatException($"separadores de miles mal ubicados: \"{numStr}\"");
                    }
                }

                digits = string.Concat(groups);
            }
            else
            {
                digits = numStr;
            }

            if (!ulong.TryParse(digits, NumberStyles.None, CultureInfo.InvariantCulture, out ulong uval))
            {
                throw new FormatException("desborde o error numérico al parsear monto CLP");
            }

            if (isNegative || hasParen)
            {
                const ulong minInt64Abs = (ulong)long.MaxValue + 1UL;
                if (uval > minInt64Abs)
                {
                    throw new FormatException("desborde o error numérico al parsear monto CLP");
                }

                return uval == minInt64Abs ? long.MinValue : -(long)uval;
            }

            if (uval > (ulong)long.MaxValue)
            {
                throw new FormatException("desborde o error numérico al parsear monto CLP");
            }

            return (long)uval;
        }
    }
}
