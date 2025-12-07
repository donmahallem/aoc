#!/usr/bin/env bash

if [ -z "$DATA_PW" ]; then
  echo "Error: DATA_PW environment variable is not set."
  exit 1
fi

# Use --batch --yes to avoid interactive prompts
# Use --passphrase-fd 0 to read password from stdin (safer) or keep using --passphrase
tar -cvzf - ./data | gpg --batch --yes --pinentry-mode loopback --passphrase "$DATA_PW" --symmetric --cipher-algo AES256 > data.tar.gz.gpg