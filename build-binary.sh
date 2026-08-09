#!/bin/bash
set -euo pipefail
source "$(dirname "$0")/packaging/common.sh"

build_binary
echo "==> Binário pronto: ./$NAME"
