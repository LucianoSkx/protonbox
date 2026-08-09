#!/bin/bash
cd "$(dirname "$0")" || exit 1
set -u

echo "==> build.sh: gerando todos os formatos"
FAILED=""
./build-binary.sh || FAILED="$FAILED binario"
./build-appimage.sh || FAILED="$FAILED appimage"
./build-deb.sh || FAILED="$FAILED deb"
./build-rpm.sh || FAILED="$FAILED rpm"

if [ -n "$FAILED" ]; then
  echo ""
  echo "Atenção: não foi possível gerar:$FAILED"
  echo "Rode o script individual (ex.: ./build-deb.sh) para ver o motivo."
  echo "Você também pode gerar apenas um formato:"
  echo "  ./build-binary.sh   ./build-appimage.sh   ./build-deb.sh   ./build-rpm.sh"
fi

echo ""
echo "==> Artefatos em dist/:"
ls -lh dist/ 2>/dev/null || echo "(dist/ vazio)"
