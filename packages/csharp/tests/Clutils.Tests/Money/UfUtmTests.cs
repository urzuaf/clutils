using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Serialization;
using Xunit;

namespace Clutils.Tests.MoneyTests
{
    public class UfUtmTests
    {
        private sealed class UfCase
        {
            [JsonPropertyName("amount")]
            public double Amount { get; set; }

            [JsonPropertyName("decimals")]
            public int Decimals { get; set; }

            [JsonPropertyName("expected_default")]
            public string ExpectedDefault { get; set; } = string.Empty;

            [JsonPropertyName("expected_no_symbol")]
            public string ExpectedNoSymbol { get; set; } = string.Empty;

            [JsonPropertyName("description")]
            public string Description { get; set; } = string.Empty;
        }

        private sealed class UtmCase
        {
            [JsonPropertyName("amount")]
            public long Amount { get; set; }

            [JsonPropertyName("expected_default")]
            public string ExpectedDefault { get; set; } = string.Empty;

            [JsonPropertyName("expected_no_symbol")]
            public string ExpectedNoSymbol { get; set; } = string.Empty;

            [JsonPropertyName("description")]
            public string Description { get; set; } = string.Empty;
        }

        private sealed class MoneySuite
        {
            [JsonPropertyName("uf_formatting")]
            public List<UfCase> UfFormatting { get; set; } = new();

            [JsonPropertyName("utm_formatting")]
            public List<UtmCase> UtmFormatting { get; set; } = new();
        }

        public static IEnumerable<object[]> UfCases()
        {
            var json = SpecLoader.ReadSpecFile("money_formatting.json");
            var suite = JsonSerializer.Deserialize<MoneySuite>(json)!;
            foreach (var tc in suite.UfFormatting)
            {
                yield return new object[]
                {
                    tc.Description,
                    tc.Amount,
                    tc.Decimals,
                    tc.ExpectedDefault,
                    tc.ExpectedNoSymbol,
                };
            }
        }

        public static IEnumerable<object[]> UtmCases()
        {
            var json = SpecLoader.ReadSpecFile("money_formatting.json");
            var suite = JsonSerializer.Deserialize<MoneySuite>(json)!;
            foreach (var tc in suite.UtmFormatting)
            {
                yield return new object[]
                {
                    tc.Description,
                    tc.Amount,
                    tc.ExpectedDefault,
                    tc.ExpectedNoSymbol,
                };
            }
        }

        [Theory]
        [MemberData(nameof(UfCases))]
        public void FormatUF_MatchesCanonicalSpec(
            string description,
            double amount,
            int decimals,
            string expectedDefault,
            string expectedNoSymbol)
        {
            _ = description;
            Assert.Equal(expectedDefault, Money.FormatUF(amount, new UfFormatOptions { Decimals = decimals }));
            Assert.Equal(expectedNoSymbol, Money.FormatUF(amount, new UfFormatOptions { Decimals = decimals, NoSymbol = true }));
        }

        [Fact]
        public void FormatUF_HandlesEdgeCasesAndExtremeValues()
        {
            Assert.Equal("UF 0,00", Money.FormatUF(double.NaN));
            Assert.Equal("UF 0,00", Money.FormatUF(double.PositiveInfinity));
            Assert.Equal("UF 0,00", Money.FormatUF(double.NegativeInfinity));
            Assert.Equal("UF 100.000.000.000.000.000.000,00", Money.FormatUF(1e20));
            Assert.Equal("UF 1,38", Money.FormatUF(1.375));
            Assert.StartsWith("UF 1,50000", Money.FormatUF(1.5, new UfFormatOptions { Decimals = 100 }));
        }

        [Theory]
        [MemberData(nameof(UtmCases))]
        public void FormatUTM_MatchesCanonicalSpec(
            string description,
            long amount,
            string expectedDefault,
            string expectedNoSymbol)
        {
            _ = description;
            Assert.Equal(expectedDefault, Money.FormatUTM(amount));
            Assert.Equal(expectedNoSymbol, Money.FormatUTM(amount, new UtmFormatOptions { NoSymbol = true }));
        }

        [Fact]
        public void FormatUTM_HandlesLongMinValueSafely()
        {
            Assert.Equal("UTM -9.223.372.036.854.775.808", Money.FormatUTM(long.MinValue));
        }
    }
}
