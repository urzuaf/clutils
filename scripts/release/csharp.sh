#!/usr/bin/env bash
set -euo pipefail

# Release independiente para el paquete C# / NuGet (Clutils) vía Trusted Publishing
# Uso:
#   bash scripts/release/csharp.sh 1.0.1
#   bash scripts/release/csharp.sh 1.0.1 --dry-run

ROOT_DIR="$(cd "$(dirname "$0")/../.." && pwd)"
CS_DIR="$ROOT_DIR/packages/csharp"
CSPROJ="$CS_DIR/src/Clutils/Clutils.csproj"

if [ $# -lt 1 ]; then
  echo "Uso: bash scripts/release/csharp.sh <version> [--dry-run]"
  echo "Ejemplo: bash scripts/release/csharp.sh 1.0.1"
  exit 1
fi

RAW_VERSION="$1"
VERSION="${RAW_VERSION#v}"
DRY_RUN=false
if [ "${2:-}" = "--dry-run" ]; then
  DRY_RUN=true
fi

if ! [[ "$VERSION" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  echo "❌ Versión inválida: '$RAW_VERSION'. Debe seguir formato SemVer X.Y.Z (ej. 1.0.1)."
  exit 1
fi

TAG="packages/csharp/v$VERSION"

echo "=================================================="
echo "  clutils [C#/NuGet] release: $TAG (dry-run: $DRY_RUN)"
echo "=================================================="

cd "$ROOT_DIR"

CURRENT_BRANCH="$(git rev-parse --abbrev-ref HEAD)"
if [ "$CURRENT_BRANCH" != "main" ]; then
  if [ "$DRY_RUN" = true ]; then
    echo "⚠️  [dry-run] Estás en rama '$CURRENT_BRANCH' (para release real debe ser 'main')."
  else
    echo "❌ Debes estar en la rama 'main' para publicar (actual: '$CURRENT_BRANCH')."
    exit 1
  fi
fi

if [ -n "$(git status --porcelain)" ]; then
  if [ "$DRY_RUN" = true ]; then
    echo "⚠️  [dry-run] Hay cambios sin commitear en el repositorio."
  else
    echo "❌ El repositorio tiene cambios sin commitear. Haz commit antes de publicar."
    git status -s
    exit 1
  fi
fi

if git rev-parse "$TAG" >/dev/null 2>&1; then
  echo "❌ El tag '$TAG' ya existe localmente."
  exit 1
fi

CSPROJ_VERSION="$(sed -n 's/.*<Version>\(.*\)<\/Version>.*/\1/p' "$CSPROJ" | head -n 1)"
NEED_BUMP=false
if [ "$CSPROJ_VERSION" != "$VERSION" ]; then
  NEED_BUMP=true
  echo "ℹ️  Clutils.csproj está en $CSPROJ_VERSION (se actualizará a $VERSION)."
else
  echo "✅ Clutils.csproj ya está en $VERSION."
fi

echo ""
echo "🧪 Validando paquete C# (dotnet test + dotnet pack)..."
TMP_PACK="$(mktemp -d /tmp/clutils-nupkg-XXXXXX)"
trap 'rm -rf "$TMP_PACK"' EXIT
(cd "$CS_DIR" && dotnet test Clutils.sln --configuration Release && dotnet pack src/Clutils/Clutils.csproj --configuration Release --no-build --output "$TMP_PACK")

if [ "$DRY_RUN" = true ]; then
  echo ""
  echo "🔎 [dry-run] En un release real se ejecutaría:"
  if [ "$NEED_BUMP" = true ]; then
    echo "   - Actualizar <Version> en packages/csharp/src/Clutils/Clutils.csproj a $VERSION"
    echo "   - git add packages/csharp/src/Clutils/Clutils.csproj"
    echo "   - git commit -m \"chore(csharp): release v$VERSION\""
    echo "   - git push origin main"
  fi
  echo "   - git tag -a \"$TAG\" -m \"Release C# NuGet package $TAG\""
  echo "   - git push origin \"$TAG\"  (dispara .github/workflows/publish-nuget.yml vía OIDC)"
  echo "✅ Dry-run [C#/NuGet] completado sin modificaciones."
  exit 0
fi

echo ""
echo "Acciones irreversibles:"
if [ "$NEED_BUMP" = true ]; then
  echo "  • Actualizar <Version> en Clutils.csproj a $VERSION y hacer commit + push a main"
fi
echo "  • Crear y pushear tag: $TAG (dispara publicación automática en NuGet.org vía OIDC)"
echo ""
read -r -p "¿Continuar con el release C# v$VERSION? [y/N]: " CONFIRM
if [[ ! "$CONFIRM" =~ ^[Yy]$ ]]; then
  echo "Abortado por el usuario."
  exit 0
fi

if [ "$NEED_BUMP" = true ]; then
  python3 -c "
p = '$CSPROJ'
s = open(p, encoding='utf-8').read()
import re
s = re.sub(r'<Version>[^<]+</Version>', '<Version>$VERSION</Version>', s, count=1)
open(p, 'w', encoding='utf-8').write(s)
"
  git add packages/csharp/src/Clutils/Clutils.csproj
  git commit -m "chore(csharp): release v$VERSION"
  git push origin main
fi

git tag -a "$TAG" -m "Release C# NuGet package $TAG"
git push origin "$TAG"

echo ""
echo "🚀 ¡Tag $TAG subido! GitHub Actions publicará Clutils@$VERSION en NuGet.org."
echo "   • Actions: https://github.com/urzuaf/clutils/actions"
echo "   • NuGet:   https://www.nuget.org/packages/Clutils"
