#!/usr/bin/env bash
set -euo pipefail

# Script de release para clutils (Go + npm)
# Uso:
#   bash scripts/release.sh 1.0.0
#   bash scripts/release.sh 1.0.0 --dry-run

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
GO_DIR="$ROOT_DIR/packages/go"
TS_DIR="$ROOT_DIR/packages/ts"

if [ $# -lt 1 ]; then
  echo "Uso: bash scripts/release.sh <version> [--dry-run]"
  echo "Ejemplo: bash scripts/release.sh 1.0.0"
  exit 1
fi

RAW_VERSION="$1"
VERSION="${RAW_VERSION#v}"
DRY_RUN=false

if [ "${2:-}" = "--dry-run" ]; then
  DRY_RUN=true
fi

# 1. Validar formato SemVer (X.Y.Z)
if ! [[ "$VERSION" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  echo "❌ Versión inválida: '$RAW_VERSION'. Debe seguir formato SemVer X.Y.Z (ej. 1.0.0)."
  exit 1
fi

TAG_ROOT="v$VERSION"
TAG_GO="packages/go/v$VERSION"

echo "=================================================="
echo "  clutils release: v$VERSION (dry-run: $DRY_RUN)"
echo "=================================================="

cd "$ROOT_DIR"

# 2. Verificar rama main
CURRENT_BRANCH="$(git rev-parse --abbrev-ref HEAD)"
if [ "$CURRENT_BRANCH" != "main" ]; then
  if [ "$DRY_RUN" = true ]; then
    echo "⚠️  [dry-run] Estás en rama '$CURRENT_BRANCH' (para release real debe ser 'main')."
  else
    echo "❌ Debes estar en la rama 'main' para publicar (actual: '$CURRENT_BRANCH')."
    exit 1
  fi
fi

# 3. Verificar working tree limpio
if [ -n "$(git status --porcelain)" ]; then
  if [ "$DRY_RUN" = true ]; then
    echo "⚠️  [dry-run] Hay cambios sin commitear en el repositorio (en release real debes commitear antes)."
  else
    echo "❌ El repositorio tiene cambios sin commitear. Haz commit antes de correr el release."
    git status -s
    exit 1
  fi
fi

# 4. Verificar que los tags no existan ya
if git rev-parse "$TAG_GO" >/dev/null 2>&1 || git rev-parse "$TAG_ROOT" >/dev/null 2>&1; then
  echo "❌ El tag '$TAG_GO' o '$TAG_ROOT' ya existe localmente."
  exit 1
fi

# 5. Verificar versión en packages/ts/package.json
PKG_VERSION="$(node -p "require('$TS_DIR/package.json').version")"
NEED_BUMP=false
if [ "$PKG_VERSION" != "$VERSION" ]; then
  NEED_BUMP=true
  echo "ℹ️  packages/ts/package.json está en $PKG_VERSION (se actualizará a $VERSION)."
else
  echo "✅ packages/ts/package.json ya está en $VERSION."
fi

# 6. Ejecutar validaciones completas (Go + TypeScript)
echo ""
echo "🧪 [1/2] Validando módulo Go (go vet + go test -race)..."
(cd "$GO_DIR" && go vet ./... && go test -race ./...)

echo ""
echo "🧪 [2/2] Validando paquete TypeScript (tsc + vitest + build)..."
(cd "$TS_DIR" && npx tsc --noEmit && npm test && npm run build)

# 7. Modo dry-run: mostrar pasos y salir sin tocar git ni npm
if [ "$DRY_RUN" = true ]; then
  echo ""
  echo "🔎 [dry-run] Todo pasó correctamente. En un release real se ejecutaría:"
  if [ "$NEED_BUMP" = true ]; then
    echo "   - cd packages/ts && npm version $VERSION --no-git-tag-version"
    echo "   - git add packages/ts/package.json packages/ts/package-lock.json"
    echo "   - git commit -m \"chore(release): prepare v$VERSION\""
    echo "   - git push origin main"
  fi
  echo "   - git tag -a \"$TAG_GO\" -m \"Release Go module $TAG_GO\""
  echo "   - git tag -a \"$TAG_ROOT\" -m \"Release $TAG_ROOT\""
  echo "   - git push origin \"$TAG_GO\" \"$TAG_ROOT\""
  echo "   - cd packages/ts && npm publish --access public"
  echo ""
  echo "✅ Dry-run completado sin modificaciones."
  exit 0
fi

# 8. Confirmación interactiva
echo ""
echo "Se realizarán las siguientes acciones irreversibles:"
if [ "$NEED_BUMP" = true ]; then
  echo "  • Actualizar packages/ts/package.json a $VERSION y hacer commit + push a main"
fi
echo "  • Crear y pushear tags: $TAG_GO y $TAG_ROOT"
echo "  • Publicar @urzuaf/clutils@$VERSION en npm"
echo ""
read -r -p "¿Continuar con la publicación de v$VERSION? [y/N]: " CONFIRM
if [[ ! "$CONFIRM" =~ ^[Yy]$ ]]; then
  echo "Abortado por el usuario."
  exit 0
fi

# 9. Bump de versión si corresponde
if [ "$NEED_BUMP" = true ]; then
  (cd "$TS_DIR" && npm version "$VERSION" --no-git-tag-version && npm run build)
  git add packages/ts/package.json packages/ts/package-lock.json
  git commit -m "chore(release): prepare v$VERSION"
  git push origin main
fi

# 10. Crear y pushear tags (Go + raíz)
git tag -a "$TAG_GO" -m "Release Go module $TAG_GO"
git tag -a "$TAG_ROOT" -m "Release $TAG_ROOT"
git push origin "$TAG_GO" "$TAG_ROOT"

# 11. Publicar en npm (soporta 2FA / OTP interactivo)
echo ""
echo "📦 Publicando @urzuaf/clutils@$VERSION en npm..."
(cd "$TS_DIR" && npm publish --access public)

echo ""
echo "¡Release v$VERSION publicado con éxito!"
echo "   • npm:        https://www.npmjs.com/package/@urzuaf/clutils"
echo "   • pkg.go.dev: https://pkg.go.dev/github.com/urzuaf/clutils/packages/go@$TAG_ROOT"
echo "   (Abre el link de pkg.go.dev para disparar la indexación inmediata)"
