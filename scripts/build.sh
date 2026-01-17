#!/usr/bin/env bash
set -euo pipefail
source scripts/env.sh

mkdir -p build

echo "==> build orca"
go build \
  -ldflags="-s -w" \
  -o build/orca 

echo "✔ build/orca"