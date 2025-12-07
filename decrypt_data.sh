#!/usr/bin/env bash

if [ -z "$DATA_PW" ]; then
  echo "Error: DATA_PW environment variable is not set."
  exit 1
fi

mkdir -p ./data

cat data.tar.gz.gpg | gpg --batch --yes --pinentry-mode loopback --passphrase "$DATA_PW" --decrypt | tar -xvzf - -C ./