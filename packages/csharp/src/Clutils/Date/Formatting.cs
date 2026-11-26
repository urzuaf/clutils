using System;
using System.Globalization;

namespace Clutils
{
    public static partial class Date
    {
        private static readonly string[] SpanishMonths =
        {
            "",
            "enero",
            "febrero",
            "marzo",
            "abril",
            "mayo",
            "junio",
            "julio",
            "agosto",
            "septiembre",
            "octubre",
            "noviembre",
            "diciembre",
        };

        private static readonly string[] SpanishWeekdays =
        {
            "Domingo",
            "Lunes",
            "Martes",
            "Miércoles",
            "Jueves",
            "Viernes",
            "Sábado",
        };

        private static bool IsLeapYear(int year) =>
            year % 4 == 0 && (year % 100 != 0 || year % 400 == 0);

        private static int DaysInMonth(int year, int month)
        {
            switch (month)
            {
                case 1:
                case 3:
                case 5:
                case 7:
                case 8:
                case 10:
                case 12:
                    return 31;
                case 4:
                case 6:
                case 9:
                case 11:
                    return 30;
                case 2:
                    return IsLeapYear(year) ? 29 : 28;
                default:
                    return 0;
            }
        }

        private static bool TryParseStrictDigits(string s, out int value)
        {
            value = 0;
            if (s.Length == 0)
            {
                return false;
            }

            int acc = 0;
            for (int i = 0; i < s.Length; i++)
            {
                char ch = s[i];
                if (ch < '0' || ch > '9')
                {
                    return false;
                }

                acc = (acc * 10) + (ch - '0');
            }

            value = acc;
            return true;
        }

        private static string FormatComponents(int year, int month, int day, DayOfWeek dayOfWeek, DateFormatStyle style)
        {
            string dayPad = day.ToString("D2", CultureInfo.InvariantCulture);
            string monthPad = month.ToString("D2", CultureInfo.InvariantCulture);
            string yearPad = year.ToString("D4", CultureInfo.InvariantCulture);

            switch (style)
            {
                case DateFormatStyle.NumericHyphen:
                    return dayPad + "-" + monthPad + "-" + yearPad;
                case DateFormatStyle.Formal:
                    return day.ToString(CultureInfo.InvariantCulture) + " de " + SpanishMonths[month] + " de " + yearPad;
                case DateFormatStyle.Full:
                    return SpanishWeekdays[(int)dayOfWeek] + ", " + day.ToString(CultureInfo.InvariantCulture) + " de " + SpanishMonths[month] + " de " + yearPad;
                case DateFormatStyle.Numeric:
                default:
                    return dayPad + "/" + monthPad + "/" + yearPad;
            }
        }

        /// <summary>
        /// Formatea una fecha según los estándares y convenciones locales chilenas.
        /// </summary>
        /// <example>
        /// <code>
        /// var d = Date.ParseDate("18/09/2026");
        /// Date.FormatDate(d);                       // "18/09/2026"
        /// Date.FormatDate(d, DateFormatStyle.Full); // "Viernes, 18 de septiembre de 2026"
        /// </code>
        /// </example>
        public static string FormatDate(DateTime date, DateFormatStyle style = DateFormatStyle.Numeric) =>
            FormatComponents(date.Year, date.Month, date.Day, date.DayOfWeek, style);

        /// <summary>
        /// Formatea una fecha con offset según los estándares y convenciones locales chilenas.
        /// </summary>
        public static string FormatDate(DateTimeOffset date, DateFormatStyle style = DateFormatStyle.Numeric) =>
            FormatComponents(date.Year, date.Month, date.Day, date.DayOfWeek, style);

        /// <summary>
        /// Interpreta una cadena en formato chileno estándar ("DD/MM/YYYY" o "DD-MM-YYYY"),
        /// garantizando que el primer segmento sea el día y el segundo el mes.
        /// Valida estrictamente dígitos, días por mes, años bisiestos y rechaza formatos ISO o malformados.
        /// </summary>
        /// <exception cref="FormatException">Si el formato es inválido o los componentes están fuera de rango.</exception>
        /// <example>
        /// <code>
        /// var dt = Date.ParseDate("18/09/2026"); // 2026-09-18 UTC
        /// </code>
        /// </example>
        public static DateTime ParseDate(string? text)
        {
            if (string.IsNullOrWhiteSpace(text))
            {
                throw new FormatException("el texto de fecha está vacío");
            }

            string trimmed = text!.Trim();

            char sep;
            if (trimmed.IndexOf('/') >= 0)
            {
                sep = '/';
            }
            else if (trimmed.IndexOf('-') >= 0)
            {
                sep = '-';
            }
            else
            {
                throw new FormatException("formato no reconocido: se requiere separador '/' o '-'");
            }

            string[] parts = trimmed.Split(sep);
            if (parts.Length != 3)
            {
                throw new FormatException($"formato de fecha inválido, se esperaba DD{sep}MM{sep}YYYY");
            }

            string dayStr = parts[0];
            string monthStr = parts[1];
            string yearStr = parts[2];

            if (dayStr.Length < 1 || dayStr.Length > 2 || monthStr.Length < 1 || monthStr.Length > 2)
            {
                throw new FormatException("componentes de día o mes inválidos");
            }

            if (yearStr.Length != 4)
            {
                throw new FormatException($"el año debe tener 4 dígitos: \"{yearStr}\"");
            }

            if (!TryParseStrictDigits(dayStr, out int day) ||
                !TryParseStrictDigits(monthStr, out int month) ||
                !TryParseStrictDigits(yearStr, out int year))
            {
                throw new FormatException("componentes numéricos inválidos");
            }

            if (year <= 0)
            {
                throw new FormatException($"año inválido: {year}");
            }

            if (month < 1 || month > 12)
            {
                throw new FormatException($"mes fuera de rango: {month} (debe ser 1 a 12)");
            }

            int maxDays = DaysInMonth(year, month);
            if (day < 1 || day > maxDays)
            {
                throw new FormatException($"día fuera de rango: {day} para mes {month} del año {year}");
            }

            return new DateTime(year, month, day, 0, 0, 0, DateTimeKind.Utc);
        }

        /// <summary>Retorna true si la fecha corresponde a sábado o domingo.</summary>
        public static bool IsWeekend(DateTime date) =>
            date.DayOfWeek == DayOfWeek.Saturday || date.DayOfWeek == DayOfWeek.Sunday;

        /// <summary>Retorna true si la fecha con offset corresponde a sábado o domingo.</summary>
        public static bool IsWeekend(DateTimeOffset date) =>
            date.DayOfWeek == DayOfWeek.Saturday || date.DayOfWeek == DayOfWeek.Sunday;
    }
}
