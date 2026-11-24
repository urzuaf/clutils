using Xunit;

namespace Clutils.Tests.RutTests
{
    public class RutGenerationTests
    {
        [Fact]
        public void Generate_ProducesValidRutsInGeneralRange()
        {
            for (int i = 0; i < 100; i++)
            {
                var generated = Rut.Generate();
                Assert.True(Rut.Validate(generated), $"RUT generado inválido: {generated}");

                var body = Rut.Split(generated).BodyInt();
                Assert.InRange(body, RutRanges.DefaultGeneralMin, RutRanges.DefaultGeneralMax);
            }
        }

        [Fact]
        public void GeneratePerson_ProducesValidRutsInPersonRange()
        {
            for (int i = 0; i < 100; i++)
            {
                var generated = Rut.GeneratePerson();
                Assert.True(Rut.Validate(generated));

                var body = Rut.Split(generated).BodyInt();
                Assert.InRange(body, RutRanges.DefaultPersonMin, RutRanges.DefaultPersonMax);
            }
        }

        [Fact]
        public void GenerateBusiness_ProducesValidRutsInBusinessRange()
        {
            for (int i = 0; i < 100; i++)
            {
                var generated = Rut.GenerateBusiness();
                Assert.True(Rut.Validate(generated));

                var body = Rut.Split(generated).BodyInt();
                Assert.InRange(body, RutRanges.DefaultBusinessMin, RutRanges.DefaultBusinessMax);
            }
        }

        [Fact]
        public void GenerateBetween_RespectsRangeAndSwapsInvertedBounds()
        {
            for (int i = 0; i < 50; i++)
            {
                var generated = Rut.GenerateBetween(20_000_000, 10_000_000, new RutFormatOptions { Dots = false, Hyphen = true });
                Assert.True(Rut.Validate(generated));
                Assert.DoesNotContain(".", generated);
                Assert.Contains("-", generated);

                var body = Rut.Split(generated).BodyInt();
                Assert.InRange(body, 10_000_000, 20_000_000);
            }
        }
    }
}
