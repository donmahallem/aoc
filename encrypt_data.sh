#!/usr/bin/env bash
tar -cvzf - ./data | gpg  --batch --yes --pinentry-mode loopback --passphrase "$DATA_PW" --symmetric > data.tar.gz.gpg 