#!/usr/bin/env bash
set -Eeuo pipefail
IMAGE_NAME_PREFIX=szpp-judge-image

if [[ $# -ne 1 ]]; then
  echo >&2 "Usage: $0 DOCKER_DIR"
  exit 1
fi

dir="$1"

if [[ ! -f "$dir/Dockerfile" ]]; then
  echo >&2 "Not a docker dir: File not found: $dir/Dockerfile"
  exit 1
fi

name=$(basename "$dir")
name="${name#_}"  # trim prefix '_' (e.g. "_szpprun" -> "szpprun")

cmd="docker build -t ${IMAGE_NAME_PREFIX}-${name}:latest $dir"
echo "$cmd"
$cmd
