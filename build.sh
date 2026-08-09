#!/bin/bash
set -euo pipefail

NAME="protonbox"
VERSION="0.1.0"
DESKTOP="packaging/ProtonBox.desktop"
ICON="assets/icon.png"
DIST="dist"

echo "==> Compilando binário..."
go build -trimpath -ldflags "-s -w" -o "$NAME" .

echo "==> Montando AppDir (AppImage)..."
rm -rf AppDir "$DIST"
mkdir -p "$DIST"
mkdir -p AppDir/usr/bin
mkdir -p AppDir/usr/share/applications
mkdir -p AppDir/usr/share/icons/hicolor/256x256/apps
cp "$NAME" AppDir/usr/bin/
cp "$DESKTOP" AppDir/ProtonBox.desktop
cp "$DESKTOP" AppDir/usr/share/applications/ProtonBox.desktop
cp "$ICON" AppDir/protonbox.png
cp "$ICON" AppDir/usr/share/icons/hicolor/256x256/apps/protonbox.png
cat > AppDir/AppRun <<'EOF'
#!/bin/sh
SELF=$(readlink -f "$0")
HERE=${SELF%/*}
exec "$HERE/usr/bin/protonbox" "$@"
EOF
chmod +x AppDir/AppRun
cp "$ICON" AppDir/usr/share/icons/hicolor/256x256/apps/"$NAME".png

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

echo "==> Gerando AppImage..."
APPIMAGE_EXTRACT_AND_RUN=1 "$TOOL" AppDir "$DIST/${NAME}-x86_64.AppImage"

echo "==> Gerando .deb..."
DEB_DIR="$DIST/deb"
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

echo "==> Gerando .rpm..."
RPM_DIR="build/rpm"
mkdir -p "$RPM_DIR"/{BUILD,RPMS,SOURCES,SPECS,SRPMS}
SRCTREE="$RPM_DIR/srctree/${NAME}-${VERSION}"
rm -rf "$RPM_DIR/srctree"
mkdir -p "$SRCTREE"
cp "$NAME" "$SRCTREE/"
cp "$DESKTOP" "$SRCTREE/ProtonBox.desktop"
cp "$ICON" "$SRCTREE/icon.png"
tar -C "$RPM_DIR/srctree" -czf "$RPM_DIR/SOURCES/${NAME}-${VERSION}.tar.gz" "${NAME}-${VERSION}"
cp packaging/protonbox.spec "$RPM_DIR/SPECS/"
rpmbuild --define "_topdir $(pwd)/$RPM_DIR" -bb "$RPM_DIR/SPECS/protonbox.spec"
cp "$RPM_DIR"/RPMS/x86_64/*.rpm "$DIST/"

echo ""
echo "==> Artefatos em $DIST/:"
ls -lh "$DIST"
