namespace Clutils
{
    /// <summary>
    /// Utilidades canónicas para la gestión de fechas y los 3 husos horarios oficiales de Chile
    /// (Continental, Magallanes e Insular).
    /// </summary>
    public static partial class Date
    {
        /// <summary>Identificador IANA oficial para Chile Continental ("America/Santiago").</summary>
        public const string TzSantiago = "America/Santiago";

        /// <summary>Identificador IANA oficial para la Región de Magallanes y Antártica Chilena ("America/Punta_Arenas").</summary>
        public const string TzPuntaArenas = "America/Punta_Arenas";

        /// <summary>Identificador IANA oficial para Chile Insular Occidental / Rapa Nui ("Pacific/Easter").</summary>
        public const string TzEaster = "Pacific/Easter";
    }
}
