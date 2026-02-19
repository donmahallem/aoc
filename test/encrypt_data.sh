#!/usr/bin/env bash

set -euo pipefail

if [ -z "${DATA_PW:-}" ]; then
  echo "Error: DATA_PW environment variable is not set." >&2
  exit 1
fi

script_dir="$(cd -- "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
data_dir="$script_dir/data/full"
archive="$script_dir/data.tar.gz.gpg"

# Create a temporary copy of the data, normalize text file line endings (CRLF -> LF),
# then archive & encrypt that normalized copy. This avoids packaging CRLF files.
tmpdir="$(mktemp -d)"
trap 'rm -rf "$tmpdir"' EXIT
cp -a "$script_dir/data/full" "$tmpdir/full"

# Normalize common text files to LF only. Adjust pattern as needed.
find "$tmpdir/full" -type f -name "*.txt" -print0 \
  | xargs -0 -n1 sh -c 'awk "{ sub(/\r$/, \"\" ); print }" "$0" > "$0.tmp" && mv "$0.tmp" "$0"' \
  || true

# Use --batch --yes to avoid interactive prompts. Store only the full/ tree relative to test/data.
tar -C "$tmpdir" -cvzf - full | gpg --batch --yes --pinentry-mode loopback --passphrase "$DATA_PW" --symmetric --cipher-algo AES256 > "$archive"