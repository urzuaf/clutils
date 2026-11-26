using System.Collections.Generic;

namespace Clutils
{
    public static partial class Money
    {
        private static readonly string[] UnitsWords =
        {
            "", "un", "dos", "tres", "cuatro", "cinco", "seis", "siete", "ocho", "nueve",
            "diez", "once", "doce", "trece", "catorce", "quince", "dieciséis", "diecisiete", "dieciocho", "diecinueve",
            "veinte", "veintiún", "veintidós", "veintitrés", "veinticuatro", "veinticinco", "veintiséis", "veintisiete", "veintiocho", "veintinueve",
        };

        private static readonly string[] TensWords =
        {
            "", "", "", "treinta", "cuarenta", "cincuenta", "sesenta", "setenta", "ochenta", "noventa",
        };

        private static readonly string[] HundredsWords =
        {
            "", "ciento", "doscientos", "trescientos", "cuatrocientos", "quinientos", "seiscientos", "setecientos", "ochocientos", "novecientos",
        };

        private static string ConvertGroup(int n)
        {
            if (n == 0) return string.Empty;
            if (n == 100) return "cien";

            var parts = new List<string>(3);
            int c = n / 100;
            int d = (n % 100) / 10;
            int u = n % 10;

            if (c > 0)
            {
                parts.Add(HundredsWords[c]);
            }

            int du = n % 100;
            if (du > 0)
            {
                if (du < 30)
                {
                    parts.Add(UnitsWords[du]);
                }
                else if (u == 0)
                {
                    parts.Add(TensWords[d]);
                }
                else
                {
                    parts.Add(TensWords[d] + " y " + UnitsWords[u]);
                }
            }

            return string.Join(" ", parts);
        }

        private static string NumberToWords(ulong val)
        {
            if (val == 0) return "cero";
            if (val == 1) return "un";

            ulong trillones = val / 1_000_000_000_000_000_000UL;
            ulong remTrillones = val % 1_000_000_000_000_000_000UL;

            ulong billonesTotal = remTrillones / 1_000_000_000_000UL;
            ulong remBillones = remTrillones % 1_000_000_000_000UL;

            ulong millonesTotal = remBillones / 1_000_000UL;
            ulong remMillones = remBillones % 1_000_000UL;

            ulong miles = remMillones / 1_000UL;
            ulong unidades = remMillones % 1_000UL;

            var parts = new List<string>(8);

            // 1. Trillones (10^18)
            if (trillones > 0)
            {
                if (trillones == 1)
                {
                    parts.Add("un trillón");
                }
                else
                {
                    parts.Add(ConvertGroup((int)trillones) + " trillones");
                }
            }

            // 2. Billones (10^12)
            if (billonesTotal > 0)
            {
                if (billonesTotal == 1)
                {
                    parts.Add("un billón");
                }
                else
                {
                    ulong milesDeBillones = billonesTotal / 1_000UL;
                    ulong restoBillones = billonesTotal % 1_000UL;

                    var billParts = new List<string>(3);
                    if (milesDeBillones > 0)
                    {
                        if (milesDeBillones == 1)
                        {
                            billParts.Add("mil");
                        }
                        else
                        {
                            billParts.Add(ConvertGroup((int)milesDeBillones));
                            billParts.Add("mil");
                        }
                    }

                    if (restoBillones > 0)
                    {
                        billParts.Add(ConvertGroup((int)restoBillones));
                    }

                    parts.Add(string.Join(" ", billParts) + " billones");
                }
            }

            // 3. Millones (10^6)
            if (millonesTotal > 0)
            {
                if (millonesTotal == 1)
                {
                    parts.Add("un millón");
                }
                else
                {
                    ulong milesDeMillones = millonesTotal / 1_000UL;
                    ulong restoMillones = millonesTotal % 1_000UL;

                    var millParts = new List<string>(3);
                    if (milesDeMillones > 0)
                    {
                        if (milesDeMillones == 1)
                        {
                            millParts.Add("mil");
                        }
                        else
                        {
                            millParts.Add(ConvertGroup((int)milesDeMillones));
                            millParts.Add("mil");
                        }
                    }

                    if (restoMillones > 0)
                    {
                        millParts.Add(ConvertGroup((int)restoMillones));
                    }

                    parts.Add(string.Join(" ", millParts) + " millones");
                }
            }

            // 4. Miles (10^3)
            if (miles > 0)
            {
                if (miles == 1)
                {
                    parts.Add("mil");
                }
                else
                {
                    parts.Add(ConvertGroup((int)miles) + " mil");
                }
            }

            // 5. Unidades (10^0)
            if (unidades > 0)
            {
                parts.Add(ConvertGroup((int)unidades));
            }

            return string.Join(" ", parts);
        }

        /// <summary>
        /// Convierte un monto entero en pesos chilenos a su representación textual formal en español,
        /// respetando las normas gramaticales, ortográficas y bancarias chilenas (cheques, pagarés y contratos).
        /// </summary>
        /// <example>
        /// <code>
        /// Money.ToWords(1500000); // "un millón quinientos mil pesos"
        /// Money.ToWords(2000000); // "dos millones de pesos"
        /// Money.ToWords(100, new ToWordsOptions { NoCurrency = true }); // "cien"
        /// </code>
        /// </example>
        public static string ToWords(long amount, ToWordsOptions? options = null)
        {
            bool noCurrency = options?.NoCurrency ?? false;
            bool isNegative = amount < 0;
            ulong absVal = AbsUnsigned(amount);

            string text = NumberToWords(absVal);

            if (noCurrency)
            {
                return isNegative ? "menos " + text : text;
            }

            string currencySuffix;
            if (absVal == 1)
            {
                currencySuffix = "peso";
            }
            else if (absVal >= 1_000_000UL && absVal % 1_000_000UL == 0)
            {
                currencySuffix = "de pesos";
            }
            else
            {
                currencySuffix = "pesos";
            }

            return isNegative
                ? "menos " + text + " " + currencySuffix
                : text + " " + currencySuffix;
        }
    }
}
