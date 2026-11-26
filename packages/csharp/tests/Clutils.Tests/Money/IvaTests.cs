using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Serialization;
using Xunit;

namespace Clutils.Tests.MoneyTests
{
    public class IvaTests
    {
        private sealed class IvaCase
        {
            [JsonPropertyName("input_amount")]
            public long InputAmount { get; set; }

            [JsonPropertyName("is_bruto")]
            public bool IsBruto { get; set; }

            [JsonPropertyName("expected_neto")]
            public long ExpectedNeto { get; set; }

            [JsonPropertyName("expected_iva")]
            public long ExpectedIva { get; set; }

            [JsonPropertyName("expected_bruto")]
            public long ExpectedBruto { get; set; }

            [JsonPropertyName("description")]
            public string Description { get; set; } = string.Empty;
        }

        private sealed class MoneySuite
        {
            [JsonPropertyName("iva_calculation")]
            public List<IvaCase> IvaCalculation { get; set; } = new();
        }

        public static IEnumerable<object[]> IvaCases()
        {
            var json = SpecLoader.ReadSpecFile("money_formatting.json");
            var suite = JsonSerializer.Deserialize<MoneySuite>(json)!;
            foreach (var tc in suite.IvaCalculation)
            {
                yield return new object[]
                {
                    tc.Description,
                    tc.InputAmount,
                    tc.IsBruto,
                    tc.ExpectedNeto,
                    tc.ExpectedIva,
                    tc.ExpectedBruto,
                };
            }
        }

        [Theory]
        [MemberData(nameof(IvaCases))]
        public void CalculateIVA_MatchesCanonicalSpec(
            string description,
            long inputAmount,
            bool isBruto,
            long expectedNeto,
            long expectedIva,
            long expectedBruto)
        {
            _ = description;
            var res = Money.CalculateIVA(inputAmount, isBruto);

            Assert.Equal(expectedNeto, res.Neto);
            Assert.Equal(expectedIva, res.Iva);
            Assert.Equal(expectedBruto, res.Bruto);
            Assert.Equal(res.Bruto, res.Neto + res.Iva);
        }

        [Fact]
        public void CalculateIVA_CustomRateAndExactRounding()
        {
            var fromNeto = Money.CalculateIVA(1000, false, new IvaOptions { Rate = 0.10 });
            Assert.Equal(1000, fromNeto.Neto);
            Assert.Equal(100, fromNeto.Iva);
            Assert.Equal(1100, fromNeto.Bruto);

            var fromBruto = Money.CalculateIVA(1100, true, new IvaOptions { Rate = 0.10 });
            Assert.Equal(1000, fromBruto.Neto);
            Assert.Equal(100, fromBruto.Iva);
            Assert.Equal(1100, fromBruto.Bruto);

            // Tasa +Infinity debe caer al 19% por defecto sin desbordar
            var infRate = Money.CalculateIVA(1000, false, new IvaOptions { Rate = double.PositiveInfinity });
            Assert.Equal(190, infRate.Iva);
            Assert.Equal(1190, infRate.Bruto);

            // Verificar redondeo AwayFromZero contra aritmética entera exacta en 1..10000
            for (long n = 1; n <= 10000; n++)
            {
                long expectedIva = (n * 19 + 50) / 100;
                var breakdown = Money.CalculateIVA(n, false);
                Assert.Equal(expectedIva, breakdown.Iva);
                Assert.Equal(n + expectedIva, breakdown.Bruto);
            }
        }
    }
}
