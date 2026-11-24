using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Serialization;
using Xunit;

namespace Clutils.Tests.RutTests
{
    public class RutValidationTests
    {
        private sealed class ValidationCase
        {
            [JsonPropertyName("input")]
            public string Input { get; set; } = string.Empty;

            [JsonPropertyName("expected_clean")]
            public string ExpectedClean { get; set; } = string.Empty;

            [JsonPropertyName("expected_body")]
            public string? ExpectedBody { get; set; }

            [JsonPropertyName("expected_dv")]
            public string? ExpectedDv { get; set; }

            [JsonPropertyName("is_valid")]
            public bool IsValid { get; set; }

            [JsonPropertyName("description")]
            public string Description { get; set; } = string.Empty;
        }

        private sealed class ValidationSuite
        {
            [JsonPropertyName("cases")]
            public List<ValidationCase> Cases { get; set; } = new();
        }

        public static IEnumerable<object[]> ValidationCases()
        {
            var json = SpecLoader.ReadSpecFile("rut_validation.json");
            var suite = JsonSerializer.Deserialize<ValidationSuite>(json)!;
            foreach (var tc in suite.Cases)
            {
                yield return new object[]
                {
                    tc.Description,
                    tc.Input,
                    tc.ExpectedClean,
                    tc.ExpectedBody ?? string.Empty,
                    tc.ExpectedDv ?? string.Empty,
                    tc.IsValid,
                };
            }
        }

        [Theory]
        [MemberData(nameof(ValidationCases))]
        public void Clean_MatchesCanonicalSpec(
            string description,
            string input,
            string expectedClean,
            string expectedBody,
            string expectedDv,
            bool isValid)
        {
            _ = (description, expectedBody, expectedDv, isValid);
            Assert.Equal(expectedClean, Rut.Clean(input));
        }

        [Theory]
        [MemberData(nameof(ValidationCases))]
        public void Validate_MatchesCanonicalSpec(
            string description,
            string input,
            string expectedClean,
            string expectedBody,
            string expectedDv,
            bool isValid)
        {
            _ = (description, expectedClean, expectedBody, expectedDv);
            Assert.Equal(isValid, Rut.Validate(input));
        }

        [Theory]
        [MemberData(nameof(ValidationCases))]
        public void CalculateDv_MatchesCanonicalSpec(
            string description,
            string input,
            string expectedClean,
            string expectedBody,
            string expectedDv,
            bool isValid)
        {
            _ = (description, input, expectedClean, isValid);
            if (!string.IsNullOrEmpty(expectedBody) && !string.IsNullOrEmpty(expectedDv))
            {
                Assert.Equal(expectedDv, Rut.CalculateDv(expectedBody));
                Assert.Equal(expectedDv, Rut.CalculateDv(int.Parse(expectedBody)));
            }
        }

        [Fact]
        public void ValidateDetailed_DiagnosesSpecificErrorCodes()
        {
            Assert.Null(Rut.ValidateDetailed("12.345.678-5"));

            Assert.Equal(RutErrorCode.Empty, Rut.ValidateDetailed("")?.Code);
            Assert.Equal(RutErrorCode.Empty, Rut.ValidateDetailed("   ")?.Code);
            Assert.Equal(RutErrorCode.Empty, Rut.ValidateDetailed(null)?.Code);

            Assert.Equal(RutErrorCode.MalformedFormat, Rut.ValidateDetailed("12..345..678-5")?.Code);
            Assert.Equal(RutErrorCode.MalformedFormat, Rut.ValidateDetailed("12345678--5")?.Code);
            Assert.Equal(RutErrorCode.MalformedFormat, Rut.ValidateDetailed("12-345-678-5")?.Code);
            Assert.Equal(RutErrorCode.MalformedFormat, Rut.ValidateDetailed("12345678-.5")?.Code);
            Assert.Equal(RutErrorCode.MalformedFormat, Rut.ValidateDetailed("18123458-K.")?.Code);
            Assert.Equal(RutErrorCode.MalformedFormat, Rut.ValidateDetailed(".12345678-5")?.Code);
            Assert.Equal(RutErrorCode.MalformedFormat, Rut.ValidateDetailed("1.2.345678-5")?.Code);

            Assert.Equal(RutErrorCode.InvalidLength, Rut.ValidateDetailed("1")?.Code);
            Assert.Equal(RutErrorCode.InvalidLength, Rut.ValidateDetailed("K")?.Code);
            Assert.Equal(RutErrorCode.InvalidLength, Rut.ValidateDetailed("12345678901-5")?.Code);

            Assert.Equal(RutErrorCode.InvalidCharacters, Rut.ValidateDetailed("18.12A.458-K")?.Code);
            Assert.Equal(RutErrorCode.InvalidCharacters, Rut.ValidateDetailed("abcdefgh-K")?.Code);
            Assert.Equal(RutErrorCode.InvalidCharacters, Rut.ValidateDetailed("+12345678-5")?.Code);

            Assert.Equal(RutErrorCode.InvalidDv, Rut.ValidateDetailed("12.345.678-0")?.Code);
            Assert.Equal(RutErrorCode.InvalidDv, Rut.ValidateDetailed("18.123.458-1")?.Code);
        }

        [Fact]
        public void CalculateDv_RejectsInvalidInputs()
        {
            Assert.Equal("", Rut.CalculateDv(0));
            Assert.Equal("", Rut.CalculateDv(-12345678));
            Assert.Equal("", Rut.CalculateDv((string?)null));
            Assert.Equal("", Rut.CalculateDv(""));
            Assert.Equal("", Rut.CalculateDv("12345678abc"));
            Assert.Equal("", Rut.CalculateDv("+12345678"));
        }
    }
}
