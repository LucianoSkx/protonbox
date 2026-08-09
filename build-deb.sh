#!/bin/bash
cd "$(dirname "$0")" || exit 1
set -euo pipefail
source packaging/common.sh

build_binary

echo "==> Gerando .deb..."
if ! command -v dpkg-deb >/dev/null 2>&1; then
  echo "dpkg-deb não encontrado. Instale com: sudo apt install dpkg-dev (Debian/Ubuntu) ou sudo pacman -S dpkg (Arch)" >&2
  exit 1
fi

DEB_DIR="$DIST/deb"
rm -rf "$DIST/deb"
mkdir -p "$DEB_DIR/DEBIAN"
mkdir -p "$DEB_DIR/usr/bin"
mkdir -p "$DEB_DIR/usr/share/applications"
mkdir -p "$DEB_DIR/usr/share/icons/hicolor/256x256/apps"
cp "$NAME" "$DEB_DIR/usr/bin/"
cp "$DESKTOP" "$DEB_DIR/usr/share/applications/"
cp "$ICON" "$DEB_DIR/usr/share/icons/hicolor/256x256/apps/$NAME.png"
cat > "$DEB_DIR/DEBIAN/control" <<EOF
Package: $NAME
Version: $VERSION
Section: games
Priority: optional
Architecture: amd64
Maintainer: Luciano Oliveira <lucianoskx@gmail.com>
Depends: libc6, libgl1
Description: Useful Proton launch commands manager
 A simple GUI to browse, copy and combine useful Proton launch
 commands (standard Proton, Proton-GE and Proton-CachyOS) for Steam.
EOF
dpkg-deb --build --root-owner-group "$DEB_DIR" "$DIST/${NAME}_${VERSION}_amd64.deb"
echo "==> .deb pronto: $DIST/${NAME}_${VERSION}_amd64.deb"
