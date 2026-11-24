using System;
using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Serialization;
using Xunit;

namespace Clutils.Tests.RutTests
{
    public class RutFormattingTests
    {
        private sealed class FormatCase
        {
            [JsonPropertyName("input")]
            public string Input { get; set; } = string.Empty;

            [JsonPropertyName("expected_default")]
            public string ExpectedDefault { get; set; } = string.Empty;

            [JsonPropertyName("expected_no_dots")]
            public string ExpectedNoDots { get; set; } = string.Empty;

            [JsonPropertyName("expected_no_hyphen")]
            public string ExpectedNoHyphen { get; set; } = string.Empty;

            [JsonPropertyName("expected_plain")]
            public string ExpectedPlain { get; set; } = string.Empty;

            [JsonPropertyName("description")]
            public string Description { get; set; } = string.Empty;
        }

        private sealed class FormatSuite
        {
            [JsonPropertyName("cases")]
            public List<FormatCase> Cases { get; set; } = new();
        }

        private sealed class SplitCase
        {
            [JsonPropertyName("input")]
            public string Input { get; set; } = string.Empty;

            [JsonPropertyName("expected_body")]
            public string? ExpectedBody { get; set; }

            [JsonPropertyName("expected_dv")]
            public string? ExpectedDv { get; set; }

            [JsonPropertyName("expected_int")]
            public int? ExpectedInt { get; set; }

            [JsonPropertyName("is_error")]
            public bool IsError { get; set; }

            [JsonPropertyName("description")]
            public string Description { get; set; } = string.Empty;
        }

        private sealed class SplitSuite
        {
            [JsonPropertyName("cases")]
            public List<SplitCase> Cases { get; set; } = new();
        }

        private sealed class MaskCase
        {
            [JsonPropertyName("input")]
            public string Input { get; set; } = string.Empty;

            [JsonPropertyName("expected_default")]
            public string ExpectedDefault { get; set; } = string.Empty;

            [JsonPropertyName("expected_plain")]
            public string ExpectedPlain { get; set; } = string.Empty;

            [JsonPropertyName("description")]
            public string Description { get; set; } = string.Empty;
        }

        private sealed class MaskSuite
        {
            [JsonPropertyName("cases")]
            public List<MaskCase> Cases { get; set; } = new();
        }

        public static IEnumerable<object[]> FormatCases()
        {
            var json = SpecLoader.ReadSpecFile("rut_formatting.json");
            var suite = JsonSerializer.Deserialize<FormatSuite>(json)!;
            foreach (var tc in suite.Cases)
            {
                yield return new object[]
                {
                    tc.Description,
                    tc.Input,
                    tc.ExpectedDefault,
                    tc.ExpectedNoDots,
                    tc.ExpectedNoHyphen,
                    tc.ExpectedPlain,
                };
            }
        }

        public static IEnumerable<object[]> SplitCases()
        {
            var json = SpecLoader.ReadSpecFile("rut_split.json");
            var suite = JsonSerializer.Deserialize<SplitSuite>(json)!;
            foreach (var tc in suite.Cases)
            {
                yield return new object[]
                {
                    tc.Description,
                    tc.Input,
                    tc.ExpectedBody ?? string.Empty,
                    tc.ExpectedDv ?? string.Empty,
                    tc.ExpectedInt ?? -1,
                    tc.IsError,
                };
            }
        }

        public static IEnumerable<object[]> MaskCases()
        {
            var json = SpecLoader.ReadSpecFile("rut_mask.json");
            var suite = JsonSerializer.Deserialize<MaskSuite>(json)!;
            foreach (var tc in suite.Cases)
            {
                yield return new object[]
                {
                    tc.Description,
                    tc.Input,
                    tc.ExpectedDefault,
                    tc.ExpectedPlain,
                };
            }
        }

        [Theory]
        [MemberData(nameof(FormatCases))]
        public void Format_MatchesCanonicalSpec(
            string description,
            string input,
            string expectedDefault,
            string expectedNoDots,
            string expectedNoHyphen,
            string expectedPlain)
        {
            _ = description;
            Assert.Equal(expectedDefault, Rut.Format(input));
            Assert.Equal(expectedNoDots, Rut.Format(input, new RutFormatOptions { Dots = false, Hyphen = true }));
            Assert.Equal(expectedNoHyphen, Rut.Format(input, new RutFormatOptions { Dots = true, Hyphen = false }));
            Assert.Equal(expectedPlain, Rut.Format(input, new RutFormatOptions { Dots = false, Hyphen = false }));
        }

        [Theory]
        [MemberData(nameof(SplitCases))]
        public void Split_MatchesCanonicalSpec(
            string description,
            string input,
            string expectedBody,
            string expectedDv,
            int expectedInt,
            bool isError)
        {
            _ = description;
            if (isError)
            {
                Assert.Throws<ArgumentException>(() => Rut.Split(input));
            }
            else
            {
                var parts = Rut.Split(input);
                Assert.Equal(expectedBody, parts.Body);
                Assert.Equal(expectedDv, parts.Dv);
                if (expectedInt >= 0)
                {
                    Assert.Equal(expectedInt, parts.BodyInt());
                }
            }
        }

        [Fact]
        public void Split_BodyInt_ThrowsOnNonNumericBody()
        {
            var parts = Rut.Split("12A34-K");
            Assert.Throws<FormatException>(() => parts.BodyInt());
        }

        [Theory]
        [MemberData(nameof(MaskCases))]
        public void Mask_MatchesCanonicalSpec(
            string description,
            string input,
            string expectedDefault,
            string expectedPlain)
        {
            _ = description;
            Assert.Equal(expectedDefault, Rut.Mask(input));
            Assert.Equal(expectedPlain, Rut.Mask(input, new RutMaskOptions { Formatted = false }));
        }

        [Fact]
        public void Mask_CustomOptionsAndNegativeVisibleCount()
        {
            Assert.Equal("XX.XXX.458-K", Rut.Mask("18123458K", new RutMaskOptions { MaskChar = 'X', VisibleCount = 4, Formatted = true }));
            Assert.Equal("**.***.**8-K", Rut.Mask("18123458K", new RutMaskOptions { VisibleCount = 2, Formatted = true }));
            Assert.Equal("*********", Rut.Mask("18123458K", new RutMaskOptions { VisibleCount = -5, Formatted = false }));
            Assert.Equal("**.***.***-*", Rut.Mask("18123458K", new RutMaskOptions { VisibleCount = -5, Formatted = true }));
        }
    }
}
