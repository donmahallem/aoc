#!/usr/bin/env bash

set -euo pipefail

if [ -z "${DATA_PW:-}" ]; then
  echo "Error: DATA_PW environment variable is not set." >&2
  exit 1
fi

script_dir="$(cd -- "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
data_dir="$script_dir/data/full"
archive="$script_dir/data.tar.gz.gpg"

# Use --batch --yes to avoid interactive prompts. Store only the full/ tree relative to test/data.
tar -C "$script_dir/data" -cvzf - full | gpg --batch --yes --pinentry-mode loopback --passphrase "$DATA_PW" --symmetric --cipher-algo AES256 > "$archive"