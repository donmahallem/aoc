# Test Generator for several languages

Data from `data.json` is used to auto‑generate test files for the various
languages in this repository. The generator helps avoid duplicating the
same test setup across implementations.

## Generating the tests

The generator reads `test/data.json` and writes the test files for Go, C, and
C++.

Run either of the following:

- From `test/testgen`: `go run ./ --config ./config.json` regenerates the
  tests for all languages.

Use `--dry-run` to log the actions without writing any files.

## Test data

Full test inputs are stored encrypted in `data.tar.gz.gpg`. Decrypt them with
`bash ./test/decrypt_data.sh` (requires `DATA_PW`).
Go and Python tests whose input file is missing are skipped, not failed.
