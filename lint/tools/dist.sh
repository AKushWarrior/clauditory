#!/usr/bin/env bash
# Builds clauditory-lint for every supported platform into DIR, compressed
# with xz, with a SHA256SUMS file the launcher uses to name its cache entry.
# Usage: lint/tools/dist.sh DIR
set -euo pipefail
out="${1:?usage: dist.sh DIR}"
root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
mkdir -p "$out"
rm -f "$out"/clauditory-lint-* "$out/SHA256SUMS"
targets="linux/amd64 linux/arm64 darwin/arm64 darwin/amd64 windows/amd64"
for t in $targets; do
  os=${t%/*} arch=${t#*/}
  name="clauditory-lint-$os-$arch"
  [ "$os" = windows ] && name="$name.exe"
  (cd "$root" && CGO_ENABLED=0 GOOS=$os GOARCH=$arch go build -trimpath -ldflags "-s -w" -o "$out/$name" ./cmd/clauditory-lint)
  xz -9e -T1 -f "$out/$name"
done
(cd "$out" && sha256sum clauditory-lint-*.xz > SHA256SUMS)
