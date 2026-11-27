#!/usr/bin/env bash
set -euo pipefail

# Release independiente para el paquete TypeScript / npm (@urzuaf/clutils)
# Uso:
#   bash scripts/release/ts.sh 1.0.1
#   bash scripts/release/ts.sh 1.0.1 --dry-run

ROOT_DIR="$(cd "$(dirname "$0")/../.." && pwd)"
TS_DIR="$ROOT_DIR/packages/ts"

if [ $# -lt 1 ]; then
  echo "Uso: bash scripts/release/ts.sh <version> [--dry-run]"
  echo "Ejemplo: bash scripts/release/ts.sh 1.0.1"
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

TAG="packages/ts/v$VERSION"

echo "=================================================="
echo "  clutils [TS/npm] release: v$VERSION (dry-run: $DRY_RUN)"
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

# Verificar sesión activa de npm ANTES de tocar commits/tags
if [ "$DRY_RUN" = false ]; then
  echo ""
  echo "🔐 Verificando autenticación en npm..."
  if NPM_USER="$(npm whoami 2>/dev/null)"; then
    echo "✅ Autenticado en npm como: $NPM_USER"
  else
    echo "⚠️  No hay sesión activa en npm (o expiró). Iniciando 'npm login'..."
    npm login
    NPM_USER="$(npm whoami)"
    echo "✅ Autenticado en npm como: $NPM_USER"
  fi
fi

PKG_VERSION="$(node -p "require('$TS_DIR/package.json').version")"
NEED_BUMP=false
if [ "$PKG_VERSION" != "$VERSION" ]; then
  NEED_BUMP=true
  echo "ℹ️  packages/ts/package.json está en $PKG_VERSION (se actualizará a $VERSION)."
else
  echo "✅ packages/ts/package.json ya está en $VERSION."
fi

echo ""
echo "🧪 Validando paquete TypeScript (tsc + vitest + build)..."
(cd "$TS_DIR" && npx tsc --noEmit && npm test && npm run build)

if [ "$DRY_RUN" = true ]; then
  echo ""
  echo "🔎 [dry-run] En un release real se ejecutaría:"
  echo "   - npm whoami (o npm login si expiró)"
  if [ "$NEED_BUMP" = true ]; then
    echo "   - cd packages/ts && npm version $VERSION --no-git-tag-version && npm run build"
    echo "   - git add packages/ts/package.json packages/ts/package-lock.json"
    echo "   - git commit -m \"chore(ts): release v$VERSION\""
    echo "   - git push origin main"
  fi
  echo "   - cd packages/ts && npm publish --access public"
  echo "   - git tag -a \"$TAG\" -m \"Release TypeScript package $TAG\""
  echo "   - git push origin \"$TAG\""
  echo "✅ Dry-run [TS/npm] completado sin modificaciones."
  exit 0
fi

echo ""
echo "Acciones irreversibles:"
if [ "$NEED_BUMP" = true ]; then
  echo "  • Actualizar packages/ts/package.json a $VERSION y hacer commit + push a main"
fi
echo "  • Publicar @urzuaf/clutils@$VERSION en npm"
echo "  • Crear y pushear tag: $TAG"
echo ""
read -r -p "¿Continuar con el release npm v$VERSION? [y/N]: " CONFIRM
if [[ ! "$CONFIRM" =~ ^[Yy]$ ]]; then
  echo "Abortado por el usuario."
  exit 0
fi

if [ "$NEED_BUMP" = true ]; then
  (cd "$TS_DIR" && npm version "$VERSION" --no-git-tag-version && npm run build)
  git add packages/ts/package.json packages/ts/package-lock.json
  git commit -m "chore(ts): release v$VERSION"
  git push origin main
fi

echo ""
echo "📦 Publicando @urzuaf/clutils@$VERSION en npm..."
(cd "$TS_DIR" && npm publish --access public)

git tag -a "$TAG" -m "Release TypeScript package $TAG"
git push origin "$TAG"

echo ""
echo "🚀 ¡Release TypeScript v$VERSION publicado en npm!"
echo "   • npm: https://www.npmjs.com/package/@urzuaf/clutils"
