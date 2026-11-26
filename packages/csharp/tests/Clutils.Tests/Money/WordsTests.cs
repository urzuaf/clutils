using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Serialization;
using Xunit;

namespace Clutils.Tests.MoneyTests
{
    public class WordsTests
    {
        private sealed class ToWordsCase
        {
            [JsonPropertyName("amount")]
            public long Amount { get; set; }

            [JsonPropertyName("expected")]
            public string Expected { get; set; } = string.Empty;

            [JsonPropertyName("expected_no_currency")]
            public string ExpectedNoCurrency { get; set; } = string.Empty;

            [JsonPropertyName("description")]
            public string Description { get; set; } = string.Empty;
        }

        private sealed class MoneySuite
        {
            [JsonPropertyName("to_words")]
            public List<ToWordsCase> ToWords { get; set; } = new();
        }

        public static IEnumerable<object[]> ToWordsCases()
        {
            var json = SpecLoader.ReadSpecFile("money_formatting.json");
            var suite = JsonSerializer.Deserialize<MoneySuite>(json)!;
            foreach (var tc in suite.ToWords)
            {
                yield return new object[]
                {
                    tc.Description,
                    tc.Amount,
                    tc.Expected,
                    tc.ExpectedNoCurrency,
                };
            }
        }

        [Theory]
        [MemberData(nameof(ToWordsCases))]
        public void ToWords_MatchesCanonicalSpec(
            string description,
            long amount,
            string expected,
            string expectedNoCurrency)
        {
            _ = description;
            Assert.Equal(expected, Money.ToWords(amount));
            Assert.Equal(expectedNoCurrency, Money.ToWords(amount, new ToWordsOptions { NoCurrency = true }));
        }
    }
}
