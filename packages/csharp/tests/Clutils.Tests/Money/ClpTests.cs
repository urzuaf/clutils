using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Serialization;
using Xunit;

namespace Clutils.Tests.MoneyTests
{
    public class ClpTests
    {
        private sealed class ClpFormatCase
        {
            [JsonPropertyName("amount")]
            public long Amount { get; set; }

            [JsonPropertyName("expected_default")]
            public string ExpectedDefault { get; set; } = string.Empty;

            [JsonPropertyName("expected_space")]
            public string ExpectedSpace { get; set; } = string.Empty;

            [JsonPropertyName("expected_symbol_clp")]
            public string ExpectedSymbolClp { get; set; } = string.Empty;

            [JsonPropertyName("expected_no_symbol")]
            public string ExpectedNoSymbol { get; set; } = string.Empty;

            [JsonPropertyName("expected_parentheses")]
            public string? ExpectedParentheses { get; set; }

            [JsonPropertyName("description")]
            public string Description { get; set; } = string.Empty;
        }

        private sealed class ClpParseCase
        {
            [JsonPropertyName("input")]
            public string Input { get; set; } = string.Empty;

            [JsonPropertyName("expected")]
            public long Expected { get; set; }

            [JsonPropertyName("valid")]
            public bool Valid { get; set; }

            [JsonPropertyName("description")]
            public string Description { get; set; } = string.Empty;
        }

        private sealed class MoneySuite
        {
            [JsonPropertyName("clp_formatting")]
            public List<ClpFormatCase> ClpFormatting { get; set; } = new();

            [JsonPropertyName("clp_parsing")]
            public List<ClpParseCase> ClpParsing { get; set; } = new();
        }

        public static IEnumerable<object[]> FormatCases()
        {
            var json = SpecLoader.ReadSpecFile("money_formatting.json");
            var suite = JsonSerializer.Deserialize<MoneySuite>(json)!;
            foreach (var tc in suite.ClpFormatting)
            {
                yield return new object[]
                {
                    tc.Description,
                    tc.Amount,
                    tc.ExpectedDefault,
                    tc.ExpectedSpace,
                    tc.ExpectedSymbolClp,
                    tc.ExpectedNoSymbol,
                    tc.ExpectedParentheses ?? string.Empty,
                };
            }
        }

        public static IEnumerable<object[]> ParseCases()
        {
            var json = SpecLoader.ReadSpecFile("money_formatting.json");
            var suite = JsonSerializer.Deserialize<MoneySuite>(json)!;
            foreach (var tc in suite.ClpParsing)
            {
                yield return new object[]
                {
                    tc.Description,
                    tc.Input,
                    tc.Expected,
                    tc.Valid,
                };
            }
        }

        [Theory]
        [MemberData(nameof(FormatCases))]
        public void FormatCLP_MatchesCanonicalSpec(
            string description,
            long amount,
            string expectedDefault,
            string expectedSpace,
            string expectedSymbolClp,
            string expectedNoSymbol,
            string expectedParentheses)
        {
            _ = description;
            Assert.Equal(expectedDefault, Money.FormatCLP(amount));
            Assert.Equal(expectedSpace, Money.FormatCLP(amount, new ClpFormatOptions { Space = true }));
            Assert.Equal(expectedSymbolClp, Money.FormatCLP(amount, new ClpFormatOptions { Symbol = "CLP", Space = true }));
            Assert.Equal(expectedNoSymbol, Money.FormatCLP(amount, new ClpFormatOptions { NoSymbol = true }));

            if (!string.IsNullOrEmpty(expectedParentheses))
            {
                Assert.Equal(expectedParentheses, Money.FormatCLP(amount, new ClpFormatOptions { NegativeFormat = NegativeFormat.Parentheses }));
            }
        }

        [Theory]
        [MemberData(nameof(ParseCases))]
        public void ParseCLP_MatchesCanonicalSpec(
            string description,
            string input,
            long expected,
            bool valid)
        {
            _ = description;
            if (valid)
            {
                Assert.Equal(expected, Money.ParseCLP(input));
            }
            else
            {
                Assert.Throws<FormatException>(() => Money.ParseCLP(input));
            }
        }

        [Fact]
        public void FormatCLP_HandlesLongMinValueSafely()
        {
            string formatted = Money.FormatCLP(long.MinValue);
            Assert.Equal("-$9.223.372.036.854.775.808", formatted);
            Assert.Equal(long.MinValue, Money.ParseCLP(formatted));
        }

        [Fact]
        public void ParseCLP_ThrowsOnNull()
        {
            Assert.Throws<FormatException>(() => Money.ParseCLP(null));
        }
    }
}
