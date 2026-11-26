namespace Clutils
{
    /// <summary>
    /// Representa una de las 16 regiones oficiales de Chile según SUBDERE e INE.
    /// </summary>
    public sealed class Region
    {
        /// <summary>Código numérico oficial INE de 2 dígitos (ej: "13", "16", "05").</summary>
        public string Code { get; }

        /// <summary>Código estándar internacional ISO 3166-2:CL (ej: "CL-RM", "CL-NB").</summary>
        public string IsoCode { get; }

        /// <summary>Nombre oficial completo de la región.</summary>
        public string Name { get; }

        /// <summary>Abreviatura oficial (ej: "RM", "NUBLE", "VALPO").</summary>
        public string Abbreviation { get; }

        /// <summary>Ciudad capital regional.</summary>
        public string Capital { get; }

        /// <summary>Posición geográfica canónica de Norte a Sur (1 a 16).</summary>
        public int Order { get; }

        /// <summary>Inicializa una instancia inmutable de <see cref="Region"/>.</summary>
        public Region(string code, string isoCode, string name, string abbreviation, string capital, int order)
        {
            Code = code;
            IsoCode = isoCode;
            Name = name;
            Abbreviation = abbreviation;
            Capital = capital;
            Order = order;
        }
    }

    /// <summary>
    /// Representa una de las 56 provincias oficiales de Chile según SUBDERE e INE.
    /// </summary>
    public sealed class Provincia
    {
        /// <summary>Código numérico oficial INE de 3 dígitos (ej: "131", "161").</summary>
        public string Code { get; }

        /// <summary>Nombre oficial de la provincia.</summary>
        public string Name { get; }

        /// <summary>Código numérico oficial INE de la región a la que pertenece (ej: "13", "16").</summary>
        public string RegionCode { get; }

        /// <summary>Inicializa una instancia inmutable de <see cref="Provincia"/>.</summary>
        public Provincia(string code, string name, string regionCode)
        {
            Code = code;
            Name = name;
            RegionCode = regionCode;
        }
    }

    /// <summary>
    /// Representa una de las 346 comunas oficiales de Chile según SUBDERE e INE.
    /// </summary>
    public sealed class Comuna
    {
        /// <summary>Código Único Territorial oficial del INE de 5 dígitos (ej: "13120", "16101").</summary>
        public string Cut { get; }

        /// <summary>Nombre oficial de la comuna.</summary>
        public string Name { get; }

        /// <summary>Código oficial de la provincia de 3 dígitos (ej: "131", "161").</summary>
        public string ProvinciaCode { get; }

        /// <summary>Código oficial de la región de 2 dígitos (ej: "13", "16").</summary>
        public string RegionCode { get; }

        /// <summary>Inicializa una instancia inmutable de <see cref="Comuna"/>.</summary>
        public Comuna(string cut, string name, string provinciaCode, string regionCode)
        {
            Cut = cut;
            Name = name;
            ProvinciaCode = provinciaCode;
            RegionCode = regionCode;
        }
    }
}
