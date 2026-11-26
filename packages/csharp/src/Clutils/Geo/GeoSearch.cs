using System;
using System.Collections.Generic;
using System.Globalization;
using System.Text;

namespace Clutils
{
    public static partial class Geo
    {
        internal static string NormalizeSearch(string? s)
        {
            if (string.IsNullOrWhiteSpace(s))
            {
                return string.Empty;
            }

            string decomposed = s!.Trim().Normalize(NormalizationForm.FormD);
            var sb = new StringBuilder(decomposed.Length);
            bool prevSpace = false;

            for (int i = 0; i < decomposed.Length; i++)
            {
                char ch = decomposed[i];
                if (CharUnicodeInfo.GetUnicodeCategory(ch) == UnicodeCategory.NonSpacingMark)
                {
                    continue;
                }

                if (ch == '\'' || ch == '.' || ch == '-' || ch == '_' || ch == '/')
                {
                    continue;
                }

                if (char.IsWhiteSpace(ch))
                {
                    if (sb.Length > 0 && !prevSpace)
                    {
                        sb.Append(' ');
                        prevSpace = true;
                    }

                    continue;
                }

                sb.Append(char.ToLowerInvariant(ch));
                prevSpace = false;
            }

            if (sb.Length > 0 && sb[sb.Length - 1] == ' ')
            {
                sb.Length--;
            }

            return sb.ToString();
        }

        private static T? SearchEntity<T>(
            string? query,
            T[] items,
            Func<T, string[]> getIdents,
            Func<T, string> getName) where T : class
        {
            if (string.IsNullOrWhiteSpace(query))
            {
                return null;
            }

            string trimmed = query!.Trim();
            string normQuery = NormalizeSearch(trimmed);

            string padded = trimmed.Length == 1 && trimmed[0] >= '0' && trimmed[0] <= '9'
                ? "0" + trimmed
                : trimmed;

            // Paso 1: Coincidencia por identificadores exactos
            for (int i = 0; i < items.Length; i++)
            {
                string[] idents = getIdents(items[i]);
                for (int j = 0; j < idents.Length; j++)
                {
                    string id = idents[j];
                    if (string.IsNullOrEmpty(id))
                    {
                        continue;
                    }

                    if (string.Equals(id, trimmed, StringComparison.OrdinalIgnoreCase) ||
                        string.Equals(id, padded, StringComparison.OrdinalIgnoreCase) ||
                        (normQuery.Length > 0 && NormalizeSearch(id) == normQuery))
                    {
                        return items[i];
                    }
                }
            }

            if (normQuery.Length == 0)
            {
                return null;
            }

            // Paso 2: Coincidencia exacta por nombre normalizado
            for (int i = 0; i < items.Length; i++)
            {
                if (NormalizeSearch(getName(items[i])) == normQuery)
                {
                    return items[i];
                }
            }

            // Paso 3: Coincidencia por prefijo (coherente con SearchAllEntities[0])
            for (int i = 0; i < items.Length; i++)
            {
                if (NormalizeSearch(getName(items[i])).StartsWith(normQuery, StringComparison.Ordinal))
                {
                    return items[i];
                }
            }

            // Paso 4: Coincidencia parcial contenida (subcadena)
            for (int i = 0; i < items.Length; i++)
            {
                if (NormalizeSearch(getName(items[i])).IndexOf(normQuery, StringComparison.Ordinal) >= 0)
                {
                    return items[i];
                }
            }

            return null;
        }

        private static IReadOnlyList<T> SearchAllEntities<T>(
            string? query,
            T[] items,
            Func<T, string[]> getIdents,
            Func<T, string> getName) where T : class
        {
            if (string.IsNullOrWhiteSpace(query))
            {
                return Array.Empty<T>();
            }

            string trimmed = query!.Trim();
            string normQuery = NormalizeSearch(trimmed);

            string padded = trimmed.Length == 1 && trimmed[0] >= '0' && trimmed[0] <= '9'
                ? "0" + trimmed
                : trimmed;

            var exactMatches = new List<T>();
            var prefixMatches = new List<T>();
            var substringMatches = new List<T>();
            var seen = new HashSet<int>();

            // 1. Identificadores exactos
            for (int i = 0; i < items.Length; i++)
            {
                string[] idents = getIdents(items[i]);
                for (int j = 0; j < idents.Length; j++)
                {
                    string id = idents[j];
                    if (string.IsNullOrEmpty(id))
                    {
                        continue;
                    }

                    if (string.Equals(id, trimmed, StringComparison.OrdinalIgnoreCase) ||
                        string.Equals(id, padded, StringComparison.OrdinalIgnoreCase) ||
                        (normQuery.Length > 0 && NormalizeSearch(id) == normQuery))
                    {
                        if (seen.Add(i))
                        {
                            exactMatches.Add(items[i]);
                        }

                        break;
                    }
                }
            }

            if (normQuery.Length == 0)
            {
                return exactMatches.Count == 0 ? Array.Empty<T>() : exactMatches.ToArray();
            }

            // 2. Coincidencia exacta por nombre
            for (int i = 0; i < items.Length; i++)
            {
                if (seen.Contains(i)) continue;
                if (NormalizeSearch(getName(items[i])) == normQuery)
                {
                    seen.Add(i);
                    exactMatches.Add(items[i]);
                }
            }

            // 3. Coincidencia por prefijo del nombre
            for (int i = 0; i < items.Length; i++)
            {
                if (seen.Contains(i)) continue;
                if (NormalizeSearch(getName(items[i])).StartsWith(normQuery, StringComparison.Ordinal))
                {
                    seen.Add(i);
                    prefixMatches.Add(items[i]);
                }
            }

            // 4. Coincidencia por subcadena
            for (int i = 0; i < items.Length; i++)
            {
                if (seen.Contains(i)) continue;
                if (NormalizeSearch(getName(items[i])).IndexOf(normQuery, StringComparison.Ordinal) >= 0)
                {
                    seen.Add(i);
                    substringMatches.Add(items[i]);
                }
            }

            var result = new List<T>(exactMatches.Count + prefixMatches.Count + substringMatches.Count);
            result.AddRange(exactMatches);
            result.AddRange(prefixMatches);
            result.AddRange(substringMatches);
            return result.ToArray();
        }
    }
}
