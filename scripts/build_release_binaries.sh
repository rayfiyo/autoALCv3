#!/usr/bin/env bash
set -euo pipefail

mkdir -p bin

targets=(
  "windows 386 .exe"
  "windows amd64 .exe"
  "windows arm .exe"
  "windows arm64 .exe"
  "linux 386"
  "linux amd64"
  "linux arm"
  "linux arm64"
  "darwin amd64"
  "darwin arm64"
)

for target in "${targets[@]}"; do
  read -r goos goarch ext <<<"${target}"
  output="bin/autoALCv3_${goos}_${goarch}${ext:-}"
  echo "building ${output}"
  GOOS="${goos}" GOARCH="${goarch}" go build -o "${output}" .
done
