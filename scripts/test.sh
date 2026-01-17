#!/usr/bin/env bash
set -euo pipefail
source scripts/env.sh

PKG="${1:-./...}"

echo "==> test $PKG"
go test \
  -count=1 \
  -race \
  "$PKG"