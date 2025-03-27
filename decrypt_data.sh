#!/usr/bin/env bash
mkdir -p ./data

cat data.tar.gz.gpg | gpg --batch --yes --pinentry-mode loopback --passphrase "$DATA_PW" --decrypt | tar -xvzf - -C ./