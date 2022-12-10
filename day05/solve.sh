#!/usr/bin/env bash

set -o errexit -o nounset -o pipefail
set -o xtrace

(
  cd day05
  go build organize-supplies.go

  ./get-input.sh
  cat puzzle-input.txt | ./organize-supplies
  cat puzzle-input.txt | ./organize-supplies --instrver 9001
)
