#!/bin/bash
set -euo pipefail
source "$(dirname "$0")/packaging/common.sh"

build_binary

echo "==> Gerando .rpm..."
if ! command -v rpmbuild >/dev/null 2>&1; then
  echo "rpmbuild não encontrado. Instale com: sudo dnf install rpm-build (Fedora) ou sudo apt install rpm (Debian/Ubuntu)" >&2
  exit 1
fi

RPM_DIR="build/rpm"
SRCTREE="$RPM_DIR/srctree/${NAME}-${VERSION}"
rm -rf "$RPM_DIR"
mkdir -p "$RPM_DIR"/{BUILD,RPMS,SOURCES,SPECS,SRPMS}
mkdir -p "$SRCTREE"
cp "$NAME" "$SRCTREE/"
cp "$DESKTOP" "$SRCTREE/ProtonBox.desktop"
cp "$ICON" "$SRCTREE/icon.png"
tar -C "$RPM_DIR/srctree" -czf "$RPM_DIR/SOURCES/${NAME}-${VERSION}.tar.gz" "${NAME}-${VERSION}"
cp packaging/protonbox.spec "$RPM_DIR/SPECS/"
rpmbuild --define "_topdir $(pwd)/$RPM_DIR" -bb "$RPM_DIR/SPECS/protonbox.spec"
mkdir -p "$DIST"
cp "$RPM_DIR"/RPMS/x86_64/*.rpm "$DIST/"
echo "==> .rpm pronto em $DIST/"
