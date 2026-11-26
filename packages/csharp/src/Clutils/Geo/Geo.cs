using System;
using System.Collections.Generic;

namespace Clutils
{
    /// <summary>
    /// División Político-Administrativa (DPA) oficial de Chile según SUBDERE e INE:
    /// 16 regiones, 56 provincias y 346 comunas con sus Códigos Únicos Territoriales (CUT).
    /// </summary>
    public static partial class Geo
    {
        private static readonly Dictionary<string, string> RegionCodeToRoman = new Dictionary<string, string>(StringComparer.Ordinal)
        {
            ["01"] = "I",
            ["02"] = "II",
            ["03"] = "III",
            ["04"] = "IV",
            ["05"] = "V",
            ["06"] = "VI",
            ["07"] = "VII",
            ["08"] = "VIII",
            ["09"] = "IX",
            ["10"] = "X",
            ["11"] = "XI",
            ["12"] = "XII",
            ["13"] = "RM",
            ["14"] = "XIV",
            ["15"] = "XV",
            ["16"] = "XVI",
        };

        // ====================================================================
        // REGIONES
        // ====================================================================

        /// <summary>Retorna las 16 regiones oficiales de Chile ordenadas geográficamente de Norte a Sur.</summary>
        public static IReadOnlyList<Region> GetRegiones() => Array.AsReadOnly(RegionesData);

        /// <summary>
        /// Busca una región por código INE ("13", "16", "1"), código ISO ("CL-RM", "cl-nb"),
        /// abreviatura ("RM", "NUBLE"), número romano ("VIII", "XVI") o nombre tolerante a tildes y eñes.
        /// </summary>
        public static Region? FindRegion(string? query) =>
            SearchEntity(
                query,
                RegionesData,
                r => RegionCodeToRoman.TryGetValue(r.Code, out string roman)
                    ? new[] { r.Code, r.IsoCode, r.Abbreviation, roman }
                    : new[] { r.Code, r.IsoCode, r.Abbreviation },
                r => r.Name);

        // ====================================================================
        // PROVINCIAS
        // ====================================================================

        /// <summary>Retorna las 56 provincias oficiales de Chile.</summary>
        public static IReadOnlyList<Provincia> GetProvincias() => Array.AsReadOnly(ProvinciasData);

        /// <summary>
        /// Retorna las provincias pertenecientes a una región específica (acepta código INE, ISO o nombre).
        /// </summary>
        public static IReadOnlyList<Provincia> GetProvinciasByRegion(string? regionQuery)
        {
            Region? reg = FindRegion(regionQuery);
            if (reg == null)
            {
                return Array.Empty<Provincia>();
            }

            var result = new List<Provincia>();
            for (int i = 0; i < ProvinciasData.Length; i++)
            {
                if (ProvinciasData[i].RegionCode == reg.Code)
                {
                    result.Add(ProvinciasData[i]);
                }
            }

            return result.ToArray();
        }

        /// <summary>Busca una provincia por código INE de 3 dígitos ("131", "161") o nombre oficial.</summary>
        public static Provincia? FindProvincia(string? query) =>
            SearchEntity(query, ProvinciasData, p => new[] { p.Code }, p => p.Name);

        /// <summary>Busca todas las provincias que coincidan con el término, ordenadas por relevancia.</summary>
        public static IReadOnlyList<Provincia> SearchProvincias(string? query) =>
            SearchAllEntities(query, ProvinciasData, p => new[] { p.Code }, p => p.Name);

        /// <summary>Resuelve la región oficial a la que pertenece una provincia dada.</summary>
        public static Region? GetRegionByProvincia(string? provinciaQuery)
        {
            Provincia? prov = FindProvincia(provinciaQuery);
            return prov == null ? null : FindRegion(prov.RegionCode);
        }

        // ====================================================================
        // COMUNAS
        // ====================================================================

        /// <summary>Retorna las 346 comunas oficiales de Chile.</summary>
        public static IReadOnlyList<Comuna> GetComunas() => Array.AsReadOnly(ComunasData);

        /// <summary>
        /// Retorna las comunas pertenecientes a una región específica (acepta código INE, ISO o nombre).
        /// </summary>
        public static IReadOnlyList<Comuna> GetComunasByRegion(string? regionQuery)
        {
            Region? reg = FindRegion(regionQuery);
            if (reg == null)
            {
                return Array.Empty<Comuna>();
            }

            var result = new List<Comuna>();
            for (int i = 0; i < ComunasData.Length; i++)
            {
                if (ComunasData[i].RegionCode == reg.Code)
                {
                    result.Add(ComunasData[i]);
                }
            }

            return result.ToArray();
        }

        /// <summary>
        /// Retorna las comunas pertenecientes a una provincia específica (acepta código INE o nombre).
        /// </summary>
        public static IReadOnlyList<Comuna> GetComunasByProvincia(string? provinciaQuery)
        {
            Provincia? prov = FindProvincia(provinciaQuery);
            if (prov == null)
            {
                return Array.Empty<Comuna>();
            }

            var result = new List<Comuna>();
            for (int i = 0; i < ComunasData.Length; i++)
            {
                if (ComunasData[i].ProvinciaCode == prov.Code)
                {
                    result.Add(ComunasData[i]);
                }
            }

            return result.ToArray();
        }

        /// <summary>
        /// Busca una comuna por Código Único Territorial (CUT de 5 dígitos, ej: "13120")
        /// o por nombre oficial (tolerante a mayúsculas, tildes, eñes y Unicode NFD).
        /// </summary>
        public static Comuna? FindComuna(string? query) =>
            SearchEntity(query, ComunasData, c => new[] { c.Cut }, c => c.Name);

        /// <summary>
        /// Busca todas las comunas que coincidan con el término, ordenadas por relevancia (autocompletado).
        /// </summary>
        public static IReadOnlyList<Comuna> SearchComunas(string? query) =>
            SearchAllEntities(query, ComunasData, c => new[] { c.Cut }, c => c.Name);

        /// <summary>Resuelve la región oficial a la que pertenece una comuna dada.</summary>
        public static Region? GetRegionByComuna(string? comunaQuery)
        {
            Comuna? comuna = FindComuna(comunaQuery);
            return comuna == null ? null : FindRegion(comuna.RegionCode);
        }

        /// <summary>Resuelve la provincia oficial a la que pertenece una comuna dada.</summary>
        public static Provincia? GetProvinciaByComuna(string? comunaQuery)
        {
            Comuna? comuna = FindComuna(comunaQuery);
            return comuna == null ? null : FindProvincia(comuna.ProvinciaCode);
        }
    }
}
