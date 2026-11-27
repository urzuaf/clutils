#!/usr/bin/env bash
set -euo pipefail

# Release independiente para el módulo Go (github.com/urzuaf/clutils/packages/go)
# Uso:
#   bash scripts/release/go.sh 1.0.1
#   bash scripts/release/go.sh 1.0.1 --dry-run

ROOT_DIR="$(cd "$(dirname "$0")/../.." && pwd)"
GO_DIR="$ROOT_DIR/packages/go"

if [ $# -lt 1 ]; then
  echo "Uso: bash scripts/release/go.sh <version> [--dry-run]"
  echo "Ejemplo: bash scripts/release/go.sh 1.0.1"
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

TAG="packages/go/v$VERSION"

echo "=================================================="
echo "  clutils [Go] release: $TAG (dry-run: $DRY_RUN)"
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

echo ""
echo "🧪 Validando módulo Go (go vet + go test -race)..."
(cd "$GO_DIR" && go vet ./... && go test -race ./...)

if [ "$DRY_RUN" = true ]; then
  echo ""
  echo "🔎 [dry-run] En un release real se ejecutaría:"
  echo "   - git tag -a \"$TAG\" -m \"Release Go module $TAG\""
  echo "   - git push origin \"$TAG\""
  echo "✅ Dry-run [Go] completado sin modificaciones."
  exit 0
fi

echo ""
echo "Acciones irreversibles:"
echo "  • Crear y pushear tag: $TAG"
echo ""
read -r -p "¿Continuar con el release Go v$VERSION? [y/N]: " CONFIRM
if [[ ! "$CONFIRM" =~ ^[Yy]$ ]]; then
  echo "Abortado por el usuario."
  exit 0
fi

git tag -a "$TAG" -m "Release Go module $TAG"
git push origin "$TAG"

echo ""
echo "🚀 ¡Release Go v$VERSION publicado!"
echo "   • pkg.go.dev: https://pkg.go.dev/github.com/urzuaf/clutils/packages/go@v$VERSION"
