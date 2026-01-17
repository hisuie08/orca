#!/usr/bin/env bash
set -euo pipefail

echo "==> debug orca (dlv)"
dlv debug . -- \
  init