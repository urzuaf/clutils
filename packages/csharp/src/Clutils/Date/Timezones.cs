using System;

namespace Clutils
{
    public static partial class Date
    {
        private static readonly Lazy<TimeZoneInfo> SantiagoZone = new Lazy<TimeZoneInfo>(
            () => ResolveZone(TzSantiago, "Pacific SA Standard Time", "CLT", TimeSpan.FromHours(-4)));

        private static readonly Lazy<TimeZoneInfo> PuntaArenasZone = new Lazy<TimeZoneInfo>(
            () => ResolveZone(TzPuntaArenas, "Magallanes Standard Time", "-03", TimeSpan.FromHours(-3)));

        private static readonly Lazy<TimeZoneInfo> EasterZone = new Lazy<TimeZoneInfo>(
            () => ResolveZone(TzEaster, "Easter Island Standard Time", "EAST", TimeSpan.FromHours(-6)));

        private static TimeZoneInfo ResolveZone(string ianaId, string windowsId, string fallbackId, TimeSpan fallbackOffset)
        {
            try
            {
                return TimeZoneInfo.FindSystemTimeZoneById(ianaId);
            }
            catch (TimeZoneNotFoundException)
            {
            }
            catch (InvalidTimeZoneException)
            {
            }

            try
            {
                return TimeZoneInfo.FindSystemTimeZoneById(windowsId);
            }
            catch (TimeZoneNotFoundException)
            {
            }
            catch (InvalidTimeZoneException)
            {
            }

            return TimeZoneInfo.CreateCustomTimeZone(fallbackId, fallbackOffset, fallbackId, fallbackId);
        }

        private static DateTimeOffset ToOffset(DateTime dt, TimeZoneInfo zone)
        {
            DateTimeOffset dto = dt.Kind == DateTimeKind.Unspecified
                ? new DateTimeOffset(DateTime.SpecifyKind(dt, DateTimeKind.Utc))
                : new DateTimeOffset(dt);
            return TimeZoneInfo.ConvertTime(dto, zone);
        }

        /// <summary>Retorna el <see cref="TimeZoneInfo"/> oficial de Chile Continental ("America/Santiago").</summary>
        public static TimeZoneInfo TimeZoneSantiago() => SantiagoZone.Value;

        /// <summary>Retorna el <see cref="TimeZoneInfo"/> oficial de la Región de Magallanes (UTC-3 permanente).</summary>
        public static TimeZoneInfo TimeZonePuntaArenas() => PuntaArenasZone.Value;

        /// <summary>Retorna el <see cref="TimeZoneInfo"/> oficial de Rapa Nui / Isla de Pascua ("Pacific/Easter").</summary>
        public static TimeZoneInfo TimeZoneEaster() => EasterZone.Value;

        /// <summary>Retorna la fecha y hora actual en Chile Continental ("America/Santiago").</summary>
        public static DateTimeOffset NowChile() => TimeZoneInfo.ConvertTime(DateTimeOffset.UtcNow, SantiagoZone.Value);

        /// <summary>Retorna la fecha y hora actual en la Región de Magallanes ("America/Punta_Arenas", UTC-3 permanente).</summary>
        public static DateTimeOffset NowMagallanes() => TimeZoneInfo.ConvertTime(DateTimeOffset.UtcNow, PuntaArenasZone.Value);

        /// <summary>Retorna la fecha y hora actual en Rapa Nui ("Pacific/Easter").</summary>
        public static DateTimeOffset NowEasterIsland() => TimeZoneInfo.ConvertTime(DateTimeOffset.UtcNow, EasterZone.Value);

        /// <summary>Convierte cualquier <see cref="DateTimeOffset"/> a la zona oficial de Chile Continental.</summary>
        public static DateTimeOffset ToChileTime(DateTimeOffset dt) => TimeZoneInfo.ConvertTime(dt, SantiagoZone.Value);

        /// <summary>Convierte cualquier <see cref="DateTime"/> a la zona oficial de Chile Continental.</summary>
        public static DateTimeOffset ToChileTime(DateTime dt) => ToOffset(dt, SantiagoZone.Value);

        /// <summary>Convierte cualquier <see cref="DateTimeOffset"/> a la zona oficial de Magallanes.</summary>
        public static DateTimeOffset ToMagallanesTime(DateTimeOffset dt) => TimeZoneInfo.ConvertTime(dt, PuntaArenasZone.Value);

        /// <summary>Convierte cualquier <see cref="DateTime"/> a la zona oficial de Magallanes.</summary>
        public static DateTimeOffset ToMagallanesTime(DateTime dt) => ToOffset(dt, PuntaArenasZone.Value);

        /// <summary>Convierte cualquier <see cref="DateTimeOffset"/> a la zona oficial de Rapa Nui.</summary>
        public static DateTimeOffset ToEasterIslandTime(DateTimeOffset dt) => TimeZoneInfo.ConvertTime(dt, EasterZone.Value);

        /// <summary>Convierte cualquier <see cref="DateTime"/> a la zona oficial de Rapa Nui.</summary>
        public static DateTimeOffset ToEasterIslandTime(DateTime dt) => ToOffset(dt, EasterZone.Value);
    }
}
