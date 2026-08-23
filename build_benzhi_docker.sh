#!/bin/sh
set -eu
platform=${1:-linux/arm64}
tag=${2:-field-observatory:local}
docker buildx build --platform "$platform" -f benzhi.Dockerfile -t "$tag" --load .
