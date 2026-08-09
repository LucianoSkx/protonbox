#!/bin/bash
cd "$(dirname "$0")" || exit 1
set -euo pipefail
source packaging/common.sh

build_binary
echo "==> Binário pronto: ./$NAME"
