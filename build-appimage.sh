#!/bin/bash
cd "$(dirname "$0")" || exit 1
set -euo pipefail
source packaging/common.sh

build_binary

echo "==> Montando AppDir (AppImage)..."
rm -rf AppDir
mkdir -p AppDir/usr/bin
mkdir -p AppDir/usr/share/applications
mkdir -p AppDir/usr/share/icons/hicolor/256x256/apps
cp "$NAME" AppDir/usr/bin/
cp FyneApp.toml AppDir/usr/bin/FyneApp.toml
cp "$DESKTOP" AppDir/protoncommand.desktop
cp "$DESKTOP" AppDir/usr/share/applications/protoncommand.desktop
cp "$ICON" AppDir/protoncommand.png
cp "$ICON" AppDir/usr/share/icons/hicolor/256x256/apps/protoncommand.png
cat > AppDir/AppRun <<'EOF'
#!/bin/sh
SELF=$(readlink -f "$0")
HERE=${SELF%/*}
exec "$HERE/usr/bin/protoncommand" "$@"
EOF
chmod +x AppDir/AppRun

if command -v appimagetool >/dev/null 2>&1; then
  TOOL="appimagetool"
elif [ -x ./appimagetool ]; then
  TOOL="./appimagetool"
else
  echo "appimagetool não encontrado. Baixando..."
  curl -L -o appimagetool \
    https://github.com/AppImage/appimagetool/releases/download/continuous/appimagetool-x86_64.AppImage
  chmod +x appimagetool
  TOOL="./appimagetool"
fi

mkdir -p "$DIST"
echo "==> Gerando AppImage..."
APPIMAGE_EXTRACT_AND_RUN=1 "$TOOL" AppDir "$DIST/${ARTIFACT}.AppImage"
