#!/bin/bash
set -euo pipefail

NAME="protoncommand"
VERSION="0.1.0"
DESKTOP="packaging/protoncommand.desktop"
ICON="assets/icon.png"
DIST="dist"

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

build_binary() {
  echo "==> Compilando binário..."
  go build -trimpath -ldflags "-s -w" -o "$NAME" .
}

check_binary() {
  if [ ! -x "$NAME" ]; then
    echo "Binário $NAME não encontrado. Execute build-binary.sh antes." >&2
    exit 1
  fi
}
