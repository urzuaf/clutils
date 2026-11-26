using System;
using System.Collections.Generic;
using System.Globalization;
using System.Text.Json;
using System.Text.Json.Serialization;
using Xunit;

namespace Clutils.Tests.DateTests
{
    public class DateFormattingAndTimezoneTests
    {
        private sealed class ParseDateCase
        {
            [JsonPropertyName("input")]
            public string Input { get; set; } = string.Empty;

            [JsonPropertyName("expected_year")]
            public int ExpectedYear { get; set; }

            [JsonPropertyName("expected_month")]
            public int ExpectedMonth { get; set; }

            [JsonPropertyName("expected_day")]
            public int ExpectedDay { get; set; }

            [JsonPropertyName("valid")]
            public bool Valid { get; set; }

            [JsonPropertyName("description")]
            public string Description { get; set; } = string.Empty;
        }

        private sealed class FormatDateCase
        {
            [JsonPropertyName("year")]
            public int Year { get; set; }

            [JsonPropertyName("month")]
            public int Month { get; set; }

            [JsonPropertyName("day")]
            public int Day { get; set; }

            [JsonPropertyName("expected_numeric")]
            public string ExpectedNumeric { get; set; } = string.Empty;

            [JsonPropertyName("expected_numeric_hyphen")]
            public string ExpectedNumericHyphen { get; set; } = string.Empty;

            [JsonPropertyName("expected_formal")]
            public string ExpectedFormal { get; set; } = string.Empty;

            [JsonPropertyName("expected_full")]
            public string ExpectedFull { get; set; } = string.Empty;

            [JsonPropertyName("description")]
            public string Description { get; set; } = string.Empty;
        }

        private sealed class IsWeekendCase
        {
            [JsonPropertyName("year")]
            public int Year { get; set; }

            [JsonPropertyName("month")]
            public int Month { get; set; }

            [JsonPropertyName("day")]
            public int Day { get; set; }

            [JsonPropertyName("expected")]
            public bool Expected { get; set; }

            [JsonPropertyName("description")]
            public string Description { get; set; } = string.Empty;
        }

        private sealed class TimezoneCase
        {
            [JsonPropertyName("identifier")]
            public string Identifier { get; set; } = string.Empty;

            [JsonPropertyName("expected_iana")]
            public string ExpectedIana { get; set; } = string.Empty;

            [JsonPropertyName("utc_timestamp")]
            public string UtcTimestamp { get; set; } = string.Empty;

            [JsonPropertyName("expected_hour")]
            public int ExpectedHour { get; set; }

            [JsonPropertyName("description")]
            public string Description { get; set; } = string.Empty;
        }

        private sealed class DateSuite
        {
            [JsonPropertyName("parse_date")]
            public List<ParseDateCase> ParseDate { get; set; } = new();

            [JsonPropertyName("format_date")]
            public List<FormatDateCase> FormatDate { get; set; } = new();

            [JsonPropertyName("is_weekend")]
            public List<IsWeekendCase> IsWeekend { get; set; } = new();

            [JsonPropertyName("timezones")]
            public List<TimezoneCase> Timezones { get; set; } = new();
        }

        private static DateSuite LoadSuite()
        {
            var json = SpecLoader.ReadSpecFile("date_formatting.json");
            return JsonSerializer.Deserialize<DateSuite>(json)!;
        }

        public static IEnumerable<object[]> ParseDateCases()
        {
            foreach (var tc in LoadSuite().ParseDate)
            {
                yield return new object[]
                {
                    tc.Description,
                    tc.Input,
                    tc.ExpectedYear,
                    tc.ExpectedMonth,
                    tc.ExpectedDay,
                    tc.Valid,
                };
            }
        }

        public static IEnumerable<object[]> FormatDateCases()
        {
            foreach (var tc in LoadSuite().FormatDate)
            {
                yield return new object[]
                {
                    tc.Description,
                    tc.Year,
                    tc.Month,
                    tc.Day,
                    tc.ExpectedNumeric,
                    tc.ExpectedNumericHyphen,
                    tc.ExpectedFormal,
                    tc.ExpectedFull,
                };
            }
        }

        public static IEnumerable<object[]> IsWeekendCases()
        {
            foreach (var tc in LoadSuite().IsWeekend)
            {
                yield return new object[]
                {
                    tc.Description,
                    tc.Year,
                    tc.Month,
                    tc.Day,
                    tc.Expected,
                };
            }
        }

        public static IEnumerable<object[]> TimezoneCases()
        {
            foreach (var tc in LoadSuite().Timezones)
            {
                yield return new object[]
                {
                    tc.Description,
                    tc.Identifier,
                    tc.ExpectedIana,
                    tc.UtcTimestamp,
                    tc.ExpectedHour,
                };
            }
        }

        [Theory]
        [MemberData(nameof(ParseDateCases))]
        public void ParseDate_MatchesCanonicalSpec(
            string description,
            string input,
            int expectedYear,
            int expectedMonth,
            int expectedDay,
            bool valid)
        {
            _ = description;
            if (valid)
            {
                var dt = Date.ParseDate(input);
                Assert.Equal(expectedYear, dt.Year);
                Assert.Equal(expectedMonth, dt.Month);
                Assert.Equal(expectedDay, dt.Day);
            }
            else
            {
                Assert.Throws<FormatException>(() => Date.ParseDate(input));
            }
        }

        [Theory]
        [MemberData(nameof(FormatDateCases))]
        public void FormatDate_MatchesCanonicalSpec(
            string description,
            int year,
            int month,
            int day,
            string expectedNumeric,
            string expectedNumericHyphen,
            string expectedFormal,
            string expectedFull)
        {
            _ = description;
            var dt = new DateTime(year, month, day, 0, 0, 0, DateTimeKind.Utc);
            var dto = new DateTimeOffset(dt);

            Assert.Equal(expectedNumeric, Date.FormatDate(dt));
            Assert.Equal(expectedNumeric, Date.FormatDate(dt, DateFormatStyle.Numeric));
            Assert.Equal(expectedNumericHyphen, Date.FormatDate(dt, DateFormatStyle.NumericHyphen));
            Assert.Equal(expectedFormal, Date.FormatDate(dt, DateFormatStyle.Formal));
            Assert.Equal(expectedFull, Date.FormatDate(dt, DateFormatStyle.Full));
            Assert.Equal(expectedFull, Date.FormatDate(dto, DateFormatStyle.Full));
        }

        [Fact]
        public void FormatDate_ZeroPadsHistoricalYears()
        {
            var dt = Date.ParseDate("05/01/0500");
            Assert.Equal("05/01/0500", Date.FormatDate(dt, DateFormatStyle.Numeric));
            Assert.Equal("05-01-0500", Date.FormatDate(dt, DateFormatStyle.NumericHyphen));
            Assert.Equal("5 de enero de 0500", Date.FormatDate(dt, DateFormatStyle.Formal));
        }

        [Theory]
        [MemberData(nameof(IsWeekendCases))]
        public void IsWeekend_MatchesCanonicalSpec(
            string description,
            int year,
            int month,
            int day,
            bool expected)
        {
            _ = description;
            var dt = new DateTime(year, month, day, 0, 0, 0, DateTimeKind.Utc);
            Assert.Equal(expected, Date.IsWeekend(dt));
            Assert.Equal(expected, Date.IsWeekend(new DateTimeOffset(dt)));
        }

        [Theory]
        [MemberData(nameof(TimezoneCases))]
        public void Timezones_MatchesCanonicalSpec(
            string description,
            string identifier,
            string expectedIana,
            string utcTimestamp,
            int expectedHour)
        {
            _ = description;
            var utc = DateTimeOffset.Parse(utcTimestamp, CultureInfo.InvariantCulture, DateTimeStyles.AssumeUniversal);

            DateTimeOffset converted;
            string ianaConst;
            switch (identifier)
            {
                case "santiago":
                    ianaConst = Date.TzSantiago;
                    converted = Date.ToChileTime(utc);
                    break;
                case "punta_arenas":
                    ianaConst = Date.TzPuntaArenas;
                    converted = Date.ToMagallanesTime(utc);
                    break;
                case "easter_island":
                    ianaConst = Date.TzEaster;
                    converted = Date.ToEasterIslandTime(utc);
                    break;
                default:
                    throw new ArgumentException($"Identificador desconocido: {identifier}");
            }

            Assert.Equal(expectedIana, ianaConst);
            Assert.Equal(expectedHour, converted.Hour);
        }

        [Fact]
        public void NowHelpers_ReturnValidOffsets()
        {
            var chile = Date.NowChile();
            var magallanes = Date.NowMagallanes();
            var easter = Date.NowEasterIsland();

            Assert.True(chile.Offset == TimeSpan.FromHours(-3) || chile.Offset == TimeSpan.FromHours(-4));
            Assert.Equal(TimeSpan.FromHours(-3), magallanes.Offset);
            Assert.True(easter.Offset == TimeSpan.FromHours(-5) || easter.Offset == TimeSpan.FromHours(-6));
        }
    }
}
