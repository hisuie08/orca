#!/usr/bin/env bash
set -euo pipefail

PKG="${1:-./...}"

echo "==> watching tests: $PKG"
find . -name "*.go" | entr -c go test "$PKG"