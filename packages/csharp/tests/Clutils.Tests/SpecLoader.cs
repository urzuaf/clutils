using System;
using System.IO;

namespace Clutils.Tests
{
    internal static class SpecLoader
    {
        public static string ReadSpecFile(string fileName)
        {
            var dir = new DirectoryInfo(AppContext.BaseDirectory);
            while (dir != null)
            {
                var candidate = Path.Combine(dir.FullName, "specs", "tests", fileName);
                if (File.Exists(candidate))
                {
                    return File.ReadAllText(candidate);
                }
                dir = dir.Parent;
            }

            throw new FileNotFoundException($"No se encontró el archivo de spec: specs/tests/{fileName}");
        }
    }
}
