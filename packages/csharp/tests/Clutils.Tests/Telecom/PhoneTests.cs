using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Serialization;
using Xunit;

namespace Clutils.Tests.TelecomTests
{
    public class PhoneTests
    {
        private sealed class NormalizationCase
        {
            [JsonPropertyName("input")]
            public string Input { get; set; } = string.Empty;

            [JsonPropertyName("expected_e164")]
            public string ExpectedE164 { get; set; } = string.Empty;

            [JsonPropertyName("valid")]
            public bool Valid { get; set; }

            [JsonPropertyName("description")]
            public string Description { get; set; } = string.Empty;
        }

        private sealed class FormattingCase
        {
            [JsonPropertyName("input")]
            public string Input { get; set; } = string.Empty;

            [JsonPropertyName("expected_international")]
            public string ExpectedInternational { get; set; } = string.Empty;

            [JsonPropertyName("expected_national")]
            public string ExpectedNational { get; set; } = string.Empty;

            [JsonPropertyName("expected_e164")]
            public string ExpectedE164 { get; set; } = string.Empty;

            [JsonPropertyName("description")]
            public string Description { get; set; } = string.Empty;
        }

        private sealed class PhoneTypeCase
        {
            [JsonPropertyName("input")]
            public string Input { get; set; } = string.Empty;

            [JsonPropertyName("expected_type")]
            public string ExpectedType { get; set; } = string.Empty;

            [JsonPropertyName("description")]
            public string Description { get; set; } = string.Empty;
        }

        private sealed class TelecomSuite
        {
            [JsonPropertyName("normalization")]
            public List<NormalizationCase> Normalization { get; set; } = new();

            [JsonPropertyName("formatting")]
            public List<FormattingCase> Formatting { get; set; } = new();

            [JsonPropertyName("phone_types")]
            public List<PhoneTypeCase> PhoneTypes { get; set; } = new();
        }

        private static TelecomSuite LoadSuite()
        {
            var json = SpecLoader.ReadSpecFile("telecom_phone.json");
            return JsonSerializer.Deserialize<TelecomSuite>(json)!;
        }

        private static PhoneType ParseExpectedType(string raw) => raw switch
        {
            "mobile" => PhoneType.Mobile,
            "landline" => PhoneType.Landline,
            "voip" => PhoneType.VoIP,
            "toll_free" => PhoneType.TollFree,
            "shared_cost" => PhoneType.SharedCost,
            "emergency" => PhoneType.Emergency,
            _ => PhoneType.Unknown,
        };

        public static IEnumerable<object[]> NormalizationCases()
        {
            foreach (var tc in LoadSuite().Normalization)
            {
                yield return new object[] { tc.Description, tc.Input, tc.ExpectedE164, tc.Valid };
            }
        }

        public static IEnumerable<object[]> FormattingCases()
        {
            foreach (var tc in LoadSuite().Formatting)
            {
                yield return new object[] { tc.Description, tc.Input, tc.ExpectedInternational, tc.ExpectedNational, tc.ExpectedE164 };
            }
        }

        public static IEnumerable<object[]> PhoneTypeCases()
        {
            foreach (var tc in LoadSuite().PhoneTypes)
            {
                yield return new object[] { tc.Description, tc.Input, tc.ExpectedType };
            }
        }

        [Theory]
        [MemberData(nameof(NormalizationCases))]
        public void NormalizeAndValidatePhone_MatchesCanonicalSpec(
            string description,
            string input,
            string expectedE164,
            bool valid)
        {
            _ = description;
            Assert.Equal(valid, Telecom.ValidatePhone(input));

            if (valid)
            {
                Assert.Equal(expectedE164, Telecom.NormalizePhone(input));
            }
            else
            {
                Assert.Throws<FormatException>(() => Telecom.NormalizePhone(input));
            }
        }

        [Theory]
        [MemberData(nameof(FormattingCases))]
        public void FormatPhone_MatchesCanonicalSpec(
            string description,
            string input,
            string expectedInternational,
            string expectedNational,
            string expectedE164)
        {
            _ = description;
            Assert.Equal(expectedInternational, Telecom.FormatPhone(input));
            Assert.Equal(expectedInternational, Telecom.FormatPhone(input, PhoneFormatStyle.International));
            Assert.Equal(expectedNational, Telecom.FormatPhone(input, PhoneFormatStyle.National));
            Assert.Equal(expectedE164, Telecom.FormatPhone(input, PhoneFormatStyle.E164));
        }

        [Theory]
        [MemberData(nameof(PhoneTypeCases))]
        public void GetPhoneType_MatchesCanonicalSpec(
            string description,
            string input,
            string expectedType)
        {
            _ = description;
            Assert.Equal(ParseExpectedType(expectedType), Telecom.GetPhoneType(input));
        }

        [Fact]
        public void NormalizeMobile_AcceptsMobileAndRejectsNonMobile()
        {
            Assert.Equal("+56912345678", Telecom.NormalizeMobile("912345678"));
            Assert.Equal("+56956123456", Telecom.NormalizeMobile("5612 3456"));
            Assert.Throws<FormatException>(() => Telecom.NormalizeMobile("221234567"));
            Assert.Throws<FormatException>(() => Telecom.NormalizeMobile("12345"));
            Assert.Throws<FormatException>(() => Telecom.NormalizeMobile(null));
        }

        [Fact]
        public void IsEmergencyNumber_DetectsOfficialShortCodes()
        {
            foreach (var code in new[] { "130", "131", "132", "133", "134", "136", "143", "149", "911", "1404", "1407", "1411", "1412", "1488" })
            {
                Assert.True(Telecom.IsEmergencyNumber(code), $"Esperaba emergencia para {code}");
            }

            Assert.False(Telecom.IsEmergencyNumber("123"));
            Assert.False(Telecom.IsEmergencyNumber("999"));
            Assert.False(Telecom.IsEmergencyNumber("1499"));
            Assert.False(Telecom.IsEmergencyNumber(null));
        }

        [Fact]
        public void FormatPhone_ThrowsOnInvalid()
        {
            Assert.Throws<FormatException>(() => Telecom.FormatPhone("invalid"));
        }
    }
}
