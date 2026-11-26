using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Serialization;
using Xunit;

namespace Clutils.Tests.GeoTests
{
    public class DpaTests
    {
        private sealed class RegionCountCase
        {
            [JsonPropertyName("region_code")]
            public string RegionCode { get; set; } = string.Empty;

            [JsonPropertyName("iso_code")]
            public string IsoCode { get; set; } = string.Empty;

            [JsonPropertyName("name")]
            public string Name { get; set; } = string.Empty;

            [JsonPropertyName("expected_communes")]
            public int ExpectedCommunes { get; set; }

            [JsonPropertyName("expected_provinces")]
            public int ExpectedProvinces { get; set; }
        }

        private sealed class CountsSpec
        {
            [JsonPropertyName("expected_regions")]
            public int ExpectedRegions { get; set; }

            [JsonPropertyName("expected_provinces")]
            public int ExpectedProvinces { get; set; }

            [JsonPropertyName("expected_communes")]
            public int ExpectedCommunes { get; set; }

            [JsonPropertyName("region_commune_counts")]
            public List<RegionCountCase> RegionCommuneCounts { get; set; } = new();
        }

        private sealed class FilteringCase
        {
            [JsonPropertyName("description")]
            public string Description { get; set; } = string.Empty;

            [JsonPropertyName("region_query")]
            public string? RegionQuery { get; set; }

            [JsonPropertyName("expected_count")]
            public int? ExpectedCount { get; set; }

            [JsonPropertyName("sample_commune_cuts")]
            public List<string>? SampleCommuneCuts { get; set; }

            [JsonPropertyName("provinces_region_query")]
            public string? ProvincesRegionQuery { get; set; }

            [JsonPropertyName("expected_provinces_count")]
            public int? ExpectedProvincesCount { get; set; }

            [JsonPropertyName("expected_province_names")]
            public List<string>? ExpectedProvinceNames { get; set; }
        }

        private sealed class FindComunaCase
        {
            [JsonPropertyName("query")]
            public string Query { get; set; } = string.Empty;

            [JsonPropertyName("expected_cut")]
            public string? ExpectedCut { get; set; }

            [JsonPropertyName("expected_name")]
            public string? ExpectedName { get; set; }

            [JsonPropertyName("expected_region_code")]
            public string? ExpectedRegionCode { get; set; }

            [JsonPropertyName("description")]
            public string Description { get; set; } = string.Empty;
        }

        private sealed class FindRegionCase
        {
            [JsonPropertyName("query")]
            public string Query { get; set; } = string.Empty;

            [JsonPropertyName("expected_code")]
            public string? ExpectedCode { get; set; }

            [JsonPropertyName("expected_iso")]
            public string? ExpectedIso { get; set; }

            [JsonPropertyName("expected_name")]
            public string? ExpectedName { get; set; }

            [JsonPropertyName("description")]
            public string Description { get; set; } = string.Empty;
        }

        private sealed class ReverseLookupCase
        {
            [JsonPropertyName("comuna_query")]
            public string ComunaQuery { get; set; } = string.Empty;

            [JsonPropertyName("expected_region_code")]
            public string? ExpectedRegionCode { get; set; }

            [JsonPropertyName("expected_region_name")]
            public string? ExpectedRegionName { get; set; }

            [JsonPropertyName("description")]
            public string Description { get; set; } = string.Empty;
        }

        private sealed class GeoSuite
        {
            [JsonPropertyName("counts")]
            public CountsSpec Counts { get; set; } = new();

            [JsonPropertyName("filtering")]
            public List<FilteringCase> Filtering { get; set; } = new();

            [JsonPropertyName("find_comuna")]
            public List<FindComunaCase> FindComuna { get; set; } = new();

            [JsonPropertyName("find_region")]
            public List<FindRegionCase> FindRegion { get; set; } = new();

            [JsonPropertyName("get_region_by_comuna")]
            public List<ReverseLookupCase> GetRegionByComuna { get; set; } = new();
        }

        private static GeoSuite LoadSuite()
        {
            var json = SpecLoader.ReadSpecFile("geo_dpa.json");
            return JsonSerializer.Deserialize<GeoSuite>(json)!;
        }

        public static IEnumerable<object[]> FindComunaCases()
        {
            foreach (var tc in LoadSuite().FindComuna)
            {
                yield return new object[]
                {
                    tc.Description,
                    tc.Query,
                    tc.ExpectedCut ?? string.Empty,
                    tc.ExpectedName ?? string.Empty,
                    tc.ExpectedRegionCode ?? string.Empty,
                };
            }
        }

        public static IEnumerable<object[]> FindRegionCases()
        {
            foreach (var tc in LoadSuite().FindRegion)
            {
                yield return new object[]
                {
                    tc.Description,
                    tc.Query,
                    tc.ExpectedCode ?? string.Empty,
                    tc.ExpectedIso ?? string.Empty,
                    tc.ExpectedName ?? string.Empty,
                };
            }
        }

        public static IEnumerable<object[]> ReverseLookupCases()
        {
            foreach (var tc in LoadSuite().GetRegionByComuna)
            {
                yield return new object[]
                {
                    tc.Description,
                    tc.ComunaQuery,
                    tc.ExpectedRegionCode ?? string.Empty,
                    tc.ExpectedRegionName ?? string.Empty,
                };
            }
        }

        [Fact]
        public void CanonicalCountsAndDistribution_MatchSpec()
        {
            var suite = LoadSuite();
            var regiones = Geo.GetRegiones();
            var provincias = Geo.GetProvincias();
            var comunas = Geo.GetComunas();

            Assert.Equal(suite.Counts.ExpectedRegions, regiones.Count);
            Assert.Equal(1, regiones[0].Order);
            Assert.Equal(16, regiones[15].Order);

            Assert.Equal(suite.Counts.ExpectedProvinces, provincias.Count);
            Assert.Equal(suite.Counts.ExpectedCommunes, comunas.Count);

            foreach (var r in suite.Counts.RegionCommuneCounts)
            {
                Assert.Equal(r.ExpectedCommunes, Geo.GetComunasByRegion(r.RegionCode).Count);
                Assert.Equal(r.ExpectedProvinces, Geo.GetProvinciasByRegion(r.RegionCode).Count);
            }
        }

        [Fact]
        public void HierarchicalFiltering_MatchesSpec()
        {
            var suite = LoadSuite();
            foreach (var tc in suite.Filtering)
            {
                if (tc.RegionQuery != null)
                {
                    var comunas = Geo.GetComunasByRegion(tc.RegionQuery);
                    Assert.Equal(tc.ExpectedCount!.Value, comunas.Count);
                    if (tc.SampleCommuneCuts != null)
                    {
                        var cuts = comunas.Select(c => c.Cut).ToHashSet();
                        foreach (var cut in tc.SampleCommuneCuts)
                        {
                            Assert.Contains(cut, cuts);
                        }
                    }
                }

                if (tc.ProvincesRegionQuery != null)
                {
                    var provs = Geo.GetProvinciasByRegion(tc.ProvincesRegionQuery);
                    Assert.Equal(tc.ExpectedProvincesCount!.Value, provs.Count);
                    if (tc.ExpectedProvinceNames != null)
                    {
                        var names = provs.Select(p => p.Name).OrderBy(n => n).ToList();
                        var expected = tc.ExpectedProvinceNames.OrderBy(n => n).ToList();
                        Assert.Equal(expected, names);
                    }
                }
            }

            Assert.Equal(32, Geo.GetComunasByProvincia("131").Count);
            Assert.Equal(9, Geo.GetComunasByProvincia("diguillin").Count);
            Assert.Empty(Geo.GetComunasByProvincia("999"));
        }

        [Theory]
        [MemberData(nameof(FindComunaCases))]
        public void FindComuna_MatchesCanonicalSpec(
            string description,
            string query,
            string expectedCut,
            string expectedName,
            string expectedRegionCode)
        {
            _ = description;
            var result = Geo.FindComuna(query);
            if (string.IsNullOrEmpty(expectedCut))
            {
                Assert.Null(result);
            }
            else
            {
                Assert.NotNull(result);
                Assert.Equal(expectedCut, result!.Cut);
                Assert.Equal(expectedName, result.Name);
                Assert.Equal(expectedRegionCode, result.RegionCode);
            }
        }

        [Theory]
        [MemberData(nameof(FindRegionCases))]
        public void FindRegion_MatchesCanonicalSpec(
            string description,
            string query,
            string expectedCode,
            string expectedIso,
            string expectedName)
        {
            _ = description;
            var result = Geo.FindRegion(query);
            if (string.IsNullOrEmpty(expectedCode))
            {
                Assert.Null(result);
            }
            else
            {
                Assert.NotNull(result);
                Assert.Equal(expectedCode, result!.Code);
                Assert.Equal(expectedIso, result.IsoCode);
                Assert.Equal(expectedName, result.Name);
            }
        }

        [Theory]
        [MemberData(nameof(ReverseLookupCases))]
        public void GetRegionByComuna_MatchesCanonicalSpec(
            string description,
            string comunaQuery,
            string expectedRegionCode,
            string expectedRegionName)
        {
            _ = description;
            var result = Geo.GetRegionByComuna(comunaQuery);
            if (string.IsNullOrEmpty(expectedRegionCode))
            {
                Assert.Null(result);
            }
            else
            {
                Assert.NotNull(result);
                Assert.Equal(expectedRegionCode, result!.Code);
                Assert.Equal(expectedRegionName, result.Name);
            }
        }

        [Fact]
        public void ProvinciasAndSearchAndQARegressions_WorkCorrectly()
        {
            // FindProvincia & GetRegionByProvincia
            var santiago = Geo.FindProvincia("131");
            Assert.NotNull(santiago);
            Assert.Equal("Santiago", santiago!.Name);
            Assert.Equal("13", santiago.RegionCode);

            var diguillin = Geo.FindProvincia("diguillin");
            Assert.NotNull(diguillin);
            Assert.Equal("Diguillín", diguillin!.Name);

            Assert.Null(Geo.FindProvincia("narnia"));
            Assert.Null(Geo.FindProvincia(null));

            var regNuble = Geo.GetRegionByProvincia("161");
            Assert.NotNull(regNuble);
            Assert.Equal("Ñuble", regNuble!.Name);
            Assert.Null(Geo.GetRegionByProvincia("inexistente"));

            // GetProvinciaByComuna
            var provNunoa = Geo.GetProvinciaByComuna("Ñuñoa");
            Assert.NotNull(provNunoa);
            Assert.Equal("131", provNunoa!.Code);
            Assert.Null(Geo.GetProvinciaByComuna("inexistente"));

            // SearchProvincias & SearchComunas
            var provMatches = Geo.SearchProvincias("San");
            Assert.Contains(provMatches, p => p.Name == "San Antonio");
            Assert.Contains(provMatches, p => p.Name == "Santiago");
            Assert.Empty(Geo.SearchProvincias("xyz999"));

            var comMatches = Geo.SearchComunas("San");
            Assert.True(comMatches.Count > 10);
            Assert.Contains(comMatches, c => c.Name == "San Bernardo");
            Assert.Contains(comMatches, c => c.Name == "San Miguel");
            Assert.Empty(Geo.SearchComunas("---"));
            Assert.Empty(Geo.SearchComunas(null));

            // Regresiones QA: Unicode NFD + coherencia Find == Search[0] + padding dígito simple
            Assert.Equal("13120", Geo.FindComuna("N\u0303un\u0303oa")?.Cut);
            var foundMa = Geo.FindComuna("ma");
            var searchedMa = Geo.SearchComunas("ma");
            Assert.NotNull(foundMa);
            Assert.Equal(searchedMa[0].Cut, foundMa!.Cut);
            Assert.Equal("María Elena", foundMa.Name);

            Assert.Equal("01", Geo.FindRegion("1")?.Code);
        }
    }
}
