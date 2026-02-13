#!/usr/bin/env bash

set -euo pipefail

if [ -z "${DATA_PW:-}" ]; then
  echo "Error: DATA_PW environment variable is not set." >&2
  exit 1
fi

script_dir="$(cd -- "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
archive="$script_dir/data.tar.gz.gpg"
target_dir="$script_dir/data"

mkdir -p "$target_dir"

gpg --batch --yes --pinentry-mode loopback --passphrase "$DATA_PW" --decrypt "$archive" \
  | tar -xvzf - -C "$target_dir"